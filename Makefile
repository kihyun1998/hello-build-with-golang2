build:
	go build -o wsub .

test:
	go test -v -cover ./calc

.PHONY: build, test