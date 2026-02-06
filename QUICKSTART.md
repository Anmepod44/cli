# Quick Start Guide

## Installation & Setup

### 1. Build the Application

```bash
go build -o cli-assistant ./cmd/cli-assistant
```

### 2. Run Setup Script (Recommended)

The setup script will guide you through configuration:

```bash
./setup.sh
```

This will:
- Create the configuration directory (`~/.cli-assistant/`)
- Ask if you want to enable LLM mode
- Prompt for your OpenAI API key (if enabling LLM)
- Check for clipboard tools and offer to install them
- Create a configuration file with your preferences

### 3. Manual Setup (Alternative)

If you prefer manual setup:

```bash
# Create config directory
mkdir -p ~/.cli-assistant

# Copy example config
cp config.example.yaml ~/.cli-assistant/config.yaml

# Edit configuration
nano ~/.cli-assistant/config.yaml
```

## Getting Your OpenAI API Key

To use LLM-powered mode:

1. Go to [OpenAI Platform](https://platform.openai.com/api-keys)
2. Sign in or create an account
3. Click "Create new secret key"
4. Copy the key (starts with `sk-`)
5. Add it to your configuration

### Option A: Configuration File

Edit `~/.cli-assistant/config.yaml`:

```yaml
usellm: true
openaiapikey: "sk-your-actual-api-key-here"
```

### Option B: Environment Variable

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
export OPENAI_API_KEY="sk-your-actual-api-key-here"
```

Then edit `~/.cli-assistant/config.yaml`:

```yaml
usellm: true
openaiapikey: ""  # Will use environment variable
```

## Running the Application

```bash
./cli-assistant
```

## First Commands to Try

### Pattern-Based Mode (Works Offline)

```
> list all files
> show disk usage
> find files larger than 10MB
> find process nginx
```

### LLM-Powered Mode (Better Understanding)

```
> show me all log files modified in the last week
> find the largest files in my home directory
> check if nginx is running and show its memory usage
> compress all pdf files in the current directory
```

## Switching Between Modes

You can switch modes anytime by editing `~/.cli-assistant/config.yaml`:

**Enable LLM Mode:**
```yaml
usellm: true
```

**Disable LLM Mode (Pattern-Based):**
```yaml
usellm: false
```

Changes take effect on next application start.

## Troubleshooting

### "API error" when using LLM mode

- Check your API key is correct
- Verify you have credits in your OpenAI account
- Check your internet connection
- The app will automatically fall back to pattern-based mode

### "Clipboard unavailable"

Install xclip:
```bash
sudo apt install xclip
```

### Commands not generating correctly

- Try rephrasing your request
- Use more specific language
- Check the examples in the help menu
- If using pattern mode, try enabling LLM mode for better understanding

## Configuration Tips

### For Daily Use (Pattern Mode)
```yaml
usellm: false           # Fast, offline, no API costs
autoexecute: false      # Review commands before running
confirmdangerous: true  # Extra safety for dangerous commands
```

### For Advanced Use (LLM Mode)
```yaml
usellm: true            # Better natural language understanding
autoexecute: false      # Still review AI-generated commands
confirmdangerous: true  # Always confirm dangerous operations
openaiapikey: "sk-..."  # Your API key
```

### For Power Users
```yaml
usellm: true
autoexecute: true       # Skip confirmations (use with caution!)
confirmdangerous: true  # But still confirm dangerous commands
maxconcurrent: 10       # Run more commands in parallel
```

## Cost Considerations (LLM Mode)

- Uses GPT-3.5-turbo by default (cost-effective)
- Typical cost: ~$0.001-0.002 per command generation
- Pattern-based mode is always free and offline
- LLM mode automatically falls back if API fails

## Next Steps

- Read the full [README.md](README.md) for detailed documentation
- Check [examples](#) for more command patterns
- Explore the configuration options
- Join the community for tips and tricks

## Getting Help

- Type `help` or `?` in the application
- Check command history with `history`
- View logs at `~/.cli-assistant/errors.log`
- Open an issue on GitHub for bugs or feature requests
