test:
	@GOCACHE=off go test -race -cover -timeout=5s ./...
