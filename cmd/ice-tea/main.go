package main

import (
	"os"

	"github.com/zakirkun/ice-tea/internal/cli"
)

// Build info set via ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.SetBuildInfo(version, commit, date)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
