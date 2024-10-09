# Define the application name
APP_NAME := auto-clicker

# Define the output directory
OUTPUT_DIR := build

# Define the Go build command
GO_BUILD := go build

# Define the build flags
BUILD_FLAGS := -ldflags "-s -w"

# Define the default OS and architecture combinations
OS_ARCH := windows/amd64 linux/amd64 darwin/amd64

# Default target
all: build

# Build target
build:
	@$(MAKE) $(OS_ARCH)

# Build for specific OS/ARCH
$(OS_ARCH):
ifeq ($(OS),Windows_NT)
	if not exist "$(OUTPUT_DIR)\$(word 1, $(subst /, ,$@))\$(word 2, $(subst /, ,$@))" mkdir "$(OUTPUT_DIR)\$(word 1, $(subst /, ,$@))\$(word 2, $(subst /, ,$@))"
	set "GOOS=$(word 1, $(subst /, ,$@))" && set "GOARCH=$(word 2, $(subst /, ,$@))" && $(GO_BUILD) -ldflags "-H=windowsgui -s -w" -o "$(OUTPUT_DIR)\$(word 1, $(subst /, ,$@))\$(word 2, $(subst /, ,$@))\$(APP_NAME).exe"
else
	@if [ ! -d "$(OUTPUT_DIR)/$(word 1, $(subst /, ,$@))/$(word 2, $(subst /, ,$@))" ]; then mkdir -p "$(OUTPUT_DIR)/$(word 1, $(subst /, ,$@))/$(word 2, $(subst /, ,$@))"; fi
	GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) $(GO_BUILD) $(BUILD_FLAGS) -o "$(OUTPUT_DIR)/$(word 1, $(subst /, ,$@))/$(word 2, $(subst /, ,$@))/$(APP_NAME)"
endif

# Clean target
clean:
ifeq ($(OS),Windows_NT)
	if exist "$(OUTPUT_DIR)" rmdir /s /q "$(OUTPUT_DIR)"
else
	@if [ -d "$(OUTPUT_DIR)" ]; then rm -rf "$(OUTPUT_DIR)"; fi
endif
	@echo "Clean completed."

# PHONY targets
.PHONY: all build clean $(OS_ARCH)