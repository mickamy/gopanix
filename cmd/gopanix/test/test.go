package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/gopanix"
	"github.com/mickamy/gopanix/internal/browser"
	"github.com/mickamy/gopanix/internal/panics"
)

var (
	errGoTestFailed = errors.New("go test failed")
)

var Cmd = &cobra.Command{
	Use:   "test [packages...]",
	Short: "Run go test and visualize any panic as HTML",
	Long:  "Executes `go test -json` and captures panic output to display it in the browser via gopanix.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			args = []string{"./..."}
		}

		if err := Run(args); err != nil {
			if errors.Is(err, errGoTestFailed) {
				fmt.Println("⚠️ Go test failed.")
				os.Exit(1)
			}
			return err
		}

		return nil
	},
}

func Run(packages []string) error {
	allArgs := append([]string{"test", "-json"}, packages...)
	cmd := exec.Command("go", allArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		out := stderr.String()
		if panics.Contains(out) {
			stacks := panics.Extract(out)
			if len(stacks) == 0 {
				fmt.Println("⚠️ Panic detected but failed to extract stack trace.")
				// fallback to full text
				stacks = [][]string{strings.Split(out, "\n")}
			}

			// Write each stack trace to a separate HTML file
			paths := make([]string, len(stacks))
			for i, stack := range stacks {
				title := fmt.Sprintf("panic #%d", i+1)
				path, err := gopanix.Write([]byte(strings.Join(stack, "\n")), title, "")
				if err != nil {
					return fmt.Errorf("failed to write report: %w", err)
				}

				fmt.Printf("📄 HTML report #%d written to: file://%s\n", i+1, path)
				paths[i] = path
			}

			// Open the first panic report in the browser
			fmt.Println("🌐 Opening in browser...")
			_ = browser.Open(paths[0])

			return nil
		}

		decodeJSONAndPrint(&stdout)
		return errGoTestFailed
	}

	decodeJSONAndPrint(&stdout)
	return nil
}

func decodeJSONAndPrint(buf *bytes.Buffer) {
	decoder := json.NewDecoder(buf)
	for decoder.More() {
		var event map[string]any
		_ = decoder.Decode(&event)
		if output, ok := event["Output"].(string); ok {
			fmt.Print(output)
		}
	}
}
