GOPATH=`pwd`/vendor

build:
	env GOPATH=$(GOPATH) CGO_ENABLED=0 go build -o dmarc-report
