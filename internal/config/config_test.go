package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadDefaults(t *testing.T) {
	cfg, err := Load("")
	require.NoError(t, err)
	require.NotNil(t, cfg)

	// Verify defaults
	assert.Equal(t, ".", cfg.Scan.Target)
	assert.Equal(t, "medium", cfg.Scan.Severity)
	assert.Equal(t, "medium", cfg.Scan.Confidence)
	assert.Equal(t, 4, cfg.Scan.Concurrency)
	assert.Equal(t, "console", cfg.Output.Format)
	assert.False(t, cfg.LLM.Enabled)
	assert.Equal(t, "openai", cfg.LLM.Provider)
	assert.Equal(t, "./skills", cfg.Skills.Dir)
	assert.False(t, cfg.Skills.TrustExternal)

	// Verify default exclusions
	assert.Contains(t, cfg.Exclude.Dirs, "vendor")
	assert.Contains(t, cfg.Exclude.Dirs, "node_modules")
	assert.Contains(t, cfg.Exclude.Dirs, ".git")
	assert.Contains(t, cfg.Exclude.Files, "*.min.js")
	assert.Contains(t, cfg.Exclude.Extensions, ".md")
}

func TestValidation(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*Config)
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid defaults",
			modify:  func(c *Config) {},
			wantErr: false,
		},
		{
			name:    "invalid severity",
			modify:  func(c *Config) { c.Scan.Severity = "super-high" },
			wantErr: true,
			errMsg:  "invalid severity",
		},
		{
			name:    "invalid confidence",
			modify:  func(c *Config) { c.Scan.Confidence = "ultra" },
			wantErr: true,
			errMsg:  "invalid confidence",
		},
		{
			name:    "invalid format",
			modify:  func(c *Config) { c.Output.Format = "xml" },
			wantErr: true,
			errMsg:  "invalid format",
		},
		{
			name: "invalid LLM provider",
			modify: func(c *Config) {
				c.LLM.Enabled = true
				c.LLM.Provider = "gpt-free"
			},
			wantErr: true,
			errMsg:  "invalid LLM provider",
		},
		{
			name:    "concurrency clamped to 1",
			modify:  func(c *Config) { c.Scan.Concurrency = 0 },
			wantErr: false,
		},
		{
			name:    "concurrency clamped to 32",
			modify:  func(c *Config) { c.Scan.Concurrency = 100 },
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Scan:   ScanConfig{Severity: "medium", Confidence: "medium", Concurrency: 4},
				Output: OutputConfig{Format: "console"},
				LLM:    LLMConfig{Provider: "openai"},
			}
			tt.modify(cfg)

			err := validate(cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestConcurrencyClamping(t *testing.T) {
	cfg := &Config{
		Scan:   ScanConfig{Severity: "medium", Confidence: "medium", Concurrency: 0},
		Output: OutputConfig{Format: "console"},
		LLM:    LLMConfig{Provider: "openai"},
	}

	err := validate(cfg)
	assert.NoError(t, err)
	assert.Equal(t, 1, cfg.Scan.Concurrency)

	cfg.Scan.Concurrency = 100
	err = validate(cfg)
	assert.NoError(t, err)
	assert.Equal(t, 32, cfg.Scan.Concurrency)
}
