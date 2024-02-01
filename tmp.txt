build:
	go build -o -buildvcs=false $(TAG_NAME)_wsub .

test:
	go test -v -cover ./calc

.PHONY: build, test