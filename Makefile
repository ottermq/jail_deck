BINARY_NAME=jaildeck
MAIN_PATH=cmd/jaildeck
BUILD_DIR=bin

run: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

build:
	@mkdir -p $(BUILD_DIR)
	@VERSION=$$(git describe --tags --always 2>/dev/null || echo dev);\
		go build -ldflags "-X main.version=$$VERSION" -o ./$(BUILD_DIR)/$(BINARY_NAME) ./$(MAIN_PATH)
