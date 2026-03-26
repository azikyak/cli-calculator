APP_NAME ?= calc
BUILD_DIR := bin
ARGS ?=

.PHONY: help build run install uninstall clean

ifeq ($(OS),Windows_NT)
EXE_EXT := .exe
PREFIX ?= $(USERPROFILE)/AppData/Local/Programs/cli-calculator/bin
BINARY := $(BUILD_DIR)/$(APP_NAME)$(EXE_EXT)
MKDIR_BUILD_CMD = powershell -NoProfile -ExecutionPolicy Bypass -Command "New-Item -ItemType Directory -Force '$(BUILD_DIR)' | Out-Null"
INSTALL_CMD = powershell -NoProfile -ExecutionPolicy Bypass -File .\install.ps1 -Prefix "$(PREFIX)" -AppName "$(APP_NAME)" -BuildDir "$(BUILD_DIR)"
UNINSTALL_CMD = powershell -NoProfile -ExecutionPolicy Bypass -Command "Remove-Item -Force '$(PREFIX)/$(APP_NAME)$(EXE_EXT)' -ErrorAction SilentlyContinue"
CLEAN_CMD = powershell -NoProfile -ExecutionPolicy Bypass -Command "Remove-Item -Recurse -Force '$(BUILD_DIR)' -ErrorAction SilentlyContinue"
else
EXE_EXT :=
PREFIX ?= /usr/local/bin
BINARY := $(BUILD_DIR)/$(APP_NAME)
MKDIR_BUILD_CMD = mkdir -p $(BUILD_DIR)
INSTALL_CMD = PREFIX=$(PREFIX) APP_NAME=$(APP_NAME) BUILD_DIR=$(BUILD_DIR) ./install.sh
UNINSTALL_CMD = rm -f $(PREFIX)/$(APP_NAME)
CLEAN_CMD = rm -rf $(BUILD_DIR)
endif

help:
	@printf '%s\n' \
		'Available targets:' \
		'  make build                    Build the calculator binary into bin/' \
		'  make run ARGS="\"2+3*4\""      Run the calculator module with a custom expression' \
		'  make install                  Install the binary into PREFIX (OS-aware)' \
		'  make uninstall                Remove the installed binary from PREFIX' \
		'  make clean                    Remove build artifacts'

build:
	@$(MKDIR_BUILD_CMD)
	go build -o $(BINARY) .

run:
	go run . $(ARGS)

install:
	$(INSTALL_CMD)

uninstall:
	$(UNINSTALL_CMD)

clean:
	$(CLEAN_CMD)
