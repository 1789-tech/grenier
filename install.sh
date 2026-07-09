#!/usr/bin/env sh
# grenier installer — copies the grenier Skill into your Claude skills dir.
#
# Inspect this script before you run it. It does exactly one thing: fetch this
# repo's `skills/` folder and copy the Skill into ~/.claude/skills/.
# No sudo, no network calls beyond the clone/tarball, no state left behind.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/1789-tech/grenier/main/install.sh | sh
#   # or, override the destination:
#   CLAUDE_SKILLS_DIR=/path/to/skills sh install.sh
#
# Destination resolution (first match wins):
#   1. $CLAUDE_SKILLS_DIR            — explicit override
#   2. $CLAUDE_CONFIG_DIR/skills     — follows a relocated Claude config dir
#   3. $HOME/.claude/skills          — default
# Note: Claude Code keys its config off CLAUDE_CONFIG_DIR, NOT XDG_CONFIG_HOME,
# so we deliberately do not read XDG here — that would install where Claude
# never looks for users who set XDG_CONFIG_HOME (most Linux desktops).
#
# Idempotent: re-running replaces each Skill cleanly (safe to upgrade).

set -eu

REPO="https://github.com/1789-tech/grenier"
DEST="${CLAUDE_SKILLS_DIR:-${CLAUDE_CONFIG_DIR:-$HOME/.claude}/skills}"
SKILLS="grenier"

printf 'grenier → installing skills into: %s\n' "$DEST"
mkdir -p "$DEST"

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT INT TERM

if command -v git >/dev/null 2>&1; then
  git clone --depth 1 "$REPO" "$tmp/grenier" >/dev/null 2>&1
  src="$tmp/grenier/skills"
else
  tarball="$REPO/archive/refs/heads/main.tar.gz"
  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$tarball" | tar -xz -C "$tmp"
  elif command -v wget >/dev/null 2>&1; then
    wget -qO- "$tarball" | tar -xz -C "$tmp"
  else
    echo "grenier: need git, curl or wget to fetch the skills" >&2
    exit 1
  fi
  src="$tmp/grenier-main/skills"
fi

for s in $SKILLS; do
  if [ ! -d "$src/$s" ]; then
    echo "grenier: expected skill '$s' not found in download — aborting" >&2
    exit 1
  fi
  rm -rf "$DEST/$s"
  cp -R "$src/$s" "$DEST/$s"
  printf '  ✓ %s\n' "$s"
done

cat <<'EOF'

done. restart Claude Code (or your Skills-compatible agent), then just describe
your situation — the right Skill triggers itself:

  "my garage is a mess and I don't know where to start"
  "where do I sell or donate this old sofa?"
  "what price should I list this bike at, or should I just donate it?"
  "write me the ad for my bike, Trek FX 2, size M"

happy decluttering. — grenier · https://grenier.1789.tech
EOF
