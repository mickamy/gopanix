package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/internal/cli/report"
	"github.com/mickamy/gopanix/internal/cli/run"
	"github.com/mickamy/gopanix/internal/cli/test"
	"github.com/mickamy/gopanix/internal/cli/version"
)

var cmd = &cobra.Command{
	Use:   "gopanix",
	Short: "Catch and visualize Go panics in your browser",
	Long: `gopanix is a CLI and library tool that catches panics in Go programs and opens a detailed stack trace report in your browser.

It helps you debug crashes more comfortably by saving panic information as HTML, with syntax highlighting and timestamps.`,
}

func init() {
	cmd.AddCommand(report.Cmd)
	cmd.AddCommand(run.Cmd)
	cmd.AddCommand(test.Cmd)
	cmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
