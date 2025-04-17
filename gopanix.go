package gopanix

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/mickamy/gopanix/internal/browser"
)

// Handle is intended to be used as a top-level defer.
// It recovers from panic and generates an HTML report.
// If `open` is true, it opens the report in the browser.
func Handle(open bool) {
	if r := recover(); r != nil {
		filename, err := Report(r)
		if err != nil {
			fmt.Printf("⚠️ Failed to generate report: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("📄 panic file written to: %s\n", filename)
		if open {
			fmt.Println("🌐 Opening in browser...")
			_ = browser.Open(filename)
		}
		os.Exit(1)
	}
}

// Report creates an HTML file with panic details.
func Report(r any) (string, error) {
	stack := debug.Stack()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	filename, err := Write(stack, fmt.Sprint(r), timestamp)
	if err != nil {
		return "", err
	}

	return filename, nil
}
