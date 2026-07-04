# Reference — pricing & the one best move

Load this when the person asks "what's this worth?", "is it worth selling?",
"what price should I choose?", or wants the simplest good-deal decision for an
item or lot. This sits between disposition (*which route*) and listing (*write the
ad*): **for how much, and is selling worth it at all?**

Your job is not to show every possible resale path. It's to reduce mental load:
**one price, one move, maybe one alternative only if the trade-off is genuinely
worth seeing.**

## First fork: ask once, early

Before recommending anything, determine the person's priority:

1. **max cash** — willing to wait, take messages, ship, or split lots if it pays.
2. **balanced** — wants decent money without turning this into a project.
3. **gone fast** — wants the house cleared; money is a bonus.

If they didn't say, ask one short question:

> Do you want max cash, a balanced outcome, or the fastest way to get it gone?

If the person is already mid-flow and clearly stressed, default to **balanced**
and say so in one sentence.

## Localize

Pricing channels are local. Read the `countries/<iso2>.md` file for the user's
country (its **Pricing** section) and price from it. No country file? Use the
generic fallback below and say the advice isn't country-tailored.

## Generic pricing fallback

- Price from **sold or realistic used listings**, not new retail price.
- For common second-hand goods, start at roughly **25–50% of new price** in good
  condition, lower for bulky, dated, damaged, or hard-to-ship items.
- Below **10–15 units of local currency net**, single-item selling is usually not
  worth the photo + listing + message + handover loop.
- Similar low-value items should be **bundled**: books by genre/author, kids'
  clothes by size, kitchenware as a starter lot, small toys as a box.
- Broken electronics, stained textiles, incomplete toys, and unsafe items should
  usually go to **recycling / donation / trash**, not a listing.

## Effort score

Always price with effort in the equation. Mentally score each item:

- **Value** — likely net money after a realistic negotiation.
- **Friction** — photos, listing, questions, no-shows, shipping, transport.
- **Speed** — how likely it is to leave in 48 hours, a week, or not at all.
- **Volume** — whether a lot/bulk route beats item-by-item sales.

When effort beats value, the recommendation is allowed to be:

> Donate / recycle — not worth selling.

That is a successful answer, not a failure.

## Decision rules by priority

### Max cash
- Pick the channel and price with the best realistic net return.
- Accept slower sale time and more messages if the difference is material.
- Split lots only when the extra money is clearly worth the extra handling.
- Still refuse false optimization: don't suggest a 5-unit item as a standalone
  sale just because it might sell someday.

### Balanced
- Optimize for the best money-to-hassle ratio.
- Prefer one clean listing, one bundle, one local pickup, or one buy-back shipment
  over many small decisions.
- Price to sell in days, not to sit for weeks.

### Gone fast
- Prefer immediate routes: local pickup at a friendly price, buy-back services,
  donation, giving apps, recycling, or bulk clearance.
- Recommend selling only if the item has obvious demand and can leave quickly.
- Drop the price rather than spend mental energy chasing the top of the range.

## Output format

For each item or lot, return this compact structure:

```text
PRIORITY: <max cash | balanced | gone fast>
ESTIMATE: <realistic resale range> — <one-line reasoning / comp logic>
MOVE: <one clear recommendation with price/channel/action>
WHY: <one sentence: money vs effort vs speed>
OPTIONAL ALT: <only if genuinely useful; otherwise "none">
```

For multiple items, group into **lots / routes**, not a giant table of
micro-decisions. End with a one-line "do this first" next step.

## Hard guardrails

- **One recommended move.** Never give 3+ options. Collapse into one primary move
  plus at most one alternative.
- **Realistic, not aspirational.** Don't anchor on retail price or the highest
  unsold listing.
- **No fake live comps.** If you didn't actually browse live listings, say
  "ballpark" and explain the comp logic.
- **No router rebuild.** Don't enumerate every platform — use the country recipe
  to pick the one that best fits the item and priority.
- **No listing draft unless asked.** Hand off to `references/listing.md` for
  title, description, photos, and copy-paste listing.
- **Respect hassle.** A lower-cash recommendation can be correct when it removes
  decisions, messages, transport, or delays.

## Examples of good recommendations

- "Balanced: list the kids' bike at 45 euros on Leboncoin; accept 35–40 euros for
  same-week pickup. Alternative: 30 euros if you want it gone this weekend."
- "Gone fast: scan the 30 paperbacks into Momox/RecycLivre and donate rejects.
  Do not list them one by one."
- "Max cash: sell the art books individually at 18–35 euros each; bundle the
  ordinary paperbacks separately or use buy-back."
- "Donate: the 8 euro Ikea lamp isn't worth a standalone listing unless it's
  already photographed and pickup is effortless."
