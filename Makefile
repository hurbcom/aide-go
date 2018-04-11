.DEFAULT_GOAL := help
.PHONY: help

welcome:
	@printf "\n"
	@printf "\033[33m Aide go \n"
	@printf "\n"
	@printf "\033[0m"

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/matm/gocov-html

dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure -v

setup: dep ## Used to develop

test:
	go test ./...

test-dev:
	# -@go test ./legacy -run ^TestGetSupplier$$
	# -@go test ./... | grep -v level
	-@go test ./...

test-coverage:
	go test ./... -cover

run:
	go run main.go --verbose

install-os:
	apt-get update
	apt-get -y install `cat /tmp/requirements.apt`

bin: dep
	GOOS=linux GOARCH=amd64 go build

ci: install-os dep gocov
	gocov test ./... | gocov-xml > coverage.xml

docker-test:
	@docker run --rm -v ${PWD}/gitconfig:/root/.gitconfig -v ${PWD}/requirements.apt:/tmp/requirements.apt -v ${PWD}:/go/src/github.com/hotelurbano/aide-go -w /go/src/github.com/hotelurbano/aide-go --name docker-test golang:1 make ci

format:
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
