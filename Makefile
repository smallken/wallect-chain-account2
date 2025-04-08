GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

wallect-chain-account2:
	env GO111MODULE=on go build -v $(LDFLAGS)

clean:
	rm wallect-chain-account2

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	wallect-chain-account2 \
	clean \
	test \
	lint