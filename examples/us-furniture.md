# Example — US user selling furniture (country adaptation)

**Stages:** disposition → listing · one Skill: `grenier` *(shows how the country recipe kicks in)*

## What you say

> "I'm in Seattle and I've got a couch, a coffee table and an old dresser I need
> gone before I move. What's the move?"

## What Claude does (abbreviated)

grenier reads the country from context (US) → opens `countries/us.md` and routes
with **US channels**, not French ones:

- **Couch** → *Facebook Marketplace*, local pickup, "OBO". Big, so price to move;
  if it doesn't sell in a week, *Buy Nothing group* to give it away fast.
- **Coffee table** → *Facebook Marketplace* or *Craigslist*. Solid resale value.
  → draft the ad in the listing stage.
- **Old dresser** → if decent, *Facebook Marketplace*; if rough,
  *Habitat ReStore* (they often pick up) or a *Buy Nothing* give-away.
- **If any of it is dead** → municipal bulk pickup, not the curb.

Then the listing stage (US recipe) drafts the coffee table ad in dollars:

```
TITLE: Solid Oak Coffee Table - 44x24 in - Great Condition
PRICE: $70 (range: $50–$95) OBO
PLATFORM: Facebook Marketplace — local, free, big reach
CATEGORY: Furniture > Tables
...
Handover: porch pickup, Seattle 98103
```

Same Skill, same logic as a French user — different marketplaces, currency and
conventions, pulled from the local recipe. Prices are illustrative.
