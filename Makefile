build:
	@go build -o bin/vitshop cmd/main.go

run: build
	@./bin/vitshop
