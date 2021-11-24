SHELL=/usr/bin/env bash

APPLICATION_NAME=rosenchat

# Go tools.
GO_IMPORTS=$(shell which goimports || echo '')
GO_LINT=$(shell which golint || echo '')
STATIC_CHECK=$(shell which staticcheck || echo '')

# Installation URLs for Go tools.
GO_IMPORTS_INSTALL_URL=golang.org/x/tools/cmd/goimports
GO_LINT_INSTALL_URL=golang.org/x/lint/golint
STATIC_CHECK_INSTALL_URL=honnef.co/go/tools/cmd/staticcheck

# Listing all .go files.
GO_FILES=$(shell find . -type f -name '*.go' | grep -v /vendor/)

# For ProtoBuf code generation.
PROTO_GEN_PARENT_PATH=.
PROTO_DEF_PATH=src/protogen/*.proto

tools:
	@echo "+$@"
	$(if $(GO_IMPORTS), , go install $(GO_IMPORTS_INSTALL_URL))
	$(if $(GO_LINT), , go install $(GO_LINT_INSTALL_URL))
	$(if $(STATIC_CHECK), , go install $(STATIC_CHECK_INSTALL_URL))

check: tools
	@echo "+$@"
	@$(STATIC_CHECK) ./...

vet:
	@echo "+$@"
	@go vet ./...

imports: tools
	@echo "+$@"
	@$(GO_IMPORTS) -w $(GO_FILES)

fmt:
	@echo "+$@"
	@go fmt ./...

lint: tools
	@echo "+$@"
	@$(GO_LINT) ./...

build: check vet imports fmt lint
	@echo "+$@"
	@go build -o bin/main

test:
	@echo "+$@"
	@go test ./...

image: check vet imports fmt lint test
	@echo "+$@"
	@docker build -t $(APPLICATION_NAME):latest .

run: build
	@echo "+$@"
	@echo "#######################################################"
	@bin/main

container:
	@echo "+$@"
	@echo "############### Removing old container ################"
	@docker rm -f $(APPLICATION_NAME)
	@echo "################ Running new container ################"
	@docker run \
		--detach \
		--name $(APPLICATION_NAME) \
		--restart unless-stopped \
		--net host \
		--env-file prod.env \
		--volume $(HOME)/docker/volumes/$(APPLICATION_NAME)/app-logs:/var/log \
		$(APPLICATION_NAME):latest

proto:
	@echo "+$@"
	@protoc \
		--go_out=$(PROTO_GEN_PARENT_PATH) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_GEN_PARENT_PATH) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DEF_PATH)
