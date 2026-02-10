# Quick Start Guide

## For Complete Beginners

### Step 1: Download

```bash
git clone https://github.com/yourusername/cli-command-assistant
cd cli-command-assistant
```

### Step 2: Run Setup

```bash
./setup.sh
```

The wizard will ask you a few questions. Just press Enter to accept defaults!

### Step 3: Start Using It

```bash
cli-assistant
```

Then type things like:
- `list all files`
- `show disk usage`
- `find large files`

### That's It!

You're done! The app will generate Linux commands for you.

## Using the `?` Shortcut

If you enabled shell integration during setup, you can use `?` anywhere:

```bash
$ ? list all files
→ ls -la
Run this command? [Y/n/e(dit)]: 
```

Just type `?` followed by what you want to do!

## Examples

```bash
? find files larger than 10MB
? show memory usage
? search for text in files
? find process nginx
? copy file.txt to backup.txt
```

## Need AI Mode?

1. Get a free API key from [platform.openai.com/api-keys](https://platform.openai.com/api-keys)
2. Run `./setup.sh` again
3. Choose "yes" for AI mode
4. Paste your API key

## Help

- Type `help` or `?` in the app for commands
- Type `history` to see previous commands
- Type `exit` or `quit` to leave

## Troubleshooting

**"Permission denied"**
```bash
chmod +x setup.sh
./setup.sh
```

**"Go not found"**
Install Go first: [golang.org/doc/install](https://golang.org/doc/install)

**Need more help?**
Check [INSTALL.md](INSTALL.md) or [README.md](README.md)

---

**That's all you need to know to get started!** 🚀
