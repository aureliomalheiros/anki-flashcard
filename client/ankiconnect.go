package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	ankiConnectVersion = 6
	requestTimeout     = 10 * time.Second
)

type AnkiRequest struct {
	Action  string                 `json:"action"`
	Version int                    `json:"version"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type AnkiResponse struct {
	Result interface{} `json:"result"`
	Error  *string     `json:"error"`
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: requestTimeout,
		},
		baseURL: baseURL,
	}
}

func (c *Client) makeRequest(ctx context.Context, action string, params map[string]interface{}) (*AnkiResponse, error) {
	request := AnkiRequest{
		Action:  action,
		Version: ankiConnectVersion,
		Params:  params,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var ankiResp AnkiResponse
	if err := json.NewDecoder(resp.Body).Decode(&ankiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if ankiResp.Error != nil {
		return nil, fmt.Errorf("anki error: %s", *ankiResp.Error)
	}

	return &ankiResp, nil
}

func (c *Client) TestConnection(ctx context.Context) error {
	_, err := c.makeRequest(ctx, "version", nil)
	return err
}

func (c *Client) GetDeckNames(ctx context.Context) ([]string, error) {
	resp, err := c.makeRequest(ctx, "deckNames", nil)
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		return []string{}, nil
	}

	deckNames, ok := resp.Result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	var decks []string
	for _, name := range deckNames {
		if deckName, ok := name.(string); ok {
			decks = append(decks, deckName)
		}
	}

	return decks, nil
}

func (c *Client) CreateDeck(ctx context.Context, deckName string) error {
	params := map[string]interface{}{
		"deck": deckName,
	}

	_, err := c.makeRequest(ctx, "createDeck", params)
	return err
}

func (c *Client) AddNote(ctx context.Context, deckName, modelName string, fields map[string]string, tags []string) error {
	noteParams := map[string]interface{}{
		"deckName":  deckName,
		"modelName": modelName,
		"fields":    fields,
	}

	if len(tags) > 0 {
		noteParams["tags"] = tags
	}

	params := map[string]interface{}{
		"note": noteParams,
	}

	_, err := c.makeRequest(ctx, "addNote", params)
	return err
}

func (c *Client) GetVersion(ctx context.Context) (int, error) {
	resp, err := c.makeRequest(ctx, "version", nil)
	if err != nil {
		return 0, err
	}

	if version, ok := resp.Result.(float64); ok {
		return int(version), nil
	}

	return 0, fmt.Errorf("unexpected version format")
}

func (c *Client) AddNoteFromYamlCard(ctx context.Context, card interface{}, deckName string) error {
	cardValue := reflect.ValueOf(card)

	toAnkiFieldsMethod := cardValue.MethodByName("ToAnkiFields")
	if !toAnkiFieldsMethod.IsValid() {
		return fmt.Errorf("card type %T does not have ToAnkiFields method", card)
	}

	getTagsMethod := cardValue.MethodByName("GetTags")
	if !getTagsMethod.IsValid() {
		return fmt.Errorf("card type %T does not have GetTags method", card)
	}

	fieldsResult := toAnkiFieldsMethod.Call(nil)
	if len(fieldsResult) != 1 {
		return fmt.Errorf("ToAnkiFields method should return exactly one value")
	}

	fields, ok := fieldsResult[0].Interface().(map[string]string)
	if !ok {
		return fmt.Errorf("ToAnkiFields method should return map[string]string")
	}

	tagsResult := getTagsMethod.Call(nil)
	if len(tagsResult) != 1 {
		return fmt.Errorf("GetTags method should return exactly one value")
	}

	tags, ok := tagsResult[0].Interface().([]string)
	if !ok {
		return fmt.Errorf("GetTags method should return []string")
	}

	return c.AddNote(ctx, deckName, "Basic", fields, tags)
}

func (c *Client) BulkAddNotes(ctx context.Context, notes []map[string]interface{}) error {
	params := map[string]interface{}{
		"notes": notes,
	}

	_, err := c.makeRequest(ctx, "addNotes", params)
	return err
}



func (c *Client) StoreAudioFromURL(ctx context.Context, url, filename string) (string, error) {
	params := map[string]interface{}{
		"url":      url,
		"filename": filename,
	}

	result, err := c.makeRequest(ctx, "downloadFile", params)
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", fmt.Errorf("anki error: %v", result.Error)
	}

	if audioFile, ok := result.Result.(string); ok {
		return audioFile, nil
	}

	return "", fmt.Errorf("unexpected audio filename format")
}
