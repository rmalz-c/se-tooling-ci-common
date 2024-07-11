VERSION = $(shell git describe --tags 2> /dev/null || echo "unknown")

.PHONY: all
all: lint build test

.PHONY: build
build:
	go build -v -o my_app -ldflags="-X main.commit=$(VERSION)" cmd/my_app/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	@fmt_result=$$(gofmt -d $$(find . -name '*.go' -print)); \
		if [ -n "$${fmt_result}" ]; then \
				echo "gofmt checking failed"; \
				echo; \
				echo "$${fmt_result}"; \
				false; \
		fi
