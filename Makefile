
.DEFAULT_GOAL := build

precommit:
	pre-commit autoupdate && pre-commit install && pre-commit run -a
.PHONY:precommit

lint:
	golangci-lint run
.PHONY:lint

lint-fix:
	golangci-lint run --fix
.PHONY:lint

test:
	go test
.PHONY:test

build:
	go build -mod vendor -o dist/roller *.go
.PHONY:build

vhs: build
	vhs demo.tape
.PHONY: vhs

clean:
	$(RM) -rf dist/
.PHONY:clean
