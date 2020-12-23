.PHONY: test
test:
	go test -v ./. -count 1 -race

.PHONY: cover
cover:
	go test -v ./. -covermode=count -coverprofile=coverage.txt

.PHONY: view-cover
view-cover: cover
	go tool cover -html coverage.txt

setup:
	rm go.sum
	go mod tidy
	go mod edit -fmt go.mod
	cd testdata/src/a
	rm go.sum
	go mod tidy
	go mod edit -fmt go.mod

setup/tools:
	GO111MODULE=off go get -u \
		github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run ./.
