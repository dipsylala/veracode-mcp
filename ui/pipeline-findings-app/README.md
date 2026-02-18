# Pipeline Findings MCP App

This directory contains the interactive UI for the Veracode Pipeline Findings tool, built with React and the MCP Apps SDK.

## Structure

- `src/mcp-app.tsx` - Main React component
- `src/types.ts` - TypeScript type definitions matching Go structs
- `src/mcp-app.module.css` - Component styles
- `src/global.css` - Global styles
- `mcp-app.html` - HTML template
- `vite.config.ts` - Vite build configuration
- `dist/mcp-app.html` - Built single-file HTML (generated)

## Building

From this directory:

```powershell
# Install dependencies (first time only)
npm install

# Build the UI
npm run build

# Watch mode (rebuild on changes)
npm run watch
```

Or from the repository root:

```powershell
# Build UI and Go server
.\build-all.ps1

# Build UI only
.\build-all.ps1 -UIOnly
```

## How It Works

1. **Tool Registration**: The Go server registers the `pipeline-findings` tool with UI metadata (`_meta.ui.resourceUri`)
2. **Resource Registration**: The server also registers the UI resource at `ui://pipeline-findings/app.html`
3. **Tool Execution**: When Claude calls the tool, it receives both the data (as JSON) and the UI resource URI
4. **UI Rendering**: The MCP host (Claude Desktop) renders the UI in an iframe and passes the tool result to it
5. **Data Display**: The React app receives the data via `ontoolresult` handler and displays it in a table

## Development

The UI uses:

- **MCP Apps SDK** (`@modelcontextprotocol/ext-apps`) for connecting to the MCP host
- **React** with TypeScript for the UI
- **Vite** with `vite-plugin-singlefile` to bundle everything into a single HTML file
- **Host CSS variables** for theme integration with Claude Desktop

## Customization

To modify the table display, edit `src/mcp-app.tsx`. The data structure is defined in `src/types.ts` and matches the Go `MCPFindingsResponse` struct.
