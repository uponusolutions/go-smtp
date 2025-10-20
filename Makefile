
test:
	go test ./...

race:
	go test  ./... -race

cover:
	go test ./... -race -coverprofile=coverage.out
	go tool -modfile=go.tool.mod cover -html=coverage.out -o coverage.html
	go tool -modfile=go.tool.mod cover -func=coverage.out

lint:
	@echo "Linting go"
	@go tool -modfile=go.tool.mod golangci-lint run

bench:
	@go test ./... -bench=.

pprof:
	@go test ./internal/benchmark -cpuprofile cpu.pprof -memprofile mem.pprof -bench ^Benchmark/^SmallWithChunkingSameConnection$
	@go tool pprof -http=":8000" cpu.pprof

vet:
	@go vet ./...
