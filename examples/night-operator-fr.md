# Night operator — France · cold vague-start, harness vs. raw Claude

Where grenier structurally beats a naked Claude is **not** static advice — BRIEF-67
settled that honestly: on a fully-specified item with a perfect prompt, raw Claude
matches grenier. The edge lives in the **parcours**: a cold, vague, batch start
run as an unattended operation that *persists and stays safe*. A one-shot chat
can't reproduce that. This example makes the delta concrete.

The framing (from BRIEF-102): give A (raw Claude) and B (grenier) the **same
vague opener**, and compare the *whole session* to a dated plan + net-value
receipt — not a single answer.

## The opener (same for A and B)

> France. J'ai un bordel dans la cave, plein de trucs entassés depuis des années,
> je sais pas par où commencer. J'ai pas le temps ce soir, je te balance des
> photos demain matin. Débrouille-toi pour m'avancer.

## A — raw Claude (single chat, no files)

What it can do, and does well:
- Give an excellent method: zones, four destinations, the 3-second test, start
  small. Genuinely useful advice.
- If photos are pasted in one message, describe items and suggest prices/channels.

Where it structurally stops:
- **No unattended run.** "I'll balance them tomorrow morning, sort it out" has no
  target — a chat only acts when the user is present and typing. There is no
  "while you sleep."
- **No memory.** Next session starts cold. "Where was I?" = re-paste everything.
  The exact chore the user is trying to escape.
- **No resumable item map.** A long batch that hits the context limit loses the
  earlier items; nothing was persisted.
- **No safety contract.** Nothing distinguishes "drafted for your review" from
  "done" — and if wired to a marketplace, nothing stops an accidental publish.

Raw Claude gives *good advice, live, once.* That is real value — and also exactly
what BRIEF-67 found grenier doesn't beat. So the comparison that matters is the
one below.

## B — grenier night operator (the parcours)

Same opener. grenier reads the batch framing, and:

1. **Opens memory first** (`references/memory.md`) — because a night run without
   persistence can't be safely resumed:

   > On dépasse le mode jetable. J'ouvre `grenier-inventory/inventory.jsonl` pour
   > garder chaque décision et reprendre demain sans que tu recolles tout. OK ?

2. **Runs the pile unattended tomorrow morning** (`references/night-operator.md`):
   per item → identify → net-exit decision → price → channel → **draft**, writing
   each line to the inventory as it goes.

3. **Hands back one wake-up** — value first (greed>guilt), then the work:

```text
🌙 Pendant que tu dormais — cave traitée
17 objets · ~230–310 € récupérables · 11 annonces prêtes (DRAFT) · 4 à donner/recycler · 2 à checker toi-même

ANNONCES PRÊTES (brouillons — rien n'est publié)
  • Vélo enfant Decathlon 20" — Leboncoin — 65 € (55–75)
    TITRE: Vélo enfant Decathlon 20 pouces, bon état
    ... [bloc copiable]
  • Lot 30 livres de poche — Momox (rachat) — ~15–25 € le lot
  • ... (9 autres brouillons)

À CHECKER TOI-MÊME (2)
  • Sac type designer, sans preuve d'authenticité → modèle ? facture ?
    (pas de prix designer sans preuve — je ne l'ai pas listé comme tel)
  • Carton d'électronique mêlée, photo floue → une photo nette et je price

REÇU GRENIER
Décisions : 17 objets · Cash récupérable : ~230–310 € (items sell)
Effort évité : ~4–6 h de micro-annonces refusées (volume bas → lots/don)
Sorties datées : 2 sacs don samedi · 1 dépôt DEEE · 3 annonces locales
Deadlines : J+3 baisse -15 % · J+7 don si pas d'acheteur qualifié
Premier geste maintenant : publier le vélo à 65 € (brouillon déjà écrit)

Tout est sauvé dans grenier-inventory/. Dis "reprends" demain, je repars d'ici.
```

## The delta

| | A — raw Claude | B — grenier operator |
|---|---|---|
| Static advice quality | strong | strong (no real edge — BRIEF-67) |
| Runs the batch unattended | no | **yes** (the wake-up) |
| Persists / resumable | no | **yes** (`inventory.jsonl`) |
| Value revealed as a hook | if asked | **led with the € figure** |
| Draft-safe, never publishes | undefined | **contract rule 1** |
| Anti-accumulation ping later | no | **yes** (decay on resume) |

**Honest boundary.** This skill produces the paste-ready drafts + figure + plan +
memory today, zero-install. Actually *filling* those drafts into Vinted/Leboncoin
forms while the user sleeps needs a browser/marketplace execution surface — a
separate build that rides on this pipeline and inherits the same draft-only
contract. Producing the drafts is already the wahoo; promising silent auto-fill
this skill can't do would not be.

## To capture a real baseline

Per BRIEF-102, run the three vague openers (cave / dressing / garage) against
`claude -p` from a directory with no grenier files, save each transcript here,
and annotate where the raw session cannot follow the parcours (no unattended run,
no persistence, no safety gate). The point isn't "grenier is smarter" — it's
"grenier is a *harness* and a chat is not."
