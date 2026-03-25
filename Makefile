APP_NAME ?= calc
MAIN_FILE := main.go
BUILD_DIR := bin
BINARY := $(BUILD_DIR)/$(APP_NAME)
PREFIX ?= /usr/local/bin
ARGS ?=

.PHONY: help build run install uninstall clean

help:
	@printf '%s\n' \
		'Available targets:' \
		'  make build                    Build the calculator binary into bin/' \
		'  make run ARGS="\"2+3*4\""      Run the calculator with a custom expression' \
		'  make install                  Install the binary into PREFIX' \
		'  make uninstall                Remove the installed binary from PREFIX' \
		'  make clean                    Remove build artifacts'

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY) $(MAIN_FILE)

run:
	go run $(MAIN_FILE) $(ARGS)

install:
	PREFIX=$(PREFIX) APP_NAME=$(APP_NAME) ./install.sh

uninstall:
	rm -f $(PREFIX)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
