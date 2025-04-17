package report

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gopanix/gopanix"
	"github.com/mickamy/gopanix/internal/browser"
	"github.com/mickamy/gopanix/internal/panics"
)

var Cmd = &cobra.Command{
	Use:   "report",
	Short: "Read panic output from stdin and generate HTML report",
	Long: `gopanix report reads stderr/stdout output from a Go program (like go test)
and generates a formatted HTML report for panic stack traces.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read stdin: %w", err)
		}

		text := string(input)
		return Run(text)
	},
}

func Run(input string) error {
	if len(input) == 0 {
		fmt.Println("⚠️ No input received. Did you forget to pipe from go test?")
		return nil
	}

	if !panics.Contains(input) {
		fmt.Println("✅ No panic detected in input.")
		return nil
	}

	stacks := panics.Extract(input)
	if len(stacks) == 0 {
		fmt.Println("⚠️ Panic detected but failed to extract stack trace.")
		// fallback to full text
		stacks = [][]string{strings.Split(input, "\n")}
	}

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

	fmt.Println("🌐 Opening in browser...")
	_ = browser.Open(paths[0])

	return nil
}
