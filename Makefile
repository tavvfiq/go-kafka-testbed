.PHONY: build clean

clean:
	@echo cleaning project
	@rm -rf build

build:
	@echo building binary
	@go build -o build/go-kafka-testbed ./main.go