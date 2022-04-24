coverage.out: *.go
	go test -race -covermode=atomic -coverprofile=$@
	go tool cover -html=coverage.out

benchmark:
	go test -bench=. -count=5
