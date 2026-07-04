# Example — one price, one move

**Skill:** `price-and-decide`

## What you say

> "France. J'ai un vélo enfant Decathlon 20 pouces, bon état, quelques rayures.
> Je veux un choix équilibré."

## What Claude does (abbreviated)

`price-and-decide` asks or infers the priority, then gives a compact decision:

```text
PRIORITÉ : équilibré
ESTIMATION : 25-45 € — vélo enfant Decathlon 20 pouces en bon état, demande locale correcte.
MOVE : Mets-le à 35 € sur Leboncoin, remise en main propre; accepte 30 € si la personne vient vite.
POURQUOI : assez de valeur pour vendre, mais pas assez pour optimiser pendant deux semaines.
OPTION : aucune.
```

For the same item with **gone fast**, the move can flip to "30 € this weekend or
donate to Emmaüs if no pickup". For **max cash**, it can hold nearer 45 € and
wait longer. The point is not perfect pricing; it is one clear decision.
