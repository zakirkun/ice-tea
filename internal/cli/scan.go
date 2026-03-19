package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zakirkun/ice-tea/internal/analyzer/llm/providers"
	"github.com/zakirkun/ice-tea/internal/finding"
	"github.com/zakirkun/ice-tea/internal/notifier/telegram"
	"github.com/zakirkun/ice-tea/internal/parser/goparser"
	"github.com/zakirkun/ice-tea/internal/parser/textparser"
	"github.com/zakirkun/ice-tea/internal/parser/treesitter"
	"github.com/zakirkun/ice-tea/internal/reporter"
	"github.com/zakirkun/ice-tea/internal/scanner"
)

var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Scan code for security vulnerabilities",
	Long: `Scan source code files or directories for security vulnerabilities.

Uses a multi-engine approach combining static pattern matching,
data flow analysis, and optional LLM deep reasoning.

Examples:
  ice-tea scan .
  ice-tea scan ./src --format sarif --output results.sarif
  ice-tea scan main.go --severity high
  ice-tea scan . --enable-llm --exclude-dir vendor,node_modules`,
	Args: cobra.MaximumNArgs(1),
	RunE: runScan,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Scan-specific flags
	scanCmd.Flags().StringP("format", "f", "console", "output format (console, sarif, gitlab, json)")
	scanCmd.Flags().StringP("output", "o", "", "output file path (default: stdout)")
	scanCmd.Flags().StringP("severity", "s", "medium", "minimum severity threshold (critical, high, medium, low, info)")
	scanCmd.Flags().String("confidence", "medium", "minimum confidence threshold (high, medium, low)")
	scanCmd.Flags().StringSlice("exclude-dir", nil, "directories to exclude (comma-separated)")
	scanCmd.Flags().StringSlice("exclude-file", nil, "file patterns to exclude (comma-separated)")
	scanCmd.Flags().StringSlice("language", nil, "languages to scan (default: auto-detect)")
	scanCmd.Flags().IntP("concurrency", "c", 4, "number of concurrent workers")
	scanCmd.Flags().Bool("enable-llm", false, "enable LLM deep reasoning engine")
	scanCmd.Flags().String("skills-dir", "", "custom skills directory")
	scanCmd.Flags().Bool("notify-telegram", false, "send Telegram notification after scan (requires notify.telegram config)")
}

func runScan(cmd *cobra.Command, args []string) error {
	log := getLogger()
	conf := getConfig()
	scanStart := time.Now()

	// Determine target
	target := "."
	if len(args) > 0 {
		target = args[0]
	}

	// Validate target exists
	info, err := os.Stat(target)
	if err != nil {
		return fmt.Errorf("target not found: %s", target)
	}

	// Get flags (with Viper fallbacks via config)
	format, _ := cmd.Flags().GetString("format")
	severity, _ := cmd.Flags().GetString("severity")
	confidence, _ := cmd.Flags().GetString("confidence")
	concurrency, _ := cmd.Flags().GetInt("concurrency")
	enableLLM, _ := cmd.Flags().GetBool("enable-llm")
	skillsDir, _ := cmd.Flags().GetString("skills-dir")
	excludeDirs, _ := cmd.Flags().GetStringSlice("exclude-dir")
	excludeFiles, _ := cmd.Flags().GetStringSlice("exclude-file")
	languages, _ := cmd.Flags().GetStringSlice("language")

	// Use config values as fallback
	if !cmd.Flags().Changed("severity") {
		severity = conf.Scan.Severity
	}
	if !cmd.Flags().Changed("concurrency") {
		concurrency = conf.Scan.Concurrency
	}
	if !cmd.Flags().Changed("enable-llm") {
		enableLLM = conf.LLM.Enabled
	}
	if !cmd.Flags().Changed("skills-dir") {
		skillsDir = conf.Skills.Dir
	}
	if !cmd.Flags().Changed("format") {
		format = conf.Output.Format
	}

	// Merge exclude dirs and files from config
	if len(excludeDirs) == 0 {
		excludeDirs = conf.Exclude.Dirs
	}
	if len(excludeFiles) == 0 {
		excludeFiles = conf.Exclude.Files
	}

	// Override config with flags
	conf.Output.Format = format
	conf.Output.File, _ = cmd.Flags().GetString("output")
	conf.Scan.Severity = severity
	conf.Scan.Confidence = confidence
	conf.Scan.Concurrency = concurrency
	conf.Scan.Languages = languages
	conf.LLM.Enabled = enableLLM
	conf.Skills.Dir = skillsDir
	conf.Exclude.Dirs = excludeDirs
	conf.Exclude.Files = excludeFiles

	// Print scan banner
	printScanBanner(target, info.IsDir(), format, severity, concurrency, enableLLM)

	log.Infow("Starting scan",
		"target", target,
		"format", format,
		"severity", severity,
		"confidence", confidence,
		"concurrency", concurrency,
		"enableLLM", enableLLM,
		"skillsDir", skillsDir,
		"excludeDirs", excludeDirs,
		"excludeFiles", excludeFiles,
		"languages", languages,
	)

	// Initialize and run scan engine
	engine := scanner.NewEngine(conf, log)
	
	// Register parsers
	engine.RegisterParser(goparser.New())
	engine.RegisterParser(treesitter.New())
	engine.RegisterParser(textparser.New())

	// Register reporters
	engine.RegisterReporter(reporter.NewConsoleReporter(!conf.Output.Color))
	engine.RegisterReporter(reporter.NewJSONReporter())
	engine.RegisterReporter(reporter.NewSarifReporter())
	engine.RegisterReporter(reporter.NewGitLabReporter())
	engine.RegisterReporter(reporter.NewPDFReporter())

	// Initialize LLM provider if enabled
	if conf.LLM.Enabled {
		provider, err := providers.NewOpenAIProvider("OPENAI_API_KEY", "gpt-4o-mini")
		if err != nil {
			log.Warnw("LLM enabled but provider initialization failed", "error", err)
		} else {
			engine.SetLLMProvider(provider)
			log.Info("LLM reasoning engine initialized")
		}
	}

	if err := engine.Init(); err != nil {
		return fmt.Errorf("scan engine init failed: %w", err)
	}

	findings, err := engine.Run(cmd.Context(), target)
	if err != nil {
		log.Errorw("Scan failed with error", "error", err)
		return err
	}

	scanDuration := time.Since(scanStart)

	// Send Telegram notification if enabled
	notifyTelegram, _ := cmd.Flags().GetBool("notify-telegram")
	if conf.Notify.Telegram.Enabled || notifyTelegram {
		// Override config if flag was set
		telegramCfg := conf.Notify.Telegram
		if notifyTelegram && !telegramCfg.Enabled {
			telegramCfg.Enabled = true
		}
		if telegramCfg.BotToken == "" {
			if envToken := os.Getenv("ICE_TEA_TELEGRAM_BOT_TOKEN"); envToken != "" {
				telegramCfg.BotToken = envToken
			}
		}
		if telegramCfg.ChatID == "" {
			if envChatID := os.Getenv("ICE_TEA_TELEGRAM_CHAT_ID"); envChatID != "" {
				telegramCfg.ChatID = envChatID
			}
		}
		if telegramCfg.BotToken != "" && telegramCfg.ChatID != "" {
			n := telegram.New(telegramCfg)
			if err := n.Notify(cmd.Context(), target, finding.Summarize(findings), findings, scanDuration); err != nil {
				log.Warnw("Telegram notification failed", "error", err)
			} else {
				log.Info("Telegram notification sent")
			}
		} else {
			log.Warnw("Telegram notification enabled but bot_token or chat_id is missing")
		}
	}

	// Exit with code 1 if findings exist (typical CI behavior)
	if len(findings) > 0 {
		os.Exit(1)
	}

	return nil
}

func printScanBanner(target string, isDir bool, format, severity string, concurrency int, enableLLM bool) {
	cyan := color.New(color.FgCyan, color.Bold)
	white := color.New(color.FgWhite)
	green := color.New(color.FgGreen)

	cyan.Println("🍵 Ice Tea Security Scanner")
	fmt.Println(strings.Repeat("─", 40))

	targetType := "File"
	if isDir {
		targetType = "Directory"
	}

	white.Printf("  Target:      %s (%s)\n", target, targetType)
	white.Printf("  Format:      %s\n", format)
	white.Printf("  Severity:    %s+\n", severity)
	white.Printf("  Workers:     %d\n", concurrency)

	if enableLLM {
		green.Println("  LLM Engine:  ✓ Enabled")
	} else {
		white.Println("  LLM Engine:  ✗ Disabled")
	}

	fmt.Println(strings.Repeat("─", 40))
}
