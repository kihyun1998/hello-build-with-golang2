build:
	go build -o $(TAG_NAME)_wsub .

test:
	go test -v -cover ./calc

.PHONY: build, test