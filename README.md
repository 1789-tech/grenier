# grenier 🏠📦

[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
![skills](https://img.shields.io/badge/skills-4-ffb000.svg)
![evals](https://img.shields.io/badge/evals-harness-3aa675.svg)
[![made by 1789](https://img.shields.io/badge/made%20by-1789-111111.svg)](https://1789.tech)

> Un compagnon Claude pour **désencombrer** sa maison, façon collection de
> *Skills*. Inspiré de [paperasse](https://github.com/romainsimon/paperasse) — mais pour le tri,
> la revente et le recyclage. **France-first**, open source.

`grenier` est une collection de **Claude Skills** qui t'aident à :

1. **Trier** — méthode, exercices, étapes, techniques pour t'attaquer au bazar.
2. **Décider quoi en faire** — pour chaque objet : vendre (et où), donner,
   recycler, jeter. Plateformes françaises (Leboncoin, Vinted, Emmaüs, Geev,
   Momox/Recyclivre pour les livres…).
3. **Préparer la vente** — suggérer un prix, rédiger l'annonce, grouper les lots.
4. **Négocier sans y passer sa vie** — répondre aux acheteurs, contre-offres,
   réservations, no-shows.

On commence petit : des Skills purement *prompt* (niveaux 1–2), utiles dès le
premier jour. Les parties agentiques (vision photo, scraping de prix) viennent
ensuite.

## Skills

| Skill | Niveau | État |
|-------|--------|------|
| [`declutter-coach`](skills/declutter-coach/) | 1 · Coaching de tri | ✅ v0 |
| [`where-to-sell-fr`](skills/where-to-sell-fr/) | 2 · Quoi en faire | ✅ v0 |
| [`listing-drafter`](skills/listing-drafter/) | 3 · Préparer la vente | ✅ v0 |
| [`haggle`](skills/haggle/) | 4 · Répondre aux acheteurs | ✅ v0 |

## Installation

Ces Skills suivent le format [Claude Agent Skills](https://docs.claude.com). Trois
portes d'entrée, selon ton niveau de contrôle voulu :

**Option A — demander à ton agent (le plus simple)**

```text
Installe les Skills grenier depuis https://github.com/1789-tech/grenier
Puis lance une session de tri de 15 minutes pour une zone de ma maison.
Commence par me demander la zone, mon énergie, et ce que je veux obtenir à la fin.
```

**Option B — marketplace Claude Code (versionné)**

```text
/plugin marketplace add 1789-tech/grenier
/plugin install grenier@grenier
```

**Option C — script inspectable (CLI/CI)**

```bash
curl -fsSL https://raw.githubusercontent.com/1789-tech/grenier/main/install.sh | sh
```

Le script copie les skills dans `${CLAUDE_SKILLS_DIR}`, `${CLAUDE_CONFIG_DIR}/skills`,
`${XDG_CONFIG_HOME}/claude/skills` si ce dossier existe, sinon `~/.claude/skills`.

Relance ton agent, puis lance-toi — pas besoin de commande spéciale, décris ta
situation et la Skill se déclenche toute seule :

> « J'ai un garage bordélique et je sais pas par où commencer. »
> → `declutter-coach` prend la main : une zone, un créneau de 15 min, on trie.

> « J'ai un vieux vélo, un carton de romans et une imprimante à me débarrasser. »
> → `where-to-sell-fr` arbitre pour chaque objet : vendre (et où), donner,
> recycler ou jeter.

> « Rédige une annonce Leboncoin pour ce vélo enfant 20 pouces, quelques rayures. »
> → `listing-drafter` prépare titre, prix, description, photos et logistique.

> « L'acheteur propose 30 € alors que j'ai mis 80 €, je réponds quoi ? »
> → `haggle` propose une réponse courte, ferme et polie.

**Option dev — pointer sur le repo**

Ouvre Claude Code (ou tout harnais compatible Skills) dans le dossier `grenier/`
cloné : les Skills du dossier `skills/` sont disponibles directement.

### Les Skills en un coup d'œil

| Tu veux… | Dis quelque chose comme… | Skill |
|----------|--------------------------|-------|
| T'attaquer au bazar sans stress | « aide-moi à trier ma chambre » | `declutter-coach` |
| Savoir quoi faire d'un objet | « où je revends / donne ça ? » | `where-to-sell-fr` |
| Publier une annonce propre | « rédige une annonce Vinted/Leboncoin » | `listing-drafter` |
| Répondre aux acheteurs | « il négocie trop bas, je dis quoi ? » | `haggle` |

Pas besoin de coder. Si tu sais discuter avec Claude, tu sais utiliser grenier.

### Parcours complet

> « J'ai un garage plein, 45 minutes, et je veux que ça parte. »

grenier enchaîne : `declutter-coach` découpe une micro-zone, `where-to-sell-fr`
route chaque objet, `listing-drafter` écrit les annonces pour ce qui mérite
d'être vendu, puis `haggle` aide à répondre sans perdre la soirée.

### Compatible avec quoi ?

Le format est celui des Skills lisibles par agent : Claude Code en premier, mais
aussi tout harnais capable de charger un dossier `skills/` et d'injecter un
`SKILL.md` dans le contexte (Codex, Cline, Aider, Cursor, Windsurf, etc.).

## Evals

Un harness minimal est disponible :

```bash
make test
make run
```

Par défaut, il valide les fixtures et les chemins de Skills. Pour un vrai
benchmark `with skill` / `without skill`, branche ton générateur :

```bash
EVAL_GENERATE_CMD='mon-agent --stdin' make run
```

Le delta obtenu remplacera le badge `evals-harness`.

## Licence

MIT — voir [LICENSE](LICENSE). Contributions bienvenues, surtout les recettes
locales (autres pays, autres plateformes).

---

*Un produit [1789](https://1789.tech) · vitrine : [grenier.1789.tech](https://grenier.1789.tech)*
