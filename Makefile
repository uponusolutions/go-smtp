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
	go test ./... -bench . -count=10 -run ^$$ > .bench/$$(date +%Y%m%dT%H%M%S).txt
    # go test ./... -bench '^Benchmark$/^SmallWithChunking$' -benchtime=10s -run ^$

stats:
	go tool -modfile=go.tool.mod benchstat .bench/*.txt

pprof:
	go test . -bench ^Benchmark/^SmallWithChunking$$ -benchtime=10s -run ^$$ -cpuprofile cpu.pprof -memprofile mem.pprof
	go tool pprof -http=":8000" cpu.pprof

vet:
	go vet ./...
