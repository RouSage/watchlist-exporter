package main

import (
	"fmt"
	"log"
	"os"

	"watchlist-exporter/config"
	"watchlist-exporter/internal/notion"
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
	if len(parsedWatchlist) == 0 {
		fmt.Println("Empy watchlist")
		os.Exit(0)
	}

	notionClient := notion.New(cfg.NotionKey)
	database, err := notionClient.CreateDatabase(cfg.NotionPageID, "Watchlist DB", true)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	pagesCreated, err := notionClient.ExportWathlist(database.ID, parsedWatchlist)
	if err != nil {
		log.Fatalf("Error creating page: %v. Created %d pages", err, pagesCreated)
	}

	fmt.Printf("Created %d pages\n", pagesCreated)
}
