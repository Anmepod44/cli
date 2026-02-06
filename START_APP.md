# How to Start Your App

## Option 1: Run Directly (Quickest)

```bash
# Build the application
go build -o cli-assistant ./cmd/cli-assistant

# Run setup (first time only)
./setup.sh

# Start the app
./cli-assistant
```

## Option 2: Install to System Path

```bash
# Build
go build -o cli-assistant ./cmd/cli-assistant

# Install to /usr/local/bin (requires sudo)
sudo cp cli-assistant /usr/local/bin/

# Now you can run it from anywhere
cli-assistant
```

## Option 3: Create an Alias

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
alias clia='/path/to/your/cli-assistant'
```

Then reload:
```bash
source ~/.bashrc
```

Now run with:
```bash
clia
```

## First Time Setup

Before first use, run the setup script:

```bash
./setup.sh
```

This will:
- Create config directory
- Set up OpenAI API key (optional)
- Check clipboard support

## Verify Installation

```bash
# Check if installed
which cli-assistant

# Check version/help
cli-assistant
```

## Troubleshooting

**"command not found"**
- Make sure you built the app: `go build -o cli-assistant ./cmd/cli-assistant`
- Check the binary exists: `ls -la cli-assistant`
- If installed to system, check PATH: `echo $PATH`

**"permission denied"**
- Make executable: `chmod +x cli-assistant`

**App doesn't start**
- Check logs: `cat ~/.cli-assistant/errors.log`
- Run setup: `./setup.sh`
