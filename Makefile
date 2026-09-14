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

benchlog:
	GOMAXPROCS=1 go test ./... -bench . -benchmem -count=10 -run ^$$ | tee .bench/$$(date +%Y%m%dT%H%M%S).txt

BENCH	?= .
TIME	?= 1s

bench:
	GOMAXPROCS=1 go test ./... -bench '$(BENCH)' -vet=off -benchmem -benchtime=$(TIME) -run ^$$

bench-mailer:
	GOMAXPROCS=1 go test . -bench '$(BENCH)' -benchmem -benchtime=$(TIME) -run ^$$

bench-dot:
	GOMAXPROCS=1 go test ./internal/textsmtp -bench '$(BENCH)' -benchmem -benchtime=$(TIME) -run ^$$

stats:
	go tool -modfile=go.tool.mod benchstat .bench/*.txt

pprof:
	go test . -bench ^Benchmark/^SmallWithChunking$$ -benchtime=10s -run ^$$ -cpuprofile cpu.pprof -memprofile mem.pprof
	# go test ./mailer -run TestClient_SendMailDirectManyRcptsPipelining -cpuprofile cpu.pprof -memprofile mem.pprof
	go tool pprof -http=":8000" cpu.pprof

vet:
	go vet ./...
