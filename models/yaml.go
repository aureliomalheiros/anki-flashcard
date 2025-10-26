package models

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadCardsFromYAML(filePath string) (*CardSet, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var cardSet CardSet
	if err := yaml.Unmarshal(data, &cardSet); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &cardSet, nil
}

func SaveCardsToYAML(cardSet *CardSet, filePath string) error {
	data, err := yaml.Marshal(cardSet)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write YAML file: %w", err)
	}

	return nil
}

func ValidateCardSet(cardSet *CardSet) error {
	if cardSet == nil {
		return fmt.Errorf("card set is nil")
	}

	if len(cardSet.Cards) == 0 {
		return fmt.Errorf("no cards found in card set")
	}

	for i, card := range cardSet.Cards {
		if card.Front == "" {
			return fmt.Errorf("card %d: front field is required", i+1)
		}
		if card.Meaning == "" {
			return fmt.Errorf("card %d: meaning field is required", i+1)
		}
	}

	return nil
}
