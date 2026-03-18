package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zakirkun/ice-tea/internal/mcp"
	"github.com/zakirkun/ice-tea/internal/scanner"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the Ice Tea MCP Server",
	Long:  `Starts a Model Context Protocol (MCP) compatible server over HTTP JSON-RPC for integration with AI assistants.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, _ := cmd.Flags().GetString("addr")

		// We need to initialize the engine to pass it to the MCP server
		// Some config overrides
		conf := getConfig()
		log := getLogger()

		engine := scanner.NewEngine(conf, log)
		if err := engine.Init(); err != nil {
			return fmt.Errorf("scan engine init failed: %w", err)
		}

		server := mcp.NewServer(conf, log, engine)
		return server.Start(addr)
	},
}

func init() {
	mcpCmd.Flags().StringP("addr", "a", "127.0.0.1:8080", "Address to bind the MCP server to")
	rootCmd.AddCommand(mcpCmd)
}
