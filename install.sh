#!/bin/bash
# Minecraft Launcher Installer - Quick Install Script
# Usage: curl -fsSL https://raw.githubusercontent.com/mj41/mc-desktop/main/install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Minecraft Launcher Quick Installer ===${NC}"
echo

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        BINARY="mc-installer-amd64"
        ;;
    aarch64|arm64)
        BINARY="mc-installer-arm64"
        ;;
    *)
        echo -e "${RED}Error: Unsupported architecture: $ARCH${NC}"
        echo "Supported: x86_64, aarch64/arm64"
        exit 1
        ;;
esac

echo -e "Detected architecture: ${GREEN}$ARCH${NC}"
echo -e "Downloading: ${GREEN}$BINARY${NC}"
echo

# Download URL
DOWNLOAD_URL="https://github.com/mj41/mc-desktop/releases/latest/download/$BINARY"
TEMP_FILE="/tmp/$BINARY"

# Download the binary
if ! curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_FILE"; then
    echo -e "${RED}Error: Failed to download installer${NC}"
    echo "URL: $DOWNLOAD_URL"
    exit 1
fi

# Make it executable
chmod +x "$TEMP_FILE"

echo -e "${GREEN}Download complete!${NC}"
echo -e "${YELLOW}Running installer...${NC}"
echo

# Run the installer
"$TEMP_FILE" "$@"

# Clean up
rm -f "$TEMP_FILE"

echo
echo -e "${GREEN}Installation script complete!${NC}"
