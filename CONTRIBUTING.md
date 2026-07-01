# Contribuer à grenier

Merci d'aider ! grenier est une collection de **Claude Skills** pour désencombrer
et revendre. Les contributions les plus utiles : des **recettes locales** (autres
pays, autres plateformes) et l'amélioration des Skills existantes.

## Anatomie d'une Skill

Une Skill = un dossier sous `skills/<nom>/` avec un fichier `SKILL.md` :

```
skills/<nom>/
  SKILL.md        # obligatoire
  references/     # optionnel : fichiers de référence chargés à la demande
```

Le `SKILL.md` commence par un frontmatter YAML :

```yaml
---
name: mon-skill                 # kebab-case, unique
description: >-                  # QUAND l'utiliser (Claude s'en sert pour déclencher)
  Une-deux phrases : ce que fait la Skill ET dans quelles situations la lancer.
metadata:
  level: 2                      # 1 tri · 2 décider · 3 annonce · 4 suivi
  status: v0                    # v0 / v1 / expérimental
  last_updated: 2026-07-01      # date de dernière révision
  locale: fr-FR
---
```

Puis, juste après le titre `#`, une ligne **« Comment m'utiliser »** en langage
non-codeur (ce que la personne dit pour déclencher la Skill).

## Règles de style

- **France-first**, mais structure « cœur global + recette locale » quand c'est
  possible. Une recette `-fr`, `-be`, `-ca`… est bienvenue.
- **Anti-perfectionnisme** : le biais du produit est de *faire partir les choses*,
  pas d'optimiser chaque euro.
- **Honnêteté** : pas de fausses promesses de prix, pas de superlatifs creux.
- **Vie privée** : jamais d'inventaire, de photos, de prix ou de données perso
  réels dans le repo — uniquement du contenu de prompt et de la doc.
- Teste que la Skill **bat le fait de demander à Claude sans elle**. Sinon, ce
  n'est pas un produit.

## Proposer

1. Fork + branche.
2. Ajoute/édite la Skill, mets un exemple dans `examples/` si pertinent.
3. Mets `last_updated` à jour.
4. Ouvre une PR en expliquant le cas d'usage et le pays/plateforme visé.

Licence des contributions : MIT (voir [LICENSE](LICENSE)).
