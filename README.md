# grenier 🏠📦

> A Claude companion for **decluttering** your home — as a collection of *Skills*.
> Sort it, decide what to do with it, sell or donate or recycle it. **English by
> default, and it knows your local marketplaces.** Open source.

`grenier` is a collection of **Claude Skills** that help you:

1. **Sort** — a method, exercises, steps and techniques to take on the clutter.
2. **Decide what to do with each thing** — sell (and where), donate, recycle, or
   trash — with the right effort/value trade-off.
3. **Prepare the sale** — from a photo or a line of text: identify the item,
   suggest a price, write the listing, bundle lots.
4. **Track the sale** — answer buyers, handle logistics (coming later).

We start small: pure *prompt* Skills (levels 1–2) that are useful from day one.
The agentic parts (photo vision, price scraping, tracking) come next.

## Skills

| Skill | Level | Status |
|-------|-------|--------|
| [`declutter`](skills/declutter/) | 1 · Sorting coach | ✅ v0 |
| [`sell`](skills/sell/) | 2 · What to do with it | ✅ v0 |
| [`listing`](skills/listing/) | 3 · Draft the ad | ✅ v0 |
| `sale-tracker` | 4 · Follow-up | 🔭 later |

The first three chain together: **sort → decide → draft the listing.** See
[`examples/`](examples/) for concrete end-to-end cases.

## Works anywhere

The Skills are written in **English** and country-agnostic at their core. The
`sell` and `listing` Skills load a small **local recipe** for your country —
which marketplaces, donation networks, recycling routes and listing conventions
apply where you live.

**France, US and UK today** — and adding yours is one file
(`countries/<iso2>.md`). If there's no recipe for your country yet, the Skill
falls back to sensible global guidance and tells you so.

## Install

These Skills follow the [Claude Agent Skills](https://docs.claude.com) format.

**Copy them into Claude Code:**

```bash
git clone https://github.com/1789-tech/grenier
cp -r grenier/skills/* ~/.claude/skills/
```

Restart Claude Code, then just go — no special command. Describe your situation
and the right Skill triggers itself.

**Or point Claude at the repo:** open Claude Code (or any Skills-compatible
harness) inside the cloned `grenier/` folder and the Skills in `skills/` are
available directly.

## Say something like…

No coding needed. If you can chat with Claude, you can use grenier. Try:

| You say… | Skill that triggers |
|----------|---------------------|
| "My garage is a total mess and I don't know where to start." | `declutter` |
| "Help me sort my bedroom, I've got 20 minutes." | `declutter` |
| "I've got an old bike, a box of novels and a broken printer to get rid of." | `sell` |
| "Where do I sell or donate this old sofa?" | `sell` |
| "I'm in the US — what do I do with a bunch of furniture I don't want?" | `sell` (US recipe) |
| "UK here — I've got dead electronics and e-waste to deal with." | `sell` (UK recipe) |
| "Write me the ad for my bike: Trek FX 2, size M, good condition." | `listing` |
| "Draft a listing for this coffee table." *(paste a photo)* | `listing` |
| "What do I even do with this single weird item?" | `sell` |
| "Turn my box of 40 books into the simplest possible outcome." | `sell` → `listing` |

## Contributing

The most useful contributions: **local recipes** (add your country) and improving
the existing Skills. See [`CONTRIBUTING.md`](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE). Contributions welcome, especially local recipes for
new countries and platforms.

---

*A product of [1789](https://1789.tech) · grenier.1789.tech*
