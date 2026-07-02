# Evals

`grenier` est une collection de Skills prompt-only. Les evals mesurent donc la
qualite des reponses attendues par assertions simples, avec deux modes :

- sans commande : validation rapide des fixtures et des chemins de Skills ;
- avec `EVAL_GENERATE_CMD` : benchmark `with skill` / `without skill` contre le
  modele ou harnais de ton choix.

```bash
make run
EVAL_GENERATE_CMD='mon-agent --stdin' make run
```

Le runner envoie au generateur soit `SKILL.md + prompt`, soit le prompt seul,
puis compare les sorties avec les assertions `must_include` / `must_avoid` de
`evals.json`. Le delta est le chiffre a publier en badge une fois le generateur
stabilise.
