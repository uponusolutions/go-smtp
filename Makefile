test:
	go test ./...

race:
	go test  ./... -race

cover:
	go test ./... -tags cover -race -coverprofile=coverage.out
	go tool -modfile=go.tool.mod cover -html=coverage.out -o coverage.html

lint:
	echo "Linting go"
	go tool -modfile=go.tool.mod golangci-lint run

bench:
	go test ./... -bench . -benchtime=10s -run ^$$

pprof:
	go test . -bench ^Benchmark/^SmallWithChunking$$ -benchtime=10s -run ^$$ -cpuprofile cpu.pprof -memprofile mem.pprof
	go tool pprof -http=":8000" cpu.pprof

vet:
	go vet ./...
