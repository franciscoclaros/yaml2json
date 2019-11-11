ifeq ($(shell uname -s), Darwin)
    shasum=shasum -a256
else
    shasum=sha256sum
endif

repo=github.com/ssuareza/yaml2json
version=$(shell git describe --all --dirty --long | awk -F"-|/" '/^heads/ {print $$2 "-" substr($$4, 2) "-" $$5}; /^tags/ { print $$2 }')
build_args=-ldflags "-X main.versionString=$(version)" ./cmd/yaml2json
files=$(shell find cmd -type f)

.PHONY: test

all: test build checksums

build: build-linux build-darwin

build-linux: build/yaml2json-$(version)-linux-amd64

build/yaml2json-$(version)-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/yaml2json-$(version)-darwin-amd64
build/yaml2json-$(version)-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)

checksums: build
	cd build/ && ${shasum} * > $(version)-SHA256SUMS