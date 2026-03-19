package main

import (
	"os"

	"github.com/zakirkun/ice-tea/internal/cli"
)

// Build info set via ldflags
var (
	version = "1.0.0"
	commit  = "14cff13"
	date    = "2025-03-19"
)

func main() {
	cli.SetBuildInfo(version, commit, date)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
