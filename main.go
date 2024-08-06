package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"watchlist-exporter/internal/watchlist"
)

func main() {
	file, err := os.Open("./data/movies.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	parsedWatchlist := watchlist.ReadWatchlist(file)
	fmt.Printf("%+v", parsedWatchlist)
	fmt.Println(parsedWatchlist[0].Created.Format(time.DateOnly))
}
