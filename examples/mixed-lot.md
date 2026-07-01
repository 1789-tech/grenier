# Example — a mixed lot to deal with

**Skill:** `offload` *(France recipe, for illustration)*

## What you say

> "I've pulled a bunch of stuff from sorting: an old bike that still rides, a box
> of 30 novels, a broken microwave, and a bag of decent clothes. What do I do
> with it all?"

## What Claude does (abbreviated)

`offload` detects the country (French context here → reads `countries/fr.md`) and
**groups by channel** instead of item by item:

- **Working bike** → *Leboncoin*, local pickup. Real value, sells fast locally.
  → draft the ad via `sell`.
- **30 novels** → *Momox / RecycLivre*: scan the ISBNs, they buy and pay postage.
  Zero friction for low-per-item volume.
- **Dead microwave** → *déchèterie / e-waste (DEEE)* or in-store take-back
  (1-for-1). Never the trash. Not sellable.
- **Bag of decent clothes** → *Vinted* if you want a few euros, otherwise *Geev /
  Le Relais bin* so it's gone today.

Rule applied: when in doubt sell vs. donate → **donate** (it moves fast).

> In the US this would route to Facebook Marketplace / Decluttr / Goodwill / Buy
> Nothing instead — same logic, local channels from `countries/us.md`.
