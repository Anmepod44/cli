# 🔒 Security Guidelines

## Protecting Your API Keys

### ⚠️ NEVER Commit These Files:
- `~/.cli-assistant/config.yaml` (contains your API key)
- `.env` files
- Any file with API keys or secrets

### ✅ Safe Practices:

1. **Use Environment Variables**
   ```bash
   export OPENAI_API_KEY="sk-your-key-here"
   ```

2. **Check Before Committing**
   ```bash
   git status
   # Make sure config.yaml is NOT listed
   ```

3. **Use .env.example as Template**
   ```bash
   cp .env.example .env
   # Edit .env with your keys
   # .env is automatically ignored by git
   ```

4. **Verify .gitignore is Working**
   ```bash
   git check-ignore ~/.cli-assistant/config.yaml
   # Should show the file is ignored
   ```

## 🛡️ What's Protected

The `.gitignore` file automatically excludes:
- ✅ `config.yaml` files
- ✅ `.env` files
- ✅ `.cli-assistant/` directory
- ✅ `*.key` and `*_secret` files
- ✅ Log files
- ✅ Build artifacts

## 🚨 If You Accidentally Commit a Key

1. **Revoke the key immediately**
   - Go to https://platform.openai.com/api-keys
   - Delete the compromised key

2. **Generate a new key**
   - Create a new API key
   - Update your config

3. **Remove from git history** (if needed)
   ```bash
   git filter-branch --force --index-filter \
     "git rm --cached --ignore-unmatch config.yaml" \
     --prune-empty --tag-name-filter cat -- --all
   ```

4. **Force push** (use with caution)
   ```bash
   git push origin --force --all
   ```

## 📋 Pre-Commit Checklist

Before every commit:
- [ ] Check `git status` for sensitive files
- [ ] Verify no API keys in code
- [ ] Ensure .gitignore is working
- [ ] Review changes with `git diff`

## 🔐 Best Practices

1. **Rotate Keys Regularly** - Change API keys every 90 days
2. **Use Separate Keys** - Different keys for dev/prod
3. **Set Usage Limits** - Configure spending limits in OpenAI dashboard
4. **Monitor Usage** - Check API usage regularly
5. **Never Share Keys** - Don't share keys in chat, email, or forums

## 📞 Report Security Issues

If you find a security vulnerability, please email:
security@yourproject.com

**Do NOT open a public issue for security vulnerabilities.**

---

**Remember: Your API key is like a password. Keep it secret, keep it safe!** 🔒
