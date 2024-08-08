package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NotionKey    string
	NotionPageID string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	config := &Config{}

	config.NotionKey = os.Getenv("NOTION_KEY")
	if config.NotionKey == "" {
		return nil, fmt.Errorf("NOTION_KEY is not set")
	}

	config.NotionPageID = os.Getenv("NOTION_PAGE_ID")
	if config.NotionPageID == "" {
		return nil, fmt.Errorf("NOTION_PAGE_ID is not set")
	}

	return config, nil
}
