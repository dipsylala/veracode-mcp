package server

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tools "github.com/dipsylala/veracode-mcp/internal/tool_registry"
	"github.com/dipsylala/veracode-mcp/internal/types"
)

// TestHandleCallTool_TaskRequired_WithoutTask_Rejected exercises the router's
// "required" branch. No shipped tool currently declares taskSupport: "required"
// (pipeline-scan is "optional" so it still works for clients, like Claude's,
// that don't send a "task" param — see pipeline-scan's fallback to its own
// background+poll mode), so this test temporarily forces pipeline-scan's
// declared taskSupport to "required" via the tools.json fixture to reach that
// branch; the underlying handler is never invoked since the router rejects
// the call before dispatch.
func TestHandleCallTool_TaskRequired_WithoutTask_Rejected(t *testing.T) {
	toolsJSONPath := filepath.Join("..", "..", "tools.json")
	// #nosec G304 -- toolsJSONPath is a fixed test fixture path, not user input
	originalToolsJSON, err := os.ReadFile(toolsJSONPath)
	if err != nil {
		t.Fatalf("Failed to read tools.json fixture: %v", err)
	}
	defer tools.SetToolsJSON(originalToolsJSON)

	forcedRequiredJSON := strings.Replace(string(originalToolsJSON),
		`"taskSupport": "optional"`, `"taskSupport": "required"`, 1)
	if forcedRequiredJSON == string(originalToolsJSON) {
		t.Fatal("Failed to force pipeline-scan's taskSupport to \"required\" in the tools.json fixture")
	}
	tools.SetToolsJSON([]byte(forcedRequiredJSON))

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	params := types.CallToolParams{
		Name:      "pipeline-scan",
		Arguments: map[string]interface{}{"application_path": t.TempDir()},
	}
	paramsJSON, _ := json.Marshal(params)

	_, err = server.handleCallTool(paramsJSON)
	if err == nil {
		t.Fatal("Expected an error when calling a task-required tool without task augmentation")
	}

	var coded *rpcCodedError
	if !errors.As(err, &coded) {
		t.Fatalf("Expected *rpcCodedError, got %T: %v", err, err)
	}
	if coded.code != -32601 {
		t.Errorf("Expected code -32601, got %d", coded.code)
	}
}

func TestHandleCallTool_TaskForbidden_WithTask_Rejected(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// api-health declares no execution.taskSupport, so it defaults to "forbidden".
	params := types.CallToolParams{
		Name:      "api-health",
		Arguments: map[string]interface{}{},
		Task:      &types.TaskParams{},
	}
	paramsJSON, _ := json.Marshal(params)

	_, err = server.handleCallTool(paramsJSON)
	if err == nil {
		t.Fatal("Expected an error when task-augmenting a tool that doesn't support it")
	}

	var coded *rpcCodedError
	if !errors.As(err, &coded) {
		t.Fatalf("Expected *rpcCodedError, got %T: %v", err, err)
	}
	if coded.code != -32601 {
		t.Errorf("Expected code -32601, got %d", coded.code)
	}
}

func TestHandleCallTool_TaskAugmented_PipelineScan_CreatesTaskAndCompletes(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	params := types.CallToolParams{
		Name:      "pipeline-scan",
		Arguments: map[string]interface{}{"application_path": t.TempDir()},
		Task:      &types.TaskParams{},
	}
	paramsJSON, _ := json.Marshal(params)

	rawResult, err := server.handleCallTool(paramsJSON)
	if err != nil {
		t.Fatalf("handleCallTool failed: %v", err)
	}

	created, ok := rawResult.(*types.CreateTaskResult)
	if !ok {
		t.Fatalf("Expected *types.CreateTaskResult, got %T", rawResult)
	}
	if created.Task.TaskID == "" {
		t.Error("Expected a non-empty task ID")
	}
	if created.Task.Status != types.TaskStatusWorking {
		t.Errorf("Expected initial status %q, got %q", types.TaskStatusWorking, created.Task.Status)
	}

	// No packaged artifact or (likely) veracode CLI in the test environment, so the
	// handler should fail fast rather than hang — verifying the task still reaches
	// a terminal state and tasks/result doesn't block forever.
	result, err := server.taskManager.WaitResult(created.Task.TaskID)
	if err != nil {
		t.Fatalf("WaitResult failed: %v", err)
	}
	if !result.IsError {
		t.Error("Expected an error result (no veracode CLI / packaged artifact in test env)")
	}
}

func TestHandleTasksGetRequest_NotFound(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := &types.JSONRPCRequest{Params: json.RawMessage(`{"taskId":"does-not-exist"}`)}
	resp := &types.JSONRPCResponse{}
	server.handleTasksGetRequest(req, resp)

	if resp.Error == nil {
		t.Fatal("Expected an error for unknown task ID")
	}
	if resp.Error.Code != -32602 {
		t.Errorf("Expected code -32602, got %d", resp.Error.Code)
	}
}

func TestHandleTasksGetRequest_MissingTaskID(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := &types.JSONRPCRequest{Params: json.RawMessage(`{}`)}
	resp := &types.JSONRPCResponse{}
	server.handleTasksGetRequest(req, resp)

	if resp.Error == nil {
		t.Fatal("Expected an error for missing taskId")
	}
	if resp.Error.Code != -32602 {
		t.Errorf("Expected code -32602, got %d", resp.Error.Code)
	}
}

func TestHandleTasksGetThenResult(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	entry := server.taskManager.Create(nil)
	taskID := entry.snapshot().TaskID
	server.taskManager.Complete(taskID, &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: "hello"}},
	})

	getReq := &types.JSONRPCRequest{Params: json.RawMessage(`{"taskId":"` + taskID + `"}`)}
	getResp := &types.JSONRPCResponse{}
	server.handleTasksGetRequest(getReq, getResp)
	if getResp.Error != nil {
		t.Fatalf("tasks/get failed: %v", getResp.Error)
	}
	task, ok := getResp.Result.(types.Task)
	if !ok {
		t.Fatalf("Expected types.Task, got %T", getResp.Result)
	}
	if task.Status != types.TaskStatusCompleted {
		t.Errorf("Expected status %q, got %q", types.TaskStatusCompleted, task.Status)
	}

	resultReq := &types.JSONRPCRequest{Params: json.RawMessage(`{"taskId":"` + taskID + `"}`)}
	resultResp := &types.JSONRPCResponse{}
	server.handleTasksResultRequest(resultReq, resultResp)
	if resultResp.Error != nil {
		t.Fatalf("tasks/result failed: %v", resultResp.Error)
	}
	result, ok := resultResp.Result.(*types.CallToolResult)
	if !ok {
		t.Fatalf("Expected *types.CallToolResult, got %T", resultResp.Result)
	}
	if result.Content[0].Text != "hello" {
		t.Errorf("Expected 'hello', got %q", result.Content[0].Text)
	}
}

func TestHandleTasksCancelRequest_AlreadyTerminal(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	entry := server.taskManager.Create(nil)
	taskID := entry.snapshot().TaskID
	server.taskManager.Complete(taskID, &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: "done"}},
	})

	req := &types.JSONRPCRequest{Params: json.RawMessage(`{"taskId":"` + taskID + `"}`)}
	resp := &types.JSONRPCResponse{}
	server.handleTasksCancelRequest(req, resp)

	if resp.Error == nil {
		t.Fatal("Expected an error cancelling an already-terminal task")
	}
	if resp.Error.Code != -32602 {
		t.Errorf("Expected code -32602, got %d", resp.Error.Code)
	}
}

func TestHandleTasksCancelRequest_Working(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	entry := server.taskManager.Create(nil)
	taskID := entry.snapshot().TaskID
	called := false
	entry.setCancelFunc(func() error {
		called = true
		return nil
	})

	req := &types.JSONRPCRequest{Params: json.RawMessage(`{"taskId":"` + taskID + `"}`)}
	resp := &types.JSONRPCResponse{}
	server.handleTasksCancelRequest(req, resp)

	if resp.Error != nil {
		t.Fatalf("tasks/cancel failed: %v", resp.Error)
	}
	if !called {
		t.Error("Expected the registered cancel func to be invoked")
	}
	task, ok := resp.Result.(types.Task)
	if !ok {
		t.Fatalf("Expected types.Task, got %T", resp.Result)
	}
	if task.Status != types.TaskStatusCancelled {
		t.Errorf("Expected status %q, got %q", types.TaskStatusCancelled, task.Status)
	}
}
