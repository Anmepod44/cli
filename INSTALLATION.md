# Installation Guide

## Quick Start (No Installation)

```bash
# Build and run directly
go build -o cli-assistant ./cmd/cli-assistant
./setup.sh
./cli-assistant
```

---

## Method 1: Install to System Path (Recommended)

This makes the app available system-wide:

```bash
# Build the application
go build -o cli-assistant ./cmd/cli-assistant

# Install to /usr/local/bin
sudo cp cli-assistant /usr/local/bin/

# Run from anywhere
cli-assistant
```

**Uninstall:**
```bash
sudo rm /usr/local/bin/cli-assistant
```

---

## Method 2: Debian Package (.deb)

For Ubuntu, Debian, and derivatives:

### Build the Package

```bash
# Install build dependencies
sudo apt install dpkg-dev

# Build the .deb package
./build-deb.sh
```

This creates: `cli-command-assistant_1.0.0_amd64.deb`

### Install the Package

```bash
# Install
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb

# Fix dependencies if needed
sudo apt-get install -f

# Run the app
cli-assistant
```

### Uninstall

```bash
sudo apt remove cli-command-assistant
```

**Benefits:**
- ✅ Proper system integration
- ✅ Automatic dependency management
- ✅ Easy updates and removal
- ✅ Installs to standard locations

---

## Method 3: Snap Package

For any Linux distribution with snap support:

### Build the Snap

```bash
# Install snapcraft
sudo snap install snapcraft --classic

# Build the snap
snapcraft

# Install locally
sudo snap install cli-command-assistant_1.0.0_amd64.snap --dangerous
```

### Install from Snap Store (Future)

Once published to the Snap Store:

```bash
sudo snap install cli-command-assistant
```

### Uninstall

```bash
sudo snap remove cli-command-assistant
```

**Benefits:**
- ✅ Works on all Linux distributions
- ✅ Automatic updates
- ✅ Sandboxed security
- ✅ Easy distribution

---

## Method 4: Build from Source

For developers or custom installations:

```bash
# Clone the repository
git clone https://github.com/yourusername/cli-command-assistant
cd cli-command-assistant

# Build
go build -o cli-assistant ./cmd/cli-assistant

# Run setup
./setup.sh

# Run the app
./cli-assistant
```

---

## Post-Installation Setup

After installing by any method:

### 1. Run Setup Script

```bash
# If installed to system path
cli-assistant

# Or if running locally
./setup.sh
```

### 2. Configure OpenAI API (Optional)

For AI-powered mode:

```bash
# Edit configuration
nano ~/.cli-assistant/config.yaml

# Add your API key
usellm: true
openaiapikey: "sk-your-key-here"
```

See [OPENAI_SETUP.md](docs/OPENAI_SETUP.md) for detailed instructions.

### 3. Install Clipboard Support (Optional)

```bash
sudo apt install xclip
```

---

## Verification

Check if installation was successful:

```bash
# Check if command is available
which cli-assistant

# Run the app
cli-assistant

# Check version (in app, type 'help')
```

---

## File Locations

After installation, files are located at:

- **Binary**: `/usr/local/bin/cli-assistant` (or snap location)
- **Config**: `~/.cli-assistant/config.yaml`
- **History**: `~/.cli-assistant/history.json`
- **Logs**: `~/.cli-assistant/errors.log`
- **Docs**: `/usr/share/doc/cli-command-assistant/` (deb package)

---

## Troubleshooting

### "command not found"

**If installed to system:**
```bash
# Check if binary exists
ls -la /usr/local/bin/cli-assistant

# Check PATH
echo $PATH | grep /usr/local/bin
```

**If running locally:**
```bash
# Make sure you built it
go build -o cli-assistant ./cmd/cli-assistant

# Make it executable
chmod +x cli-assistant

# Run with ./
./cli-assistant
```

### "permission denied"

```bash
# Make executable
chmod +x cli-assistant

# Or reinstall with sudo
sudo cp cli-assistant /usr/local/bin/
```

### Build errors

```bash
# Update Go
go version  # Should be 1.22+

# Clean and rebuild
go clean
go mod tidy
go build -o cli-assistant ./cmd/cli-assistant
```

### Snap confinement issues

```bash
# Check snap connections
snap connections cli-command-assistant

# Connect required interfaces
sudo snap connect cli-command-assistant:home
sudo snap connect cli-command-assistant:network
```

---

## Distribution

### Share Your .deb Package

```bash
# Build the package
./build-deb.sh

# Share the file
cli-command-assistant_1.0.0_amd64.deb
```

Users can install with:
```bash
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
```

### Publish to Snap Store

1. Create a Snapcraft account: https://snapcraft.io/
2. Build the snap: `snapcraft`
3. Login: `snapcraft login`
4. Upload: `snapcraft upload cli-command-assistant_1.0.0_amd64.snap`
5. Release: `snapcraft release cli-command-assistant 1.0.0 stable`

### Create a PPA (Ubuntu)

For automatic updates via apt:

1. Create Launchpad account
2. Set up PPA
3. Upload source package
4. Users add PPA: `sudo add-apt-repository ppa:yourname/cli-assistant`

---

## Updates

### Manual Installation

```bash
# Rebuild and reinstall
go build -o cli-assistant ./cmd/cli-assistant
sudo cp cli-assistant /usr/local/bin/
```

### Debian Package

```bash
# Build new version
./build-deb.sh

# Install update
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
```

### Snap Package

```bash
# Automatic updates (if from Snap Store)
# Or manually:
sudo snap refresh cli-command-assistant
```

---

## Uninstallation

### System Installation

```bash
sudo rm /usr/local/bin/cli-assistant
```

### Debian Package

```bash
sudo apt remove cli-command-assistant
```

### Snap Package

```bash
sudo snap remove cli-command-assistant
```

### Remove User Data (Optional)

```bash
rm -rf ~/.cli-assistant
```

---

## Next Steps

- Read the [Quick Start Guide](QUICKSTART.md)
- Set up [OpenAI API](docs/OPENAI_SETUP.md)
- Check the [README](README.md) for full documentation
- Start using: `cli-assistant`

---

## Getting Help

- Check logs: `cat ~/.cli-assistant/errors.log`
- Run setup again: `./setup.sh`
- Open an issue on GitHub
- Read the documentation
