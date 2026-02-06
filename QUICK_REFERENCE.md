# Quick Reference Card

## 🚀 Getting Started (Choose One)

### Option 1: Quick Run (No Installation)
```bash
make build
./setup.sh
./cli-assistant
```

### Option 2: System Install
```bash
make install    # Installs to /usr/local/bin
cli-assistant   # Run from anywhere
```

### Option 3: Debian Package
```bash
make deb
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
cli-assistant
```

---

## 📋 Common Commands

### Building
```bash
make build      # Build the app
make clean      # Clean build files
make test       # Run tests
make help       # Show all commands
```

### Installation
```bash
make install    # Install system-wide
make uninstall  # Remove from system
make setup      # Run setup wizard
```

### Running
```bash
./cli-assistant              # Run locally
cli-assistant                # Run if installed
make run                     # Build and run
```

---

## ⚙️ Configuration

### Location
```bash
~/.cli-assistant/config.yaml
```

### Quick Edit
```bash
nano ~/.cli-assistant/config.yaml
```

### Enable LLM Mode
```yaml
usellm: true
openaiapikey: "sk-your-key-here"
```

### Disable LLM Mode
```yaml
usellm: false
```

---

## 🔑 OpenAI API Setup

### Get API Key
1. Visit: https://platform.openai.com/api-keys
2. Create new secret key
3. Copy the key (starts with `sk-`)

### Add to Config
```bash
nano ~/.cli-assistant/config.yaml
```

```yaml
usellm: true
openaiapikey: "sk-your-actual-key-here"
```

### Or Use Environment Variable
```bash
export OPENAI_API_KEY="sk-your-key-here"
```

---

## 💡 Example Commands

### Pattern Mode (Offline)
```
> list all files
> show disk usage
> find files larger than 10MB
> find process nginx
> check memory usage
```

### LLM Mode (AI-Powered)
```
> show me all log files modified in the last week
> find the largest files in my home directory
> check if nginx is running and show memory usage
> compress all pdf files in the current directory
```

---

## 🛠️ Troubleshooting

### App Won't Start
```bash
# Check if built
ls -la cli-assistant

# Rebuild
make clean && make build

# Check logs
cat ~/.cli-assistant/errors.log
```

### Command Not Found
```bash
# If installed
which cli-assistant

# Reinstall
make install
```

### API Errors
```bash
# Check config
cat ~/.cli-assistant/config.yaml

# Verify API key
echo $OPENAI_API_KEY

# Test with pattern mode
# Set usellm: false in config
```

---

## 📁 File Locations

```
Binary:  /usr/local/bin/cli-assistant
Config:  ~/.cli-assistant/config.yaml
History: ~/.cli-assistant/history.json
Logs:    ~/.cli-assistant/errors.log
```

---

## 🔄 Updates

### Manual
```bash
git pull
make clean
make build
make install
```

### Debian Package
```bash
make deb
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
```

---

## 🗑️ Uninstall

### Remove Binary
```bash
make uninstall
# or
sudo rm /usr/local/bin/cli-assistant
```

### Remove User Data (Optional)
```bash
rm -rf ~/.cli-assistant
```

---

## 📚 Documentation

- Full Guide: [README.md](README.md)
- Installation: [INSTALLATION.md](INSTALLATION.md)
- Quick Start: [QUICKSTART.md](QUICKSTART.md)
- OpenAI Setup: [docs/OPENAI_SETUP.md](docs/OPENAI_SETUP.md)
- Your Setup: [YOUR_SETUP_GUIDE.md](YOUR_SETUP_GUIDE.md)

---

## 🎯 In-App Commands

```
help or ?    - Show help
history      - View command history
exit or quit - Exit application
```

---

## 💰 Cost Estimate (LLM Mode)

- Per command: ~$0.001-0.002
- 100 commands: ~$0.10-0.20
- Monthly (typical): $1-5

Set budget at: https://platform.openai.com/account/billing/limits

---

## ⚡ Quick Tips

1. **First time?** Run `./setup.sh`
2. **Want AI?** Get OpenAI API key
3. **Offline?** Pattern mode works without internet
4. **Dangerous command?** App will warn you
5. **Need help?** Type `help` in the app

---

## 🆘 Get Help

- Logs: `cat ~/.cli-assistant/errors.log`
- Setup: `./setup.sh`
- Docs: Check markdown files in repo
- Issues: Open on GitHub
