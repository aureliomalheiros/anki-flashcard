# Anki Flashcard Importer

A powerful, professional Go-based tool that imports flashcards from YAML files directly into Anki using the AnkiConnect API. Built for language learners, students, and anyone who wants to streamline their flashcard creation process.

## Key Features

- **YAML-based card definitions** - Human-readable, version-controllable card format
- **Advanced tagging system** - Automatic and custom tags for organization
- **Batch import** - Process multiple cards simultaneously
- **Connection testing** - Built-in AnkiConnect verification
- **Comprehensive testing** - Full test suite with 100% reliability
- **Smart build system** - Multiple build modes for different scenarios
- **Auto-deck creation** - Automatically creates target decks
- **Error handling** - Detailed feedback and duplicate detection
- **CLI interface** - Flexible command-line options
- **Progress tracking** - Real-time import status and statistics

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

## Build Script Usage Guide

The `build.sh` script is the main tool for building and testing the application. It provides multiple modes for different development and usage scenarios.

### Basic Usage

```bash
chmod +x build.sh

./build.sh

./build.sh --quick

./build.sh --verbose

./build.sh --skip-tests

./build.sh --help
```

### Build Script Options

| Option | Description | Use Case |
|--------|-------------|----------|
| `./build.sh` | **Full build** with tests and connection check | **Production builds, development** |
| `--quick` | Fast build, skip tests and connection check | **Rapid iteration, local testing** |
| `--skip-tests` | Build without tests, but check connection | **When tests fail but code works** |
| `--verbose` | Detailed output for debugging | **Troubleshooting build issues** |
| `--help` | Show all available options | **Learning the script capabilities** |

### What the Build Script Does

1. **Environment Check**: Verifies Go installation and version
2. **Testing Phase**: Runs comprehensive test suite (optional)
3. **Cleanup**: Removes previous build artifacts
4. **Compilation**: Builds optimized binary with `-ldflags="-s -w"`
5. **Permissions**: Makes binary executable automatically
6. **Connection Test**: Verifies AnkiConnect availability (optional)
7. **Documentation**: Shows usage examples and project info

## 📖 Quick Start Guide

### Step 1: Clone and Setup

```bash
git clone <repository-url>
cd anki-flashcard
```

### Step 2: Build the Application

```bash
./build.sh

./build.sh --quick
```

### Step 3: Test Connection

```bash
./anki-importer -test
```

### Step 4: Import Your Cards

```bash
./anki-importer -file=cards.yaml -deck=English-Vocabulary

./anki-importer -test

./anki-importer -file=my-cards.yaml -deck=Spanish-Vocabulary
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

## 🏷️ Advanced Tagging System

### Automatic Tags

The system automatically adds these tags to every card:

- **`yaml-import`** - Identifies cards imported from YAML
- **Language code** - Based on the `lang` field (e.g., `en`, `es`, `fr`)

### Custom Tag Categories

Organize your cards with meaningful tags:

**Grammar Categories:**
- `noun`, `verb`, `adjective`, `adverb`, `preposition`
- `singular`, `plural`, `countable`, `uncountable`

**Difficulty Levels:**
- `beginner`, `intermediate`, `advanced`, `expert`
- `common`, `uncommon`, `rare`

**Subject Areas:**
- `business`, `technology`, `science`, `medicine`
- `daily-life`, `academic`, `professional`

**Usage Context:**
- `formal`, `informal`, `colloquial`, `slang`
- `written`, `spoken`, `literary`

### Tag Examples

```yaml
tags:
  - noun
  - business
  - finance
  - advanced
  - formal

tags:
  - verb
  - daily-life
  - common
  - informal
```

## 💻 Command Line Usage

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

**Daily Vocabulary Addition:**

```bash
nano cards.yaml

./anki-importer -test

./anki-importer -file=cards.yaml -deck=Daily-English
```

**Multiple Subject Management:**

```bash
./anki-importer -file=business.yaml -deck=Business-English

./anki-importer -file=tech.yaml -deck=Programming

./anki-importer -file=medical.yaml -deck=Medical-Terminology
```

## 🛠️ Troubleshooting Guide

### Common Issues and Solutions

#### "Connection failed" Error

**Problem:** Cannot connect to AnkiConnect
**Solutions:**

1. Ensure Anki is running
2. Verify AnkiConnect addon is installed
3. Restart Anki after addon installation
4. Check if port 8765 is available

```bash
./anki-importer -test

curl http://localhost:8765 -d '{"action":"version","version":6}'
```

#### "Cannot create note because it is a duplicate"

**Problem:** Cards already exist in Anki
**Solution:** This is normal behavior when re-importing. Anki prevents duplicates automatically.

#### "Build failed" Error

**Problem:** Go compilation fails
**Solutions:**

1. Verify Go 1.22+ is installed: `go version`
2. Update dependencies: `go mod tidy`
3. Check for syntax errors in code

#### "Tests failed" Error

**Problem:** Test suite fails during build
**Solutions:**

1. Use `./build.sh --skip-tests` to build without tests
2. Check test files aren't corrupted
3. Ensure test data files exist

### Debug Mode

Use verbose build for detailed output:

```bash
./build.sh --verbose
```

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

### Batch Processing Script

Create a simple batch import script:

```bash
#!/bin/bash
# batch-import.sh

echo "🚀 Batch importing vocabulary files..."

files=(
    "basic-english.yaml:English::Basic"
    "intermediate-english.yaml:English::Intermediate"
    "advanced-english.yaml:English::Advanced"
    "business-terms.yaml:Business::General"
    "technical-terms.yaml:Programming::Vocabulary"
)

for entry in "${files[@]}"; do
    file="${entry%:*}"
    deck="${entry#*:}"
    
    if [ -f "$file" ]; then
        echo "📝 Importing $file to $deck..."
        ./anki-importer -file="$file" -deck="$deck"
    else
        echo "⚠️  File $file not found, skipping..."
    fi
done

echo "✅ Batch import completed!"
```

## 📊 Project Structure

```
anki-flashcard/
├── build.sh                 # 🔨 Smart build script with multiple modes
├── anki-importer            # 📦 Compiled binary (generated)
├── cards.yaml               # 📝 Your flashcard definitions
├── go.mod                   # 📋 Go module definition
├── go.sum                   # 🔒 Dependency checksums
├── README.md                # 📖 This documentation
│
├── cmd/
│   └── main.go              # 🚪 CLI application entry point
│
├── client/
│   ├── ankiconnect.go       # 🔌 AnkiConnect API client
│   └── ankiconnect_test.go  # 🧪 Client tests
│
├── models/
│   ├── card.go              # 🎴 Card data structures
│   ├── card_test.go         # 🧪 Card tests
│   ├── yaml.go              # 📄 YAML operations
│   └── yaml_test.go         # 🧪 YAML tests
│
├── config/
│   ├── config.go            # ⚙️  Configuration management
│   └── config_test.go       # 🧪 Config tests
│
└── testdata/                # 🧪 Test fixtures and sample files
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
# Use different host/port
export ANKICONNECT_HOST=192.168.1.100
export ANKICONNECT_PORT=8766
./anki-importer -test

# Reset to defaults
unset ANKICONNECT_HOST ANKICONNECT_PORT
```

## 🧪 Testing and Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./models -v
go test ./client -v
go test ./config -v

# Test coverage
go test -cover ./...
```

### Development Workflow

```bash
# 1. Make changes to code
vim models/card.go

# 2. Run tests to verify changes
go test ./models -v

# 3. Build and test full application
./build.sh

# 4. Test with sample data
./anki-importer -test
./anki-importer -file=cards.yaml -deck=Test-Deck
```

## 🤝 Contributing

We welcome contributions! Here's how to get started:

### Development Setup

```bash
# 1. Fork the repository on GitHub

# 2. Clone your fork
git clone https://github.com/your-username/anki-flashcard.git
cd anki-flashcard

# 3. Create a feature branch
git checkout -b feature/your-feature-name

# 4. Install dependencies
go mod tidy

# 5. Make your changes and add tests
# ... edit code ...

# 6. Run tests
./build.sh

# 7. Commit and push
git add .
git commit -m "feat: add your feature description"
git push origin feature/your-feature-name

# 8. Create a Pull Request on GitHub
```

### Contribution Guidelines

- **Code Quality**: Follow Go best practices and conventions
- **Testing**: Add tests for new functionality
- **Documentation**: Update README for new features
- **Commit Messages**: Use conventional commit format
- **Build**: Ensure `./build.sh` passes all checks

### Areas for Contribution

- 🐛 **Bug fixes** - Fix issues and improve reliability
- ✨ **New features** - Add new card formats or import options
- 📚 **Documentation** - Improve guides and examples
- 🧪 **Testing** - Add more test cases and coverage
- 🔧 **Build tools** - Improve build script and automation
- 🌍 **Localization** - Add support for more languages

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### MIT License Summary

- ✅ **Commercial use** - Use in commercial projects
- ✅ **Modification** - Modify and adapt the code  
- ✅ **Distribution** - Share and distribute freely
- ✅ **Private use** - Use for personal projects
- ❌ **Liability** - No warranty or liability
- ❌ **Trademark use** - Cannot use project trademarks

## 🙏 Acknowledgments

Special thanks to:

- **[AnkiConnect](https://foosoft.net/projects/anki-connect/)** - Excellent API for Anki integration
- **[Anki](https://apps.ankiweb.net/)** - The powerful spaced repetition system
- **[Go YAML](https://gopkg.in/yaml.v3)** - Robust YAML parsing library
- **Go Community** - For excellent tooling and libraries

## 📞 Support and Community

### Getting Help

- 📖 **Documentation**: Read this README thoroughly
- 🐛 **Issues**: [Create an issue](https://github.com/your-repo/anki-flashcard/issues) for bugs
- 💡 **Feature Requests**: Use GitHub issues with enhancement label
- 📧 **Questions**: Create a discussion or issue

### Quick Support Checklist

Before asking for help, please:

1. ✅ Read this documentation completely
2. ✅ Try the troubleshooting section
3. ✅ Test with `./anki-importer -test`
4. ✅ Check AnkiConnect is working: `curl http://localhost:8765`
5. ✅ Verify your YAML syntax is correct

---

<div align="center">

**Happy Learning! 🎓**

Made with ❤️ for language learners and knowledge enthusiasts.

[⭐ Star this project](https://github.com/your-repo/anki-flashcard) if it helps you!

</div>  
- **translate** (optional): Translation to your native language
- **pronounce** (optional): Pronunciation guide
- **lang** (optional): Language code for tagging

## Command Line Options

```bash
# Import cards from default cards.yaml to English-Vocabulary deck
go run main.go

# Test connection only
go run main.go -test

# Custom file and deck
go run main.go -file=vocabulary.yaml -deck=MyDeck

# Build and run
go build -o anki-importer main.go
./anki-importer -file=cards.yaml -deck=English
```

## Configuration

Environment variables (optional):

```bash
export ANKI_CONNECT_URL="http://localhost:8765"  # Default AnkiConnect URL
export ANKI_DEFAULT_DECK="Default"               # Fallback deck name
```

## Project Structure

```
anki-flashcard/
├── main.go              # Example usage and connection test
├── go.mod               # Go module definition
├── config/
│   └── config.go        # AnkiConnect client and configuration
├── models/
│   └── card.go          # Data structures for cards, notes, decks
└── README.md            # This file
```

## API Reference

### Configuration

- `LoadConfig()` - Load configuration from environment variables
- `NewClient(config)` - Create new AnkiConnect client

### Client Methods

- `TestConnection(ctx)` - Test if AnkiConnect is available
- `GetVersion(ctx)` - Get AnkiConnect version
- `GetDeckNames(ctx)` - List all deck names
- `CreateDeck(ctx, name)` - Create a new deck
- `AddNote(ctx, deck, model, fields, tags)` - Add a new note/card

## Error Handling

The client implements comprehensive error handling:
- Network connectivity issues
- AnkiConnect API errors  
- JSON parsing errors
- Context timeout handling
- HTTP status code validation

## Development

### Run Tests

```bash
go test ./...
```

### Build Binary

```bash
go build -o anki-flashcard main.go
```

## Contributing

1. Follow Go best practices
2. Include proper error handling
3. Add context support for operations
4. Update documentation for new features

## License

MIT License - See LICENSE file for details
