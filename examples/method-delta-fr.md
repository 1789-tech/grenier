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

## Plain-Claude baseline captured 2026-07-15

Captured from `/tmp`, outside this repository, with no grenier files in the
working directory:

```bash
claude -p --model opus "<case prompt>"
```

### Case 1 — lampe Ikea basique

Prompt:

> France. Je désencombre une lampe Ikea basique. Je veux un résultat honnête,
> pas maximiser chaque euro. Que dois-je faire ?

Plain Claude:

```text
Pour une lampe IKEA basique en France, l'honnête et rapide :

Option la plus simple : don
- Emmaüs (ils passent chez toi gratuitement pour les gros lots, sinon dépôt en
  boutique)
- Le Bon Coin en "gratuit à venir chercher" — partie souvent en 24-48h
- Geev (appli de don entre particuliers) — très actif pour ce genre d'objet

Si personne n'en veut :
- Déchèterie, benne "tout-venant" ou "DEEE" si elle est électrique (avec
  ampoule/câble)
- Ou dépôt gratuit dans n'importe quel magasin IKEA (bac de reprise à l'entrée)

Ce que je ne ferais pas à ta place : essayer de la vendre. Une lampe IKEA
basique se vend 3-5€, pour 30 min de photos + messages + rendez-vous. Le ratio
temps/argent n'en vaut pas la peine si l'objectif est de désencombrer.

Ma reco : poste-la sur Geev ce soir, si pas de preneur d'ici 2-3 jours, dépôt
IKEA ou déchèterie au prochain passage.
```

Delta: thin. Grenier adds the named 20/20/20 doctrine and receipt format, but
plain Claude already makes the core refusal, effort trade-off, channels and
deadline.

### Case 2 — 30 livres ordinaires + 3 beaux livres d'art

Prompt:

> France. Je désencombre 30 livres de poche ordinaires et 3 beaux livres d art.
> Je veux un résultat honnête, pas maximiser chaque euro. Que dois-je faire ?

Plain Claude:

```text
Voici un plan honnête et proportionné à l'effort.

Les 30 poches : ne perdez pas de temps à les vendre à l'unité. Les rachats en
ligne (Momox, Recyclivre) refusent la moitié des poches courants et paient 10-30
centimes le reste — c'est du tri pour rien. Deux options simples :
- Boîte à livres de quartier : dépôt gratuit, ça circule vite, usage réel.
- Emmaüs / Recyclerie / bibliothèque municipale : don en bloc, un seul
  déplacement.

Les 3 beaux livres d'art : ça vaut le petit effort de vérifier. Cherchez chaque
titre sur Leboncoin, Rakuten et Momox pour voir le prix d'occasion réel.
- Si un livre se revend >20-30 € → Leboncoin (moins de commission que Vinted
  pour ce format, acheteurs plus ciblés). Une seule annonce par livre, photo
  correcte, prix aligné, c'est vendu en quelques semaines.
- Si c'est un livre confidentiel ou abîmé → même destin que les poches, avec une
  préférence pour une bibliothèque ou une école d'art à proximité (usage plus
  juste que la benne d'un tri).

En pratique : 20 minutes pour les 3 beaux livres (recherche + annonces), et un
carton à déposer pour les 30 autres. Le désencombrement est fait dans la
journée, sans culpabiliser sur les quelques euros laissés.
```

Delta: thin. Grenier can package this as a class split and derive a receipt, but
the raw answer already gives the split, threshold, channels, time budget and
anti-micro-sale stance.

### Case 3 — vélo enfant Decathlon

Prompt:

> France. Je désencombre un vélo enfant Decathlon en bon état. Je veux un
> résultat honnête, pas maximiser chaque euro. Que dois-je faire ?

Plain Claude:

```text
Pour un vélo Decathlon en bon état, en France, deux voies honnêtes et sans
friction :

1. Decathlon Occasion (reprise en magasin) — c'est le plus cohérent : tu
apportes le vélo, ils l'évaluent en quelques minutes et te font une offre en bon
d'achat (ou parfois en cash selon la période). Prix un peu en dessous du marché
particulier, mais zéro annonce à gérer, zéro rendez-vous, zéro négociation.
Idéal pour "désencombrer honnêtement". → https://occasion.decathlon.fr rubrique
Vendre, ou directement au comptoir Occasion du magasin.

2. Don à une structure locale — Emmaüs, une ressourcerie, une Maison du vélo, ou
une association type Un vélo pour tous / Solicycle. Un vélo enfant part très vite
et sert à une famille qui en a besoin. Pas d'argent, mais résultat clair et
utile.

Le compromis :
- Reprise Decathlon = tu récupères ~30-50 % de la valeur neuve sous forme de
  bon, en 15 min.
- Don = 0 € mais impact social et zéro logistique.
- Le Bon Coin / Vinted rapporteraient 20-40 % de plus qu'une reprise, mais c'est
  exactement ce que tu dis vouloir éviter.

Ma reco par défaut : reprise Decathlon si tu comptes racheter du matériel sport
à terme, sinon don Emmaüs / ressourcerie du coin. Dans les deux cas, c'est réglé
dans la semaine.
```

Delta: not proven. Grenier's expected Leboncoin move is more optimized for cash,
but plain Claude may be better aligned with "not maximizing every euro" because
it suggests Decathlon Occasion or donation.

### Case 4 — sac designer sans preuve

Prompt:

> France. Je désencombre un vieux sac qui ressemble à du designer mais sans
> preuve d authenticité. Je veux un résultat honnête, pas maximiser chaque euro.
> Que dois-je faire ?

Plain Claude:

```text
Voici l'approche honnête, dans l'ordre :

1. Déterminer d'abord ce que c'est
Regarde les indices concrets : coutures régulières, tampon "Made in…", numéro de
série/date code intérieur, qualité de la doublure et de la quincaillerie, poids.
Compare à des guides d'authentification en ligne pour la marque supposée. Ce
n'est pas une preuve, mais ça oriente vers "probablement authentique" ou
"probablement copie".

2. Si c'est probablement une contrefaçon
En France, la revente de contrefaçon est un délit (art. L716-9 et suivants du
Code de la propriété intellectuelle), même entre particuliers, même sans
intention de tromper. Donc :
- Ne le mets pas en vente sous le nom de la marque, ni en "inspiré de", ni
  "style X" — les plateformes et les douanes considèrent ça comme contrefaçon.
- Garde-le pour ton usage, ou détruis-le. Les associations (Emmaüs, Croix-Rouge)
  refusent normalement les contrefaçons.

3. Si c'est probablement authentique mais sans facture
Passe par une plateforme qui authentifie elle-même : Vestiaire Collective,
Rebelle, ou un dépôt-vente physique près de chez toi. Ils rejettent si c'est faux
— ce qui est exactement le filtre honnête que tu cherches. Tu toucheras moins
qu'en direct, mais l'acheteur a une garantie et toi tu es couverte juridiquement.

4. Si tu ne veux pas t'embêter
Dépôt-vente physique : tu déposes, ils décident, tu récupères ton dû ou l'objet.
Zéro effort, zéro risque.

La règle courte : jamais vendre en solo un sac "griffé" sans preuve — soit tu
passes par un tiers qui authentifie, soit tu le sors du circuit.
```

Delta: thin to negative. Plain Claude is stricter on legal/authenticity risk
than the expected grenier shape and already refuses the unsafe sale path.

## Delta verdict

The current method is useful, but these four comparison prompts do **not** prove
the "flagrant" gap the board brief requires. The strongest remaining grenier
advantages are packaging and reproducibility: named cutoff, item-level receipt
math, exit deadlines and multi-item orchestration. To justify the skill vs raw
Claude, the comparison set should be changed to cases where those advantages are
load-bearing, especially async Vinted draft production, mixed piles requiring a
derived receipt, interruption-safe resumption, or platform-specific listing
workflows.
