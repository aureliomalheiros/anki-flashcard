package models

import "time"

type YamlCard struct {
	Front     string   `yaml:"front"`
	Meaning   string   `yaml:"meaning"`
	Examples  []string `yaml:"examples"`
	Translate string   `yaml:"translate"`
	Pronounce string   `yaml:"pronounce"`
	Lang      string   `yaml:"lang"`
	Tags      []string `yaml:"tags,omitempty"`
}

type CardSet struct {
	Cards []YamlCard `yaml:"cards"`
}

type Card struct {
	ID       int64             `json:"id,omitempty"`
	Front    string            `json:"front"`
	Back     string            `json:"back"`
	DeckName string            `json:"deck_name"`
	Tags     []string          `json:"tags,omitempty"`
	Fields   map[string]string `json:"fields,omitempty"`
}

type Note struct {
	ID        int64             `json:"id,omitempty"`
	ModelName string            `json:"model_name"`
	Fields    map[string]string `json:"fields"`
	Tags      []string          `json:"tags,omitempty"`
	DeckName  string            `json:"deck_name"`
}

type Deck struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	CardCount     int    `json:"card_count,omitempty"`
	ReviewCount   int    `json:"review_count,omitempty"`
	NewCount      int    `json:"new_count,omitempty"`
	LearningCount int    `json:"learning_count,omitempty"`
}

type StudyStats struct {
	DeckName     string        `json:"deck_name"`
	CardsStudied int           `json:"cards_studied"`
	TimeSpent    time.Duration `json:"time_spent"`
	Accuracy     float64       `json:"accuracy"`
	LastStudied  time.Time     `json:"last_studied"`
}

type CreateNoteRequest struct {
	DeckName  string            `json:"deckName"`
	ModelName string            `json:"modelName"`
	Fields    map[string]string `json:"fields"`
	Tags      []string          `json:"tags,omitempty"`
}

type AddCardOptions struct {
	AllowDuplicate bool     `json:"allowDuplicate,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}

func (yc *YamlCard) ToAnkiFields() map[string]string {
	examples := ""
	for i, example := range yc.Examples {
		if i > 0 {
			examples += "<br>"
		}
		examples += "• " + example
	}

	back := yc.Meaning
	if yc.Translate != "" {
		back += "<br><br><b>Translation:</b> " + yc.Translate
	}
	if yc.Pronounce != "" {
		back += "<br><b>Pronunciation:</b> " + yc.Pronounce
	}
	if examples != "" {
		back += "<br><br><b>Examples:</b><br>" + examples
	}

	return map[string]string{
		"Front": yc.Front,
		"Back":  back,
	}
}

func (yc *YamlCard) GetTags() []string {
	tags := []string{"yaml-import"}

	if yc.Lang != "" {
		tags = append(tags, yc.Lang)
	}

	if len(yc.Tags) > 0 {
		tags = append(tags, yc.Tags...)
	}

	return tags
}
