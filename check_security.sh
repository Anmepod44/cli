#!/bin/bash

# Security check script - verifies sensitive files are ignored

echo "🔒 Security Check - Verifying Git Ignore Rules"
echo "================================================"
echo ""

# Check if .gitignore exists
if [ ! -f .gitignore ]; then
    echo "❌ .gitignore file not found!"
    exit 1
fi

echo "✓ .gitignore file exists"
echo ""

# Files that should be ignored
SENSITIVE_FILES=(
    "config.yaml"
    ".env"
    ".cli-assistant/config.yaml"
    "~/.cli-assistant/config.yaml"
)

echo "Checking if sensitive files are ignored:"
echo ""

ALL_IGNORED=true

for file in "${SENSITIVE_FILES[@]}"; do
    if git check-ignore -q "$file" 2>/dev/null; then
        echo "  ✓ $file - IGNORED (safe)"
    else
        echo "  ⚠ $file - NOT IGNORED (check .gitignore)"
        ALL_IGNORED=false
    fi
done

echo ""

# Check for accidentally staged sensitive files
echo "Checking for staged sensitive files:"
echo ""

STAGED_SENSITIVE=$(git diff --cached --name-only | grep -E '(config\.yaml|\.env|.*key|.*secret)' || true)

if [ -z "$STAGED_SENSITIVE" ]; then
    echo "  ✓ No sensitive files staged"
else
    echo "  ❌ WARNING: Sensitive files are staged:"
    echo "$STAGED_SENSITIVE"
    echo ""
    echo "  Run: git reset HEAD <file> to unstage"
    ALL_IGNORED=false
fi

echo ""

# Check for API keys in tracked files
echo "Scanning for API keys in tracked files:"
echo ""

API_KEY_PATTERN='sk-[a-zA-Z0-9]{20,}'
FOUND_KEYS=$(git grep -n "$API_KEY_PATTERN" 2>/dev/null || true)

if [ -z "$FOUND_KEYS" ]; then
    echo "  ✓ No API keys found in tracked files"
else
    echo "  ❌ WARNING: Possible API keys found:"
    echo "$FOUND_KEYS"
    echo ""
    echo "  URGENT: Remove these keys and revoke them!"
    ALL_IGNORED=false
fi

echo ""
echo "================================================"

if [ "$ALL_IGNORED" = true ]; then
    echo "✅ Security check passed!"
    echo ""
    echo "Your sensitive files are protected."
    exit 0
else
    echo "⚠️  Security issues found!"
    echo ""
    echo "Please review the warnings above."
    exit 1
fi
