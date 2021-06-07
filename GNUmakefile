NAME=web-init
BIN=packer-plugin-${NAME}
SOURCES := $(shell find . -name '*.go')

COUNT?=1
TEST?=$(shell go list ./...)

.PHONY: dev

build:
	go build -o ${BIN}

install:
	mkdir -p ~/.packer.d/plugins/
	mv ${BIN} ~/.packer.d/plugins/${BIN}

dev: create-hcl build install clean

build-example: dev
	@packer build ./example/build.pkr.hcl

test:
	@go test -count $(COUNT) $(TEST) -timeout=3m

testacc: dev
	@PACKER_ACC=1 go test -count $(COUNT) -v $(TEST) -timeout=120m

create-hcl: 
	GOBIN=$(shell pwd) go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@latest
	PATH=$(shell pwd):${PATH} go generate builder/web/config.go

clean:
	rm -f packer-sdc