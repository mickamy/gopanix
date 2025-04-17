APP_NAME = gopanix
VERSION ?= dev
BUILD_DIR = bin
GORELEASER ?= go tool goreleaser
VERSION_VARIABLE = github.com/mickamy/gopanix/internal/cli/version/version.version

.PHONY: all build install uninstall clean version test fmt

all: build

build:
	@echo "🔨 Building $(APP_NAME)..."
	go build -ldflags "-X $(VERSION_VARIABLE)=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/gopanix

install:
	@echo "📦 Installing $(APP_NAME)..."
	go install -ldflags "-X $(VERSION_VARIABLE)=$(VERSION)" ./cmd/gopanix

uninstall:
	@echo "🗑️  Uninstalling $(APP_NAME)..."
	@bin_dir=$$(go env GOBIN); \
	if [ -z "$$bin_dir" ]; then \
		bin_dir=$$(go env GOPATH)/bin; \
	fi; \
	echo "Removing $$bin_dir/$(APP_NAME)"; \
	rm -f $$bin_dir/$(APP_NAME)

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BUILD_DIR)

version:
	@echo "🔖 Version: $(VERSION)"

test:
	@echo "🧪 Running tests..."
	go test ./...

test-panic:
	@echo "🧪 Testing: expected panic (gopanix run)..."
	@$(APP_NAME) run ./testdata/panic.go --open || echo "💥 Panic detected and reported"

test-ok:
	@echo "🧪 Testing: no panic expected (gopanix run)..."
	@$(APP_NAME) run ./testdata/no_panic.go --open

test-lib:
	@echo "🧪 Testing: panic using gopanix.Handle() (embedded)..."
	@go run ./testdata/handle.go || echo "💥 Panic detected and reported"

fmt:
	@echo "📝 Formatting code..."
	gofmt -w -l .

release:
	@echo "🚀 Running release..."
	$(GORELEASER) release --clean

snapshot:
	@echo "🔍 Running snapshot release (dry run)..."
	$(GORELEASER) release --snapshot --clean
