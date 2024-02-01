build:
	go build -o $(FILE_NAME) .

test:
	go test -v -cover ./calc

.PHONY: build, test