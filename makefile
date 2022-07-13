run: |
	gofmt -w .
	go run ./cmd/main.go

mock:
	mockgen -source=github.com/decadevs/lunch-api/internal/ports/repository.go -destination=internal/adapters/repository/mocks/db_mock.go -package=mocks