# grenier 🏠📦

> A Claude companion for **decluttering** your home — **one Skill** that walks the
> whole arc. Sort it, decide what to do with it, sell or donate or recycle it.
> **English by default, and it knows your local marketplaces.** Open source.

`grenier` is **one Claude Skill** you invoke once ("help me declutter"). It walks
the whole journey, going deeper only when your situation reaches that stage:

1. **Sort** — a method, exercises, steps and techniques to take on the clutter.
2. **Decide what to do with each thing** — sell (and where), donate, recycle, or
   trash — with the right effort/value trade-off.
3. **Price the decision** — estimate realistic resale value, bundle when useful,
   and decide whether selling is worth the hassle.
4. **Prepare the sale** — from a photo or a line of text: identify the item,
   suggest a price, write the listing, bundle lots.
5. **Track the sale** — answer buyers, handle logistics (coming later).

One name, one trigger, one mental model. Under the hood it's a lean entry point
that **loads deeper knowledge only when the flow reaches it** — the native Claude
Skills [progressive-disclosure](https://docs.claude.com) pattern. No picking the
"right" sub-skill first.

## How it's built

| Piece | Role |
|-------|------|
| [`skills/grenier/SKILL.md`](skills/grenier/) | The coach/orchestrator — the whole journey + the sorting stage |
| [`references/method.md`](skills/grenier/references/method.md) | Signature method — net exit value, 20/20/20 cutoff, receipt contract |
| [`references/disposition.md`](skills/grenier/references/disposition.md) | Sell / donate / recycle / trash arbitrage — loaded when items head out |
| [`references/pricing.md`](skills/grenier/references/pricing.md) | Realistic value + the one best move — loaded on "what's it worth?" |
| [`references/listing.md`](skills/grenier/references/listing.md) | Draft the ad — loaded when it's time to sell |
| [`countries/<iso2>.md`](skills/grenier/countries/) | Local marketplaces, donation/recycling routes, pricing & listing conventions |
| `sale-tracker` | Follow-up (answer buyers, logistics) — 🔭 later |

The stages flow: **sort → decide route → price/one move → draft the listing.** See
[`examples/`](examples/) for concrete end-to-end cases.

## Works anywhere

The Skill is written in **English** and country-agnostic at its core. When the
flow reaches disposition, pricing or listing it loads a small **local recipe**
for your country (`countries/<iso2>.md`) — which marketplaces, donation networks,
recycling routes, pricing norms and listing conventions apply where you live.

**France, US and UK today** — and adding yours is one file
(`countries/<iso2>.md`). If there's no recipe for your country yet, the Skill
falls back to sensible global guidance and tells you so.

## Install

This Skill follows the [Claude Agent Skills](https://docs.claude.com) format.
Pick whichever door you like — they all end with the same Skill in
`~/.claude/skills/`.

### Door 1 — just ask your agent (no clone, no commands)

Paste this into [Claude Code](https://claude.com/claude-code) (or any
Skills-compatible agent):

```
Install the grenier skill from https://github.com/1789-tech/grenier
into ~/.claude/skills/, then help me start decluttering.
```

The agent fetches the `skills/` folder, drops the Skill into `~/.claude/skills/`,
and you're ready. This is the easiest path if you're already chatting with Claude.

### Door 2 — native plugin (Claude Code marketplace)

grenier is a Claude Code **plugin**. Add the marketplace once, then install:

```
/plugin marketplace add 1789-tech/grenier
/plugin install grenier@grenier
```

Claude Code handles updates and keeps the Skill in sync with the repo. This is
the recommended path for Claude Code users.

### Door 3 — one-line install script (`curl | sh`)

For terminals and CI, a small, inspectable, idempotent installer:

```bash
curl -fsSL https://raw.githubusercontent.com/1789-tech/grenier/main/install.sh | sh
```

It clones this repo (or downloads the tarball if you have no `git`) and copies the
Skill into `~/.claude/skills/`. If you've relocated Claude's config with
`CLAUDE_CONFIG_DIR`, the installer follows it (`$CLAUDE_CONFIG_DIR/skills`); or
override the target explicitly with `CLAUDE_SKILLS_DIR=/path sh install.sh`. Read
[`install.sh`](install.sh) first — it's ~40 lines and does nothing surprising.

**By hand, if you prefer:**

```bash
git clone https://github.com/1789-tech/grenier
cp -r grenier/skills/* ~/.claude/skills/
```

Or just open Claude Code inside the cloned `grenier/` folder — the Skill in
`skills/` is available directly.

## Say something like…

No coding needed. If you can chat with Claude, you can use grenier. It's one
Skill — the same trigger for all of these; it works out which stage you're at:

| You say… | Stage it walks into |
|----------|---------------------|
| "My garage is a total mess and I don't know where to start." | sort |
| "Help me sort my bedroom, I've got 20 minutes." | sort |
| "I've got an old bike, a box of novels and a broken printer to get rid of." | decide (disposition) |
| "Where do I sell or donate this old sofa?" | decide (disposition) |
| "I'm in the US — what do I do with a bunch of furniture I don't want?" | decide (US recipe) |
| "UK here — I've got dead electronics and e-waste to deal with." | decide (UK recipe) |
| "Write me the ad for my bike: Trek FX 2, size M, good condition." | draft the listing |
| "Draft a listing for this coffee table." *(paste a photo)* | draft the listing |
| "What do I even do with this single weird item?" | decide (disposition) |
| "Turn my box of 40 books into the simplest possible outcome." | decide → price |
| "France: what price, and is this worth selling at all?" | price (one move) |

## Contributing

The most useful contributions: **local recipes** (add your country) and improving
the Skill. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE). Contributions welcome, especially local recipes for
new countries and platforms.

---

*A product of [1789](https://1789.tech) · grenier.1789.tech*
