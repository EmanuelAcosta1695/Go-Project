# This file is being used to bootstrap some of common tasks
# with the @ the commands are not printed out
build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -b ./...