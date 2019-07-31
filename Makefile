GOLANGCI_LINT_VERSION := v1.16.0
TMP_PROJECT_NAME=my-go-project
TMP_DIR=.tmp/go/$(TMP_PROJECT_NAME)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

all: test lint

tidy:
	go mod tidy -v

build:
	go build ./...

test: build
	go test -cover -race -v ./...

test-coverage:
	go test ./... -race -coverprofile=.testCoverage.txt && go tool cover -html=.testCoverage.txt

ci-test: build
	go test -race $$(go list ./...) -v -coverprofile .testCoverage.txt

cleanup-test-dir:
	rm -rf $(TMP_DIR)
	mkdir -p $(TMP_DIR)
	cp -r * $(TMP_DIR)

test-template: cleanup-test-dir
	cd $(TMP_DIR) &&\
	HYGEN_OVERWRITE=1 ../../../$(HYGEN_BIN) template init $(TMP_PROJECT_NAME) \
		--repo_path=gitlab.appsflyer.com/rantav \
		--description="My awesome go project" \
		--include_grpc=0 \
		&&\
	make

setup: setup-git-hooks

setup-git-hooks:
	git config core.hooksPath .githooks

lint: $(GOLANGCI_LINT)
	$(GOPATH)/bin/golangci-lint run --fast --enable-all

$(GOLANGCI_LINT):
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)