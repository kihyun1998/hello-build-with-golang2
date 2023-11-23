build:
	go build -o $(TAG_NAME)_wadd

test:
	go test -v .

.PHONY: build