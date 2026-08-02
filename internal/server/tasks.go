package server

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/dipsylala/veracode-mcp/internal/types"
)

const (
	// defaultTaskTTLMs comfortably exceeds pipeline-scan's own 70-minute internal timeout
	// (see waitForScan in mcp_tools/pipeline_scan.go) so a task doesn't expire mid-scan.
	defaultTaskTTLMs      int64 = 90 * 60 * 1000
	maxTaskTTLMs          int64 = 120 * 60 * 1000
	defaultPollIntervalMs int64 = 10000
)

// ErrTaskNotFound indicates the requested task ID is unknown or has expired.
var ErrTaskNotFound = errors.New("task not found")

// ErrTaskTerminal indicates a cancel was attempted on a task that has already
// reached a terminal status.
var ErrTaskTerminal = errors.New("task is already in a terminal status")

// taskEntry holds the live state for a single task.
type taskEntry struct {
	mu         sync.Mutex
	task       types.Task
	createdAt  time.Time
	result     *types.CallToolResult
	cancelFunc func() error
	done       chan struct{}
}

// snapshot returns a copy of the task's current public state.
func (e *taskEntry) snapshot() types.Task {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.task
}

// setCancelFunc registers the function used to stop the task's underlying work.
func (e *taskEntry) setCancelFunc(fn func() error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.cancelFunc = fn
}

// transition moves the task to a new status, storing the final result and
// closing the done channel exactly once a terminal status is reached.
// A no-op if the task is already terminal.
func (e *taskEntry) transition(status, statusMessage string, result *types.CallToolResult) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if isTerminalStatus(e.task.Status) {
		return
	}
	e.task.Status = status
	e.task.StatusMessage = statusMessage
	e.task.LastUpdatedAt = time.Now().UTC().Format(time.RFC3339)
	e.result = result
	if isTerminalStatus(status) {
		close(e.done)
	}
}

func isTerminalStatus(status string) bool {
	switch status {
	case types.TaskStatusCompleted, types.TaskStatusFailed, types.TaskStatusCancelled:
		return true
	default:
		return false
	}
}

// TaskManager tracks in-flight and completed tasks for task-augmented tool calls.
// Tasks live in memory only, which matches this server's model: one persistent
// stdio session per client process, no cross-process or cross-restart durability.
type TaskManager struct {
	mu    sync.Mutex
	tasks map[string]*taskEntry
}

// NewTaskManager creates an empty task manager.
func NewTaskManager() *TaskManager {
	return &TaskManager{tasks: make(map[string]*taskEntry)}
}

// Create registers a new task in the "working" status and returns its entry.
func (tm *TaskManager) Create(requestedTTL *int64) *taskEntry {
	ttl := defaultTaskTTLMs
	if requestedTTL != nil && *requestedTTL > 0 {
		ttl = *requestedTTL
		if ttl > maxTaskTTLMs {
			ttl = maxTaskTTLMs
		}
	}
	pollInterval := defaultPollIntervalMs
	now := time.Now()

	entry := &taskEntry{
		task: types.Task{
			TaskID:        newTaskID(),
			Status:        types.TaskStatusWorking,
			CreatedAt:     now.UTC().Format(time.RFC3339),
			LastUpdatedAt: now.UTC().Format(time.RFC3339),
			TTL:           &ttl,
			PollInterval:  &pollInterval,
		},
		createdAt: now,
		done:      make(chan struct{}),
	}

	tm.mu.Lock()
	tm.tasks[entry.task.TaskID] = entry
	tm.mu.Unlock()
	return entry
}

// Get retrieves a live (non-expired) task by ID.
func (tm *TaskManager) Get(taskID string) (*taskEntry, error) {
	tm.mu.Lock()
	entry, ok := tm.tasks[taskID]
	tm.mu.Unlock()
	if !ok {
		return nil, ErrTaskNotFound
	}
	if tm.expired(entry) {
		tm.mu.Lock()
		delete(tm.tasks, taskID)
		tm.mu.Unlock()
		return nil, ErrTaskNotFound
	}
	return entry, nil
}

func (tm *TaskManager) expired(entry *taskEntry) bool {
	entry.mu.Lock()
	ttl := entry.task.TTL
	entry.mu.Unlock()
	if ttl == nil {
		return false
	}
	return time.Since(entry.createdAt) > time.Duration(*ttl)*time.Millisecond
}

// Complete marks a task finished, using the tool result's IsError flag to choose
// between the "completed" and "failed" terminal statuses (per the Tasks spec:
// a tool result with isError=true moves the task to "failed").
func (tm *TaskManager) Complete(taskID string, result *types.CallToolResult) {
	entry, err := tm.Get(taskID)
	if err != nil {
		return
	}
	status := types.TaskStatusCompleted
	statusMessage := ""
	if result != nil && result.IsError {
		status = types.TaskStatusFailed
		statusMessage = firstContentText(result)
	}
	entry.transition(status, statusMessage, result)
}

// Fail marks a task as failed due to a handler-level error (distinct from a tool
// result with isError=true, which is handled by Complete).
func (tm *TaskManager) Fail(taskID string, message string) {
	entry, err := tm.Get(taskID)
	if err != nil {
		return
	}
	entry.transition(types.TaskStatusFailed, message, &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: message}},
		IsError: true,
	})
}

// Cancel stops a task's underlying work (best-effort) and marks it cancelled.
// Returns ErrTaskTerminal if the task has already reached a terminal status.
func (tm *TaskManager) Cancel(taskID string) (types.Task, error) {
	entry, err := tm.Get(taskID)
	if err != nil {
		return types.Task{}, err
	}

	entry.mu.Lock()
	if isTerminalStatus(entry.task.Status) {
		snapshot := entry.task
		entry.mu.Unlock()
		return snapshot, ErrTaskTerminal
	}
	cancelFunc := entry.cancelFunc
	entry.mu.Unlock()

	if cancelFunc != nil {
		_ = cancelFunc() // best-effort; the task is marked cancelled regardless of the outcome
	}

	entry.transition(types.TaskStatusCancelled, "The task was cancelled by request.", nil)
	return entry.snapshot(), nil
}

// WaitResult blocks until the task reaches a terminal status, then returns the
// same result a non-task-augmented call to the underlying tool would have.
func (tm *TaskManager) WaitResult(taskID string) (*types.CallToolResult, error) {
	entry, err := tm.Get(taskID)
	if err != nil {
		return nil, err
	}
	<-entry.done

	entry.mu.Lock()
	defer entry.mu.Unlock()
	if entry.result == nil {
		// Cancelled tasks have no tool result; synthesize one for tasks/result callers.
		return &types.CallToolResult{
			Content: []types.Content{{Type: "text", Text: entry.task.StatusMessage}},
			IsError: true,
		}, nil
	}
	return entry.result, nil
}

func firstContentText(result *types.CallToolResult) string {
	if result == nil || len(result.Content) == 0 {
		return ""
	}
	return result.Content[0].Text
}

// newTaskID generates a cryptographically random task ID. This server has no
// per-requestor auth context to bind tasks to, so unguessable IDs are the only
// thing preventing one requestor from reading or cancelling another's task
// (see the MCP Tasks spec's Security Considerations).
func newTaskID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand.Read failing is effectively unreachable on supported platforms;
		// fall back rather than panic so task creation can never fail outright.
		return hex.EncodeToString([]byte(time.Now().Format(time.RFC3339Nano)))
	}
	return hex.EncodeToString(b)
}
