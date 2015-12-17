GITREV=$(shell git rev-parse --verify HEAD | cut -c 1-8)
GOPATH=$(PWD):$(PWD)/vendor

all: genbarcode

genbarcode: test
	@echo "Building..."
	@go build -ldflags "-X main.version=$(GITREV)" -o bin/genbarcode genbarcode

genbarcode-debug: test
	@echo "Building..."
	@go build -ldflags "-X main.version=$(GITREV)" -tags "debug" -o bin/genbarcode-debug genbarcode

test:
	@echo "Testing..."
	go test genbarcode

clean:
	@echo "Cleaning..."
	-@rm -rf pkg/* bin/* 2>/dev/null || true
