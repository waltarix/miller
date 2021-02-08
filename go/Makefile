GOOS_GOARCH := $(word 4, $(shell go version))
GOOS        ?= $(word 1, $(subst /, , $(GOOS_GOARCH)))
GOARCH      ?= $(word 2, $(subst /, , $(GOOS_GOARCH)))

GO_MILLER_VERSION ?= v6.0.0-custom
BUILD_FLAGS       := -ldflags "-s -w -X miller/version.STRING=$(GO_MILLER_VERSION)"

gmlr:
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -o $@
