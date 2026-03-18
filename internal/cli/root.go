package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/config"
)

var (
	cfgFile string
	verbose bool
	noColor bool
	logger  *zap.SugaredLogger
	cfg     *config.Config
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "ice-tea",
	Short: "AI-powered DevOps security scanner",
	Long: `Ice Tea is an AI-powered Static Application Security Testing (SAST) tool
that combines AST-based code analysis with LLM deep reasoning.

It scans source code for security vulnerabilities using a multi-engine approach:
  • Engine 1: Static pattern matching (AST + regex)
  • Engine 2: Data flow / taint analysis
  • Engine 3: LLM deep reasoning (chain-of-thought)

Designed for CI/CD integration (GitHub Actions, GitLab Runner) and
agentic AI workflows via Model Context Protocol (MCP).`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeApp()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .ice-tea.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")

	// Bind flags to viper
	viper.BindPFlag("output.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("output.color", rootCmd.PersistentFlags().Lookup("no-color"))
}

func initializeApp() error {
	// Handle color
	if noColor {
		color.NoColor = true
	}

	// Initialize config
	var err error
	cfg, err = config.Load(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize logger
	var zapCfg zap.Config
	if verbose || cfg.Output.Verbose {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
	}

	zapLogger, err := zapCfg.Build()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	logger = zapLogger.Sugar()

	return nil
}

// getLogger returns the application logger
func getLogger() *zap.SugaredLogger {
	if logger == nil {
		// Fallback logger for commands that don't run PersistentPreRun
		zapLogger, _ := zap.NewProduction()
		logger = zapLogger.Sugar()
	}
	return logger
}

// getConfig returns the application config
func getConfig() *config.Config {
	return cfg
}

// exitWithError prints an error and exits
func exitWithError(msg string, err error) {
	red := color.New(color.FgRed, color.Bold)
	red.Fprintf(os.Stderr, "Error: %s\n", msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  %v\n", err)
	}
	os.Exit(1)
}
