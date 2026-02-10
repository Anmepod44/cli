# Installation Guide

## Quick Install (Recommended)

```bash
# 1. Clone the repository
git clone https://github.com/yourusername/cli-command-assistant
cd cli-command-assistant

# 2. Run the setup wizard
./setup.sh
```

That's it! The setup wizard will guide you through:
- Dependency checks
- Configuration setup
- API key configuration (optional)
- Building the application
- System-wide installation (optional)
- Shell integration setup (optional)

## Manual Installation

### Prerequisites

- Go 1.22 or later
- Make
- xclip (optional, for clipboard support)

```bash
# Ubuntu/Debian
sudo apt install golang-go build-essential xclip

# Fedora/RHEL
sudo dnf install golang make xclip
```

### Build and Install

```bash
# Build
make build

# Install system-wide (optional)
sudo make install

# Or run locally
./cli-assistant
```

### Configuration

Create `~/.cli-assistant/config.yaml`:

```yaml
autoexecute: false
maxconcurrent: 5
historysize: 1000
confirmdangerous: true
clipboardenabled: true
usellm: false  # Set to true for AI mode
openaiapikey: "your-api-key-here"  # Optional
```

### Shell Integration (Optional)

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
source /path/to/cli-command-assistant/shell-integration.sh
```

Then reload your shell:

```bash
source ~/.bashrc  # or source ~/.zshrc
```

## Getting an OpenAI API Key (Optional)

1. Visit [platform.openai.com/api-keys](https://platform.openai.com/api-keys)
2. Create a new API key
3. Add it to `~/.cli-assistant/config.yaml`
4. Set `usellm: true` in the config

**Cost:** ~$0.001 per command (~$1-5/month for typical use)

## Verification

Test the installation:

```bash
# Interactive mode
cli-assistant

# Quick mode
echo "list all files" | cli-assistant --quick

# Shell integration (if enabled)
? show disk usage
```

## Troubleshooting

### "Go not found"
Install Go from [golang.org/doc/install](https://golang.org/doc/install)

### "Make not found"
```bash
# Ubuntu/Debian
sudo apt install build-essential

# Fedora/RHEL
sudo dnf install make
```

### "Clipboard not working"
```bash
sudo apt install xclip
```

### "Permission denied"
```bash
chmod +x setup.sh
chmod +x cli-assistant
```

## Uninstall

```bash
# Remove binary
sudo rm /usr/local/bin/cli-assistant

# Remove configuration
rm -rf ~/.cli-assistant

# Remove shell integration
# Edit ~/.bashrc or ~/.zshrc and remove the lines:
# # CLI Command Assistant Shell Integration
# source /path/to/shell-integration.sh
```

## Next Steps

- Read the [README.md](README.md) for usage examples
- Check `~/.cli-assistant/config.yaml` for customization options
- Try the shell integration with `? <your command>`

## Support

- Check `~/.cli-assistant/errors.log` for debugging
- Open an issue on GitHub for support
- Run `./setup.sh` again to reconfigure
