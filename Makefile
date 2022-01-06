build:
	goreleaser --rm-dist --snapshot --skip-publish -p1

test:
	CGO_ENABLED=0 go test -v github.com/johnkerl/miller/internal/pkg/...

clean:
	$(RM) -r dist

.PHONY: build test clean
