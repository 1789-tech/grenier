#!/bin/sh
set -eu

repo_url="${GRENIER_REPO_URL:-https://github.com/1789-tech/grenier.git}"
dest="${CLAUDE_SKILLS_DIR:-}"

if [ -z "$dest" ]; then
  if [ -n "${CLAUDE_CONFIG_DIR:-}" ]; then
    dest="$CLAUDE_CONFIG_DIR/skills"
  elif [ -n "${XDG_CONFIG_HOME:-}" ] && [ -d "$XDG_CONFIG_HOME/claude" ]; then
    dest="$XDG_CONFIG_HOME/claude/skills"
  else
    dest="$HOME/.claude/skills"
  fi
fi

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cleanup=""

if [ -d "$script_dir/skills" ]; then
  src="$script_dir"
else
  tmp="${TMPDIR:-/tmp}/grenier-install.$$"
  cleanup="$tmp"
  rm -rf "$tmp"
  git clone --depth 1 "$repo_url" "$tmp" >/dev/null
  src="$tmp"
fi

trap 'if [ -n "$cleanup" ]; then rm -rf "$cleanup"; fi' EXIT INT HUP TERM

mkdir -p "$dest"

for skill in "$src"/skills/*; do
  [ -d "$skill" ] || continue
  name=$(basename "$skill")
  rm -rf "$dest/$name"
  cp -R "$skill" "$dest/$name"
done

printf 'Installed grenier skills to %s\n' "$dest"
