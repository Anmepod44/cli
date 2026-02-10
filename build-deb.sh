#!/bin/bash

# Build Debian package for CLI Command Assistant

set -e

VERSION="1.0.0"
ARCH="amd64"
PKG_NAME="cli-command-assistant"
BUILD_DIR="build/deb"

echo "Building Debian package..."

# Clean previous builds
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR/DEBIAN"
mkdir -p "$BUILD_DIR/usr/local/bin"
mkdir -p "$BUILD_DIR/usr/share/doc/$PKG_NAME"
mkdir -p "$BUILD_DIR/usr/share/$PKG_NAME"
mkdir -p "$BUILD_DIR/etc/profile.d"

# Build the application
echo "Compiling application..."
go build -o "$BUILD_DIR/usr/local/bin/cli-assistant" ./cmd/cli-assistant

# Copy shell integration to /etc/profile.d (loads for all users)
echo "Adding shell integration..."
cp shell-integration.sh "$BUILD_DIR/etc/profile.d/cli-assistant.sh"

# Create control file
cat > "$BUILD_DIR/DEBIAN/control" << EOF
Package: $PKG_NAME
Version: $VERSION
Section: utils
Priority: optional
Architecture: $ARCH
Maintainer: Your Name <your.email@example.com>
Description: CLI Command Assistant - Natural Language Linux Command Generator
 A lightweight yet powerful terminal application that translates natural
 language instructions into Linux shell commands. Features include AI-powered
 generation, command history, safety validation, and clipboard integration.
Depends: libc6 (>= 2.34)
Recommends: xclip | xsel
Homepage: https://github.com/yourusername/cli-command-assistant
EOF

# Create postinst script (runs after installation)
cat > "$BUILD_DIR/DEBIAN/postinst" << 'EOF'
#!/bin/bash
set -e

echo "Setting up CLI Command Assistant..."

# Make binary executable
chmod +x /usr/local/bin/cli-assistant

# Make shell integration executable
chmod +x /etc/profile.d/cli-assistant.sh

echo ""
echo "✓ CLI Command Assistant installed successfully!"
echo ""
echo "To get started:"
echo "  1. Reload your shell: source ~/.bashrc"
echo "  2. Use the ask command: ask list all files"
echo "  3. Or run interactively: cli-assistant"
echo ""
echo "Shell Integration:"
echo "  The 'ask' command is now available in your terminal!"
echo "  Examples:"
echo "    ask find large files"
echo "    ask show disk usage"
echo "    ask search for text in files"
echo ""
echo "For OpenAI API setup:"
echo "  Run: cli-assistant"
echo "  Or edit: ~/.cli-assistant/config.yaml"
echo ""

exit 0
EOF

chmod 755 "$BUILD_DIR/DEBIAN/postinst"

# Create prerm script (runs before removal)
cat > "$BUILD_DIR/DEBIAN/prerm" << 'EOF'
#!/bin/bash
set -e

echo "Removing CLI Command Assistant..."
echo "Note: Your configuration and history in ~/.cli-assistant will be preserved"

exit 0
EOF

chmod 755 "$BUILD_DIR/DEBIAN/prerm"

# Copy documentation
cp README.md "$BUILD_DIR/usr/share/doc/$PKG_NAME/"
if [ -f "QUICKSTART.md" ]; then
    cp QUICKSTART.md "$BUILD_DIR/usr/share/doc/$PKG_NAME/"
fi
if [ -f "docs/OPENAI_SETUP.md" ]; then
    cp docs/OPENAI_SETUP.md "$BUILD_DIR/usr/share/doc/$PKG_NAME/"
fi
cp config.example.yaml "$BUILD_DIR/usr/share/$PKG_NAME/"

# Create copyright file
cat > "$BUILD_DIR/usr/share/doc/$PKG_NAME/copyright" << EOF
Format: https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/
Upstream-Name: cli-command-assistant
Source: https://github.com/yourusername/cli-command-assistant

Files: *
Copyright: $(date +%Y) Your Name
License: MIT
 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:
 .
 The above copyright notice and this permission notice shall be included in all
 copies or substantial portions of the Software.
 .
 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 SOFTWARE.
EOF

# Build the package
echo "Creating .deb package..."
dpkg-deb --build "$BUILD_DIR" "${PKG_NAME}_${VERSION}_${ARCH}.deb"

echo ""
echo "✓ Debian package created: ${PKG_NAME}_${VERSION}_${ARCH}.deb"
echo ""
echo "To install:"
echo "  sudo dpkg -i ${PKG_NAME}_${VERSION}_${ARCH}.deb"
echo ""
echo "To uninstall:"
echo "  sudo apt remove $PKG_NAME"
echo ""
