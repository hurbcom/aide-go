FROM golang:1.11-alpine AS base

    COPY requirements.apt ./
    RUN apk update && \
        apk upgrade && \
        apk add --update --no-cache `cat requirements.apt` && \
        rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

    RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | DEP_RELEASE_TAG=v0.5.0 sh

    WORKDIR ${GOPATH}/src/github.com/hurbcom/aide-go/

    COPY Gopkg.lock Gopkg.toml ./

    RUN dep ensure -v -vendor-only


FROM base AS ready

    ENV CGO_ENABLED 0

    COPY . ./

