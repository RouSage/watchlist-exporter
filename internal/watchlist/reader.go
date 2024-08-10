package watchlist

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"
)

const (
	CREATED      = 2
	TITLE        = 5
	URL          = 7
	TYPE         = 8
	RELEASE_DATE = 14
)

type Watchlist struct {
	Title       string
	URL         string
	Created     time.Time
	Type        string
	ReleaseDate time.Time
}

func ReadWatchlist(file io.Reader) ([]Watchlist, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var watchlist []Watchlist

	for i, line := range data {
		// Skip the header row
		if i == 0 {
			fmt.Printf("Header row: [%v]. Skipping...\n", strings.Join(line, ","))
			continue
		}

		item := Watchlist{}
		for j, field := range line {
			switch j {
			case CREATED:
				item.Created, err = parseDate(field)
				if err != nil {
					return nil, err
				}
			case TITLE:
				item.Title = field
			case URL:
				item.URL = field
			case TYPE:
				item.Type = field
			case RELEASE_DATE:
				item.ReleaseDate, err = parseDate(field)
				if err != nil {
					return nil, err
				}
			}
		}

		watchlist = append(watchlist, item)
	}

	fmt.Printf("Parsed %d watchlist items\n", len(watchlist))

	return watchlist, nil
}

func parseDate(date string) (time.Time, error) {
	if len(date) == 7 {
		// Try parsing the date as a month
		return time.Parse("2006-01", date)
	}

	if len(date) == 4 {
		// Try parsing the date as a year
		return time.Parse("2006", date)
	}

	return time.Parse("2006-01-02", date)
}
