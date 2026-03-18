package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Build info (injected via ldflags)
var (
	buildVersion = "dev"
	buildCommit  = "none"
	buildDate    = "unknown"
)

// SetBuildInfo sets the build metadata from ldflags
func SetBuildInfo(version, commit, date string) {
	buildVersion = version
	buildCommit = commit
	buildDate = date
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Ice Tea version information",
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold)
		white := color.New(color.FgWhite)

		cyan.Println("🍵 Ice Tea - AI DevOps Security Scanner")
		fmt.Println()
		white.Printf("  Version:    %s\n", buildVersion)
		white.Printf("  Commit:     %s\n", buildCommit)
		white.Printf("  Built:      %s\n", buildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
