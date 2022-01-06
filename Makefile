build:
	goreleaser --rm-dist --snapshot --skip-publish -p1

test:
	go test -v github.com/johnkerl/miller/internal/pkg/...

fmt:
	go fmt ./cmd/...
	go fmt ./internal/pkg/...
	go fmt ./regression_test.go

clean:
	$(RM) -r dist

.PHONY: build test clean
