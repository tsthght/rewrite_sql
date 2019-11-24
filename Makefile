### Makefile for rewrite_sql

GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)
GOTEST   := CGO_ENABLED=1 $(GO) test -p 3

ARCH  := "`uname -s`"
LINUX := "Linux"
MAC   := "Darwin"

PROJ := rewrite_sql

build:
	$(GOBUILD) -o bin/$(PROJ) ./cmd/main.go



