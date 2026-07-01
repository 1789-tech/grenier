# Contributing to grenier

Thanks for helping! grenier is a collection of **Claude Skills** for decluttering
and reselling. It's **multi-country: an English core + local recipes.** The most
useful contributions are **local recipes** (add your country) and improvements to
the existing Skills.

## Anatomy of a Skill

A Skill = a folder under `skills/<name>/` with a `SKILL.md` file, plus an optional
`countries/` folder for local recipes:

```
skills/<name>/
  SKILL.md            # required — the country-agnostic core, in English
  countries/          # optional — one file per country
    fr.md
    us.md
    uk.md
```

`SKILL.md` starts with YAML frontmatter:

```yaml
---
name: my-skill                  # short, lowercase, unique
description: >-                  # WHAT it does + WHEN to use it (Claude triggers on this)
  One or two sentences: what the Skill does AND in which situations to fire it.
  Keep it country-agnostic; mention it adapts to the user's country.
metadata:
  level: 2                      # 1 sort · 2 decide · 3 listing · 4 tracking
  status: v0                    # v0 / v1 / experimental
  last_updated: 2026-07-01      # last revision date
  locale: en                    # the core is English; country files carry locales
---
```

Right after the `#` title, add a **"How to use me"** line in plain, non-coder
language (what the person says to trigger the Skill).

## The country-recipe pattern

The Skill core stays country-agnostic. Local data (marketplaces, donation and
recycling routes, listing conventions, currency) lives in per-country files that
the Skill reads **on demand** — the native Claude Skills progressive-disclosure
pattern.

To add your country:

1. Create `skills/<skill>/countries/<iso2>.md` (ISO 3166-1 alpha-2, lowercase —
   `de`, `es`, `ca`, `au`…).
2. Follow the shape of the existing files (`fr.md`, `us.md`, `uk.md`): concise,
   scannable, tables/bullets, honest. Include the local currency and language.
3. For `sell`: where to sell / donate / recycle by item type.
   For `listing`: title format, pricing norms, categories, what buyers expect,
   and one example block.
4. The `SKILL.md` already tells Claude to read `countries/<iso2>.md` when it
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
