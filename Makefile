.DEFAULT_GOAL := help
.PHONY: help

BASENAME=$(shell pwd | xargs basename)

welcome:
	@printf "\n"
	@printf "\033[33m Aide go \n"
	@printf "\n"
	@printf "\033[0m"

dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure -v

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml

setup: dep ## Used to develop

test:
	go test ./...

test-dev:
	# -@go test ./legacy -run ^TestGetSupplier$$
	# -@go test ./... | grep -v level
	-@go test ./...

run:
	go run main.go --verbose

install-os:
	apt-get update
	apt-get -y install `cat requirements.apt`

bin: dep
	GOOS=linux GOARCH=amd64 go build

sanitize:
	-@rm -rf vendor* _vendor* coverage.xml

ci: sanitize install-os dep gocov
	@go version
	gocov test ./... | gocov-xml > coverage.xml

docker-test:
	@echo "Running test in docker"
	@docker run --rm -v ${PWD}/gitconfig:/root/.gitconfig \
		-v ${PWD}:/go/src/github.com/hotelurbano/${BASENAME} \
		-w /go/src/github.com/hotelurbano/${BASENAME} --name "${BASENAME}-docker-test" golang:1.10.1 \
		make ci

format:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .
	gofmt -s -w .

vet: ## Reports suspicious constructs
	go tool vet ${PWD}

linter: ## Multi code verifier
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter --config=gometalinter.json ./... | grep -v comment | grep -v MixedCaps | grep -v "should be" | grep -v ALL_CAPS > gometalinter.txt

lint: ## Built-in code verifier
	go get github.com/golang/lint/golint
	golint ./... | grep -v comment | grep -v MixedCaps | grep -v "should be" | grep -v ALL_CAPS > golint.txt

help: welcome
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
