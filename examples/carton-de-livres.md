# Exemple — vendre un carton de livres

**Skills :** `where-to-sell-fr` (arbitrage) puis `listing-drafter` (si vente unité)

## Ce que tu dis

> « J'ai un carton de ~40 bouquins, surtout des romans de poche, quelques beaux
> livres d'art. J'en fais quoi le plus simplement ? »

## Ce que Claude fait (abrégé)

`where-to-sell-fr` tranche par **volume et valeur** :

- **Les ~35 romans de poche** → *Momox / RecycLivre* : scanne les ISBN via leur
  appli, ils te font une offre globale et payent le port. Prix faible à l'unité
  (souvent 0,10–2 €/livre), mais **zéro friction** — l'idéal pour du volume.
- **Les 4–5 beaux livres d'art** → valeur unitaire réelle (10–40 €) → *Leboncoin
  ou Vinted à l'unité*. Là, `listing-drafter` rédige chaque annonce :

```
TITRE : Beau livre d'art "Hokusai" - Éditions Hazan - Très bon état
PRIX : 22 € (fourchette : 15–30 €)
PLATEFORME : Vinted — livres/loisirs, envoi simple
...
```

- **Ce que Momox refuse** (trop abîmé, non racheté) → borne de recyclage papier ou
  boîte à livres de quartier.

Bilan : le gros du carton part en un envoi, les quelques pièces de valeur sont
vendues à l'unité, rien ne finit à la poubelle.
