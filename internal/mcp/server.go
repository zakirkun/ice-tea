package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/config"
	"github.com/zakirkun/ice-tea/internal/scanner"
)

// Server represents the Model Context Protocol server
type Server struct {
	config *config.Config
	logger *zap.SugaredLogger
	engine *scanner.Engine
}

// NewServer creates a new MCP server
func NewServer(cfg *config.Config, logger *zap.SugaredLogger, engine *scanner.Engine) *Server {
	return &Server{
		config: cfg,
		logger: logger,
		engine: engine,
	}
}

// Start launches the MCP HTTP server
func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	
	// standard JSON-RPC endpoint for MCP
	mux.HandleFunc("/jsonrpc", s.handleJSONRPC)

	s.logger.Infow("Starting MCP server", "address", addr)
	return http.ListenAndServe(addr, mux)
}

// handleJSONRPC processes incoming JSON-RPC requests from MCP clients
func (s *Server) handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.sendError(w, -32700, "Parse error", nil)
		return
	}

	method, _ := req["method"].(string)
	id := req["id"]

	s.logger.Debugw("Received MCP request", "method", method, "id", id)

	switch method {
	case "initialize":
		s.handleInitialize(w, id)
	case "tools/list":
		s.handleToolsList(w, id)
	case "tools/call":
		s.handleToolsCall(w, id, req["params"])
	default:
		s.sendError(w, -32601, "Method not found", id)
	}
}

func (s *Server) handleInitialize(w http.ResponseWriter, id interface{}) {
	resp := map[string]interface{}{
		"protocolVersion": "2024-11-05", // Example MCP protocol version
		"capabilities": map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		"serverInfo": map[string]interface{}{
			"name":    "ice-tea-mcp",
			"version": "1.0.0",
		},
	}
	s.sendResult(w, id, resp)
}

func (s *Server) handleToolsList(w http.ResponseWriter, id interface{}) {
	resp := map[string]interface{}{
		"tools": []map[string]interface{}{
			{
				"name":        "scan_directory",
				"description": "Run Ice Tea security scanner on a directory",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "Absolute path to the directory",
						},
					},
					"required": []string{"path"},
				},
			},
		},
	}
	s.sendResult(w, id, resp)
}

func (s *Server) handleToolsCall(w http.ResponseWriter, id interface{}, params interface{}) {
	p, ok := params.(map[string]interface{})
	if !ok {
		s.sendError(w, -32602, "Invalid params", id)
		return
	}

	name, _ := p["name"].(string)
	args, _ := p["arguments"].(map[string]interface{})

	if name == "scan_directory" {
		path, _ := args["path"].(string)
		
		s.logger.Infow("MCP requested scan", "path", path)
		findings, err := s.engine.Run(context.Background(), path)
		
		if err != nil {
			s.sendResult(w, id, map[string]interface{}{
				"content": []map[string]interface{}{
					{"type": "text", "text": fmt.Sprintf("Scan failed: %v", err)},
				},
				"isError": true,
			})
			return
		}

		// Return JSON representation of findings
		findingsJSON, _ := json.MarshalIndent(findings, "", "  ")

		s.sendResult(w, id, map[string]interface{}{
			"content": []map[string]interface{}{
				{"type": "text", "text": string(findingsJSON)},
			},
		})
		return
	}

	s.sendError(w, -32601, "Tool not found", id)
}

func (s *Server) sendResult(w http.ResponseWriter, id interface{}, result interface{}) {
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result":  result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) sendError(w http.ResponseWriter, code int, message string, id interface{}) {
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
