#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'
SKIP_TESTS=false
VERBOSE=false
QUICK=false

while [[ $# -gt 0 ]]; do
    case $1 in
        --skip-tests)
            SKIP_TESTS=true
            shift
            ;;
        --verbose)
            VERBOSE=true
            shift
            ;;
        --quick)
            QUICK=true
            SKIP_TESTS=true
            shift
            ;;
        --help)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Options:"
            echo "  --skip-tests    Skip running tests before build"
            echo "  --verbose       Enable verbose output"
            echo "  --quick         Quick build (skip tests and connection check)"
            echo "  --help          Show this help message"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

echo -e "${BLUE}Anki Flashcard Importer - Build Script${NC}"
echo "============================================="

if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed or not in PATH${NC}"
    exit 1
fi

echo -e "${BLUE}Go version:${NC} $(go version)"
if [ "$SKIP_TESTS" = false ]; then
    echo ""
    echo -e "${BLUE}Running tests...${NC}"
    if [ "$VERBOSE" = true ]; then
        go test -v ./...
    else
        go test ./...
    fi
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ All tests passed!${NC}"
    else
        echo -e "${RED}Tests failed. Fix tests before building.${NC}"
        exit 1
    fi
else
    echo -e "${YELLOW}Skipping tests...${NC}"
fi

echo ""
echo -e "${BLUE}Cleaning previous build...${NC}"
rm -f anki-importer

echo ""
echo -e "${BLUE}Building anki-flashcard importer...${NC}"
if go build -ldflags="-s -w" -o anki-importer cmd/main.go; then
    echo -e "${GREEN}Build successful!${NC}"
    
    chmod +x anki-importer
    
    echo ""
    echo -e "${BLUE}Binary info:${NC}"
    ls -lh anki-importer
else
    echo -e "${RED}Build failed${NC}"
    exit 1
fi

if [ "$QUICK" = false ]; then
    echo ""
    echo -e "${BLUE}Testing AnkiConnect connection...${NC}"
    if ./anki-importer -test > /dev/null 2>&1; then
        echo -e "${GREEN}✓ Connection test passed!${NC}"
    else
        echo -e "${YELLOW}Connection test failed. Make sure:${NC}"
        echo "   1. Anki is running"
        echo "   2. AnkiConnect addon is installed"
        echo "   3. AnkiConnect is enabled in Anki"
    fi
else
    echo -e "${YELLOW}Skipping connection test (quick mode)...${NC}"
fi

echo ""
echo -e "${GREEN}Build completed successfully!${NC}"
echo ""
echo -e "${BLUE}Usage Examples:${NC}"
echo "  Basic import:     ./anki-importer -file=cards.yaml -deck=English-Vocabulary"
echo "  Test connection:  ./anki-importer -test"
echo ""
echo -e "${BLUE}️Tags Support:${NC}"
echo "  Add 'tags:' array in YAML for custom tags"
echo "  Automatic tags: yaml-import + language"
echo ""
echo -e "${BLUE}Project Structure:${NC}"
echo "  cards.yaml        - Your flashcard definitions"
echo "  anki-importer     - Compiled binary"
echo "  cmd/              - Main application"
echo "  client/           - AnkiConnect API client"
echo "  models/           - Data structures"
echo "  config/           - Configuration management"
