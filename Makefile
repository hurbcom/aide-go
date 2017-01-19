
godep:
	go get github.com/tools/godep
	go get golang.org/x/sys/unix
	go get github.com/liudng/dogo
	godep restore -v ./...

gocov:
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/matm/gocov-html

setup: godep

test:
	godep go test ./...

test-coverage: godep
	godep go test ./... -cover

test-report-xml: godep gocov
	gocov test ./... | gocov-xml > coverage.xml

test-report-html: godep gocov
	gocov test ./... | gocov-html > coverage.html

