IMAGE ?= filestore
BINARY ?= filestore
.PHONY: go-build
go-build:
	@echo "Build project binaries..."
	GO111MODULE=on go build -o $(BINARY) -v main.go

.PHONY: run
run:
	@echo "Run..."
	GO111MODULE=on go run main.go

.PHONY: build
build: go-build 
	@echo "Build docker image..."
	docker build --tag sreethecool2/$(IMAGE):latest ./

.PHONY: push
push:
	@echo "Push docker image..."
	docker push sreethecool2/$(IMAGE):latest

.PHONY: clean
clean:
	rm -rf $(BINARY)
