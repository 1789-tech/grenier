# Contributing to grenier

Thanks for helping! grenier is **one Claude Skill** for decluttering and
reselling. It's **multi-country: an English core + local recipes.** The most
useful contributions are **local recipes** (add your country) and improvements to
the Skill itself.

## Anatomy of the Skill

grenier is a single Skill with progressive disclosure — a lean entry point that
loads deeper knowledge only when the flow reaches that stage:

```
skills/grenier/
  SKILL.md                  # the coach/orchestrator: the whole journey + the sorting stage
  references/               # deeper stage logic, loaded on demand
    disposition.md          #   sell / donate / recycle / trash arbitrage
    pricing.md              #   realistic value + the one best move
    listing.md              #   draft the ad (title, price, description)
  countries/                # one merged local recipe per country
    fr.md
    us.md
    uk.md
```

`SKILL.md` starts with YAML frontmatter:

```yaml
---
name: grenier                   # the one skill; its description carries all triggers
description: >-                  # WHAT it does + WHEN to use it (Claude triggers on this)
  Cover the whole arc AND every situation that should fire it — sorting, disposal,
  "what's it worth", drafting a listing. Keep it country-agnostic; mention it
  adapts to the user's country.
metadata:
  status: v0                    # v0 / v1 / experimental
  last_updated: 2026-07-04      # last revision date
  locale: en                    # the core is English; country files carry locales
---
```

Right after the `#` title, add a **"How to use me"** line in plain, non-coder
language (what the person says to trigger the Skill).

**Adding a capability?** Default to a new `references/*.md` loaded on demand from
the SKILL.md map — *not* a new sibling skill. Fewer moving parts is the point; the
user should never have to pick the "right" skill first. Only argue for a separate
skill if it has a genuinely distinct trigger a user would invoke on its own.

## The country-recipe pattern

The Skill core stays country-agnostic. Local data (marketplaces, donation and
recycling routes, pricing norms, listing conventions, currency) lives in one
per-country file the Skill reads **on demand** — progressive disclosure.

To add your country:

1. Create `skills/grenier/countries/<iso2>.md` (ISO 3166-1 alpha-2, lowercase —
   `de`, `es`, `ca`, `au`…).
2. Follow the shape of the existing files (`fr.md`, `us.md`, `uk.md`): concise,
   scannable, tables/bullets, honest. Include the local currency and language.
   Each file carries three sections:
   - **Channels** (for `disposition.md`): where to sell / donate / recycle by item type.
   - **Pricing** (for `pricing.md`): priority mapping, channel heuristics, bundling.
   - **Listing** (for `listing.md`): title format, pricing norms, categories, what
     buyers expect, and one example block.
3. The `SKILL.md` already tells Claude to read `countries/<iso2>.md` when it
   detects that country — no core change needed.

## Style rules

- **English core + local recipes.** The Skill itself is English and works
  anywhere; country specifics go in `countries/`.
- **Anti-perfectionism.** The product's bias is to *get things moving*, not
  optimize every last unit of currency.
- **Honesty.** No fake price promises, no empty superlatives.
- **Privacy.** Never put real inventory, photos, prices or personal data in the
  repo — only prompt content and docs.
- **It must beat asking Claude raw.** Test that the Skill genuinely does better
  than asking Claude the same thing without it. If not, it isn't a product.

## Proposing a change

1. Fork + branch.
2. Add/edit the Skill or country recipe; add an example in `examples/` if it
   helps.
3. Bump `last_updated`.
4. Open a PR explaining the use case and the country/platform you're targeting.

Contributions are licensed MIT (see [LICENSE](LICENSE)).
