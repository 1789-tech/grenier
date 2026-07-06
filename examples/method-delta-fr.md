# Method delta — France

Prompt:

> France. Je désencombre: une lampe Ikea basique, 30 livres de poche ordinaires,
> 3 beaux livres d'art, un vélo enfant Decathlon, et un vieux sac qui ressemble à
> du designer mais sans preuve d'authenticité. Je veux un résultat honnête, pas
> maximiser chaque euro.

## Expected grenier shape

```text
PRIORITÉ : équilibré

1) Lampe Ikea commune — vente dédiée refusée
Calcul net :
  cash réaliste      = +0-8 €
  taxe d'effort      = -30-45 min (photos + messages + rendez-vous)
  valeur de sortie   = +objet sorti aujourd'hui
  risque de regret   = 0
NET CALL : don / lot maison, pas Leboncoin dédié.
MOVE : Geev ou sac Emmaüs aujourd'hui.

2) Livres — split immédiat
30 poches ordinaires -> Momox/RecycLivre/Gibert ou don, pas un par un.
3 beaux livres d'art -> vendre seulement si chaque comp dépasse ~25 €, sinon lot
premium.

3) Vélo enfant Decathlon — liquide local
Calcul net :
  cash réaliste      = +55 € cible (45-65 €)
  taxe d'effort      = -30-45 min, remise locale
  valeur de sortie   = +1 objet volumineux hors cave
  risque de regret   = 0 si aucun usage daté
NET CALL : vaut une vente locale.
MOVE : Leboncoin 65 €, accepter 55 €, baisser 45 € à J+3.
SORTIE : à J+7, don famille/association si aucun acheteur qualifié.

4) Sac designer non prouvé — trust veto
Ne pas vendre comme designer sans preuve. Identifier modèle/preuves d'abord. Si
non authentifiable, vendre honnêtement comme sac sans marque confirmée, donner,
ou garder intentionnellement.

REÇU GRENIER
Décisions : 4 lots
Cash récupérable : environ 45-65 € + éventuels beaux livres authentifiés
Temps évité : environ 3-5 h de micro-annonces inutiles
Sorties : 1 don aujourd'hui, 1 scan livres, 1 annonce vélo avec deadline, 1 sac à
vérifier avant vente
Premier geste : prendre 4 photos du vélo et le publier à 65 €
```

## Baseline capture status

The issue asks for real same-day Claude-nu baselines. Command attempted on
2026-07-06:

```bash
claude --bare -p --model opus "<case prompt>"
```

Local result: `Not logged in · Please run /login`.

Do not paste invented baselines here. Once an authenticated Claude CLI or API key
is available, run the four case prompts from the board issue and paste the exact
plain-Claude answers next to this expected grenier shape before final merge.
