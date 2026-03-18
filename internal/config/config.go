package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Scan    ScanConfig    `mapstructure:"scan"`
	Exclude ExcludeConfig `mapstructure:"exclude"`
	Skills  SkillsConfig  `mapstructure:"skills"`
	LLM     LLMConfig     `mapstructure:"llm"`
	Output  OutputConfig  `mapstructure:"output"`
}

// ScanConfig holds scan-related settings
type ScanConfig struct {
	Target      string        `mapstructure:"target"`
	Severity    string        `mapstructure:"severity"`
	Confidence  string        `mapstructure:"confidence"`
	Concurrency int           `mapstructure:"concurrency"`
	Timeout     time.Duration `mapstructure:"timeout"`
	Languages   []string      `mapstructure:"languages"`
}

// ExcludeConfig holds exclusion patterns
type ExcludeConfig struct {
	Dirs       []string `mapstructure:"dirs"`
	Files      []string `mapstructure:"files"`
	Extensions []string `mapstructure:"extensions"`
}

// SkillsConfig holds SKILL-related settings
type SkillsConfig struct {
	Dir           string   `mapstructure:"dir"`
	ExternalDirs  []string `mapstructure:"external_dirs"`
	TrustExternal bool     `mapstructure:"trust_external"`
}

// LLMConfig holds LLM engine settings
type LLMConfig struct {
	Enabled   bool          `mapstructure:"enabled"`
	Provider  string        `mapstructure:"provider"`
	Model     string        `mapstructure:"model"`
	APIKeyEnv string        `mapstructure:"api_key_env"`
	MaxTokens int           `mapstructure:"max_tokens"`
	RateLimit int           `mapstructure:"rate_limit"`
	Timeout   time.Duration `mapstructure:"timeout"`
	Cache     bool          `mapstructure:"cache"`
}

// OutputConfig holds output settings
type OutputConfig struct {
	Format  string `mapstructure:"format"`
	File    string `mapstructure:"file"`
	Verbose bool   `mapstructure:"verbose"`
	Color   bool   `mapstructure:"color"`
}

// setDefaults configures all default values
func setDefaults() {
	// Scan defaults
	viper.SetDefault("scan.target", ".")
	viper.SetDefault("scan.severity", "medium")
	viper.SetDefault("scan.confidence", "medium")
	viper.SetDefault("scan.concurrency", 4)
	viper.SetDefault("scan.timeout", "5m")

	// Exclude defaults
	viper.SetDefault("exclude.dirs", []string{
		"vendor", "node_modules", ".git", "build", "dist",
		"testdata", "__pycache__", ".venv", "venv",
	})
	viper.SetDefault("exclude.files", []string{
		"*.min.js", "*.min.css", "*.generated.go", "*.pb.go",
	})
	viper.SetDefault("exclude.extensions", []string{
		".md", ".txt", ".png", ".jpg", ".jpeg", ".gif", ".svg",
		".ico", ".woff", ".woff2", ".ttf", ".eot", ".pdf",
		".zip", ".tar", ".gz", ".exe", ".dll", ".so", ".dylib",
	})

	// Skills defaults
	viper.SetDefault("skills.dir", "./skills")
	viper.SetDefault("skills.trust_external", false)

	// LLM defaults
	viper.SetDefault("llm.enabled", false)
	viper.SetDefault("llm.provider", "openai")
	viper.SetDefault("llm.model", "gpt-4o-mini")
	viper.SetDefault("llm.api_key_env", "ICE_TEA_LLM_API_KEY")
	viper.SetDefault("llm.max_tokens", 4096)
	viper.SetDefault("llm.rate_limit", 10)
	viper.SetDefault("llm.timeout", "30s")
	viper.SetDefault("llm.cache", true)

	// Output defaults
	viper.SetDefault("output.format", "console")
	viper.SetDefault("output.verbose", false)
	viper.SetDefault("output.color", true)
}

// Load reads configuration from file, env, and flags
func Load(cfgFile string) (*Config, error) {
	setDefaults()

	// Environment variable support
	viper.SetEnvPrefix("ICE_TEA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		// Use specific config file
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in working directory
		viper.AddConfigPath(".")
		viper.SetConfigName(".ice-tea")
		viper.SetConfigType("yaml")
	}

	// Read config file (ok if not found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}

// validate checks configuration values
func validate(cfg *Config) error {
	// Validate severity
	validSeverities := map[string]bool{
		"critical": true, "high": true, "medium": true, "low": true, "info": true,
	}
	if !validSeverities[strings.ToLower(cfg.Scan.Severity)] {
		return fmt.Errorf("invalid severity: %q (must be critical, high, medium, low, or info)", cfg.Scan.Severity)
	}

	// Validate confidence
	validConfidences := map[string]bool{
		"high": true, "medium": true, "low": true,
	}
	if !validConfidences[strings.ToLower(cfg.Scan.Confidence)] {
		return fmt.Errorf("invalid confidence: %q (must be high, medium, or low)", cfg.Scan.Confidence)
	}

	// Validate output format
	validFormats := map[string]bool{
		"console": true, "sarif": true, "gitlab": true, "json": true,
	}
	if !validFormats[strings.ToLower(cfg.Output.Format)] {
		return fmt.Errorf("invalid format: %q (must be console, sarif, gitlab, or json)", cfg.Output.Format)
	}

	// Validate concurrency
	if cfg.Scan.Concurrency < 1 {
		cfg.Scan.Concurrency = 1
	}
	if cfg.Scan.Concurrency > 32 {
		cfg.Scan.Concurrency = 32
	}

	// Validate LLM provider
	if cfg.LLM.Enabled {
		validProviders := map[string]bool{
			"openai": true, "anthropic": true, "ollama": true,
		}
		if !validProviders[strings.ToLower(cfg.LLM.Provider)] {
			return fmt.Errorf("invalid LLM provider: %q (must be openai, anthropic, or ollama)", cfg.LLM.Provider)
		}
	}

	return nil
}
