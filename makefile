test:
	goimports -l ./
	go vet ./...
	go test -v ./...