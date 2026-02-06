# 🚀 How to Start Your CLI Command Assistant

## The Simplest Way (Right Now!)

```bash
# 1. Build the app
make build

# 2. Run setup (configure OpenAI API if you want)
./setup.sh

# 3. Start using it!
./cli-assistant
```

That's it! You're ready to go.

---

## Want to Install System-Wide?

```bash
# Install to /usr/local/bin (requires sudo)
make install

# Now run from anywhere
cli-assistant
```

---

## All Available Methods

### Method 1: Run Locally (Quickest)
```bash
make build
./cli-assistant
```

### Method 2: System Install
```bash
make install
cli-assistant
```

### Method 3: Debian Package
```bash
make deb
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
cli-assistant
```

### Method 4: Snap Package
```bash
make snap
sudo snap install cli-command-assistant_1.0.0_amd64.snap --dangerous
cli-command-assistant
```

---

## Setting Up OpenAI API (Optional)

### During Setup
```bash
./setup.sh
# Answer 'y' when asked about LLM mode
# Paste your API key when prompted
```

### Manual Setup
```bash
# Edit config
nano ~/.cli-assistant/config.yaml

# Add these lines:
usellm: true
openaiapikey: "sk-your-actual-key-here"
```

### Get Your API Key
1. Go to: https://platform.openai.com/api-keys
2. Click "Create new secret key"
3. Copy the key (starts with `sk-`)
4. Paste it in the config

---

## Makefile Commands

```bash
make build      # Build the application
make install    # Install system-wide
make uninstall  # Remove from system
make run        # Build and run
make setup      # Run setup wizard
make test       # Run tests
make clean      # Clean build files
make deb        # Build .deb package
make snap       # Build snap package
make help       # Show all commands
```

---

## First Time Using?

1. **Build it:**
   ```bash
   make build
   ```

2. **Run setup:**
   ```bash
   ./setup.sh
   ```
   
   This will:
   - Create config directory
   - Ask about LLM mode
   - Prompt for API key (optional)
   - Check clipboard support

3. **Start the app:**
   ```bash
   ./cli-assistant
   ```

4. **Try a command:**
   ```
   > list all files
   ```

---

## Troubleshooting

### "make: command not found"
```bash
# Use go directly
go build -o cli-assistant ./cmd/cli-assistant
./cli-assistant
```

### "permission denied"
```bash
chmod +x cli-assistant
./cli-assistant
```

### "command not found" (after install)
```bash
# Check if installed
which cli-assistant

# Reinstall
make install
```

### App won't start
```bash
# Check logs
cat ~/.cli-assistant/errors.log

# Run setup again
./setup.sh
```

---

## Quick Reference

### File Locations
- Binary: `./cli-assistant` or `/usr/local/bin/cli-assistant`
- Config: `~/.cli-assistant/config.yaml`
- History: `~/.cli-assistant/history.json`
- Logs: `~/.cli-assistant/errors.log`

### In-App Commands
- `help` or `?` - Show help
- `history` - View command history
- `exit` or `quit` - Exit app

### Example Commands
```
> list all files
> show disk usage
> find files larger than 10MB
> find process nginx
> check memory usage
```

---

## What's Next?

1. ✅ Build the app: `make build`
2. ✅ Run setup: `./setup.sh`
3. ✅ Start using: `./cli-assistant`
4. 📖 Read docs: [README.md](README.md)
5. 🔑 Setup OpenAI: [docs/OPENAI_SETUP.md](docs/OPENAI_SETUP.md)
6. 📚 Quick reference: [QUICK_REFERENCE.md](QUICK_REFERENCE.md)

---

## Need Help?

- **Quick Reference**: [QUICK_REFERENCE.md](QUICK_REFERENCE.md)
- **Full Installation Guide**: [INSTALLATION.md](INSTALLATION.md)
- **OpenAI Setup**: [docs/OPENAI_SETUP.md](docs/OPENAI_SETUP.md)
- **Quick Start**: [QUICKSTART.md](QUICKSTART.md)
- **Your Personal Guide**: [YOUR_SETUP_GUIDE.md](YOUR_SETUP_GUIDE.md)

---

## TL;DR

```bash
make build && ./setup.sh && ./cli-assistant
```

Done! 🎉
