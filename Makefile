build:
	goreleaser --clean --snapshot --skip=publish -p1

test:
	go test -v github.com/johnkerl/miller/pkg/...

fmt:
	go fmt ./cmd/...
	go fmt ./pkg/...
	go fmt ./regression_test.go

clean:
	$(RM) -r dist

.PHONY: build test clean
