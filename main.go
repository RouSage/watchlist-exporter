package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"watchlist-exporter/config"
	"watchlist-exporter/internal/watchlist"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	file, err := os.Open("./data/movies.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	parsedWatchlist, err := watchlist.ReadWatchlist(file)
	if err != nil {
		log.Fatalf("Error reading watchlist: %v", err)
	}
	fmt.Printf("%+v\n", parsedWatchlist)
	fmt.Println(parsedWatchlist[0].Created.Format(time.DateOnly))
}
