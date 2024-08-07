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

	parsedWatchlist, err := watchlist.ReadWatchlist(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", parsedWatchlist)
	fmt.Println(parsedWatchlist[0].Created.Format(time.DateOnly))
}
