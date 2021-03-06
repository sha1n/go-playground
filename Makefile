GOFILES_NOVENDOR=.

ifdef IS_BUILD_AGENT
MAYBE_VERBOSE = -v
endif

default:
	make prepare
	make format
	make test

prepare:
	cd ..
	go mod tidy

format:
	gofmt -s -w $(GOFILES_NOVENDOR)

lint:
	gofmt -d $(GOFILES_NOVENDOR)

test:
	go test $(MAYBE_VERBOSE) `go list ./...`

demo:
	go run main.go