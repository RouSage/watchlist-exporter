build:
	go build -o watchlist-exporter -v

test:
	go test -v ./...

clean:
	rm -f watchlist-exporter

deps:
	go get ./...
