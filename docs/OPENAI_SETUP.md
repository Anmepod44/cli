# OpenAI API Setup Guide

This guide will help you set up OpenAI API access for the CLI Command Assistant's LLM-powered mode.

## Why Use LLM Mode?

**Pattern-Based Mode (Default):**
- ✅ Fast and offline
- ✅ No API costs
- ✅ Works for common commands
- ❌ Limited to predefined patterns
- ❌ Can't handle complex or creative queries

**LLM-Powered Mode:**
- ✅ Understands complex natural language
- ✅ Handles creative and ambiguous requests
- ✅ Better context awareness
- ✅ Automatically falls back to pattern mode if needed
- ❌ Requires internet connection
- ❌ Small API costs (~$0.001-0.002 per command)

## Step 1: Create OpenAI Account

1. Go to [OpenAI Platform](https://platform.openai.com/)
2. Click "Sign up" or "Log in"
3. Complete the registration process
4. Verify your email address

## Step 2: Add Payment Method

OpenAI requires a payment method for API access:

1. Go to [Billing Settings](https://platform.openai.com/account/billing/overview)
2. Click "Add payment method"
3. Enter your credit card information
4. Set up billing limits (recommended: $5-10/month for personal use)

**Cost Estimates:**
- GPT-3.5-turbo: ~$0.001-0.002 per command
- 100 commands ≈ $0.10-0.20
- 1000 commands ≈ $1-2

## Step 3: Generate API Key

1. Go to [API Keys](https://platform.openai.com/api-keys)
2. Click "Create new secret key"
3. Give it a name (e.g., "CLI Command Assistant")
4. Click "Create secret key"
5. **IMPORTANT:** Copy the key immediately (starts with `sk-`)
   - You won't be able to see it again!
   - Store it securely

## Step 4: Configure CLI Command Assistant

### Option A: Using Setup Script (Recommended)

```bash
./setup.sh
```

The script will:
- Ask if you want to enable LLM mode
- Prompt for your API key
- Create the configuration file automatically

### Option B: Manual Configuration

1. Create or edit the config file:

```bash
mkdir -p ~/.cli-assistant
nano ~/.cli-assistant/config.yaml
```

2. Add your API key:

```yaml
usellm: true
openaiapikey: "sk-your-actual-api-key-here"
```

### Option C: Environment Variable

1. Add to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export OPENAI_API_KEY="sk-your-actual-api-key-here"
```

2. Reload your shell:

```bash
source ~/.bashrc  # or ~/.zshrc
```

3. Edit config to enable LLM mode:

```yaml
usellm: true
openaiapikey: ""  # Will use OPENAI_API_KEY env var
```

## Step 5: Test the Setup

1. Build and run the application:

```bash
go build -o cli-assistant ./cmd/cli-assistant
./cli-assistant
```

2. You should see:

```
🤖 AI Mode: Using LLM-powered command generation
```

3. Try a complex query:

```
> show me all log files modified in the last week
```

If it works, you'll see a generated command!

## Troubleshooting

### "API error: Incorrect API key provided"

- Double-check your API key
- Make sure there are no extra spaces
- Verify the key starts with `sk-`
- Try regenerating the key

### "API error: You exceeded your current quota"

- Check your [billing settings](https://platform.openai.com/account/billing/overview)
- Add a payment method
- Increase your usage limits
- Wait for your quota to reset (if on free tier)

### "API error: Rate limit exceeded"

- You're making too many requests too quickly
- Wait a few seconds and try again
- The app will automatically fall back to pattern mode

### Application falls back to pattern mode

This is normal behavior when:
- API key is not configured
- Internet connection is unavailable
- OpenAI API is experiencing issues
- Rate limits are exceeded

The app will continue working with pattern-based generation.

## Security Best Practices

### Protect Your API Key

- ✅ Never commit API keys to version control
- ✅ Use environment variables for shared systems
- ✅ Rotate keys periodically
- ✅ Set usage limits in OpenAI dashboard
- ❌ Don't share your API key
- ❌ Don't post it in public forums

### Add to .gitignore

If you're tracking this project in git:

```bash
echo "~/.cli-assistant/config.yaml" >> .gitignore
```

### Revoke Compromised Keys

If your key is exposed:

1. Go to [API Keys](https://platform.openai.com/api-keys)
2. Click the trash icon next to the compromised key
3. Generate a new key
4. Update your configuration

## Cost Management

### Set Usage Limits

1. Go to [Billing Settings](https://platform.openai.com/account/billing/overview)
2. Set a monthly budget limit
3. Enable email notifications for usage alerts

### Monitor Usage

1. Check [Usage Dashboard](https://platform.openai.com/account/usage)
2. Review costs daily/weekly
3. Adjust usage patterns if needed

### Optimize Costs

- Use pattern mode for simple commands
- Enable LLM mode only for complex queries
- Set reasonable usage limits
- Consider using GPT-3.5-turbo (default, cost-effective)

## Switching Between Modes

You can switch anytime by editing `~/.cli-assistant/config.yaml`:

**Enable LLM:**
```yaml
usellm: true
```

**Disable LLM (Pattern Mode):**
```yaml
usellm: false
```

Changes take effect on next application start.

## FAQ

**Q: Is my API key secure?**
A: Yes, it's stored locally in your home directory with restricted permissions. Never share the config file.

**Q: Can I use a different LLM provider?**
A: Currently only OpenAI is supported. Support for other providers (Anthropic, local models) is planned.

**Q: What if I run out of credits?**
A: The app will automatically fall back to pattern-based mode. Add more credits to resume LLM mode.

**Q: Can I use the free tier?**
A: OpenAI's free tier has limited quotas. For regular use, a paid account is recommended.

**Q: How much will this cost me?**
A: For typical personal use (10-50 commands/day), expect $1-5/month.

## Getting Help

- Check the [main README](../README.md)
- Review [Quick Start Guide](../QUICKSTART.md)
- Open an issue on GitHub
- Check OpenAI's [API documentation](https://platform.openai.com/docs)

## Next Steps

- Try complex natural language queries
- Experiment with different phrasings
- Compare LLM vs pattern mode results
- Share feedback and suggestions!
