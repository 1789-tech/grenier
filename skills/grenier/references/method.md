# Reference — signature method & judgment rules

Load this when the person is blocked by a mixed pile, asks why you recommend
selling/donating/keeping, or needs the opinionated grenier method rather than a
plain channel list. This file is doctrine: use it to make sharper calls in
`disposition.md`, `pricing.md` and `listing.md`.

## Thesis

grenier maximizes **net exit value**, not resale price:

```text
net exit value = realistic cash - effort tax + exit value - regret risk
```

Make the arithmetic visible when it changes the answer. A 30-unit sale with 60
minutes of messages, photos and pickup can be worse than a zero-cash donation if
the user's priority is balanced or gone fast.

## The 20/20/20 cutoff is priority-aware

Default unit: **20 units of local currency**, **20 minutes of effort**, **20 days
expected time-to-exit**. Country files may override the money threshold.

This is the **balanced** default, not a rule that overrides the priority fork:

| Priority | Cutoff behavior |
|---|---|
| **max cash** | Loosen the cutoff. Individual sale can be worth it around 40 units, 45 minutes, or 30 days if the user knowingly wants cash and the item has real demand. Still refuse fake optimization for low-value ordinary items. |
| **balanced** | Use 20/20/20. If likely cash is under 20, effort over 20 minutes, or time-to-exit over 20 days, do not recommend standalone sale; bundle, donate, recycle or keep intentionally. |
| **gone fast** | Tighten the cutoff. Sell only if demand is obvious and the item can leave within about 7 days; otherwise donate, give away, bulk out or recycle. |

## Five object classes

| Class | Diagnostic | Default move |
|---|---|---|
| **Liquid** | Clear demand, easy to describe, easy handoff/shipping, enough value above the cutoff. | Sell individually on one channel with start / accept / deadline prices. |
| **Bulky local** | Some value, but transport and messages are the risk. | Local marketplace, aggressive price, short deadline, then donation/pickup route. |
| **Long tail** | Collector, rare model, niche buyer, or specialized proof needed. | Specialized/eBay-style channel only for max cash; otherwise keep, donate or sell honestly as unverified. |
| **Low-value volume** | Books, basics, CDs/DVDs, small goods, ordinary clothes. | Buy-back, bundle, vide-grenier, donation; never one listing per item. |
| **Disguised waste** | Broken, incomplete, unsafe, stained, expired, obsolete without demand. | Recycle, e-waste, textile bin, hazardous route or trash; "for parts" only for a clear niche. |

## Keep veto

Keeping is a correct decision if at least two are true:

- dated use in the next 90 days;
- replacement would be expensive or hard;
- sentimental value is concentrated and acknowledged;
- item is rare or hard to rebuy;
- storage is clean and already has an assigned place.

But a keep decision still needs an action: assigned home + review date. Without
those, it is only a "maybe" pile with a nicer name.

## Deadline policy

Every recommended sale needs an exit policy:

- **D+3**: no qualified message means drop 15-20%, improve the photos, or switch
  channel.
- **D+7**: if still unsold, move to donation, give-away, bundle, or gone-fast
  price.
- **D+14**: ordinary items have lost; remove them from the house.

## Receipt contract

Only produce a receipt when the underlying items or lots have been scored. If
you did not make item-level or lot-level judgments, omit totals and give a next
action instead.

Receipt numbers are mechanical:

- `cash recoverable = sum(accept price -> start price)` for items marked sell.
- `effort avoided = sum(estimated standalone sale time)` for items vetoed from
  micro-sale and routed to donation, recycling, buy-back or bundle.
- `exits scheduled = the dated sale/donation/recycling actions already
  recommended`.
- `first action now = the smallest physical action that starts the plan`.

Use ranges and labels. Do not imply precision from ballpark estimates.

```text
GRENIER RECEIPT
Decisions: 18 items
Cash recoverable: about 115-160 units (sell items only)
Effort avoided: about 3-5 h (micro-sales refused)
Scheduled exits: 2 donation bags Saturday, 1 e-waste drop-off, 1 local listing
Deadline: D+3 drop to 45; D+7 donate/give away if no qualified buyer
First move now: take 4 photos of the bike and list at 65
```

## Vinted draft-safety rule

When the task is filling marketplace drafts, the value is not advice; it is doing
the repetitive work safely. Prefer draft-first workflows:

- never publish when the user asked for drafts;
- make `save as draft` explicit before any final submit action;
- preserve a resumable item map after context interruptions;
- report count completed, count needing human review, and any fields guessed.

