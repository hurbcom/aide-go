
godep:
	go get github.com/tools/godep
	go get golang.org/x/sys/unix
	go get github.com/liudng/dogo
	godep restore -v ./...

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/matm/gocov-html

lint:
	@go get -u github.com/golang/lint/golint
	@golint ./... > golint.txt

setup: godep

format:
	@goimports -w .
	@gofmt -s -w .

test: format lint
	godep go test ./...

ci: godep gocov fmt
	gocov test ./... | gocov-xml > coverage.xml

docker-test:
	@docker run --rm -v `pwd`:/go/src/github.com/hotelurbano/aide-go -w /go/src/github.com/hotelurbano/aide-go golang:1.7.5 make ci

coverage: ci
