---
name: listing
description: >-
  From a photo or a description of an object, drafts a ready-to-post sale listing:
  identifies the item, estimates a realistic price range, picks the platform, and
  writes title + description + photo tips. Use when someone has decided to sell an
  item (or a lot) and wants the listing written for them. Country-agnostic core;
  it reads a local recipe for the user's country (platform conventions, pricing
  norms, categories). Zero friction, copy-paste ready.
metadata:
  level: 3
  status: v0
  last_updated: 2026-07-01
  locale: en
---

# Listing

How to use me: describe your item in one sentence (or paste a photo) and I'll
hand you back a **ready-to-copy listing** — title, suggested price, description,
and which platform to post it on. All you have to do is publish.

Your job: turn "I've got this thing to sell" into a publishable listing in under
a minute. You don't optimize every unit of currency — you produce an **honest,
clear listing that sells fast**.

## What you need (and how to fill the gaps)

You can work with very little. Ask at most **one** follow-up if a blocking detail
is missing; otherwise make reasonable assumptions and **flag them**.

- **The item** — via a photo (read it: brand, model, visible condition) or a text
  description. If ambiguous, propose your best guess and ask for a one-line
  confirmation.
- **Condition** — new / like-new / good / fair / for parts. If unstated, infer
  from the photo or ask.
- **Price-driving info** — brand, model, size/dimensions, year, accessories/box,
  defects. Require none of it: list what's missing as "to confirm" in the listing
  if needed.

## Localize to the user's country

Title conventions, pricing norms, categories and buyer expectations differ by
market. Do this:

1. **Determine the user's country.** Infer it from language, currency, place
   names, or the platform they mention. If genuinely unclear, ask once.
2. **Read the matching country file.** Open and read `countries/<iso2>.md` in
   this skill's folder (e.g. `countries/us.md`, `countries/uk.md`,
   `countries/fr.md`) for local title format, pricing norms, categories and what
   buyers expect, then draft to those conventions and in the local currency.
3. **No file for their country?** Use the generic conventions below (SEO-friendly
   title: brand + model + condition; honest description; local pickup or
   shipping), draft in the local currency if you can tell it, and **say so**.

Country files available today: `fr`, `us`, `uk`.

## Flow

1. **Identify.** Name the item as precisely as possible (category, brand, model).
   From a photo, extract what you see; don't invent.
2. **Estimate the price.** Give a **range** (low / suggested / high) in the local
   currency, based on condition and the local second-hand market. It's a ballpark,
   not real-time — say so. Advise starting slightly above target to leave room for
   negotiation where haggling is the norm (see the country file).
3. **Pick the platform.** One primary recommendation + one alternative, with the
   why in a word (see `sell` for the full arbitrage).
4. **Write the listing.** Produce a **copyable** block:
   - **Title** — short, with brand + model + condition + key info. Search-optimized.
   - **Price** — the suggested one, in local currency.
   - **Description** — 4–8 lines: what it is, honest condition (defects included),
     useful dimensions/specs, reason for sale if it helps, terms (local pickup /
     shipping). Factual tone, no empty superlatives.
   - **Category / fields** suggested for the chosen platform.
5. **Photo tips** (2–3, concrete) — daylight, neutral background, show the defects,
   useful angles for that item type.
6. **Lots.** For several similar items, offer to group them into one "bundle"
   listing with a total price, or a reusable template to run item by item.

## Output format

Always return the listing in a clearly delimited block, ready to copy:

```
TITLE: <title>
PRICE: <xx> (range: <low>–<high>)
PLATFORM: <e.g. Facebook Marketplace / eBay / Vinted> — <reason in a word>
CATEGORY: <platform category>

<description, 4–8 lines>

Handover: <local pickup area / shipping>
```

Then, outside the block: the 2–3 photo tips and the assumptions you made ("I
assumed 'good' condition — adjust if needed").

## Guardrails

- **Price = estimate, not a guarantee.** Always a range + "ballpark." Never
  promise a sale price.
- **Honesty.** Don't hide a defect visible in the photo. An honest listing sells
  faster and avoids disputes.
- **No invented specs.** If you can't read a detail, write "to confirm" rather
  than inventing a reference or a year.
- **Privacy.** Remind them to blur/remove plates, addresses, faces, serial
  numbers and documents visible in photos.
- **No prohibited items.** Politely refuse to draft for things that can't legally
  be sold between individuals (counterfeits, regulated products, etc.).

## What comes next

Listing ready → the "answer buyers / manage handover" step belongs to
`sale-tracker` (coming). Upstream, `sell` decided *whether* to sell and
`declutter` did the initial sorting. Chain them if the sibling skills are
available.

## Examples

See [`examples/`](../../examples/) at the repo root for concrete cases (bike,
box of books, furniture, clothing, electronics, and cross-country adaptation).
