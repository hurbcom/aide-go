
godep:
	go get github.com/tools/godep
	go get golang.org/x/sys/unix
	go get github.com/liudng/dogo
	godep restore -v ./...

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/matm/gocov-html

fmt:
	gofmt -s -w .

lint:
	go get -u github.com/golang/lint/golint
	golint ./...

setup: godep

test: fmt lint
	godep go test ./...

ci: godep gocov fmt
	gocov test ./... | gocov-xml > coverage.xml

coverage: ci
