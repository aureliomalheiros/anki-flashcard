package models

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadCardsFromYAML_ValidFile(t *testing.T) {
	cardSet, err := LoadCardsFromYAML("../testdata/valid_cards.yaml")
	if err != nil {
		t.Fatalf("LoadCardsFromYAML() error = %v, want nil", err)
	}

	if len(cardSet.Cards) != 3 {
		t.Errorf("Expected 3 cards, got %d", len(cardSet.Cards))
	}

	expectedFirst := YamlCard{
		Front:     "test",
		Meaning:   "A procedure intended to establish the quality, performance, or reliability of something",
		Examples:  []string{"We need to test this function", "The test passed successfully"},
		Translate: "Teste",
		Pronounce: "test",
		Lang:      "en",
	}

	if !reflect.DeepEqual(cardSet.Cards[0], expectedFirst) {
		t.Errorf("First card = %+v, want %+v", cardSet.Cards[0], expectedFirst)
	}

	expectedSecond := YamlCard{
		Front:     "hello",
		Meaning:   "Used as a greeting or to begin a phone conversation",
		Examples:  []string{"Hello, how are you?", "Hello world!"},
		Translate: "Olá",
		Pronounce: "hə-ˈlō",
		Lang:      "en",
	}

	if !reflect.DeepEqual(cardSet.Cards[1], expectedSecond) {
		t.Errorf("Second card = %+v, want %+v", cardSet.Cards[1], expectedSecond)
	}

	expectedThird := YamlCard{
		Front:   "minimal",
		Meaning: "A card with minimal information",
		Lang:    "en",
	}

	if !reflect.DeepEqual(cardSet.Cards[2], expectedThird) {
		t.Errorf("Third card = %+v, want %+v", cardSet.Cards[2], expectedThird)
	}
}

func TestLoadCardsFromYAML_EmptyFile(t *testing.T) {
	cardSet, err := LoadCardsFromYAML("../testdata/empty_cards.yaml")
	if err != nil {
		t.Fatalf("LoadCardsFromYAML() error = %v, want nil", err)
	}

	if len(cardSet.Cards) != 0 {
		t.Errorf("Expected 0 cards, got %d", len(cardSet.Cards))
	}
}

func TestLoadCardsFromYAML_NonExistentFile(t *testing.T) {
	_, err := LoadCardsFromYAML("../testdata/nonexistent.yaml")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestLoadCardsFromYAML_MalformedFile(t *testing.T) {
	_, err := LoadCardsFromYAML("../testdata/malformed.yaml")
	if err == nil {
		t.Error("Expected error for malformed YAML, got nil")
	}
}

func TestSaveCardsToYAML(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{
				Front:     "save",
				Meaning:   "To store data",
				Examples:  []string{"Save your work", "Save the file"},
				Translate: "Salvar",
				Pronounce: "sāv",
				Lang:      "en",
			},
			{
				Front:   "load",
				Meaning: "To read data",
				Lang:    "en",
			},
		},
	}

	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_save.yaml")

	err := SaveCardsToYAML(cardSet, tempFile)
	if err != nil {
		t.Fatalf("SaveCardsToYAML() error = %v, want nil", err)
	}

	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Error("Expected file to be created, but it doesn't exist")
	}

	loadedCardSet, err := LoadCardsFromYAML(tempFile)
	if err != nil {
		t.Fatalf("LoadCardsFromYAML() after save error = %v, want nil", err)
	}

	if len(loadedCardSet.Cards) != len(cardSet.Cards) {
		t.Errorf("Loaded cards count = %d, want %d", len(loadedCardSet.Cards), len(cardSet.Cards))
	}

	for i, card := range cardSet.Cards {
		loaded := loadedCardSet.Cards[i]

		if loaded.Front != card.Front {
			t.Errorf("Loaded card %d Front = %v, want %v", i, loaded.Front, card.Front)
		}
		if loaded.Meaning != card.Meaning {
			t.Errorf("Loaded card %d Meaning = %v, want %v", i, loaded.Meaning, card.Meaning)
		}
		if loaded.Translate != card.Translate {
			t.Errorf("Loaded card %d Translate = %v, want %v", i, loaded.Translate, card.Translate)
		}
		if loaded.Pronounce != card.Pronounce {
			t.Errorf("Loaded card %d Pronounce = %v, want %v", i, loaded.Pronounce, card.Pronounce)
		}
		if loaded.Lang != card.Lang {
			t.Errorf("Loaded card %d Lang = %v, want %v", i, loaded.Lang, card.Lang)
		}

		if len(loaded.Examples) != len(card.Examples) {
			t.Errorf("Loaded card %d Examples length = %v, want %v", i, len(loaded.Examples), len(card.Examples))
		} else {
			for j, example := range card.Examples {
				if j < len(loaded.Examples) && loaded.Examples[j] != example {
					t.Errorf("Loaded card %d Example %d = %v, want %v", i, j, loaded.Examples[j], example)
				}
			}
		}
	}
}

func TestSaveCardsToYAML_InvalidPath(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{Front: "test", Meaning: "test"},
		},
	}

	err := SaveCardsToYAML(cardSet, "/invalid/path/that/does/not/exist.yaml")
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}

func TestValidateCardSet_Valid(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{Front: "valid", Meaning: "A valid card"},
			{Front: "another", Meaning: "Another valid card"},
		},
	}

	err := ValidateCardSet(cardSet)
	if err != nil {
		t.Errorf("ValidateCardSet() error = %v, want nil", err)
	}
}

func TestValidateCardSet_NilCardSet(t *testing.T) {
	err := ValidateCardSet(nil)
	if err == nil {
		t.Error("Expected error for nil card set, got nil")
	}

	expectedMsg := "card set is nil"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %v, want %v", err.Error(), expectedMsg)
	}
}

func TestValidateCardSet_EmptyCards(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{},
	}

	err := ValidateCardSet(cardSet)
	if err == nil {
		t.Error("Expected error for empty cards, got nil")
	}

	expectedMsg := "no cards found in card set"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %v, want %v", err.Error(), expectedMsg)
	}
}

func TestValidateCardSet_EmptyFront(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{Front: "valid", Meaning: "Valid card"},
			{Front: "", Meaning: "Invalid card with empty front"},
		},
	}

	err := ValidateCardSet(cardSet)
	if err == nil {
		t.Error("Expected error for empty front field, got nil")
	}

	expectedMsg := "card 2: front field is required"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %v, want %v", err.Error(), expectedMsg)
	}
}

func TestValidateCardSet_EmptyMeaning(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{Front: "valid", Meaning: "Valid card"},
			{Front: "invalid", Meaning: ""},
		},
	}

	err := ValidateCardSet(cardSet)
	if err == nil {
		t.Error("Expected error for empty meaning field, got nil")
	}

	expectedMsg := "card 2: meaning field is required"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %v, want %v", err.Error(), expectedMsg)
	}
}

func TestValidateCardSet_MultipleErrors(t *testing.T) {
	cardSet := &CardSet{
		Cards: []YamlCard{
			{Front: "", Meaning: "Empty front"},
		},
	}

	err := ValidateCardSet(cardSet)
	if err == nil {
		t.Error("Expected error for empty front field, got nil")
	}

	expectedMsg := "card 1: front field is required"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %v, want %v", err.Error(), expectedMsg)
	}
}
