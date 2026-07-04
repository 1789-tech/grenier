# grenier 🏠📦

> A Claude companion for **decluttering** your home — as a collection of *Skills*.
> Sort it, decide what to do with it, sell or donate or recycle it. **English by
> default, and it knows your local marketplaces.** Open source.

`grenier` is a collection of **Claude Skills** that help you:

1. **Sort** — a method, exercises, steps and techniques to take on the clutter.
2. **Decide what to do with each thing** — sell (and where), donate, recycle, or
   trash — with the right effort/value trade-off.
3. **Price the decision** — estimate realistic resale value, bundle when useful,
   and decide whether selling is worth the hassle.
4. **Prepare the sale** — from a photo or a line of text: identify the item,
   suggest a price, write the listing, bundle lots.
5. **Track the sale** — answer buyers, handle logistics (coming later).

We start small: pure *prompt* Skills (levels 1–2) that are useful from day one.
The agentic parts (photo vision, price scraping, tracking) come next.

## Skills

| Skill | Level | Status |
|-------|-------|--------|
| [`declutter`](skills/declutter/) | 1 · Sorting coach | ✅ v0 |
| [`offload`](skills/offload/) | 2 · What to do with it | ✅ v0 |
| [`price-and-decide`](skills/price-and-decide/) | 2 · Price + one move | ✅ v0 |
| [`sell`](skills/sell/) | 3 · Draft the ad | ✅ v0 |
| `sale-tracker` | 4 · Follow-up | 🔭 later |

The first four chain together: **sort → decide route → price/one move → draft
the listing.** See [`examples/`](examples/) for concrete end-to-end cases.

## Works anywhere

The Skills are written in **English** and country-agnostic at their core. The
`offload` and `sell` Skills load a small **local recipe** for your country —
which marketplaces, donation networks, recycling routes and listing conventions
apply where you live.

**France, US and UK today** — and adding yours is one file
(`countries/<iso2>.md`). If there's no recipe for your country yet, the Skill
falls back to sensible global guidance and tells you so.

## Install

These Skills follow the [Claude Agent Skills](https://docs.claude.com) format.
Pick whichever door you like — they all end with the same Skills in
`~/.claude/skills/`.

### Door 1 — just ask your agent (no clone, no commands)

Paste this into [Claude Code](https://claude.com/claude-code) (or any
Skills-compatible agent):

```
Install the grenier skills from https://github.com/1789-tech/grenier
into ~/.claude/skills/, then help me start decluttering.
```

The agent fetches the `skills/` folder, drops each Skill into `~/.claude/skills/`,
and you're ready. This is the easiest path if you're already chatting with Claude.

### Door 2 — native plugin (Claude Code marketplace)

grenier is a Claude Code **plugin**. Add the marketplace once, then install:

```
/plugin marketplace add 1789-tech/grenier
/plugin install grenier@grenier
```

Claude Code handles updates and keeps the Skills in sync with the repo. This is
the recommended path for Claude Code users.

### Door 3 — one-line install script (`curl | sh`)

For terminals and CI, a small, inspectable, idempotent installer:

```bash
curl -fsSL https://raw.githubusercontent.com/1789-tech/grenier/main/install.sh | sh
```

It clones this repo (or downloads the tarball if you have no `git`) and copies the
Skills into `~/.claude/skills/`. Override the target with
`CLAUDE_SKILLS_DIR=/path sh install.sh`. Read [`install.sh`](install.sh) first —
it's ~40 lines and does nothing surprising.

**By hand, if you prefer:**

```bash
git clone https://github.com/1789-tech/grenier
cp -r grenier/skills/* ~/.claude/skills/
```

Or just open Claude Code inside the cloned `grenier/` folder — the Skills in
`skills/` are available directly.

## Say something like…

No coding needed. If you can chat with Claude, you can use grenier. Try:

| You say… | Skill that triggers |
|----------|---------------------|
| "My garage is a total mess and I don't know where to start." | `declutter` |
| "Help me sort my bedroom, I've got 20 minutes." | `declutter` |
| "I've got an old bike, a box of novels and a broken printer to get rid of." | `offload` |
| "Where do I sell or donate this old sofa?" | `offload` |
| "I'm in the US — what do I do with a bunch of furniture I don't want?" | `offload` (US recipe) |
| "UK here — I've got dead electronics and e-waste to deal with." | `offload` (UK recipe) |
| "Write me the ad for my bike: Trek FX 2, size M, good condition." | `sell` |
| "Draft a listing for this coffee table." *(paste a photo)* | `sell` |
| "What do I even do with this single weird item?" | `offload` |
| "Turn my box of 40 books into the simplest possible outcome." | `offload` → `price-and-decide` |
| "France: what price, and is this worth selling at all?" | `price-and-decide` |

## Contributing

The most useful contributions: **local recipes** (add your country) and improving
the existing Skills. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE). Contributions welcome, especially local recipes for
new countries and platforms.

---

*A product of [1789](https://1789.tech) · grenier.1789.tech*
