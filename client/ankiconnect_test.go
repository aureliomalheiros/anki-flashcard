package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	baseURL := "http://localhost:8765"
	client := NewClient(baseURL)

	if client.baseURL != baseURL {
		t.Errorf("baseURL = %v, want %v", client.baseURL, baseURL)
	}

	if client.httpClient.Timeout != requestTimeout {
		t.Errorf("timeout = %v, want %v", client.httpClient.Timeout, requestTimeout)
	}
}

func TestClient_TestConnection(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse string
		statusCode     int
		wantErr        bool
	}{
		{
			name:           "successful connection",
			serverResponse: `{"result": 6, "error": null}`,
			statusCode:     200,
			wantErr:        false,
		},
		{
			name:           "anki error response",
			serverResponse: `{"result": null, "error": "AnkiConnect not available"}`,
			statusCode:     200,
			wantErr:        true,
		},
		{
			name:           "http error",
			serverResponse: ``,
			statusCode:     500,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()

			err := client.TestConnection(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetVersion(t *testing.T) {
	tests := []struct {
		name            string
		serverResponse  string
		statusCode      int
		expectedVersion int
		wantErr         bool
	}{
		{
			name:            "successful version request",
			serverResponse:  `{"result": 6, "error": null}`,
			statusCode:      200,
			expectedVersion: 6,
			wantErr:         false,
		},
		{
			name:            "version as float",
			serverResponse:  `{"result": 6.0, "error": null}`,
			statusCode:      200,
			expectedVersion: 6,
			wantErr:         false,
		},
		{
			name:            "anki error",
			serverResponse:  `{"result": null, "error": "Version not available"}`,
			statusCode:      200,
			expectedVersion: 0,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()

			version, err := client.GetVersion(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if version != tt.expectedVersion {
				t.Errorf("GetVersion() version = %v, want %v", version, tt.expectedVersion)
			}
		})
	}
}

func TestClient_GetDeckNames(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse string
		statusCode     int
		expectedDecks  []string
		wantErr        bool
	}{
		{
			name:           "successful deck list",
			serverResponse: `{"result": ["Default", "English", "Math"], "error": null}`,
			statusCode:     200,
			expectedDecks:  []string{"Default", "English", "Math"},
			wantErr:        false,
		},
		{
			name:           "empty deck list",
			serverResponse: `{"result": [], "error": null}`,
			statusCode:     200,
			expectedDecks:  []string{},
			wantErr:        false,
		},
		{
			name:           "null result",
			serverResponse: `{"result": null, "error": null}`,
			statusCode:     200,
			expectedDecks:  []string{},
			wantErr:        false,
		},
		{
			name:           "anki error",
			serverResponse: `{"result": null, "error": "Cannot get deck names"}`,
			statusCode:     200,
			expectedDecks:  nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			client := NewClient(server.URL)
			ctx := context.Background()

			decks, err := client.GetDeckNames(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeckNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !slicesEqual(decks, tt.expectedDecks) {
				t.Errorf("GetDeckNames() decks = %v, want %v", decks, tt.expectedDecks)
			}
		})
	}
}

func TestClient_CreateDeck(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req AnkiRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.Action != "createDeck" {
			t.Errorf("Expected action 'createDeck', got %v", req.Action)
		}

		if req.Version != ankiConnectVersion {
			t.Errorf("Expected version %v, got %v", ankiConnectVersion, req.Version)
		}

		deckName, ok := req.Params["deck"].(string)
		if !ok || deckName != "TestDeck" {
			t.Errorf("Expected deck name 'TestDeck', got %v", deckName)
		}

		w.Write([]byte(`{"result": null, "error": null}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	err := client.CreateDeck(ctx, "TestDeck")
	if err != nil {
		t.Errorf("CreateDeck() error = %v, want nil", err)
	}
}

func TestClient_AddNote(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req AnkiRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.Action != "addNote" {
			t.Errorf("Expected action 'addNote', got %v", req.Action)
		}

		note, ok := req.Params["note"].(map[string]interface{})
		if !ok {
			t.Errorf("Expected note parameter")
			return
		}

		if note["deckName"] != "TestDeck" {
			t.Errorf("Expected deckName 'TestDeck', got %v", note["deckName"])
		}

		if note["modelName"] != "Basic" {
			t.Errorf("Expected modelName 'Basic', got %v", note["modelName"])
		}

		w.Write([]byte(`{"result": 1234567890, "error": null}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	fields := map[string]string{
		"Front": "Test Question",
		"Back":  "Test Answer",
	}
	tags := []string{"test", "golang"}

	err := client.AddNote(ctx, "TestDeck", "Basic", fields, tags)
	if err != nil {
		t.Errorf("AddNote() error = %v, want nil", err)
	}
}

func TestClient_makeRequest_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(`{"result": null, "error": null}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.httpClient.Timeout = 100 * time.Millisecond

	ctx := context.Background()
	_, err := client.makeRequest(ctx, "version", nil)

	if err == nil {
		t.Error("Expected timeout error, got nil")
	}

	if !strings.Contains(err.Error(), "timeout") && !strings.Contains(err.Error(), "deadline") {
		t.Errorf("Expected timeout/deadline error, got %v", err)
	}
}

func TestClient_makeRequest_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	_, err := client.makeRequest(ctx, "version", nil)
	if err == nil {
		t.Error("Expected JSON decode error, got nil")
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
