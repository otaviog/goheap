about:
	@echo Project maintenance tasks.

build:
	go build

test:
	go test 

coverage.out: *.go
	go test -race -covermode=atomic -coverprofile=$@
	go tool cover -html=coverage.out

benchmark:
	GOMAXPROCS=1 go test -bench=. -count=5
