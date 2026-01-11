# MCP Component Testing Guide

This document explains how to test the Model Context Protocol (MCP) server implementation for the Veracode MCP Go project.

## Overview

The MCP server provides a JSON-RPC 2.0 interface that allows AI assistants and other clients to interact with the Veracode API through standardized tools. This includes:

- **Protocol Handling**: JSON-RPC 2.0 request/response processing
- **Tool Discovery**: Dynamic tool registration and listing
- **Tool Execution**: Calling Veracode API tools with parameter validation
- **API Integration**: End-to-end testing with real Veracode APIs

## Test Structure

### 1. Unit Tests ([server_test.go](server_test.go))

Tests the core MCP server functionality:
- Tool definition loading from JSON
- Tool schema conversion to MCP format
- Server initialization
- Tool call handling with validation

```powershell
# Run unit tests only
go test -v -run "TestLoad|TestToMCP|TestServer" -short
```

### 2. Tool Registry Tests ([tools/registry_test.go](tools/registry_test.go))

Tests the tool registration system:
- Tool registration and retrieval
- Concurrent access safety
- Handler registration and execution

```powershell
# Run registry tests
go test -v ./tools -run "Registry"
```

### 3. Tool Integration Tests ([tools/integration_test.go](tools/integration_test.go))

Tests that actual tool implementations auto-register correctly:
- Dynamic findings tool
- Static findings tool
- API health tool

```powershell
# Run tool integration tests
go test -v ./tools -run "TestActual"
```

### 4. MCP Integration Tests ([mcp_integration_test.go](mcp_integration_test.go))

**End-to-end tests** that simulate MCP client interactions:

#### Test Coverage

| Test | Purpose | API Call |
|------|---------|----------|
| `TestMCPServerInitialization` | Verify server starts with all tools loaded | No |
| `TestMCPProtocolInitialize` | Test MCP initialize handshake | No |
| `TestMCPProtocolListTools` | Test tools/list method returns all tools | No |
| `TestMCPToolCall_APIHealth` | Call api-health tool via MCP protocol | No |
| `TestMCPToolCall_DynamicFindings` | Call get-dynamic-findings with real API | **Yes** |
| `TestMCPToolCall_StaticFindings` | Call get-static-findings with real API | **Yes** |
| `TestMCPToolCall_MissingRequiredParam` | Verify parameter validation | No |
| `TestMCPToolCall_UnknownTool` | Verify error handling for unknown tools | No |

## Running MCP Integration Tests

### Without API Credentials (Fast)

Most tests run without requiring Veracode credentials:

```powershell
# Run all MCP tests (skips API-dependent tests)
go test -v -run "TestMCP" -timeout 120s

# Results:
# ✓ TestMCPServerInitialization
# ✓ TestMCPProtocolInitialize  
# ✓ TestMCPProtocolListTools
# ✓ TestMCPToolCall_APIHealth
# ⏭️ TestMCPToolCall_DynamicFindings (skipped)
# ⏭️ TestMCPToolCall_StaticFindings (skipped)
# ✓ TestMCPToolCall_MissingRequiredParam
# ✓ TestMCPToolCall_UnknownTool
```

### With API Credentials (Full Integration)

To test actual Veracode API calls through MCP:

```powershell
# Set credentials (already configured in your environment)
# $env:VERACODE_API_ID = "your-api-id"
# $env:VERACODE_API_KEY = "your-api-key"

# Run all MCP integration tests including API calls
go test -v -run "TestMCP" -timeout 120s

# All 8 tests should run (6 pass, 2 may take 30+ seconds for API calls)
```

## Test Examples

### Example 1: Testing MCP Protocol Handshake

```go
// Create server
server, err := NewMCPServer()

// Send initialize request
req := &JSONRPCRequest{
    JSONRPC: "2.0",
    ID:      1,
    Method:  "initialize",
    Params:  initParamsJSON,
}

// Handle request
resp := server.HandleRequest(req)

// Verify protocol version and capabilities
initResult := resp.Result.(*InitializeResult)
assert.Equal(t, "2024-11-05", initResult.ProtocolVersion)
```

### Example 2: Testing Tool Listing

```go
// Request tools/list
req := &JSONRPCRequest{
    Method: "tools/list",
}

resp := server.HandleRequest(req)
listResult := resp.Result.(*ListToolsResult)

// Verify tools are present
assert.Contains(t, listResult.Tools, "api-health")
assert.Contains(t, listResult.Tools, "get-dynamic-findings")
```

### Example 3: Testing Tool Execution with Real API

```go
// Call get-static-findings tool
req := &JSONRPCRequest{
    Method: "tools/call",
    Params: CallToolParams{
        Name: "get-static-findings",
        Arguments: map[string]interface{}{
            "application_guid": "65c204e5-a74c-4b68-a62a-4bfdc08e27af",
            "size": 5,
            "severity": []interface{}{"High", "Very High"},
        },
    },
}

// Execute and verify
resp := server.HandleRequest(req)
result := resp.Result.(*CallToolResult)

// Parse findings from response
var findings map[string]interface{}
json.Unmarshal([]byte(result.Content[0].Text), &findings)

// Verify API response structure
assert.Contains(t, findings, "page")
assert.Contains(t, findings, "_embedded")
```

## What Gets Tested

### ✅ MCP Protocol Compliance

- **JSON-RPC 2.0**: Request/response format
- **Initialize handshake**: Protocol version negotiation
- **Capabilities**: Tools, resources, prompts
- **Tool listing**: Schema generation
- **Tool execution**: Parameter validation and invocation

### ✅ Tool System

- **Auto-registration**: Tools register via `init()` functions
- **Schema generation**: OpenAPI-style JSON schemas
- **Parameter validation**: Required/optional fields, type checking
- **Handler routing**: Tool name → handler function mapping
- **Error handling**: Unknown tools, missing parameters

### ✅ API Integration

- **Authentication**: HMAC signature generation
- **Request execution**: HTTP client with proper headers
- **Response parsing**: JSON deserialization
- **Error handling**: API errors, timeouts, invalid responses
- **Data validation**: Verify spec fixes (severity as int, CWE as object)

## Test Data

Tests use the following Veracode application:

- **Name**: MCPVerademo-NET
- **GUID**: `65c204e5-a74c-4b68-a62a-4bfdc08e27af`
- **Purpose**: .NET demo application with known security findings

## Continuous Integration

Add to your CI/CD pipeline:

```yaml
# GitHub Actions example
- name: Run MCP Tests
  run: |
    go test -v -run "TestMCP" -timeout 120s
  env:
    VERACODE_API_ID: ${{ secrets.VERACODE_API_ID }}
    VERACODE_API_KEY: ${{ secrets.VERACODE_API_KEY }}
```

## Debugging Failed Tests

### Enable Verbose Logging

```powershell
# See full request/response details
go test -v -run "TestMCP" -timeout 120s 2>&1 | Tee-Object test-output.log
```

### Check API Credentials

```powershell
# Verify credentials are set
Write-Host "API ID: $env:VERACODE_API_ID"
Write-Host "API Key length: $($env:VERACODE_API_KEY.Length)"
```

### Test Individual Components

```powershell
# Test just the API health tool
go test -v -run "TestMCPToolCall_APIHealth"

# Test just parameter validation
go test -v -run "TestMCPToolCall_Missing"

# Test with real API (requires credentials)
go test -v -run "TestMCPToolCall_Static"
```

## Performance

Typical test execution times:

| Test Suite | Duration | API Calls |
|------------|----------|-----------|
| Unit tests (server_test.go) | ~50ms | 0 |
| Registry tests | ~100ms | 0 |
| Tool integration tests | ~50ms | 0 |
| MCP tests (no API) | ~300ms | 0 |
| MCP tests (with API) | ~5-10s | 2-4 |

## Best Practices

1. **Run fast tests first**: Use `-short` flag to skip integration tests during development
2. **Use real test data**: The MCPVerademo-NET application is available in all Veracode environments
3. **Verify spec fixes**: Check that API responses match updated OpenAPI specs (int severity, object CWE)
4. **Test error paths**: Verify validation and error handling work correctly
5. **Isolate API tests**: Use skip conditions for tests requiring credentials

## Next Steps

To add tests for new MCP tools:

1. **Add tool implementation**: Create new file in `tools/` directory
2. **Auto-registration**: Tool registers via `init()` function
3. **Add to expected tools**: Update `mcp_integration_test.go` expectedTools list
4. **Create specific test**: Add `TestMCPToolCall_YourNewTool()` function
5. **Verify in CI**: Ensure tests run in automated pipeline

## Related Documentation

- [MCP Tools README](tools/README.md) - Tool implementation guide
- [Tool Testing Guide](tools/TESTING.md) - Detailed tool testing documentation
- [API Integration Tests](api/README.md) - Low-level API testing

## Summary

The best way to test the MCP component is through the comprehensive integration tests in `mcp_integration_test.go`. These tests:

✅ **Don't require external dependencies** for basic functionality
✅ **Test the full MCP protocol flow** from request to response  
✅ **Validate real API integration** when credentials are available
✅ **Verify data correctness** including OpenAPI spec fixes
✅ **Run fast** (under 1 second without API calls)
✅ **Easy to extend** for new tools and features

Run with:
```powershell
go test -v -run "TestMCP" -timeout 120s
```
