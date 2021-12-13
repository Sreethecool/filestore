IMAGE ?= filestore
BINARY ?= filestore
.PHONY: go-build
go-build:
	@echo "Build project binaries..."
	go build -o $(BINARY) -v main.go

.PHONY: run
run:
	@echo "Run..."
	go run main.go

.PHONY: build
build: go-build 
	@echo "Build docker image..."
	docker build --tag sreethecool2/$(IMAGE):latest $(BUILD_CONTEXT)

.PHONY: push
push:
	@echo "Push docker image..."
	docker push sreethecool2/$(IMAGE):latest

.PHONY: clean
clean:
	rm -rf $(BINARY)
