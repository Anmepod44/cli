# CLI Command Assistant

> **Transform natural language into powerful Linux commands**

Stop memorizing complex command syntax. Just describe what you want to do in plain English, and let AI generate the perfect command for you.

---

## What is this?

**CLI Command Assistant** is your intelligent terminal companion that speaks your language. Whether you're a beginner learning Linux or a pro who can't remember every flag, this tool bridges the gap between what you want and how to do it.

```
You say: "show me all log files modified in the last week"
AI generates: find /var/log -type f -mtime -7
```

Simple as that.

---

## Why You'll Love It

| Feature | What It Means For You |
|---------|----------------------|
| **AI-Powered** | Understands complex requests like a human would |
| **Lightning Fast** | Pattern mode works offline, AI mode when you need it |
| **Safety First** | Warns you before running dangerous commands |
| **Learn as You Go** | See explanations for every command generated |
| **Beautiful Interface** | Clean, colorful, easy to read |
| **Smart History** | Never lose a useful command again |

---

## See It In Action

### Simple Commands
```bash
> list all files
→ ls -la

> show disk usage
→ df -h

> find process nginx
→ ps aux | grep "nginx"
```

### Complex Queries (AI Mode)
```bash
> show me all log files modified in the last week
→ find /var/log -type f -mtime -7

> find the largest files in my home directory
→ find ~ -type f -exec du -h {} + | sort -rh | head -20

> check if nginx is running and show memory usage
→ ps aux | grep nginx | grep -v grep && ps -p $(pgrep nginx) -o %mem,rss,cmd
```

---

## Quick Start

### Install in 30 Seconds

```bash
# 1. Clone and build
git clone https://github.com/yourusername/cli-command-assistant
cd cli-command-assistant
make build

# 2. Run it
./cli-assistant
```

That's it! The app works immediately with pattern-based mode.

### Want AI Superpowers?

Get an OpenAI API key (takes 2 minutes):

1. Visit [platform.openai.com/api-keys](https://platform.openai.com/api-keys)
2. Create a key (starts with `sk-`)
3. Run: `./setup.sh` and paste your key

**Cost:** ~$0.001 per command (~$1-5/month for typical use)

---

## How It Works

### Two Modes, One Goal

**Pattern Mode** (Default - Always Free)
- Works offline
- Instant responses
- Perfect for common commands
- No API needed

**AI Mode** (Optional - Smarter)
- Powered by ChatGPT
- Understands complex requests
- Learns from context
- Handles creative queries

The app automatically falls back to pattern mode if AI is unavailable. Best of both worlds!

---

## Features That Make Life Easy

### Safety Warnings
```
⚠ CRITICAL: Command will recursively delete files without confirmation
⚠ WARNING: Command requires elevated privileges (sudo)
ℹ INFO: Command may require sudo: apt
```

### Interactive Choices
```
What would you like to do?
  1. Execute
  2. Copy to clipboard
  3. Discard
```

### Smart History
```
> history
Recent Commands:
  1. find /var/log -type f -mtime -7 [executed]
     2024-02-06 15:30:45
  2. ls -la
     2024-02-06 15:28:12
```

---

## Installation Options

### Quick Run (No Installation)
```bash
make build
./cli-assistant
```

### System-Wide Install
```bash
make install
cli-assistant  # Run from anywhere
```

### Debian Package
```bash
make deb
sudo dpkg -i cli-command-assistant_1.0.0_amd64.deb
```

### Snap Package
```bash
make snap
sudo snap install cli-command-assistant_*.snap --dangerous
```

---

## Configuration

Edit `~/.cli-assistant/config.yaml`:

```yaml
# Enable AI mode (optional)
usellm: true
openaiapikey: "sk-your-key-here"

# Or use environment variable
# export OPENAI_API_KEY="sk-your-key-here"

# Other settings
autoexecute: false        # Confirm before running
confirmdangerous: true    # Extra safety for dangerous commands
maxconcurrent: 5          # Parallel command limit
historysize: 1000         # Command history size
clipboardenabled: true    # Copy to clipboard support
```

---

## Examples

### File Operations
```
> list all files
> find files named config.yaml
> find files larger than 10MB
> copy file.txt to backup.txt
> delete old logs
```

### System Information
```
> show disk usage
> check memory usage
> show system info
> check uptime
```

### Process Management
```
> list all processes
> find process nginx
> kill process 1234
```

### Network
```
> show network connections
> ping google.com
```

---

## Built With

- **Go 1.22** - Fast, concurrent, reliable
- **OpenAI GPT-3.5** - Natural language understanding
- **Pattern Matching** - Offline command generation

---

## Contributing

Found a bug? Have an idea? Contributions are welcome!

1. Fork the repo
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

---

## License

MIT License - Use it, modify it, share it!

---

## Show Your Support

If this tool makes your life easier, give it a star on GitHub!

---

## Need Help?

- Run `./setup.sh` for interactive setup
- Check `~/.cli-assistant/errors.log` for debugging
- Open an issue on GitHub for support

---

## Perfect For

- **Students** learning Linux
- **Developers** who forget syntax
- **System Admins** managing servers
- **DevOps Engineers** automating tasks
- **Linux Beginners** getting started
- **Terminal Users** saving time

---

<div align="center">

**Made with care for the command line**

[Get Started](#quick-start) • [Report Bug](https://github.com/yourusername/cli-command-assistant/issues)

</div>
