build:
	go build -o out/watchlist-exporter main.go

test:
	go test ./...
