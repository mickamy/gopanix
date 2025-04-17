package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/cmd/gopanix/run"
	"github.com/mickamy/gopanix/cmd/gopanix/test"
	"github.com/mickamy/gopanix/cmd/gopanix/version"
)

var cmd = &cobra.Command{
	Use:   "gopanix",
	Short: "Catch and visualize Go panics in your browser",
	Long: `gopanix is a CLI and library tool that catches panics in Go programs and opens a detailed stack trace report in your browser.

It helps you debug crashes more comfortably by saving panic information as HTML, with syntax highlighting and timestamps.`,
}

func init() {
	cmd.AddCommand(run.Cmd)
	cmd.AddCommand(test.Cmd)
	cmd.AddCommand(version.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
