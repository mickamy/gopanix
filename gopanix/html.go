package gopanix

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type panicPageData struct {
	PanicMsg  string
	Timestamp string
	Stack     string
}

const pageTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Panic Report - {{ .Timestamp }}</title>
	<style>
		body { font-family: monospace; background: #111; color: #eee; padding: 2rem; }
		h1 { color: #f55; }
		pre { background: #222; padding: 1rem; border-radius: 8px; overflow-x: auto; }
	</style>
</head>
<body>
	<h1>💥 Panic!</h1>
	<p><strong>Time:</strong> {{ .Timestamp }}</p>
	<p><strong>Message:</strong> {{ .PanicMsg }}</p>
	<h2>Stack Trace:</h2>
	<pre>{{ .Stack }}</pre>
</body>
</html>
`

// Write generates an HTML file from the panic info and returns its path
func Write(stack []byte, panicMsg, timestamp string) (string, error) {
	data := panicPageData{
		PanicMsg:  panicMsg,
		Timestamp: timestamp,
		Stack:     string(stack),
	}

	tmpl, err := template.New("panic").Parse(pageTemplate)
	if err != nil {
		return "", fmt.Errorf("template parse error: %w", err)
	}

	filename := fmt.Sprintf("panic_%s.html", time.Now().Format("20060102_150405"))
	path := filepath.Join(os.TempDir(), filename)

	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("file create error: %w", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err := tmpl.Execute(f, data); err != nil {
		return "", fmt.Errorf("template exec error: %w", err)
	}

	return path, nil
}
