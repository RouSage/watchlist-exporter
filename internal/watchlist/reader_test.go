package watchlist

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReadWatchlist(t *testing.T) {
	t.Run("Successful Read", func(t *testing.T) {
		csvContent := `Position,Const,Created,Modified,Description,Title,Original Title,URL,Title Type,IMDb Rating,Runtime (mins),Year,Genres,Num Votes,Release Date,Directors,Your Rating,Date Rated
		1,tt0120611,2020-01-26,2020-01-26,,"Blade","Blade",https://www.imdb.com/title/tt0120611/,Movie,7.1,120,1998,"Action, Horror, Sci-Fi",303301,1998-08-21,"Stephen Norrington",10,2024-04-01
		2,tt0187738,2020-01-26,2020-01-26,,"Blade II","Blade II",https://www.imdb.com/title/tt0187738/,Movie,6.7,117,2002,"Action, Horror, Sci-Fi, Thriller",236854,2002-03-22,"Guillermo del Toro",8,2024-04-05
`
		reader := strings.NewReader(csvContent)
		watchlist, err := ReadWatchlist(reader)

		if err != nil {
			t.Fatalf("ReadWatchlist returned an error: %v", err)
		}

		if len(watchlist) != 2 {
			t.Errorf("Expected watchlist length 2, got %d", len(watchlist))
		}

		expectedItem := []Watchlist{
			{Created: time.Date(2020, 1, 26, 0, 0, 0, 0, time.UTC),
				Title:       "Blade",
				URL:         "https://www.imdb.com/title/tt0120611/",
				Type:        "Movie",
				ReleaseDate: time.Date(1998, 8, 21, 0, 0, 0, 0, time.UTC),
			},
			{Created: time.Date(2020, 1, 26, 0, 0, 0, 0, time.UTC),
				Title:       "Blade II",
				URL:         "https://www.imdb.com/title/tt0187738/",
				Type:        "Movie",
				ReleaseDate: time.Date(2002, 3, 22, 0, 0, 0, 0, time.UTC),
			},
		}

		for i, item := range watchlist {
			if !reflect.DeepEqual(item, expectedItem[i]) {
				t.Errorf("Expected item %+v, got %+v", expectedItem[i], item)
			}
		}
	})

	t.Run("Empty Result", func(t *testing.T) {
		tests := []struct {
			content        string
			expectedLength int
		}{
			{content: `Created,Title,URL,Type,ReleaseDate`, expectedLength: 0},
			{content: `Just a header`, expectedLength: 0},
			{content: "", expectedLength: 0},
		}

		for _, test := range tests {
			reader := strings.NewReader(test.content)
			watchlist, err := ReadWatchlist(reader)

			if err != nil {
				t.Fatalf("ReadWatchlist returned an error: %v", err)
			}

			if len(watchlist) != test.expectedLength {
				t.Errorf("Expected watchlist length %d, got %d", test.expectedLength, len(watchlist))
			}
		}
	})

	t.Run("Invalid Date Format", func(t *testing.T) {
		csvContent := `Created,Title,URL,Type,ReleaseDate
invalid-date,Movie 1,http://example.com/movie1,Movie,2023-05-15
`
		reader := strings.NewReader(csvContent)
		_, err := ReadWatchlist(reader)

		if err == nil {
			t.Error("Expected an error for invalid date format, got nil")
		}
	})

	t.Run("Missing Fields", func(t *testing.T) {
		csvContent := `Created,Title,URL,Type,ReleaseDate
2023-01-01,Movie 1,http://example.com/movie1
`
		reader := strings.NewReader(csvContent)
		_, err := ReadWatchlist(reader)

		if err == nil {
			t.Error("Expected an error for missing fields, got nil")
		}
	})

}
