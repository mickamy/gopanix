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
func Handle() {
	if r := recover(); r != nil {
		Report(r)
		os.Exit(1)
	}
}

// Report creates an HTML file with panic details and opens it in the browser.
func Report(r any) {
	stack := debug.Stack()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	filename, err := Write(stack, fmt.Sprint(r), timestamp)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write panic HTML: %v\n", err)
		return
	}

	fmt.Printf("📄 panic written to: %s\n", filename)
	fmt.Println("🌐 Opening in browser...")
	_ = browser.Open(filename)
}
