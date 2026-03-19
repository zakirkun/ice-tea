// Package telegram provides Telegram Bot API notification support for Ice Tea scan results.
package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zakirkun/ice-tea/internal/config"
	"github.com/zakirkun/ice-tea/internal/finding"
)

const (
	telegramAPIBase = "https://api.telegram.org/bot%s/sendMessage"
	maxTopFindings  = 5
)

// Notifier sends scan results to a Telegram channel or chat
type Notifier struct {
	cfg config.TelegramConfig
}

// New creates a new Telegram notifier from config
func New(cfg config.TelegramConfig) *Notifier {
	return &Notifier{cfg: cfg}
}

// severityEmoji returns an emoji for each severity level
func severityEmoji(severity string) string {
	switch strings.ToLower(severity) {
	case "critical":
		return "🔴"
	case "high":
		return "🟠"
	case "medium":
		return "🟡"
	case "low":
		return "🟢"
	default:
		return "⚪"
	}
}

// severityLevel maps severity strings to numeric levels for comparison
func severityLevel(s string) int {
	switch strings.ToLower(s) {
	case "critical":
		return 5
	case "high":
		return 4
	case "medium":
		return 3
	case "low":
		return 2
	case "info":
		return 1
	default:
		return 0
	}
}

// Notify sends scan results to the configured Telegram chat.
// It only sends if findings meet the minimum severity threshold.
func (n *Notifier) Notify(
	ctx context.Context,
	target string,
	summary finding.Summary,
	allFindings []*finding.Finding,
	duration time.Duration,
) error {
	findings := allFindings
	// Check minimum severity threshold
	minLevel := severityLevel(n.cfg.MinSeverity)
	hasQualifyingFindings := false

	if summary.Critical > 0 && severityLevel("critical") >= minLevel {
		hasQualifyingFindings = true
	}
	if summary.High > 0 && severityLevel("high") >= minLevel {
		hasQualifyingFindings = true
	}
	if summary.Medium > 0 && severityLevel("medium") >= minLevel {
		hasQualifyingFindings = true
	}
	if summary.Low > 0 && severityLevel("low") >= minLevel {
		hasQualifyingFindings = true
	}

	// Skip if only_on_findings is enabled and no qualifying findings
	if n.cfg.OnlyOnFindings && !hasQualifyingFindings {
		return nil
	}

	// Build message
	msg := buildMessage(target, summary, findings, duration, n.cfg.MinSeverity)

	return n.sendMessage(ctx, msg)
}

// buildMessage constructs the Telegram MarkdownV2 message
func buildMessage(
	target string,
	summary finding.Summary,
	findings []*finding.Finding,
	duration time.Duration,
	minSeverity string,
) string {
	var sb strings.Builder

	sb.WriteString("🍵 *Ice Tea Scan Complete*\n\n")
	sb.WriteString(fmt.Sprintf("📁 Target: `%s`\n", escapeMarkdown(target)))
	sb.WriteString(fmt.Sprintf("⏱ Duration: `%s`\n", duration.Round(time.Millisecond).String()))
	sb.WriteString(fmt.Sprintf("📅 %s UTC\n\n", time.Now().UTC().Format("2006-01-02 15:04:05")))

	sb.WriteString("📊 *Summary*\n")

	if summary.Critical > 0 {
		sb.WriteString(fmt.Sprintf("🔴 Critical: *%d*\n", summary.Critical))
	}
	if summary.High > 0 {
		sb.WriteString(fmt.Sprintf("🟠 High: *%d*\n", summary.High))
	}
	if summary.Medium > 0 {
		sb.WriteString(fmt.Sprintf("🟡 Medium: *%d*\n", summary.Medium))
	}
	if summary.Low > 0 {
		sb.WriteString(fmt.Sprintf("🟢 Low: *%d*\n", summary.Low))
	}
	if summary.Info > 0 {
		sb.WriteString(fmt.Sprintf("⚪ Info: *%d*\n", summary.Info))
	}

	sb.WriteString(fmt.Sprintf("─────────────────\n"))
	sb.WriteString(fmt.Sprintf("Total: *%d findings*\n", summary.Total))

	// Top critical/high findings
	topFindings := getTopFindings(findings, minSeverity, maxTopFindings)
	if len(topFindings) > 0 {
		sb.WriteString(fmt.Sprintf("\n🚨 *Top %s Findings*\n", strings.Title(minSeverity)+"+"))
		for i, f := range topFindings {
			emoji := severityEmoji(f.Severity)
			location := ""
			if f.File != "" {
				parts := strings.Split(f.File, "/")
				shortFile := f.File
				if len(parts) > 3 {
					shortFile = strings.Join(parts[len(parts)-3:], "/")
				}
				location = fmt.Sprintf(" — `%s:%d`", shortFile, f.StartLine)
			}
			sb.WriteString(fmt.Sprintf("%d\\. %s `%s`%s\n", i+1, emoji, escapeMarkdown(f.Type), escapeMarkdown(location)))
		}
	}

	sb.WriteString("\n_Run `ice-tea scan --format pdf --output report.pdf` for full report_")

	return sb.String()
}

// getTopFindings returns the most severe findings up to maxCount
func getTopFindings(findings []*finding.Finding, minSeverity string, maxCount int) []*finding.Finding {
	minLevel := severityLevel(minSeverity)
	var top []*finding.Finding

	for _, f := range findings {
		if severityLevel(f.Severity) >= minLevel {
			top = append(top, f)
		}
	}

	if len(top) > maxCount {
		top = top[:maxCount]
	}
	return top
}

// escapeMarkdown escapes MarkdownV2 special characters
func escapeMarkdown(s string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(s)
}

// telegramRequest is the API payload for sendMessage
type telegramRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// telegramResponse is the API response
type telegramResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
}

// sendMessage sends a message via Telegram Bot API
func (n *Notifier) sendMessage(ctx context.Context, text string) error {
	token := n.cfg.BotToken
	if token == "" {
		return fmt.Errorf("telegram bot_token is empty")
	}
	if n.cfg.ChatID == "" {
		return fmt.Errorf("telegram chat_id is empty")
	}

	payload := telegramRequest{
		ChatID:    n.cfg.ChatID,
		Text:      text,
		ParseMode: "MarkdownV2",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal telegram request: %w", err)
	}

	url := fmt.Sprintf(telegramAPIBase, token)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create telegram request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send telegram message: %w", err)
	}
	defer resp.Body.Close()

	var apiResp telegramResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("decode telegram response: %w", err)
	}

	if !apiResp.OK {
		return fmt.Errorf("telegram API error: %s", apiResp.Description)
	}

	return nil
}
