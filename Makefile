.DEFAULT_GOAL := help
.PHONY: help

APP_DIR=/go/src/github.com/hurbcom/${APP_NAME}
APP_NAME?=$(shell pwd | xargs basename)
DEP:=$(shell command -v dep 2> /dev/null)
DOCKER_IMAGE_NAME:=hu/${APP_NAME}:latest
INTERACTIVE:=$(shell [ -t 0 ] && echo i || echo d)
PROJECT_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
PWD=$(shell pwd)

welcome:
	@printf "\033[33m       _     _                         \n"
	@printf "\033[33m  __ _(_) __| | ___        __ _  ___   \n"
	@printf "\033[33m / _\` | |/ _\` |/ _ \_____ / _\` |/ _ \  \n"
	@printf "\033[33m| (_| | | (_| |  __/_____| (_| | (_) | \n"
	@printf "\033[33m \__,_|_|\__,_|\___|      \__, |\___/  \n"
	@printf "\033[33m                          |___/        \n"
	@printf "\033[0m\n"

setup: sanitize ## Used to develop
ifndef DEP
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | DEP_RELEASE_TAG=v0.5.0 sh
endif
	@GO111MODULE=off dep ensure -v

setup-docker: welcome sanitize build-docker-image ## Install dependencies to run on Docker

build-docker-image:
	@docker build . -t ${DOCKER_IMAGE_NAME}

test:
	@GO111MODULE=off go clean --testcache
	@GO111MODULE=off go test ./... -race # | grep -vE "level|Testing"

sanitize:
	-@rm -rf vendor* _vendor* coverage.xml

ci: build-docker-image ## Runs test coverage to CI
	@echo "Running test in docker"
	@docker run --rm \
		-v ${PWD}:${APP_DIR} \
		-w ${APP_DIR} \
		--name ${APP_NAME}-ci \
		${DOCKER_IMAGE_NAME} \
		sh -c "dep ensure -v -vendor-only && rm -f coverage.xml && go get github.com/axw/gocov/gocov && go get github.com/AlekSi/gocov-xml && gocov test ./... | gocov-xml > coverage.xml"

format:
	@GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	@GO111MODULE=off goimports -l -w -d ${PROJECT_FILES}
	@GO111MODULE=off gofmt -l -s -w ${PROJECT_FILES}

vet: ## Reports suspicious constructs
	@GO111MODULE=off go tool vet ${PWD}

lint: ## Built-in code verifier
	@GO111MODULE=off go get github.com/mgechev/revive
	@GO111MODULE=off revive -exclude vendor/... -formatter stylish ./...

help: welcome
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'
