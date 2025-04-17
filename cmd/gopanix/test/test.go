package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/gopanix"
	"github.com/mickamy/gopanix/internal/browser"
	"github.com/mickamy/gopanix/internal/html"
)

var Cmd = &cobra.Command{
	Use:   "test [packages...]",
	Short: "Run go test and visualize any panic as HTML",
	Long:  "Executes `go test -json` and captures panic output to display it in the browser via gopanix.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			args = []string{"./..."}
		}

		return Run(args)
	},
}

func Run(packages []string) error {
	allArgs := append([]string{"test", "-json"}, packages...)
	cmdExec := exec.Command("go", allArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmdExec.Stdout = &stdout
	cmdExec.Stderr = &stderr

	if err := cmdExec.Run(); err != nil {
		out := stderr.String()
		if strings.Contains(out, "panic:") {
			path, err := gopanix.Write([]byte(out), "panic from go test", "")
			if err != nil {
				fmt.Printf("⚠️ failed to write HTML report: %v\n", err)
			} else {
				fmt.Printf("📄 HTML report written to: file://%s\n", path)
				_ = browser.Open(path)
			}
		}
		_, _ = fmt.Fprintln(os.Stderr, out)
		return err
	}

	decoder := json.NewDecoder(&stdout)
	for decoder.More() {
		var event map[string]any
		_ = decoder.Decode(&event)
		if output, ok := event["Output"].(string); ok {
			fmt.Print(output)
		}
	}

	return nil
}
