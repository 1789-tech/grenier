# grenier 🏠📦

> Un compagnon Claude pour **désencombrer** sa maison, façon collection de
> *Skills*. Inspiré de [paperasse](https://github.com/) — mais pour le tri,
> la revente et le recyclage. **France-first**, open source.

`grenier` est une collection de **Claude Skills** qui t'aident à :

1. **Trier** — méthode, exercices, étapes, techniques pour t'attaquer au bazar.
2. **Décider quoi en faire** — pour chaque objet : vendre (et où), donner,
   recycler, jeter. Plateformes françaises (Leboncoin, Vinted, Emmaüs, Geev,
   Momox/Recyclivre pour les livres…).
3. **Préparer la vente** — d'une photo, deviner l'objet, suggérer un prix,
   rédiger l'annonce, grouper les lots.
4. **Suivre la vente** — répondre aux acheteurs, optimiser la logistique.

On commence petit : des Skills purement *prompt* (niveaux 1–2), utiles dès le
premier jour. Les parties agentiques (vision photo, scraping de prix, suivi)
viennent ensuite.

## Skills

| Skill | Niveau | État |
|-------|--------|------|
| [`declutter-coach`](skills/declutter-coach/) | 1 · Coaching de tri | ✅ v0 |
| [`where-to-sell-fr`](skills/where-to-sell-fr/) | 2 · Quoi en faire | ✅ v0 |
| `listing-drafter` | 3 · Préparer la vente | 🚧 à venir |
| `sale-tracker` | 4 · Suivi | 🔭 plus tard |

## Utilisation

Ces Skills suivent le format [Claude Agent Skills](https://docs.claude.com).
Pointe Claude Code (ou tout harnais compatible Skills) sur ce repo, ou copie un
dossier de `skills/` dans ton propre `.claude/skills/`.

## Licence

MIT — voir [LICENSE](LICENSE). Contributions bienvenues, surtout les recettes
locales (autres pays, autres plateformes).

---

*Un produit [1789](https://1789.tech) · vitrine : [grenier.1789.tech](https://grenier.1789.tech)*
