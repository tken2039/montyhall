BIN_DIR:=bin

PKG_ROOT:=github.com/tken2039/montyhall
PKG_CMD:=$(shell go list ./cmd/...)

BINS:=$(PKG_CMD:$(PKG_ROOT)/cmd/%=$(BIN_DIR)/%)
GO_FILES:=$(shell find . -type f -name '*.go' -print)

.PHONY: build
build: $(BINS)

$(BINS): $(GO_FILES)
	@go build -o $@ $(@:$(BIN_DIR)/%=$(PKG_ROOT)/cmd/%)