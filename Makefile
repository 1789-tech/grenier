.PHONY: test run install

test:
	go test ./...

run:
	go run ./cmd/grenier-evals

install:
	./install.sh
