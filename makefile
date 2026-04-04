.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -count=1

.PHONY: run
run:
	go run ./...

.PHONY: tidy
tidy:
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying and vendoring module dependencies...'
	go mod verify
	go mod vendor
	@echo 'Formatting .go files...'
	go fmt ./...