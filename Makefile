coverage.out: *.go
	go test -race -covermode=atomic -coverprofile=$@
	go tool cover -html=coverage.out