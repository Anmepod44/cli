# Your Personal Setup Guide

## Quick Setup for OpenAI Integration

Follow these steps to get your CLI Command Assistant running with AI-powered command generation:

### Step 1: Run the Setup Script

The easiest way to configure everything:

```bash
./setup.sh
```

This interactive script will:
1. Create the configuration directory
2. Ask if you want to enable LLM mode
3. Prompt for your OpenAI API key
4. Check for clipboard tools
5. Create your configuration file

### Step 2: Get Your OpenAI API Key

If you don't have one yet:

1. Visit: https://platform.openai.com/api-keys
2. Sign in or create an account
3. Click "Create new secret key"
4. Name it "CLI Command Assistant"
5. Copy the key (starts with `sk-`)

**Important:** Save this key somewhere safe - you won't be able to see it again!

### Step 3: Add Your API Key

#### Option A: During Setup Script
When the setup script asks for your API key, paste it in.

#### Option B: Manual Configuration
Edit the config file:

```bash
nano ~/.cli-assistant/config.yaml
```

Add your key:

```yaml
usellm: true
openaiapikey: "sk-your-actual-key-here"
```

#### Option C: Environment Variable (Recommended for Security)
Add to your `~/.bashrc`:

```bash
echo 'export OPENAI_API_KEY="sk-your-actual-key-here"' >> ~/.bashrc
source ~/.bashrc
```

Then edit config:

```yaml
usellm: true
openaiapikey: ""  # Will use environment variable
```

### Step 4: Test It Out

```bash
./cli-assistant
```

You should see:
```
🤖 AI Mode: Using LLM-powered command generation
```

Try a complex command:
```
> show me all log files modified in the last week
```

### Troubleshooting

**"API error: Incorrect API key"**
- Check for typos in your key
- Make sure there are no extra spaces
- Verify the key starts with `sk-`

**"API error: You exceeded your current quota"**
- Add a payment method at: https://platform.openai.com/account/billing
- Set a usage limit (recommend $5-10/month)

**App uses pattern mode instead of LLM**
- Check `usellm: true` in config
- Verify API key is set correctly
- Check internet connection

### Cost Management

Set a monthly budget limit:
1. Go to: https://platform.openai.com/account/billing/limits
2. Set "Hard limit" to $10 (or your preference)
3. Enable email notifications

**Expected costs:**
- ~$0.001-0.002 per command
- 100 commands ≈ $0.10-0.20
- For typical daily use: $1-5/month

### Switching Modes

You can switch between pattern and LLM mode anytime:

**Enable LLM:**
```bash
nano ~/.cli-assistant/config.yaml
# Set: usellm: true
```

**Disable LLM (Pattern Mode):**
```bash
nano ~/.cli-assistant/config.yaml
# Set: usellm: false
```

Restart the app for changes to take effect.

### Security Tips

✅ **DO:**
- Use environment variables for shared systems
- Set usage limits in OpenAI dashboard
- Rotate keys periodically

❌ **DON'T:**
- Commit config files to git
- Share your API key
- Post keys in public forums

### Next Steps

1. Run `./setup.sh` to configure
2. Add your OpenAI API key
3. Start the application: `./cli-assistant`
4. Try some commands!
5. Check `~/.cli-assistant/errors.log` if issues occur

### Getting Help

- Full documentation: [README.md](README.md)
- Detailed OpenAI setup: [docs/OPENAI_SETUP.md](docs/OPENAI_SETUP.md)
- Quick start: [QUICKSTART.md](QUICKSTART.md)

### Example Commands to Try

**With LLM Mode:**
```
> show me all log files modified in the last week
> find the largest files in my home directory
> check if nginx is running and show its memory usage
> compress all pdf files in the current directory
> list all python files that contain the word "import"
```

**Pattern Mode (Works Offline):**
```
> list all files
> show disk usage
> find files larger than 10MB
> find process nginx
> check memory usage
```

---

**Ready to go!** Run `./setup.sh` to get started. 🚀
