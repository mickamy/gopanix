package run

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/internal/browser"
	"github.com/mickamy/gopanix/internal/html"
)

var Cmd = &cobra.Command{
	Use:   "run [file.go | package]",
	Short: "Run a Go program and capture any panic",
	Long:  "Run a Go program using `go run` and capture panic output to visualize in HTML.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run(args)
	},
}

func Run(args []string) error {
	cmd := exec.Command("go", append([]string{"run"}, args...)...)

	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

	fmt.Println("🚀 Running:", strings.Join(cmd.Args, " "))
	err := cmd.Run()

	if err != nil {
		scanner := bufio.NewScanner(&stderr)
		var panicDetected bool
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
			if strings.Contains(line, "panic:") {
				panicDetected = true
			}
		}

		if panicDetected {
			content := strings.Join(lines, "\n")
			path, htmlErr := html.Write([]byte(content), "panic from subprocess", "")
			if htmlErr != nil {
				return fmt.Errorf("failed to write HTML report: %w", htmlErr)
			}
			fmt.Printf("\033[36m📄 Panic detected. Report written to:\033[0m file://%s\n", path)
			_ = browser.Open(path)
			return nil
		}
		return err
	}

	fmt.Println("✅ No panic detected. Program exited normally.")
	return nil
}
