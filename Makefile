GOBIN := ${CURDIR}/bin

.PHONY: build
build:
	go build \
		-o ${GOBIN}/aliasimport \
		./cmd/aliasimport

.PHONY: test
test:
	go test -v ./. -count 1 -race

.PHONY: cover
cover:
	go test -v ./. -covermode=count -coverprofile=coverage.txt

.PHONY: view-cover
view-cover: cover
	go tool cover -html coverage.txt

.PHONY: setup
setup:
	rm go.sum
	go mod tidy
	go mod edit -fmt go.mod
	cd testdata/src/a
	rm go.sum
	go mod tidy
	go mod edit -fmt go.mod
	cd testdata/src/b
	rm go.sum
	go mod tidy
	go mod edit -fmt go.mod

setup/tools:
	GO111MODULE=off go get -u \
		github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run ./.

b_dir := ${CURDIR}/testdata/src/b
.PHONY: test-run
test-run: build
	@cd $(b_dir)
	@go vet -vettool=$(GOBIN)/aliasimport -aliasimport.rule=$(b_dir)/rules.yml $(b_dir)/main.go
