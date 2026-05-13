APP=hello
TESTPKGS=./...
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

test:
	go test $(TESTPKGS)

build:
	go build -o bin/$(APP) ./cmd/$(APP)

run:
	go run ./cmd/$(APP)

coverage:
	go test -coverprofile=$(COVERAGE_FILE) -covermode=atomic $(TESTPKGS)
##	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
##	@echo "Coverage report generated: $(COVERAGE_HTML)"

format:
	gofmt -w ./...

clean:
	rm -rf bin $(COVERAGE_FILE) $(COVERAGE_HTML)

lint:
	go vet ./...