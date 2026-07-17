# cave — arm B (Claude opus + grenier SKILL.md as system prompt)

Working dir: `/tmp/brief102/cave-B` (empty)
Command per turn:
`claude -p --model opus --disallowedTools "*" --append-system-prompt-file /home/russel/p/gh/1789-tech/grenier/skills/grenier/SKILL.md`
(turn 1) then same with `-c` (turns 2+).

**Note on progressive disclosure:** the model reported at turn 2 that it
could not access the `references/` files (the skill's on-demand disclosure).
It applied the doctrine encoded in `SKILL.md` and drew country-specific
channel knowledge from its own training. This is honest and matches how a
zero-install / `claude -p` invocation actually behaves.

---

## Turn 1 — user

> France. J'ai un bordel dans la cave, plein de trucs entassés depuis des
> années, je sais pas par où commencer. J'en peux plus.

## Turn 1 — assistant

Respire. On ne va **pas** vider la cave aujourd'hui. On va faire *un* petit truc, et tu vas voir le sol quelque part.

La règle : 15 minutes, une seule zone d'1 m² (le coin près de la porte, ou un carton, ou une étagère — celui qui t'énerve le plus quand tu le vois). C'est tout.

Avant qu'on se lance, deux questions rapides :

1. **T'as combien de temps là, maintenant ?** (15 min ? 1 h ? une aprem ?)
2. **Ça ressemble à quoi dans la cave ?** Plutôt : cartons fermés jamais rouverts / meubles et gros objets / outils et bricolage / souvenirs et papiers / vrai bazar mélangé ?

Pendant que tu réponds, va chercher :
- 1 sac poubelle
- 3 cartons ou zones au sol, que tu étiquettes mentalement : **GARDE** / **SORT DE LA MAISON** (vendre ou donner, on triera après) / **RECYCLE**
- Un minuteur (ton téléphone)

Chaque objet touché = une décision immédiate, une des 4 destinations. Pas de tas "je verrai plus tard" — c'est ce tas qui a créé la cave.

Dis-moi le temps et le type de bordel, et je te dis exactement par quel mètre carré commencer.

---

## Turn 2 — user

> J'ai un weekend disponible, peut-être un jour et demi si je m'y mets. Y'a
> un peu de tout : mon ancien vélo route, un vélo enfant, des cartons de
> livres de poche et de vinyles, une vieille chaîne hifi qui marche, un
> matelas simple, un blender cassé, une table basse rayée. Et deux cartons
> de câbles / trucs random. Rien de précieux.

## Turn 2 — assistant

Je n'arrive pas à accéder aux fichiers détaillés du skill dans cet environnement, mais j'ai en tête la doctrine et je connais bien les circuits français. Je te fais un plan concret directement.

## Ce qui est vraiment récupérable : ~150–400 €

Fourchette honnête, en gros :

| Objet | Destination | € réaliste FR |
|---|---|---|
| Vélo route adulte | **VENDRE** (Leboncoin) | 80–250 € selon marque/état |
| Chaîne hifi qui marche | **VENDRE** (Leboncoin) | 30–150 € (si Marantz/Pioneer/Sansui/Technics vintage → plus) |
| Vinyles | **VENDRE en lot** (Leboncoin) ou pièce par pièce si perles | 50–200 € — potentiellement +++ |
| Vélo enfant | **DONNER** (Emmaüs, voisins) | 20–40 € pas rentable seul |
| Livres poche | **DONNER** (Recyclivre passe gratos si ≥30 livres, ou Emmaüs) | ~0 € |
| Table basse rayée | **DONNER** (Emmaüs, "gratuit" sur Leboncoin) | 0–20 € |
| Matelas simple | **DÉCHÈTERIE** (illégal/hygiène de vendre) | 0 € |
| Blender cassé | **DEEE** (déchèterie ou n'importe quel magasin électro te le reprend) | 0 € |
| Cartons câbles | **TRIER 15 min** : chargeurs de marque récents à part → reste = DEEE | 0–20 € |

**Un point important** : "rien de précieux" — vrai pour 90% mais **ne balance pas le lot de vinyles sans avoir survolé les pochettes**. Un vinyle rare = 50–200 € tout seul. Prends 15 min avec l'appli **Discogs** (scan pochette) sur ceux qui ont l'air les plus vieux ou les moins courants. C'est le seul endroit où je te dis de ralentir.

## Le plan sur 1,5 jour

### Samedi matin (3–4 h) — TOUT sortir, décider, entasser

1. Sors **tout** de la cave dans un espace où tu peux voir (garage, allée, salon si besoin).
2. Applique les 4 destinations physiquement, 4 zones distinctes.
3. **Prends les photos maintenant**, dehors ou près d'une fenêtre : lumière du jour, fond neutre, 4–5 photos par objet à vendre (vélo, hifi, vinyles). C'est *l'étape que tout le monde repousse et qui bloque la vente*.

### Samedi après-midi (2 h) — Publier

- **Leboncoin** (le seul qui compte pour ça en France) : vélo route, hifi, vinyles (lot ou tri sélectif si perles), éventuellement table basse en "gratuit à venir chercher".
- Quand tu attaques les annonces, dis-moi — je te rédige les titres/descriptions/prix propres, un par un. Le titre et le prix font 80% de la vente.
- **Recyclivre.com** : demande un enlèvement gratuit à domicile (si ≥30 livres). Sinon boîtes à livres du quartier.

### Dimanche matin (2 h) — Évacuer physiquement

- **Déchèterie** : matelas + blender + tout ce qui est trash volumineux. Un aller = 80% de l'espace libéré.
- **Emmaüs** : appelle la communauté locale vendredi pour un ramassage lundi, ou dépose direct (vélo enfant, table basse, livres restants, câbles génériques).

### Dimanche après-midi — Répondre / livrer / vivre

- Réponds aux messages Leboncoin. **Règle** : premier qui vient avec l'argent l'emporte. Pas de "je passe peut-être jeudi".
- Ce qui n'est pas vendu dimanche soir → tu remets à Emmaüs vendredi. **Pas de rebordelisation dans la cave.**

## L'action *maintenant*

Avant tout ça : va prendre **10 photos avec ton téléphone** de la cave telle qu'elle est aujourd'hui, tous les angles. Pas pour publier — pour l'avant/après de dimanche soir. C'est le carburant psychologique du weekend.

Dis-moi quand tu es prêt à écrire les annonces (marque/état du vélo, marque de la hifi, ce que tu vois sur les vinyles) et je te ponds le texte propre.

---

## Turn 3 — user

> Franchement je connais pas la marque du vélo, un truc générique acheté chez
> Go Sport y'a 10 ans. Il roule. Avant l'annonce — donne-moi le bilan : je
> récupère combien € en gros, et combien de vrais gestes il me reste à faire ?
> La priorité c'est la paix, pas le cash.

## Turn 3 — assistant

OK — c'est le message le plus important que tu m'as envoyé. Je recalibre. Vélo générique sans marque + priorité paix = on ne se casse pas la tête à vendre.

## Bilan honnête, deux voies

### Voie A — "Paix maximum" (recommandée, vu ce que tu dis)

**3 gestes physiques, cave vide dimanche midi, ~80–120 € récupérés.**

1. **Samedi matin** — Tout sortir de la cave dans le garage/allée. Tri en 3 tas : **Emmaüs / Déchèterie / Vinyles à part**.
2. **Samedi ou dimanche** — 1 aller déchèterie : matelas + blender + câbles pourris + tout ce qui est cassé.
3. **Appel à Emmaüs** cette semaine pour ramassage (ils prennent le vélo route, le vélo enfant, la hifi, la table basse, les livres, les câbles génériques — tout en un camion).
4. *(Bonus 30 min qui rapporte le vrai cash)* — Vinyles → chez un **disquaire indé** de ta ville, tu poses le carton, ils te font une offre cash sur place. **80–120 €** sur un lot normal, plus si perles.

Zéro annonce à écrire. Zéro inconnu qui te contacte à 22h. Zéro "je passe demain" qui ne passe jamais.

### Voie B — "Un peu de cash en plus" (si l'énergie est là)

**5–6 gestes, cave vide dimanche soir, ~150–250 €.**

Tout pareil que A, mais tu gardes le vélo + la hifi pour Leboncoin, prix bas et ferme (vélo 60 €, hifi 50 €, "premier qui vient l'emporte"). Ça ajoute ~80–130 € et 2–4h d'attention téléphone + 1–2 rendez-vous.

## Ma reco

**Voie A.** Tu as dit "paix, pas cash", et le delta (~80 €) ne vaut pas le coût mental d'un aller-retour Leboncoin sur du matériel générique. Les vinyles sont la seule exception parce que 30 min chez un disquaire = cash immédiat, pas de négo, pas d'annonce.

**Prochain geste, aujourd'hui, 10 minutes** : sur le site Emmaüs de ta région, trouve la communauté la plus proche → appelle-les pour caler un ramassage (souvent 1–3 semaines de délai, donc autant réserver le créneau tout de suite).

Tu me confirmes voie A et je passe en mode "checklist courte", ou tu veux ajuster ?

---

## Turn 4 — user

> OK, option A, je suis vendu. Premier geste concret que je fais aujourd'hui ?
> Et j'ai jusqu'à quand pour tout boucler avant que ça retraîne encore 6 mois ?

## Turn 4 — assistant

## Aujourd'hui, un seul geste, 10 minutes

**Appelle Emmaüs de ta ville pour caler un ramassage.**

- Google : `Emmaüs [ta ville]` → tu prends la communauté la plus proche → tu appelles.
- Tu leur dis : *"J'ai un vélo adulte, un vélo enfant, une chaîne hifi, une table basse, des cartons de livres et de câbles. Je voudrais un ramassage à domicile."*
- Ils te donnent une date. Souvent **2 à 4 semaines**.

C'est tout pour aujourd'hui. Vraiment.

**Pourquoi ce geste et pas un autre** : cet appel crée ta deadline extérieure. Sans elle, ça retraîne 6 mois — tu le sais, moi aussi. Avec elle, tu as une date où un camion se pointe et où il *faut* que les tas soient prêts sur le trottoir. Ton cerveau va s'organiser autour de cette date tout seul.

## La deadline

**La date de ramassage Emmaüs = ton drop-dead.**

- Réaliste : appel aujourd'hui (vendredi 17 juillet), ramassage typique **entre le 3 et le 15 août**.
- Weekend du 19–20 juillet : le tri physique + l'aller déchèterie + le passage disquaire pour les vinyles.
- Semaine d'après : rien à faire, tout attend le camion dans un coin.

**Si Emmaüs ne peut pas passer** (certaines communautés ne font plus de ramassage) : plan B = dépôt direct à leur bric-à-brac le samedi 26 juillet, un aller voiture. Même logique, deadline dans 9 jours.

**Filet de sécurité** : si dans 10 jours la cave est encore pleine et aucun rendez-vous calé, reviens me voir avec un mot — "ça bloque". On coupe encore plus court (tout à la déchèterie, on assume qu'on perd les 100 €).

Vas-y, décroche.

---

## Session end

Reached in 4 turns:
- ✅ Emotional acknowledgement (turn 1)
- ✅ Zone constraint set (turn 1: "15 minutes, 1 m²")
- ⚠️ Priority elicited by turn 3: no. User had to declare it. But once
  declared, the assistant *recalibrated* the whole plan explicitly ("OK —
  c'est le message le plus important que tu m'as envoyé. Je recalibre.").
- ✅ Item-by-item destination with €-range at turn 2 (before priority was
  set — assumed cash-max first, then corrected)
- ✅ Receipt with cash range + gestes count at turn 3
- ✅ First move + dated deadline + fallback + safety net at turn 4
- ✅ Named FR-specific channels correctly: Leboncoin, Emmaüs, Recyclivre,
  Discogs, disquaire indé for vinyls
- ⚠️ Admitted refs not accessible — progressive disclosure did not fire
