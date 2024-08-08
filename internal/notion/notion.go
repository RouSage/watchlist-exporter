package notion

import (
	"context"
	"fmt"
	"watchlist-exporter/internal/watchlist"

	"github.com/jomei/notionapi"
)

const (
	TITLE        = "Title"
	URL          = "URL"
	TYPE         = "Type"
	DATE_WATCHED = "Date Watched"
	RELEASE_DATE = "Release Date"
)

type notionClient struct {
	client *notionapi.Client
}

func New(token string) *notionClient {
	return &notionClient{client: notionapi.NewClient(notionapi.Token(token))}
}

func (notion *notionClient) CreateDatabase(parentPageID string, title string, isInline bool) (*notionapi.Database, error) {
	return notion.client.Database.Create(
		context.Background(),
		&notionapi.DatabaseCreateRequest{
			Parent: notionapi.Parent{
				Type:   "page_id",
				PageID: notionapi.PageID(parentPageID),
			},
			IsInline: isInline,
			Title: []notionapi.RichText{
				{
					Type: "text",
					Text: &notionapi.Text{
						Content: title,
					},
				},
			},
			Properties: notionapi.PropertyConfigs{
				TITLE: notionapi.TitlePropertyConfig{Type: "title"},
				URL:   notionapi.URLPropertyConfig{Type: "url"},
				TYPE: notionapi.SelectPropertyConfig{
					Type: "select",
					Select: notionapi.Select{
						Options: []notionapi.Option{},
					},
				},
				DATE_WATCHED: notionapi.DatePropertyConfig{Type: "date"},
				RELEASE_DATE: notionapi.DatePropertyConfig{Type: "date"},
			},
		})
}

func (notion *notionClient) ExportWathlist(databaseId notionapi.ObjectID, watchlist []watchlist.Watchlist) (int, error) {
	pagesCreated := 0

	for i, watchlist := range watchlist {
		fmt.Printf("Page %d: %s", i, watchlist.Title)

		_, err := notion.createPage(databaseId, watchlist)
		if err != nil {
			return pagesCreated, err
		}

		fmt.Println(" - Success")

		pagesCreated += 1
	}

	return pagesCreated, nil
}

func (notion *notionClient) createPage(databaseId notionapi.ObjectID, watchlist watchlist.Watchlist) (*notionapi.Page, error) {
	return notion.client.Page.Create(
		context.Background(),
		&notionapi.PageCreateRequest{
			Parent: notionapi.Parent{
				Type:       "database_id",
				DatabaseID: notionapi.DatabaseID(databaseId),
			},
			Properties: notionapi.Properties{
				TITLE: notionapi.TitleProperty{
					Type: "title",
					Title: []notionapi.RichText{
						{
							Type: "text",
							Text: &notionapi.Text{
								Content: watchlist.Title,
							},
						},
					},
				},
				URL: notionapi.URLProperty{
					Type: "url",
					URL:  watchlist.URL,
				},
				TYPE: notionapi.SelectProperty{
					Type: "select",
					Select: notionapi.Option{
						Name: watchlist.Type,
					},
				},
				DATE_WATCHED: notionapi.DateProperty{
					Type: "date",
					Date: &notionapi.DateObject{
						Start: (*notionapi.Date)(&watchlist.Created),
					},
				},
				RELEASE_DATE: notionapi.DateProperty{
					Type: "date",
					Date: &notionapi.DateObject{
						Start: (*notionapi.Date)(&watchlist.ReleaseDate),
					},
				},
			},
		})
}
