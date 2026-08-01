#!/usr/bin/env sh
set -eu

REPO="${1:-MrQwenty/fast-context-protocol}"
REMOTE="https://github.com/${REPO}.git"
BRANCH="${2:-$(git branch --show-current)}"

if git remote get-url origin >/dev/null 2>&1; then
  git remote set-url origin "$REMOTE"
else
  git remote add origin "$REMOTE"
fi

git push -u origin "$BRANCH"
echo "Published at https://github.com/$REPO"
