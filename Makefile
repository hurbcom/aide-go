
godep:
	go get github.com/tools/godep
	go get golang.org/x/sys/unix
	go get golang.org/x/tools/cmd/goimports
	godep restore -v ./...

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/matm/gocov-html

install-os:
	apt-get update
	apt-get -y install `cat requirements.apt`

lint:
	go get -u github.com/golang/lint/golint
	golint ./... > golint.txt

setup: godep

format:
	goimports -w .
	gofmt -s -w .

test: format
	godep go test ./...

ci: install-os godep gocov format
	gocov test ./... | gocov-xml > coverage.xml

docker-test:
	@docker run --rm -v `pwd`:/go/src/github.com/hotelurbano/aide-go -w /go/src/github.com/hotelurbano/aide-go golang:1 make ci

coverage: ci
