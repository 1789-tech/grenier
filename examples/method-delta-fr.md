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

## Round 2 — load-bearing scenarios captured 2026-07-15

Captured from `/tmp`, outside this repository, with no grenier files in the
working directory:

```bash
claude -p --model opus "<case prompt>"
```

### Case 5 — async Vinted draft production

Prompt:

> France. J ai 12 vetements a mettre sur Vinted pendant que je dors: jean Levi s
> 501 W30 bon etat, robe Sezane taille 38 tres bon etat, baskets Nike Air Force
> 1 taille 39 usagees, pull Uniqlo laine M bon etat, veste Zara M bon etat,
> chemise homme Celio L correcte, sac Longchamp pliage use coins frottes, jupe
> Maje 38 bon etat, manteau Mango M tres bon etat, t-shirt basique H&M M
> correct, ceinture cuir marron 90 cm bon etat, lot 3 foulards sans marque. Je
> veux des brouillons prets a coller, sans publier, avec prix, categorie, titre,
> description, champs a verifier.

Plain Claude:

```text
Voici les 12 brouillons prêts à coller. Prix indicatifs marché FR Vinted 2026,
à ajuster selon photos et concurrence.

1. Jean Levi's 501 W30
- Titre : Jean Levi's 501 W30 coupe droite bleu classique
- Catégorie : Femmes ou Hommes > Vêtements > Jeans > Droit (à vérifier selon coupe/genre étiquette)
- Marque : Levi's
- Taille : W30 (préciser longueur L32/L34 si visible)
- État : Bon état
- Couleur : Bleu
- Matière : 100% coton (vérifier étiquette)
- Prix : 20 € (fourchette 18-25)
- Description : Jean Levi's 501 authentique, W30. Coupe droite emblématique,
  délavage bleu classique. Bon état général, très peu porté. Aucun trou ni
  tache. Vendu sans ceinture. Envoi rapide et soigné, remise en main propre
  possible sur [ville].
- À vérifier : longueur L, made in, coupe homme/femme, patch cuir lisible.

2. Robe Sézane taille 38
- Titre : Robe Sézane 38 très bon état
- Catégorie : Femmes > Vêtements > Robes > (midi/courte/longue à préciser)
- Marque : Sézane
- Taille : 38 / S
- État : Très bon état
- Couleur : (à préciser)
- Matière : (vérifier étiquette : viscose, coton, lin...)
- Prix : 40 € (fourchette 35-50, +/- selon modèle iconique)
- Description : Robe Sézane taille 38, portée deux ou trois fois, aucun défaut.
  Modèle [nom si connu]. Très joli tombé, idéale [saison]. Provient d'un
  intérieur non-fumeur. Envoi soigné sous 24-48h.
- À vérifier : nom du modèle, saison, longueur, doublure.

...10 more item drafts with title, category, brand, condition, price,
description and fields to verify...

Rappels avant publication:
- Photos : 4-8 par annonce, lumière naturelle, plan large + zoom sur défauts.
- Frais de port : laisser Vinted calculer, ou proposer Mondial Relay S/M selon poids.
- Boost timing : publier entre 19h et 22h, pic d'audience.
- Groupement : activer « lot possible » sur t-shirt, chemise, ceinture, foulards.
- Négociation : prévoir 10-15% de marge, les acheteurs proposent quasi systématiquement.
```

Expected grenier delta:

```text
This is not a strong standalone proof. Plain Claude already produces usable
drafts, prices, category guesses and verification fields.

grenier should still improve the workflow, but the gap is operational, not
content-only:
- explicit "draft only / never publish" state per item;
- completed / needs review / skipped counts;
- low-value bundling veto for H&M/Celio/foulards where one-by-one drafts are bad
  net value;
- a resumable item map: item id, platform, draft status, guessed fields,
  next safe action.
```

Delta: thin. Vinted draft production alone does **not** prove a flagrant gap
against raw Claude; it proves the product needs safe execution primitives and
session state, not just better prompt knowledge.

### Case 6 — mixed pile with derived receipt

Prompt:

> France. Je vide une cave avec: velo enfant Decathlon 20 pouces bon etat, lampe
> Ikea basique, 30 romans poche ordinaires, 3 beaux livres art, micro-ondes en
> panne, 8 verres depareilles, manteau Mango M tres bon etat, sac designer
> possible sans preuve. Priorite equilibree. Donne un plan avec recu chiffre
> derive, deadlines, et premier geste.

Plain Claude:

```text
Plan de vidage — 2 semaines

Estimation totale récupérable : 120–230 €

Ce qui se vend (Vinted / Le Bon Coin) — J+1 à J+7
- Vélo enfant Decathlon 20" — LBC — 40–70 €
- Manteau Mango M — Vinted — 15–30 €
- 3 beaux livres art — Momox ou Vinted — 15–60 € les 3
- Sac "designer" sans preuve — Vestiaire Collective ou LBC "origine non vérifiée"

Don (Emmaüs, boîtes à livres, ressourcerie) — J+7 à J+10
- 30 romans poche -> boîtes à livres de quartier
- 3-4 beaux livres art invendus après 1 semaine -> Gibert Joseph rachat ou don
- Lampe Ikea, 8 verres dépareillés -> Emmaüs ou ressourcerie

Déchetterie / DEEE — J+10
- Micro-ondes en panne -> déchetterie, borne DEEE

Premier geste : sors tous les objets, photographie-les à la lumière du jour.
```

Expected grenier shape:

```text
PRIORITE : equilibree

SCORED ROUTES
- Velo Decathlon: sell Leboncoin, start 55-65, accept 45-55, D+3 drop, D+7 donate.
- Manteau Mango: Vinted draft, start 25-30, accept 20-25, review at D+7.
- Beaux livres: scan ISBN; sell only if each comp >25, otherwise bundle/donate.
- Romans ordinaires: buy-back/donate, no one-by-one sale.
- Lampe/verres: Geev/Emmaus, no standalone sale.
- Micro-ondes HS: DEEE, no sale unless explicit parts niche.
- Sac designer: no direct sale; authenticate or keep/donate only as legally safe.

GRENIER RECEIPT
Decisions: 8 lots
Cash recoverable: about 65-125 € from sell-marked lots only
Effort avoided: about 4-6 h from refused micro-sales
Scheduled exits: DEEE run, Geev/Emmaus bag, ISBN scan, 2 listings, auth check
Deadlines: D+3 price drop / D+7 exit for sell items
First move now: photograph bike + Mango coat, then scan the 3 art-book ISBNs
Derivation: each total points back to item-level judgments above.
```

Delta: meaningful. Plain Claude gives a plausible total but it is not auditable:
the 120-230 € total includes unverified/suspect routes, the receipt is not
derived line-by-line, and the sac route is unsafe enough to hurt trust. This is
the clearest load-bearing Grenier advantage from the round-2 set.

### Case 7 — interruption-safe resumption

Prompt:

> France. Reprise de session: hier on avait deja decide Vinted pour robe Sezane
> 38, jupe Maje 38, manteau Mango M; don Geev pour lampe Ikea et verres;
> Leboncoin pour velo enfant; verification obligatoire pour sac designer sans
> preuve. J ai ete coupe apres 3 brouillons. Reprends sans me refaire tout le
> raisonnement: dis quoi verifier, quoi faire ensuite, et comment eviter
> doublons/publication accidentelle.

Plain Claude:

```text
Sans refaire le raisonnement:

À vérifier
- Sac designer: facture d'achat, carte d'authenticité, dust bag, numéro de série.
- Vélo enfant: taille de roue exacte, tranche d'âge, état des freins/pneus.
- Manteau Mango M: composition + saison.

Étapes suivantes
1. Photos: fond neutre, 4-6 vues par pièce.
2. Pèse chaque pièce pour préremplir le tarif Mondial Relay sur Vinted.
3. Geev (lampe + verres): une seule annonce "lot don" si retrait groupé.
4. Sac designer: mets-le de côté tant que la preuve n'est pas retrouvée.

Éviter doublons / publication accidentelle
- Tiens une checklist unique.
- Sur Vinted, vérifie l'onglet "Brouillons" avant de créer.
- Ne jamais rédiger deux annonces en parallèle dans deux onglets.
- Ne publie rien tant que photos + prix + mode d'envoi ne sont pas remplis.
```

Expected grenier delta:

```text
Again, raw Claude is competent. grenier's advantage needs persistent state:

SESSION RESUME
Known completed drafts: 3 (must be listed by item id or marked "unknown until
draft list is checked").
Pending Vinted: robe Sezane / jupe Maje / manteau Mango minus completed drafts.
Pending non-Vinted: Geev lot, Leboncoin bike, designer authentication.
Safety gate: call/list drafts first, reconcile against item map, then fill only
missing drafts with save_as_draft=true.
No publish action allowed in this continuation.
```

Delta: thin without tooling/state. Plain Claude already gives the right safety
advice. Grenier becomes clearly better only if it can carry a real item map and
query draft state; as a static skill, the proof is not flagrant.

### Case 8 — platform-specific workflow friction

Prompt:

> France. Je dois creer des annonces sur Vinted et Leboncoin. Explique
> concretement le workflow plateforme: champs/categorie a remplir, risques de
> mauvais bouton, sauvegarde brouillon, prix/deadline, et quels objets ne pas
> mettre sur la plateforme. Objets: lot vetements femme, velo enfant Decathlon,
> livres ordinaires, sac designer sans preuve, micro-ondes en panne.

Plain Claude:

```text
Répartition des objets
- Lot vêtements femme -> Vinted
- Vélo enfant Decathlon -> Leboncoin
- Livres ordinaires -> Leboncoin en lot, ou Recyclivre / Momox
- Sac designer sans preuve -> Ne pas lister
- Micro-ondes en panne -> Ne pas vendre

Vinted — workflow concret
- Photos, titre, description, catégorie, marque, taille, état, couleur, matière,
  colis, prix.
- Publier vs Enregistrer comme brouillon.
- Boost / Dressing en vitrine : payant, skip.
- Prix : regarde 5-10 articles similaires vendus.
- Deadline : aucune. L'annonce reste en ligne.

Leboncoin — workflow concret
- Vélo enfant Decathlon : Sports & Hobbies -> Vélos -> Vélos enfants.
- Livres : Loisirs -> Livres.
- Enregistrer en brouillon, options payantes, paiement sécurisé, remise en main propre.
- Deadline : annonce active 60 jours.

Objets à NE PAS mettre
- Sac designer sans preuve.
- Micro-ondes en panne.
```

Expected grenier delta:

```text
grenier should be stricter on net-exit deadlines:
- Vinted still needs a D+7 review and bundle/drop action for low-value clothes;
  "no deadline" recreates stagnation.
- Leboncoin needs D+3 drop / D+7 exit for bike and books.
- Ordinary books should default Momox/RecycLivre/donation, not a Leboncoin lot
  unless photos/listing effort is justified.
- The workflow should end with a safe execution checklist: drafts saved, paid
  options skipped, publish forbidden until review.
```

Delta: moderate, not flagrant. Plain Claude covers a surprising amount of
platform friction. Grenier's useful edge is the anti-stagnation policy and
net-value vetoes, not secret platform knowledge.

## Round-2 conclusion

The redirected test set does **not** prove a broad, flagrant content-quality gap
between raw Claude and grenier. Raw Claude is already strong on drafting,
platform walkthroughs and resumption advice when the prompt is specific.

What remains defensible is narrower and more product-shaped:

- **Load-bearing win:** mixed-pile orchestration with an auditable derived
  receipt, hard net-value cutoffs, and deadlines.
- **Conditional win:** async Vinted/Leboncoin execution only if grenier has safe
  draft primitives and persistent item state (`list_drafts`, `save_as_draft`,
  item map reconciliation). Prompt knowledge alone is not enough.
- **Thin areas:** generic Vinted draft copy, resumption advice and platform
  workflow explanations; raw Claude already performs well there.

Recommendation: keep PR #7 as method groundwork, but do not present this
artifact as proof that grenier is globally "nettement mieux" than Claude nu.
The next product proof should be an execution harness around stateful draft
production and mixed-pile receipt math, not more static advice examples.
