run: |
	gofmt -w .
	go run ./cmd/main.go

mock:
	mockgen -source=internal/ports/repository.go -destination=internal/adapters/repository/mocks/db_mock.go -package=mocks