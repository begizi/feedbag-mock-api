build:
	@go build -o serve
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o serve.linux
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o serve.exe

test:
	@go test -v

.PHONY: build test
