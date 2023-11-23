build:
	go build -o $(TAG_NAME)_wsub

test:
	go test -v .

.PHONY: build