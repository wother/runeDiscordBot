#!/bin/bash

# RuneDiscordBot v3.0.0 Setup Script
# This script checks/installs Go, prompts for configuration, and sets up the environment

set -e

echo "🔮 RuneDiscordBot v3.0.0 Setup Script"
echo "======================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Required Go version
REQUIRED_GO_VERSION="1.19"

# Function to check if Go is installed and get version
check_go_version() {
    if command -v go >/dev/null 2>&1; then
        GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
        echo -e "${GREEN}✓${NC} Go is installed: version ${GO_VERSION}"
        
        # Compare versions (simplified - works for major.minor)
        CURRENT_MAJOR=$(echo $GO_VERSION | cut -d. -f1)
        CURRENT_MINOR=$(echo $GO_VERSION | cut -d. -f2)
        REQUIRED_MAJOR=$(echo $REQUIRED_GO_VERSION | cut -d. -f1)
        REQUIRED_MINOR=$(echo $REQUIRED_GO_VERSION | cut -d. -f2)
        
        if [ "$CURRENT_MAJOR" -gt "$REQUIRED_MAJOR" ] || 
           ([ "$CURRENT_MAJOR" -eq "$REQUIRED_MAJOR" ] && [ "$CURRENT_MINOR" -ge "$REQUIRED_MINOR" ]); then
            echo -e "${GREEN}✓${NC} Go version meets requirements (>= ${REQUIRED_GO_VERSION})"
            return 0
        else
            echo -e "${YELLOW}⚠${NC} Go version ${GO_VERSION} is below required ${REQUIRED_GO_VERSION}"
            return 1
        fi
    else
        echo -e "${RED}✗${NC} Go is not installed"
        return 1
    fi
}

# Function to install Go on Linux
install_go_linux() {
    echo -e "${BLUE}📦${NC} Installing Go ${REQUIRED_GO_VERSION}..."
    
    # Detect architecture
    ARCH=$(uname -m)
    case $ARCH in
        x86_64) GOARCH="amd64" ;;
        aarch64) GOARCH="arm64" ;;
        armv7l) GOARCH="armv6l" ;;
        *) echo -e "${RED}✗${NC} Unsupported architecture: $ARCH"; exit 1 ;;
    esac
    
    # Download and install Go
    GO_TAR="go1.21.5.linux-${GOARCH}.tar.gz"
    GO_URL="https://golang.org/dl/${GO_TAR}"
    
    echo "Downloading Go from: $GO_URL"
    
    # Create temp directory
    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    
    # Download Go
    if command -v wget >/dev/null 2>&1; then
        wget "$GO_URL"
    elif command -v curl >/dev/null 2>&1; then
        curl -LO "$GO_URL"
    else
        echo -e "${RED}✗${NC} Neither wget nor curl found. Please install one of them."
        exit 1
    fi
    
    # Remove existing Go installation
    sudo rm -rf /usr/local/go
    
    # Extract and install
    sudo tar -C /usr/local -xzf "$GO_TAR"
    
    # Add to PATH if not already there
    if ! echo "$PATH" | grep -q "/usr/local/go/bin"; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc 2>/dev/null || true
        export PATH=$PATH:/usr/local/go/bin
    fi
    
    # Clean up
    cd - >/dev/null
    rm -rf "$TEMP_DIR"
    
    echo -e "${GREEN}✓${NC} Go installed successfully!"
    go version
}

# Function to prompt for Discord token
prompt_discord_token() {
    echo
    echo -e "${BLUE}🔑${NC} Discord Bot Configuration"
    echo "You need a Discord Bot Token from https://discord.com/developers/applications"
    echo
    while true; do
        read -p "Enter your Discord Bot Token: " DISCORD_TOKEN
        if [ -n "$DISCORD_TOKEN" ]; then
            break
        else
            echo -e "${YELLOW}⚠${NC} Token cannot be empty. Please try again."
        fi
    done
}

# Function to prompt for API configuration (optional)
prompt_api_config() {
    echo
    echo -e "${BLUE}🌐${NC} API Configuration (Optional)"
    echo "If you want to run the API server, you can specify a port."
    echo "Leave empty to use the default (8888) or if you only want to run the Discord bot."
    echo
    read -p "Enter API Port (optional, default 8888): " API_PORT
}

# Function to generate .env file
generate_env_file() {
    echo
    echo -e "${BLUE}📝${NC} Generating .env file..."
    
    cat > .env << EOF
# Discord Bot Configuration
DISCORD_TOKEN=${DISCORD_TOKEN}

# API Configuration (optional)
EOF

    if [ -n "$API_PORT" ]; then
        echo "PORT=${API_PORT}" >> .env
    else
        echo "# PORT=8888" >> .env
    fi
    
    cat >> .env << EOF

# Development/Testing
# TEST_BOT_TOKEN=your-test-token-here
EOF
    
    echo -e "${GREEN}✓${NC} .env file created successfully!"
    echo -e "${YELLOW}⚠${NC} Keep your .env file secure and never commit it to git!"
}

# Function to build the project
build_project() {
    echo
    echo -e "${BLUE}🔨${NC} Building project..."
    
    # Download dependencies
    echo "Downloading Go modules..."
    go mod download
    
    # Build API binary
    echo "Building API binary..."
    go build -o runeDiscordBot-api ./cmd/api
    
    # Build Discord bot binary
    echo "Building Discord bot binary..."
    go build -o runeDiscordBot-discord ./cmd/discord
    
    echo -e "${GREEN}✓${NC} Build completed successfully!"
    echo "Binaries created: runeDiscordBot-api, runeDiscordBot-discord"
}

# Main setup process
main() {
    echo
    echo -e "${BLUE}🔍${NC} Checking Go installation..."
    
    if ! check_go_version; then
        echo
        read -p "Would you like to install/upgrade Go? (y/N): " INSTALL_GO
        if [[ $INSTALL_GO =~ ^[Yy]$ ]]; then
            if [ "$(uname)" = "Linux" ]; then
                install_go_linux
            else
                echo -e "${YELLOW}⚠${NC} Automatic Go installation is only supported on Linux."
                echo "Please install Go manually from https://golang.org/dl/"
                exit 1
            fi
        else
            echo -e "${YELLOW}⚠${NC} Go installation required. Please install Go >= ${REQUIRED_GO_VERSION}"
            exit 1
        fi
    fi
    
    # Check if .env already exists
    if [ -f ".env" ]; then
        echo
        echo -e "${YELLOW}⚠${NC} .env file already exists."
        read -p "Do you want to overwrite it? (y/N): " OVERWRITE
        if [[ ! $OVERWRITE =~ ^[Yy]$ ]]; then
            echo "Skipping .env generation."
            build_project
            echo
            echo -e "${GREEN}🎉${NC} Setup completed! Your bot is ready to run."
            echo "Start with: ./runeDiscordBot-discord or ./runeDiscordBot-api"
            return
        fi
    fi
    
    # Configuration prompts
    prompt_discord_token
    prompt_api_config
    generate_env_file
    
    # Build the project
    build_project
    
    echo
    echo -e "${GREEN}🎉${NC} Setup completed successfully!"
    echo
    echo "Next steps:"
    echo "1. Review the .env file and adjust settings if needed"
    echo "2. Make sure the media/ folder is present with rune images"
    echo "3. Run the Discord bot: ./runeDiscordBot-discord"
    echo "4. Run the API server: ./runeDiscordBot-api"
    echo
    echo "For deployment, copy these files to your server:"
    echo "  - runeDiscordBot-api (the API binary)"
    echo "  - runeDiscordBot-discord (the Discord bot binary)"
    echo "  - .env (your configuration)"
    echo "  - media/ (the rune images folder)"
}

# Run main function
main "$@"
