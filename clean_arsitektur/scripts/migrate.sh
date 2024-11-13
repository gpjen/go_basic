#!/bin/bash
# scripts/migrate.sh

set -e

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Default values
ENV="development"
ACTION="status"

# Function to load environment variables
load_env() {
    if [ -f ".env.$ENV" ]; then
        export $(cat ".env.$ENV" | grep -v '#' | xargs)
    else
        export $(cat .env | grep -v '#' | xargs)
    fi
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -e|--env)
            ENV="$2"
            shift 2
            ;;
        -u|--up)
            ACTION="up"
            shift
            ;;
        -d|--down)
            ACTION="down"
            shift
            ;;
        -s|--status)
            ACTION="status"
            shift
            ;;
        -h|--help)
            echo "Usage: migrate.sh [options]"
            echo "Options:"
            echo "  -e, --env ENV    Set environment (development/staging/production)"
            echo "  -u, --up         Run pending migrations"
            echo "  -d, --down       Rollback last batch of migrations"
            echo "  -s, --status     Show migration status"
            echo "  -h, --help       Show this help message"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Load environment variables
load_env

echo -e "${YELLOW}Running migrations for ${ENV} environment${NC}"

# Run the application in migration mode
if go run cmd/app/main.go --migrate --action=$ACTION; then
    echo -e "${GREEN}Operation completed successfully!${NC}"
else
    echo -e "${RED}Operation failed!${NC}"
    exit 1
fi