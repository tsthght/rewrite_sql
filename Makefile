### Makefile for rewrite_sql

BUILD_FLAG := -buildmode=c-shared

GO       := GO111MODULE=on go
GOBUILD_SO  := CGO_ENABLED=1 $(GO) build $(BUILD_FLAG)
GOBUILD  := CGO_ENABLED=1 $(GO) build

ARCH  := "`uname -s`"
LINUX := "Linux"
MAC   := "Darwin"

PROJ := rewrite
PROJ_SO := rewrite.so

build:
	$(GOBUILD) -o bin/$(PROJ) ./cmd/main.go
so:
	$(GOBUILD_SO) -o bin/$(PROJ_SO) ./cmd/main.go
clean:
	rm ./bin/*


