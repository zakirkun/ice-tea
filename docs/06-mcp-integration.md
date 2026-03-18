# Model Context Protocol (MCP) Integration

## Overview

The **Model Context Protocol (MCP)** is an open standard (introduced by Anthropic, November 2024) that defines how AI systems communicate with external tools and data sources. Often called "USB-C for AI," MCP provides a standardized way for LLM agents to discover and invoke tools.

Ice Tea integrates with MCP in two ways:
1. **As an MCP Server**: Exposes scan capabilities as MCP Tools for agentic workflows
2. **As an MCP Client**: Optionally consumes external MCP resources (e.g., vulnerability databases)

## MCP Architecture

```
┌─────────────────────────────────────────────┐
│  MCP Host (AI Agent / IDE / CLI)            │
│                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
│  │  Client 1 │  │  Client 2 │  │  Client 3 │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  │
└───────┼──────────────┼──────────────┼───────┘
        │              │              │
        ▼              ▼              ▼
┌───────────┐  ┌───────────┐  ┌───────────┐
│ Ice Tea   │  │ Database  │  │ Other     │
│ MCP Server│  │ MCP Server│  │ MCP Server│
└───────────┘  └───────────┘  └───────────┘
```

### MCP Concepts

| Concept | Description | Ice Tea Usage |
|---------|-------------|---------------|
| **Tools** | Functions the server exposes for the client to call | `scan`, `listRules`, `getSkills` |
| **Resources** | Structured data the client can read | SKILL definitions, rule metadata |
| **Prompts** | Pre-defined prompt templates | Security analysis prompts |

## Ice Tea as MCP Server

### Transport Options

1. **stdio**: Default for CLI integration — reads JSON-RPC from stdin, writes to stdout
2. **HTTP/SSE**: For remote/networked setups

### Exposed Tools

```json
{
  "tools": [
    {
      "name": "scan",
      "description": "Scan a file or directory for security vulnerabilities",
      "inputSchema": {
        "type": "object",
        "properties": {
          "target": {
            "type": "string",
            "description": "Path to file or directory to scan"
          },
          "languages": {
            "type": "array",
            "items": { "type": "string" },
            "description": "Filter by programming languages"
          },
          "severity": {
            "type": "string",
            "enum": ["critical", "high", "medium", "low"],
            "description": "Minimum severity threshold"
          },
          "enableLLM": {
            "type": "boolean",
            "description": "Enable LLM deep reasoning engine"
          }
        },
        "required": ["target"]
      }
    },
    {
      "name": "listRules",
      "description": "List available vulnerability detection rules",
      "inputSchema": {
        "type": "object",
        "properties": {
          "language": {
            "type": "string",
            "description": "Filter rules by language"
          },
          "category": {
            "type": "string",
            "description": "Filter by OWASP category"
          }
        }
      }
    },
    {
      "name": "analyzeSnippet",
      "description": "Analyze a code snippet for vulnerabilities",
      "inputSchema": {
        "type": "object",
        "properties": {
          "code": { "type": "string" },
          "language": { "type": "string" },
          "context": { "type": "string" }
        },
        "required": ["code", "language"]
      }
    }
  ]
}
```

### Exposed Resources

```json
{
  "resources": [
    {
      "uri": "icetea://skills",
      "name": "Available SKILLs",
      "description": "List of all loaded vulnerability detection skills",
      "mimeType": "application/json"
    },
    {
      "uri": "icetea://skills/{skillId}",
      "name": "SKILL Detail",
      "description": "Full content and instructions for a specific SKILL",
      "mimeType": "text/markdown"
    },
    {
      "uri": "icetea://rules",
      "name": "Detection Rules",
      "description": "All active detection rules with metadata",
      "mimeType": "application/json"
    }
  ]
}
```

### Exposed Prompts

```json
{
  "prompts": [
    {
      "name": "security-review",
      "description": "Deep security review of a code file",
      "arguments": [
        { "name": "file", "description": "Path to file to review", "required": true },
        { "name": "focus", "description": "Specific vulnerability type to focus on" }
      ]
    },
    {
      "name": "fix-vulnerability",
      "description": "Generate a fix for a detected vulnerability",
      "arguments": [
        { "name": "finding", "description": "JSON finding object", "required": true }
      ]
    }
  ]
}
```

## MCP Communication Protocol

Communication uses **JSON-RPC 2.0**:

### Example: Client Calls Scan Tool

**Request (Client → Server):**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "scan",
    "arguments": {
      "target": "./src",
      "severity": "high",
      "enableLLM": false
    }
  }
}
```

**Response (Server → Client):**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"findings\": [...], \"summary\": {\"total\": 5, \"critical\": 1, \"high\": 4}}"
      }
    ]
  }
}
```

## Go Implementation Approach

```go
// MCP server implementation using stdio transport
package mcp

import (
    "bufio"
    "encoding/json"
    "os"
)

type MCPServer struct {
    scanner *Scanner  // Ice Tea scan engine
    skills  *SkillManager
}

func (s *MCPServer) Start() error {
    reader := bufio.NewReader(os.Stdin)

    for {
        line, err := reader.ReadBytes('\n')
        if err != nil {
            return err
        }

        var request JSONRPCRequest
        if err := json.Unmarshal(line, &request); err != nil {
            s.sendError(request.ID, -32700, "Parse error")
            continue
        }

        response := s.handleRequest(request)
        s.sendResponse(response)
    }
}

func (s *MCPServer) handleRequest(req JSONRPCRequest) JSONRPCResponse {
    switch req.Method {
    case "initialize":
        return s.handleInitialize(req)
    case "tools/list":
        return s.handleToolsList(req)
    case "tools/call":
        return s.handleToolsCall(req)
    case "resources/list":
        return s.handleResourcesList(req)
    case "resources/read":
        return s.handleResourcesRead(req)
    case "prompts/list":
        return s.handlePromptsList(req)
    case "prompts/get":
        return s.handlePromptsGet(req)
    default:
        return s.errorResponse(req.ID, -32601, "Method not found")
    }
}
```

## Security Considerations

1. **OAuth 2.0**: For HTTP transport, implement OAuth 2.0 Resource Server pattern
2. **Input Validation**: Validate all tool arguments before processing
3. **Path Traversal Prevention**: Sanitize file paths in scan targets
4. **Rate Limiting**: Prevent abuse of scan tool in multi-user setups
5. **Consent Flow**: Request user confirmation for potentially destructive operations
