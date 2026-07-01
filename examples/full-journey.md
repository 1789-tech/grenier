# Example — the full journey, end to end

The 3 Skills chain together: **sort → decide → draft the listing.**

## 1. Sort — `declutter`

> "I'm tackling my closet, I've got 20 minutes."

Claude frames it: one shelf, 20 min on the clock, 4 destinations. Result: "18
pieces, 1 pile to sell (6), 1 bag to donate (9), 3 to recycle."

## 2. Decide — `offload`

> "In the sell pile: a leather jacket, 2 pairs of sneakers, a designer handbag."

Claude detects the country and arbitrates (UK recipe, `countries/uk.md`):
- Leather jacket + designer bag → *Vinted* (fashion, easy postage, audience).
- Recent sneakers → *Vinted* too; the worn ones → textile bank.
- The "donate" bag → *charity shop / Freegle* so it's gone this week.

## 3. Draft — `sell`

> "Write the ad for the leather jacket. Women's, size 10, black, lambskin,
> like-new, worn a few times."

```
TITLE: Black Lambskin Leather Jacket - Size 10 - Like New
PRICE: £70 (range: £50–£95)
PLATFORM: Vinted — women's fashion, audience + postage
...
```

In three messages you went from a jumbled closet to a publishable listing —
without ever typing a command. Prices are illustrative.
