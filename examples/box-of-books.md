# Example — selling a box of books

**Skills:** `offload` (arbitrage) then `sell` (if selling individually)

## What you say

> "I've got a box of ~40 books, mostly paperback novels, a few nice art books.
> What's the simplest way to deal with them?"

## What Claude does (abbreviated)

`offload` decides by **volume and value** (US recipe, `countries/us.md`):

- **The ~35 paperbacks** → *Decluttr / Ziffit*: scan the barcodes in their app,
  they make a lump-sum offer and pay shipping. Low per-item (often $0.10–$2/book),
  but **zero friction** — ideal for volume.
- **The 4–5 nice art books** → real per-item value ($10–$40) → *eBay or Facebook
  Marketplace individually*. Here `sell` writes each ad:

```
TITLE: Hokusai Art Book - Hardcover - Like New
PRICE: $24 (range: $16–$35)
PLATFORM: eBay — collectible, shippable, national buyers
...
```

- **What Decluttr won't take** (too damaged) → paper recycling or a Little Free
  Library.

Bottom line: the bulk of the box ships in one go, the few valuable pieces sell
individually, nothing hits the trash. Prices are illustrative.
