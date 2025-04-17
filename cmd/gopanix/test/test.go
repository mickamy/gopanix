package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/cmd/gopanix/report"
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
			return report.Run(out)
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
