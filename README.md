# Anki Flashcard Importer

A professional Go-based tool that imports flashcards from YAML files directly into Anki using the AnkiConnect API.

## Features

- **YAML-based card format** - Human-readable, version-controllable
- **Advanced tagging system** - Automatic and custom tags
- **Batch import** - Process multiple cards at once
- **Smart build system** - Multiple build modes
- **Auto-deck creation** - Creates target decks automatically
- **Comprehensive testing** - Full test suite included
- **CLI interface** - Flexible command-line options

## Prerequisites

Before using this tool, ensure you have:

1. **Go 1.22+** - [Download Go](https://golang.org/dl/)
2. **Anki Desktop Application** - [Download Anki](https://apps.ankiweb.net/)
3. **AnkiConnect Addon** - [Install from AnkiWeb](https://ankiweb.net/shared/info/2055492159)

### AnkiConnect Setup
1. Open Anki
2. Go to Tools → Add-ons → Get Add-ons
3. Enter code: `2055492159`
4. Restart Anki
5. Ensure AnkiConnect is listening on `localhost:8765`

## Quick Start

### 1. Build the Application

```bash
# Make script executable
chmod +x build.sh

# Full build with tests
./build.sh

# Quick build (skip tests)
./build.sh --quick
```

### 2. Test Connection

```bash
./anki-importer -test
```

### 3. Import Cards

```bash
./anki-importer -file=cards.yaml -deck=English-Vocabulary
```

## 📝 YAML Card Format

Create your flashcards in YAML format for easy editing and version control.

### Basic Structure

```yaml
cards:
  - front: word_or_question
    meaning: >
      Detailed definition or explanation of the concept.
      Use the > symbol for multi-line text.
    examples:
      - Example sentence 1
      - Example sentence 2
      - Example sentence 3
    translate: Translation in your native language
    pronounce: phonetic-pronunciation
    lang: en
    tags:
      - category1
      - category2
      - difficulty_level
```

### Field Descriptions

| Field | Required | Description | Example |
|-------|----------|-------------|---------|
| `front` | ✅ **Yes** | The front side of the card (word/question) | `"entrepreneur"` |
| `meaning` | ✅ **Yes** | Definition or detailed explanation | `"A person who starts businesses..."` |
| `examples` | ❌ Optional | List of example sentences | `["He is an entrepreneur", "..."]` |
| `translate` | ❌ Optional | Translation in your native language | `"Empreendedor"` |
| `pronounce` | ❌ Optional | Phonetic pronunciation guide | `"ahn-truh-pruh-NUR"` |
| `lang` | ❌ Optional | Language code (adds automatic tag) | `"en"`, `"es"`, `"fr"` |
| `tags` | ❌ Optional | Custom tags for organization | `["noun", "business", "advanced"]` |

### Real Example

```yaml
cards:
  - front: entrepreneur
    meaning: >
      A person who starts and operates a business venture, taking on financial 
      risks in the hope of profit. Someone who identifies opportunities and 
      creates innovative solutions.
    examples:
      - The young entrepreneur launched three successful startups.
      - She became an entrepreneur after years in corporate finance.
      - Many entrepreneurs fail several times before achieving success.
    translate: Empreendedor, empresário
    pronounce: ahn-truh-pruh-NUR
    lang: en
    tags:
      - noun
      - business
      - innovation
      - leadership

  - front: resilient
    meaning: >
      Able to recover quickly from difficult conditions; having the ability 
      to bounce back from adversity or adapt to challenging situations.
    examples:
      - The resilient community rebuilt after the disaster.
      - Children are often more resilient than adults think.
    translate: Resiliente, resistente
    pronounce: ri-ZIL-yənt
    lang: en
    tags:
      - adjective
      - psychology
      - strength
```

## Tagging System

The system automatically adds:
- `yaml-import` - All imported cards
- Language code - From `lang` field (e.g., `en`, `es`)

You can add custom tags:
```yaml
tags:
  - noun
  - business
  - advanced
```

## Command Line Usage

### Basic Commands

```bash
./anki-importer -test

./anki-importer -file=cards.yaml -deck=English-Vocabulary

./anki-importer -file=business-vocab.yaml -deck=Business-English

./anki-importer -file=technical-terms.yaml -deck=Programming-Terms
```

### Command Line Options

| Option | Description | Example |
|--------|-------------|---------|
| `-test` | Test AnkiConnect connection only | `./anki-importer -test` |
| `-file=<filename>` | Specify YAML file to import | `-file=my-cards.yaml` |
| `-deck=<deckname>` | Target deck name in Anki | `-deck=Spanish-Vocabulary` |

### Workflow Examples

## Troubleshooting

### Connection Issues
- Ensure Anki is running
- Verify AnkiConnect addon is installed
- Test with: `./anki-importer -test`

### Build Issues
- Check Go version: `go version`
- Update dependencies: `go mod tidy`
- Use verbose mode: `./build.sh --verbose`

### Duplicates
- "Duplicate note" errors are normal when re-importing

## 🎯 Advanced Usage Examples

### Custom Deck Organization

```yaml
cards:
  - front: algorithm
    meaning: >
      A process or set of rules to be followed in calculations or 
      other problem-solving operations, especially by a computer.
    examples:
      - The sorting algorithm improved performance significantly.
      - Machine learning algorithms can detect patterns in data.
    translate: Algoritmo
    pronounce: AL-guh-rith-uhm
    lang: en
    tags:
      - noun
      - computer-science
      - programming
      - technical
      - advanced

  - front: deploy
    meaning: >
      To bring into effective action; to move software from development 
      to production environment.
    examples:
      - We will deploy the new version tomorrow.
      - The deployment process was automated.
    translate: Implementar, publicar
    pronounce: dih-PLOY
    lang: en
    tags:
      - verb
      - devops
      - programming
      - technical
      - common
```

### Language Learning Workflow

```bash
touch spanish-food.yaml spanish-travel.yaml spanish-business.yaml

./build.sh
./anki-importer -test

./anki-importer -file=spanish-food.yaml -deck=Spanish::Food
./anki-importer -file=spanish-travel.yaml -deck=Spanish::Travel  
./anki-importer -file=spanish-business.yaml -deck=Spanish::Business

```

## 📊 Project Structure

```bash
anki-flashcard/
├── build.sh                 #  Smart build script with multiple modes
├── anki-importer            #  Compiled binary (generated)
├── cards.yaml               #  Your flashcard definitions
├── go.mod                   #  Go module definition
├── go.sum                   #  Dependency checksums
├── README.md                #  This documentation
│
├── cmd/
│   └── main.go              #  CLI application entry point
│
├── client/
│   ├── ankiconnect.go       #  AnkiConnect API client
│   └── ankiconnect_test.go  #  Client tests
│
├── models/
│   ├── card.go              #  Card data structures
│   ├── card_test.go         #  Card tests
│   ├── yaml.go              #  YAML operations
│   └── yaml_test.go         #  YAML tests
│
├── config/
│   ├── config.go            #  Configuration management
│   └── config_test.go       #  Config tests
│
└── testdata/                #  Test fixtures and sample files
    ├── valid_cards.yaml
    ├── invalid_cards.yaml
    └── empty_cards.yaml
```

## 🔧 Configuration Options

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `ANKICONNECT_HOST` | `localhost` | AnkiConnect server host |
| `ANKICONNECT_PORT` | `8765` | AnkiConnect server port |

### Usage with Custom Configuration

```bash
export ANKICONNECT_HOST=192.168.1.100
export ANKICONNECT_PORT=8766
./anki-importer -test

unset ANKICONNECT_HOST ANKICONNECT_PORT
```
