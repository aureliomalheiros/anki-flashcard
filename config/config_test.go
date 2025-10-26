package config

import (
	"os"
	"testing"
)

func TestLoadConfig_Defaults(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() error = %v, want nil", err)
	}

	if cfg.AnkiConnectURL != defaultAnkiConnectURL {
		t.Errorf("AnkiConnectURL = %v, want %v", cfg.AnkiConnectURL, defaultAnkiConnectURL)
	}

	if cfg.DefaultDeck != "Default" {
		t.Errorf("DefaultDeck = %v, want %v", cfg.DefaultDeck, "Default")
	}
}

func TestLoadConfig_EnvironmentVariables(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected Config
	}{
		{
			name: "custom anki connect url",
			envVars: map[string]string{
				"ANKI_CONNECT_URL": "http://custom:9999",
			},
			expected: Config{
				AnkiConnectURL: "http://custom:9999",
				DefaultDeck:    "Default",
			},
		},
		{
			name: "custom default deck",
			envVars: map[string]string{
				"ANKI_DEFAULT_DECK": "MyCustomDeck",
			},
			expected: Config{
				AnkiConnectURL: defaultAnkiConnectURL,
				DefaultDeck:    "MyCustomDeck",
			},
		},
		{
			name: "both custom values",
			envVars: map[string]string{
				"ANKI_CONNECT_URL":  "http://test:8888",
				"ANKI_DEFAULT_DECK": "TestDeck",
			},
			expected: Config{
				AnkiConnectURL: "http://test:8888",
				DefaultDeck:    "TestDeck",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			cfg, err := LoadConfig()
			if err != nil {
				t.Fatalf("LoadConfig() error = %v, want nil", err)
			}

			if cfg.AnkiConnectURL != tt.expected.AnkiConnectURL {
				t.Errorf("AnkiConnectURL = %v, want %v", cfg.AnkiConnectURL, tt.expected.AnkiConnectURL)
			}

			if cfg.DefaultDeck != tt.expected.DefaultDeck {
				t.Errorf("DefaultDeck = %v, want %v", cfg.DefaultDeck, tt.expected.DefaultDeck)
			}
		})
	}
}
