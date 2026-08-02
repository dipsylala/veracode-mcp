package server

import (
	"testing"
	"time"

	"github.com/dipsylala/veracode-mcp/internal/types"
)

func TestTaskManager_CreateAndGet(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)

	got, err := tm.Get(entry.snapshot().TaskID)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	snap := got.snapshot()
	if snap.Status != types.TaskStatusWorking {
		t.Errorf("Expected status %q, got %q", types.TaskStatusWorking, snap.Status)
	}
	if snap.TaskID == "" {
		t.Error("Expected non-empty task ID")
	}
	if snap.TTL == nil || *snap.TTL != defaultTaskTTLMs {
		t.Errorf("Expected default TTL %d, got %v", defaultTaskTTLMs, snap.TTL)
	}
}

func TestTaskManager_Get_NotFound(t *testing.T) {
	tm := NewTaskManager()
	if _, err := tm.Get("does-not-exist"); err != ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskManager_Complete_Success(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	result := &types.CallToolResult{Content: []types.Content{{Type: "text", Text: "done"}}}
	tm.Complete(taskID, result)

	snap := entry.snapshot()
	if snap.Status != types.TaskStatusCompleted {
		t.Errorf("Expected status %q, got %q", types.TaskStatusCompleted, snap.Status)
	}

	got, err := tm.WaitResult(taskID)
	if err != nil {
		t.Fatalf("WaitResult failed: %v", err)
	}
	if got.Content[0].Text != "done" {
		t.Errorf("Expected result text 'done', got %q", got.Content[0].Text)
	}
}

func TestTaskManager_Complete_ToolError_MovesToFailed(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	result := &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: "boom"}},
		IsError: true,
	}
	tm.Complete(taskID, result)

	snap := entry.snapshot()
	if snap.Status != types.TaskStatusFailed {
		t.Errorf("Expected status %q, got %q", types.TaskStatusFailed, snap.Status)
	}
	if snap.StatusMessage != "boom" {
		t.Errorf("Expected statusMessage 'boom', got %q", snap.StatusMessage)
	}
}

func TestTaskManager_Fail(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	tm.Fail(taskID, "handler exploded")

	snap := entry.snapshot()
	if snap.Status != types.TaskStatusFailed {
		t.Errorf("Expected status %q, got %q", types.TaskStatusFailed, snap.Status)
	}

	result, err := tm.WaitResult(taskID)
	if err != nil {
		t.Fatalf("WaitResult failed: %v", err)
	}
	if !result.IsError {
		t.Error("Expected IsError=true for a failed task's result")
	}
}

func TestTaskManager_Cancel_CallsCancelFuncAndTransitions(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	called := false
	entry.setCancelFunc(func() error {
		called = true
		return nil
	})

	task, err := tm.Cancel(taskID)
	if err != nil {
		t.Fatalf("Cancel failed: %v", err)
	}
	if !called {
		t.Error("Expected cancel func to be invoked")
	}
	if task.Status != types.TaskStatusCancelled {
		t.Errorf("Expected status %q, got %q", types.TaskStatusCancelled, task.Status)
	}

	// tasks/result on a cancelled task should not hang, and should report an error result.
	result, err := tm.WaitResult(taskID)
	if err != nil {
		t.Fatalf("WaitResult failed: %v", err)
	}
	if !result.IsError {
		t.Error("Expected a cancelled task's synthesized result to be an error")
	}
}

func TestTaskManager_Cancel_AlreadyTerminal(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	tm.Complete(taskID, &types.CallToolResult{Content: []types.Content{{Type: "text", Text: "done"}}})

	if _, err := tm.Cancel(taskID); err != ErrTaskTerminal {
		t.Errorf("Expected ErrTaskTerminal, got %v", err)
	}
}

func TestTaskManager_WaitResult_BlocksUntilTerminal(t *testing.T) {
	tm := NewTaskManager()
	entry := tm.Create(nil)
	taskID := entry.snapshot().TaskID

	go func() {
		time.Sleep(50 * time.Millisecond)
		tm.Complete(taskID, &types.CallToolResult{Content: []types.Content{{Type: "text", Text: "finished late"}}})
	}()

	start := time.Now()
	result, err := tm.WaitResult(taskID)
	if err != nil {
		t.Fatalf("WaitResult failed: %v", err)
	}
	if time.Since(start) < 40*time.Millisecond {
		t.Error("Expected WaitResult to block until the task completed")
	}
	if result.Content[0].Text != "finished late" {
		t.Errorf("Expected 'finished late', got %q", result.Content[0].Text)
	}
}

func TestTaskManager_Expiry(t *testing.T) {
	tm := NewTaskManager()
	ttl := int64(10) // 10ms
	entry := tm.Create(&ttl)
	taskID := entry.snapshot().TaskID

	time.Sleep(30 * time.Millisecond)

	if _, err := tm.Get(taskID); err != ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound for expired task, got %v", err)
	}
}

func TestTaskManager_TTL_ClampedToMax(t *testing.T) {
	tm := NewTaskManager()
	huge := maxTaskTTLMs * 10
	entry := tm.Create(&huge)

	snap := entry.snapshot()
	if snap.TTL == nil || *snap.TTL != maxTaskTTLMs {
		t.Errorf("Expected TTL clamped to %d, got %v", maxTaskTTLMs, snap.TTL)
	}
}
