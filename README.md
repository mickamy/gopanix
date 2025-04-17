# gopanix

> 💥 Visualize your Go panics in the browser. Because stack traces deserve better.

`gopanix` is a CLI and library for Go that turns panic stack traces into readable HTML reports.  
Use it to debug Go crashes in a more comfortable, visual way — no more squinting at walls of text.

---

## ✨ Features

- 💥 Catch Go panics and recover automatically
- 🌐 Generate an HTML report with stack trace and timestamp
- 🧪 Run any Go program with `gopanix run`
- 📦 Use as a library with `defer gopanix.Handle()`
- 🚀 Opens your browser automatically
- 🧘 Works without modifying the target program (CLI mode)

---

## 🚀 Usage

You can use `gopanix` in two ways:

---

### 📦 As a library

Add `gopanix` to your Go app and use `defer gopanix.Handle()`:

```bash
go get github.com/mickamy/gopanix@latest
```

```go
package main

import "github.com/mickamy/gopanix"

func main() {
	defer gopanix.Handle()

	panic("something went wrong!")
}
```

When a panic occurs, `gopanix` recovers it and opens a detailed HTML report in your browser.

---

### 🛠 As a CLI

You can run any Go program — even if it doesn't import gopanix.

```bash
# Install gopanix into your project
go get -tool github.com/mickamy/gopanix@latest

# or install it globally
go install github.com/mickamy/gopanix@latest
```

#### `gopanix run`

```bash
gopanix run ./main.go
```

`gopanix run` will capture the panic output from go run, format it, and open an HTML report.
It wraps `go run`, captures any panic output, and opens a detailed HTML trace in your browser — no code changes needed.

#### `gopanix report`

Reads panic output from stdin and generates a readable HTML report — perfect for piping `go test` or `go run` output.

```bash
go test 2>&1 | gopanix report
```

If any `panic:` is found, `gopanix` will extract the stack trace and open it in your browser.



## 📄 License

[MIT](./LICENSE)
