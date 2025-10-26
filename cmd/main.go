package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aureliomalheiros/anki-flashcards/client"
	"github.com/aureliomalheiros/anki-flashcards/config"
	"github.com/aureliomalheiros/anki-flashcards/models"
)

func main() {
	fmt.Println("=== Anki Flashcards - YAML Importer ===")

	var (
		yamlFile = flag.String("file", "cards.yaml", "Path to YAML file with cards")
		deckName = flag.String("deck", "English-Vocabulary", "Target deck name")
		testOnly = flag.Bool("test", false, "Test connection only, don't import cards")
	)
	flag.Parse()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ankiClient := client.NewClient(cfg.AnkiConnectURL)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("Testing connection to AnkiConnect...")
	if err := ankiClient.TestConnection(ctx); err != nil {
		log.Fatalf("Failed to connect to AnkiConnect: %v\nMake sure Anki is running with AnkiConnect addon installed", err)
	}
	fmt.Println("Connection successful!")

	version, err := ankiClient.GetVersion(ctx)
	if err != nil {
		log.Printf("Warning: Could not get version: %v", err)
	} else {
		fmt.Printf("✓ AnkiConnect version: %d\n", version)
	}

	if *testOnly {
		fmt.Println("\n✓ Connection test completed successfully!")
		return
	}

	if _, err := os.Stat(*yamlFile); os.IsNotExist(err) {
		log.Fatalf("YAML file not found: %s", *yamlFile)
	}

	fmt.Printf("\nLoading cards from: %s\n", *yamlFile)
	cardSet, err := models.LoadCardsFromYAML(*yamlFile)
	if err != nil {
		log.Fatalf("Failed to load YAML file: %v", err)
	}

	if err := models.ValidateCardSet(cardSet); err != nil {
		log.Fatalf("Invalid card set: %v", err)
	}

	fmt.Printf("Loaded %d cards from YAML\n", len(cardSet.Cards))

	fmt.Printf("\nCreating/ensuring deck exists: %s\n", *deckName)
	if err := ankiClient.CreateDeck(ctx, *deckName); err != nil {
		log.Printf("Warning: Could not create deck (may already exist): %v", err)
	} else {
		fmt.Printf("Deck '%s' ready!\n", *deckName)
	}

	fmt.Printf("\nImporting %d cards to deck '%s'...\n", len(cardSet.Cards), *deckName)

	successCount := 0
	errorCount := 0

	for i, card := range cardSet.Cards {
		fmt.Printf("Importing card %d/%d: %s", i+1, len(cardSet.Cards), card.Front)

		if err := ankiClient.AddNoteFromYamlCard(ctx, &card, *deckName); err != nil {
			fmt.Printf("Error: %v\n", err)
			errorCount++
		} else {
			fmt.Printf("\n")
			successCount++
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("\n=== Import Summary ===")
	fmt.Printf("Total cards processed: %d\n", len(cardSet.Cards))
	fmt.Printf("Successfully imported: %d\n", successCount)
	fmt.Printf("Errors: %d\n", errorCount)
	fmt.Printf("Target deck: %s\n", *deckName)

	if errorCount > 0 {
		fmt.Println("\nNote: Some cards may have failed due to duplicates or other issues.")
		fmt.Println("This is normal if you're re-importing the same file.")
	}

	fmt.Println("\n✓ Import completed!")
	fmt.Println("\nUsage examples:")
	fmt.Printf("  %s -file=my-cards.yaml -deck=MyDeck\n", os.Args[0])
	fmt.Printf("  %s -test  (connection test only)\n", os.Args[0])
}
