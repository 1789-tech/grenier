# Price and Decide — France (fr)

Currency: € · Language: French when the user writes in French.

France bias: **Leboncoin for local value**, **Vinted for fashion/small shippable
goods**, **Momox / RecycLivre / Gibert for books/media bulk**, **Emmaüs / Geev /
ressourceries for low-value useful goods**, **déchèterie / DEEE for broken or
unsafe items**.

## Priority mapping

| Priority | Default move | Pricing stance |
|----------|--------------|----------------|
| **max cash** | Best-fit resale channel, often item-by-item | Price near the realistic high end; accept slower sale and negotiation. |
| **balanced** | One clean listing or one bulk route | Price in the middle-low of the range so it moves in days. |
| **gone fast** | Pickup, buy-back, donation, Geev, recycling | Price low or skip sale if hassle beats money. |

## France channel heuristics

| Item type | Best default | When to avoid |
|-----------|--------------|---------------|
| Furniture, appliances, bikes, tools, baby gear | **Leboncoin**, local pickup | Avoid if value < 15-20 € or transport/messages will be painful. |
| Clothes, shoes, bags, accessories | **Vinted**, preferably bundles by size/style | Avoid single low-value basics unless bundled. |
| Books | **Momox / RecycLivre / Gibert** for ordinary lots; **Leboncoin/eBay** for valuable art/rare books | Avoid listing ordinary paperbacks one by one. |
| CDs/DVDs/games/consoles | **Momox / Easy Cash / Cash Converters** for fast cash; **Leboncoin/eBay** for valuable items | Avoid low-value singles. |
| Decor, small household goods | **Leboncoin** only if visually attractive or bundled; otherwise **Geev/Emmaüs** | Avoid standalone listings under 15 €. |
| Broken electronics/appliances | **DEEE / déchèterie / store take-back** | Do not sell as working; "pour pièces" only for niche items. |

## Price heuristics

- **Leboncoin**: haggling is normal. If priority is max cash, list 10-15% above
  target. If balanced, list at the target. If gone fast, list 15-30% under the
  middle of the range or donate.
- **Vinted**: bundle low-value clothes; buyers send offers. A lot priced fairly
  beats ten 3 € decisions.
- **Books/media buy-back**: expect low per-item return, often cents to a few
  euros, but it removes nearly all listing work. For ordinary paperbacks, that is
  usually the correct mental-load trade.
- **Dépôt-vente / consignment**: consider only for higher-value furniture,
  design pieces, luxury, or items where the person's priority is low hassle but
  not "gone today". Mention commission/handoff friction briefly.
- **Donation/give-away threshold**: under roughly **10-15 € net**, or when the
  item needs cleaning, transport, dimensions, many messages, or awkward pickup,
  recommend donation/Geev/recycling unless it is part of a bundle.

## Bundling rules

- Books: split into **ordinary bulk** (buy-back/donate) and **valuable few**
  (sell individually).
- Baby/kids clothes: bundle by **size + season**.
- Adult basics: bundle by **size + style**, or donate if the brands are weak.
- Kitchenware/decor: bundle as "lot premier appartement" or "lot cuisine".
- Toys: bundle by age/license only if complete and clean; otherwise donate.

## Output language examples

Use concise French when the user is French:

```text
PRIORITÉ : équilibré
ESTIMATION : 25-40 € — vélo enfant Decathlon 20 pouces en bon état, demande locale correcte.
MOVE : Mets-le à 35 € sur Leboncoin, remise en main propre; accepte 30 € si la personne vient vite.
POURQUOI : assez de valeur pour vendre, mais pas assez pour optimiser pendant deux semaines.
OPTION : aucune.
```

```text
PRIORITÉ : vite débarrassé
ESTIMATION : 0-8 € — lampe Ikea commune, faible valeur seule.
MOVE : Donne-la sur Geev ou ajoute-la à un lot maison; ne fais pas d'annonce dédiée.
POURQUOI : les messages et le rendez-vous valent plus que le prix probable.
OPTION : aucune.
```
