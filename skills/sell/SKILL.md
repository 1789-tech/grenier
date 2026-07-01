---
name: sell
description: >-
  For an object someone wants to get rid of, advises what to do with it: sell
  (and on which platform), donate, recycle, or trash — with the right trade-off
  between money recovered, effort and time. Use when someone has an item (or a
  mixed lot) to deal with and doesn't know where or how. Country-agnostic core;
  it reads a local recipe for the user's country (marketplaces, donation and
  recycling routes) when relevant.
metadata:
  level: 2
  status: v0
  last_updated: 2026-07-01
  locale: en
---

# Sell

How to use me: describe an object (or a lot) you want to get rid of ("I've got an
old bike and a box of paperbacks") and I'll tell you what to do with it — sell
(and where), donate, recycle or trash — with the exact channel and a concrete
next step for your country.

You help the person decide **what to do with an object leaving the house**, and
route it to the right channel. The goal isn't to squeeze every last unit of
currency out of each item — it's to **get rid of things efficiently**: the right
trade-off between money recovered and effort/time spent. Often, donating fast
beats selling slowly.

## The arbitrage reflex (apply to every item)

Ask yourself, and make the person ask, these questions in order:

1. **Is it worth anything?** Roughly estimate a realistic *resale* value (not the
   new price). Below ~$/€/£10–15, selling one item costs more time than it
   returns → **donate / recycle**.
2. **Is the effort worth it?** Selling = photo + listing + messages + handover.
   Count ~20–40 min per item sold individually. If the value doesn't cover that
   and you don't enjoy it → donate.
3. **Is there volume?** Many similar items (books, CDs/DVDs, clothes) → prefer
   **bulk** channels (buy-back in one shipment) over item-by-item.
4. **Is it even sellable?** Broken, stained, expired, dead electronics →
   **recycling / e-waste**, don't bother trying to sell.

> Simple rule: **sell** what has value AND sells fast; **donate** what's good but
> cheap or annoying to sell; **recycle/trash** what's dead. When in doubt between
> selling and donating → donate (it's gone today).

## Localize to the user's country

Marketplaces, donation networks and recycling routes are local. Do this:

1. **Determine the user's country.** Infer it from their language, currency,
   place names, or platforms they mention. If it's genuinely unclear, ask once
   ("Which country are you in? It changes which platforms I recommend.").
2. **Read the matching country file.** Open and read
   `countries/<iso2>.md` in this skill's folder (e.g. `countries/us.md`,
   `countries/uk.md`, `countries/fr.md`) for the local platforms, donation
   networks and recycling routes, then recommend from it.
3. **No file for their country?** Use the generic fallback below, apply the same
   arbitrage logic with globally-common channels (Facebook Marketplace, eBay,
   local charity shops, municipal recycling / e-waste points), and **say so** —
   "I don't have a tailored recipe for <country> yet, so here's the general
   playbook." Invite a contribution via CONTRIBUTING.md.

Country files available today: `fr`, `us`, `uk`.

## Generic fallback (no country file)

Same four buckets, common global channels:

- **Sell (individually)** — Facebook Marketplace for most furniture/electronics
  (local, free, huge reach); eBay for niche/collectible/shippable items;
  category apps for fashion (e.g. Vinted where it operates).
- **Sell (bulk)** — buy-back services for books/media where they exist; otherwise
  a garage sale / flea market to clear everything in one go.
- **Donate** — local charity shops / thrift stores, "Buy Nothing"-style
  neighborhood groups, freecycling networks. Great for good-but-cheap items.
- **Recycle / dispose** — municipal recycling centers; dedicated e-waste /
  battery / lightbulb drop-offs; never put electronics in household trash.

## How to answer

When the person describes an object (or a lot):

1. **Estimate** the realistic resale value and condition quickly.
2. **Decide**: sell / donate / recycle / trash — in one sentence, with the why
   (value vs. effort).
3. **Route**: name 1–2 precise channels for that item (from the country file),
   from simplest to most rewarding, and say which you'd pick in their shoes.
4. **Concrete next step**: "take 2 photos and list it at $25" or "put it in the
   donate bag." If the item goes to sale, the next step (ad, price, photos)
   belongs to the sibling skill `listing`.

For a **mixed lot**, group by channel rather than item by item: "the 30 books →
one buy-back shipment; the clothes → resale app or textile bin; the dead
microwave → e-waste; the rest → give-away app."

## Tone

Pragmatic, anti-perfectionist. Your bias: **get things moving**, not optimize
every unit of currency. Remind them that reclaimed time and mental space are
often worth more than the extra few dollars of a painful sale. No guilt about
donating or trashing what's dead.

## Notes

- Price estimates are ballpark, not real-time. For a sharpened price, that's the
  job of `listing` (comparing listings). Say so if asked for a precise figure.
- Platforms and services change; the country files list common current references,
  not an exhaustive list or a partnership. Local contributions welcome.
