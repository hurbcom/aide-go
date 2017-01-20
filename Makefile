
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
	go get golang.org/x/tools
	go get github.com/golang/lint/golint
	golint ./...

setup: godep

test: fmt lint
	godep go test ./...

test-coverage: godep
	godep go test ./... -cover

test-report-xml: godep gocov
	gocov test ./... | gocov-xml > coverage.xml

test-report-html: godep gocov
	gocov test ./... | gocov-html > coverage.html

ci: setup gocov test-report-xml fmt lint
	godep go test ./...
