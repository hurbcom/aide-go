.DEFAULT_GOAL := help
.PHONY: help

APP_DIR=/go/src/github.com/hurbcom/${APP_NAME}
APP_NAME?=$(shell pwd | xargs basename)
DEP:=$(shell command -v dep 2> /dev/null)
DOCKER_IMAGE_NAME:=hu/${APP_NAME}:latest
GITHASH=$(shell git rev-parse --verify HEAD)
GITTAG=$(shell git describe --abbrev=0 --tags)
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
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | DEP_RELEASE_TAG=v0.5.0 sh
endif
	@dep ensure -v

setup-docker: welcome build-docker-image ## Install dependencies to run on Docker
ifeq ($(shell [ -f ./config.yml ] || echo -n no),no)
	@cp config.yml.sample config.yml
endif

build-docker-image:
	docker build . -t ${DOCKER_IMAGE_NAME} ${TARGET}

test:
	@go clean --testcache
	@go test ./... -race # | grep -vE "level|Testing"

test-dev:
	@go clean --testcache
	go test ./... -run ^${TEST}$$ -failfast | grep -vE "level|Testing"

test-in-docker:
	@docker run --rm ${DOCKER_IMAGE_NAME} go test ./... | grep -vE "level|Testing"

run:
	@go run main.go --verbose

sanitize:
	-@rm -rf vendor* _vendor* coverage.xml

ci: sanitize ## Runs test coverage to CI
	@$(MAKE) TARGET="--target ready" build-docker-image
	@echo "Running test in docker"
	docker run --rm \
		-v ${PWD}:${APP_DIR} \
		-w ${APP_DIR} \
		--name ${APP_NAME}-ci \
		${DOCKER_IMAGE_NAME} \
		sh -c "dep ensure -v && rm -f coverage.xml && go get github.com/axw/gocov/gocov && go get github.com/AlekSi/gocov-xml && gocov test ./... | gocov-xml > coverage.xml"

format:
	@go get golang.org/x/tools/cmd/goimports
	goimports -l -w -d ${PROJECT_FILES}
	gofmt -l -s -w ${PROJECT_FILES}

vet: ## Reports suspicious constructs
	go tool vet ${PWD}

gometalinter: ## Multi code verifier
	@go get github.com/alecthomas/gometalinter
	@gometalinter --install
	@gometalinter --config=gometalinter.json ./... | grep -v '^vendor.*' > gometalinter.txt

lint: ## Built-in code verifier
	@go get github.com/mgechev/revive
	@revive -exclude vendor/... -formatter stylish ./...

help: welcome
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'
