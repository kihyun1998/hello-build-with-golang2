build:
	go build -o $(FILE_NAME) .

test:
	go test -v -cover ./test

.PHONY: build, test