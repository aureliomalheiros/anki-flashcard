package models

import (
	"reflect"
	"testing"
)

func TestYamlCard_ToAnkiFields(t *testing.T) {
	tests := []struct {
		name     string
		card     YamlCard
		expected map[string]string
	}{
		{
			name: "complete card with all fields",
			card: YamlCard{
				Front:     "hello",
				Meaning:   "A greeting used when meeting someone",
				Examples:  []string{"Hello, how are you?", "Hello world!"},
				Translate: "Olá",
				Pronounce: "hə-ˈlō",
				Lang:      "en",
			},
			expected: map[string]string{
				"Front": "hello",
				"Back":  "A greeting used when meeting someone<br><br><b>Translation:</b> Olá<br><b>Pronunciation:</b> hə-ˈlō<br><br><b>Examples:</b><br>• Hello, how are you?<br>• Hello world!",
			},
		},
		{
			name: "minimal card with only required fields",
			card: YamlCard{
				Front:   "test",
				Meaning: "A simple test case",
			},
			expected: map[string]string{
				"Front": "test",
				"Back":  "A simple test case",
			},
		},
		{
			name: "card with translation only",
			card: YamlCard{
				Front:     "word",
				Meaning:   "A unit of language",
				Translate: "Palavra",
			},
			expected: map[string]string{
				"Front": "word",
				"Back":  "A unit of language<br><br><b>Translation:</b> Palavra",
			},
		},
		{
			name: "card with pronunciation only",
			card: YamlCard{
				Front:     "pronunciation",
				Meaning:   "The way a word is said",
				Pronounce: "prə-ˌnən(t)-sē-ˈā-shən",
			},
			expected: map[string]string{
				"Front": "pronunciation",
				"Back":  "The way a word is said<br><b>Pronunciation:</b> prə-ˌnən(t)-sē-ˈā-shən",
			},
		},
		{
			name: "card with examples only",
			card: YamlCard{
				Front:    "example",
				Meaning:  "A representative form or pattern",
				Examples: []string{"This is an example", "Another example here"},
			},
			expected: map[string]string{
				"Front": "example",
				"Back":  "A representative form or pattern<br><br><b>Examples:</b><br>• This is an example<br>• Another example here",
			},
		},
		{
			name: "card with empty examples",
			card: YamlCard{
				Front:    "empty",
				Meaning:  "Having no content",
				Examples: []string{},
			},
			expected: map[string]string{
				"Front": "empty",
				"Back":  "Having no content",
			},
		},
		{
			name: "card with single example",
			card: YamlCard{
				Front:    "single",
				Meaning:  "Only one",
				Examples: []string{"This is a single example"},
			},
			expected: map[string]string{
				"Front": "single",
				"Back":  "Only one<br><br><b>Examples:</b><br>• This is a single example",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.card.ToAnkiFields()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToAnkiFields() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestYamlCard_GetTags(t *testing.T) {
	tests := []struct {
		name     string
		card     YamlCard
		expected []string
	}{
		{
			name: "card with language",
			card: YamlCard{
				Front: "hello",
				Lang:  "en",
			},
			expected: []string{"yaml-import", "en"},
		},
		{
			name: "card without language",
			card: YamlCard{
				Front: "hello",
			},
			expected: []string{"yaml-import"},
		},
		{
			name: "card with empty language",
			card: YamlCard{
				Front: "hello",
				Lang:  "",
			},
			expected: []string{"yaml-import"},
		},
		{
			name: "card with different language",
			card: YamlCard{
				Front: "bonjour",
				Lang:  "fr",
			},
			expected: []string{"yaml-import", "fr"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.card.GetTags()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GetTags() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCardSet_Structure(t *testing.T) {
	cardSet := CardSet{
		Cards: []YamlCard{
			{
				Front:   "test1",
				Meaning: "First test",
				Lang:    "en",
			},
			{
				Front:   "test2",
				Meaning: "Second test",
				Lang:    "en",
			},
		},
	}

	if len(cardSet.Cards) != 2 {
		t.Errorf("Expected 2 cards, got %d", len(cardSet.Cards))
	}

	if cardSet.Cards[0].Front != "test1" {
		t.Errorf("Expected first card front 'test1', got %s", cardSet.Cards[0].Front)
	}

	if cardSet.Cards[1].Front != "test2" {
		t.Errorf("Expected second card front 'test2', got %s", cardSet.Cards[1].Front)
	}
}

func TestYamlCard_EmptyFields(t *testing.T) {
	card := YamlCard{}

	fields := card.ToAnkiFields()
	if fields["Front"] != "" {
		t.Errorf("Expected empty front, got %s", fields["Front"])
	}

	if fields["Back"] != "" {
		t.Errorf("Expected empty back, got %s", fields["Back"])
	}

	tags := card.GetTags()
	expected := []string{"yaml-import"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Expected tags %v, got %v", expected, tags)
	}
}
