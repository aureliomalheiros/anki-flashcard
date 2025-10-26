package config

import "os"

const defaultAnkiConnectURL = "http://localhost:8765"

type Config struct {
	AnkiConnectURL string `json:"anki_connect_url"`
	DefaultDeck    string `json:"default_deck"`
}

func LoadConfig() (*Config, error) {
	config := &Config{
		AnkiConnectURL: defaultAnkiConnectURL,
		DefaultDeck:    "Default",
	}

	if url := os.Getenv("ANKI_CONNECT_URL"); url != "" {
		config.AnkiConnectURL = url
	}
	if deck := os.Getenv("ANKI_DEFAULT_DECK"); deck != "" {
		config.DefaultDeck = deck
	}

	return config, nil
}
