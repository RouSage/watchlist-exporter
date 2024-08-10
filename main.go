package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"watchlist-exporter/config"
	"watchlist-exporter/internal/notion"
	"watchlist-exporter/internal/watchlist"

	"github.com/jomei/notionapi"
)

var filePath = flag.String("path", "", "Path to the CSV file")
var dbName = flag.String("database-name", "Watchlist DB", "Name of the Notion database to create. No effect if --database-id is provided")
var dbId = flag.String("database-id", "", "ID of the existing Notion database. New database will be created if not provided")

func main() {
	flag.Parse()

	if *filePath == "" {
		log.Fatal("File path is required")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	fmt.Printf("Reading watchlist from %q\n", file.Name())
	parsedWatchlist, err := watchlist.ReadWatchlist(file)
	if err != nil {
		log.Fatalf("Error reading watchlist: %v", err)
	}
	if len(parsedWatchlist) == 0 {
		fmt.Println("Empy watchlist")
		os.Exit(0)
	}

	notionClient := notion.New(cfg.NotionKey)
	var database *notionapi.Database

	if *dbId == "" {
		fmt.Println("No database ID provided")

		database, err = notionClient.CreateDatabase(cfg.NotionPageID, *dbName, true)
		if err != nil {
			log.Fatalf("Error creating database: %v", err)
		}
	} else {
		fmt.Printf("Database ID provided - %q\n", *dbId)

		database, err = notionClient.GetDatabase(*dbId)
		if err != nil {
			log.Fatalf("Error retrieving database: %v", err)
		}
	}

	if pagesCreated, err := notionClient.ExportWathlist(database.ID, parsedWatchlist); err != nil {
		log.Fatalf("Error exporting watchlist: %v. Created %d pages", err, pagesCreated)
	}
}
