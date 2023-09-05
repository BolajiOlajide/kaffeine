#!/bin/bash

# Variables
VERSION="0.0.1"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
if [ "$OS" = "darwin" ] || [ "$OS" = "linux" ]; then
    echo "Detected OS: $OS"
elif [ "$OS" = "windows" ]; then
    echo "Windows is not supported."
    exit 1
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Detect Architecture
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" = "aarch64" ]; then
    ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Detected Architecture: $ARCH"

# Create the asset pattern
ASSET_PATTERN="kaffeine_${OS}_${ARCH}"

# Fetch release assets for the given version using GitHub API
ASSET_URL=$(curl -s "https://api.github.com/repos/BolajiOlajide/kaffeine/releases/tags/$VERSION" | jq -r ".assets[] | select(.name | test(\"$ASSET_PATTERN\")) | .browser_download_url")

echo $ASSET_URL

# Download the asset
if [ -z "$ASSET_URL" ]; then
    echo "Asset with pattern $ASSET_PATTERN not found for version $VERSION"
    exit 1
fi

BINARY_NAME=$(basename "$ASSET_URL")
curl -L -O "$ASSET_URL"
chmod +x "$BINARY_NAME"

mv "$BINARY_NAME" /usr/local/bin/kaffeine
