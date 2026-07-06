# Reference — draft the listing

Load this when the person has **decided to sell** an item (or a lot) and wants the
listing written for them. Your job: turn "I've got this thing to sell" into a
publishable listing in under a minute — an **honest, clear listing that sells
fast**, copy-paste ready. You don't optimize every unit of currency.

## What you need (and how to fill the gaps)

You can work with very little. Ask at most **one** follow-up if a blocking detail
is missing; otherwise make reasonable assumptions and **flag them**.

- **The item** — via a photo (read it: brand, model, visible condition) or a text
  description. If ambiguous, propose your best guess and ask for a one-line
  confirmation.
- **Condition** — new / like-new / good / fair / for parts. If unstated, infer
  from the photo or ask.
- **Price-driving info** — brand, model, size/dimensions, year, accessories/box,
  defects. Require none of it: list what's missing as "to confirm" in the listing.

## Localize

Title conventions, pricing norms, categories and buyer expectations differ by
market. Read the `countries/<iso2>.md` file for the user's country (its
**Listing** section) and draft to those conventions, in the local currency. No
country file? Use the generic conventions below (SEO-friendly title: brand + model
+ condition; honest description; local pickup or shipping), draft in the local
currency if you can tell it, and **say so**.

## Flow

1. **Identify.** Name the item as precisely as possible (category, brand, model).
   From a photo, extract what you see; don't invent.
2. **Estimate the price.** Give a **range** (low / suggested / high) in the local
   currency, based on condition and the local second-hand market. Ballpark, not
   real-time — say so. Advise starting slightly above target where haggling is the
   norm (see the country file). For the deeper "is it even worth selling"
   decision, that's `references/pricing.md`.
3. **Pick the platform.** One primary recommendation + one alternative, with the
   why in a word (see `references/disposition.md` for the full arbitrage).
4. **Write the listing.** Produce a **copyable** block:
   - **Title** — buyer keywords: brand + model + condition + key info. No poetry.
   - **Price** — the suggested one, in local currency.
   - **Description** — 4–8 lines: what it is, honest condition (defects included),
     useful dimensions/specs, reason for sale if it helps, terms (local pickup /
     shipping). Factual tone, no empty superlatives.
   - **Category / fields** suggested for the chosen platform.
5. **Photo tips** (2–3, concrete) — daylight, neutral background, show the defects,
   useful angles for that item type.
6. **Lots.** For several similar items, offer to group them into one "bundle"
   listing with a total price, or a reusable template to run item by item.

## Listing that sells fast

- Photo 1: whole object. Photo 2: brand/model/label. Photo 3: defect or wear.
  Photo 4: scale, dimensions, accessories or contents.
- Mention the visible defect before the buyer has to ask.
- Include one anti-stagnation line when culturally appropriate: "Fair price for a
  fast pickup; I'll drop it on Friday if needed."
- Handover script: propose a firm time window and area; no reservation without a
  pickup time.
- For marketplace automation, draft first. Never publish by accident; report
  which fields were guessed and which need human confirmation.

## Output format

Always return the listing in a clearly delimited block, ready to copy:

```
TITLE: <title>
PRICE: <xx> (range: <low>–<high>)
PLATFORM: <e.g. Facebook Marketplace / eBay / Vinted> — <reason in a word>
CATEGORY: <platform category>
DEADLINE: <drop price at D+3; exit route at D+7>

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
- **Privacy.** Remind them to blur/remove plates, addresses, faces, serial numbers
  and documents visible in photos.
- **No prohibited items.** Politely refuse to draft for things that can't legally
  be sold between individuals (counterfeits, regulated products, etc.).

## Where this hands off

Listing ready → the "answer buyers / manage handover" step is the sale-tracker
stage (coming later). Upstream, `references/disposition.md` decided *whether* to
sell and stage 1 (sorting, in `SKILL.md`) did the initial sorting.

## Examples

See [`examples/`](../../../examples/) at the repo root for concrete cases (bike,
box of books, furniture, clothing, electronics, and cross-country adaptation).
