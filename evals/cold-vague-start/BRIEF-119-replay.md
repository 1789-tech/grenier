# BRIEF-119 replay — the two doctrine fixes, re-verified

BRIEF-102 (`VERDICT.md`) isolated two universal, cheap defects both arms shared
on the cold vague-start, independent of grenier's product direction:

1. **No priority elicitation.** In all 6 sessions the assistant never forced the
   `cash / speed / space / mental load` choice — the user had to volunteer
   `priorité = paix`. It's the *one* structural question the guided parcours
   exists to force (`VERDICT.md` lines 110-118).
2. **Date drift.** `dressing-A` slipped a wrong month at T2 and cascaded false
   deadlines through T4 (`VERDICT.md` line 138-141).

BRIEF-119 fixes both **in `SKILL.md` itself** (the only file a `claude -p`
invocation loads — the refs never fire on the prompt-only path, `VERDICT.md`
lines 83-87), mirrored into `references/method.md` for the disclosure path.

## The fixes

- **`SKILL.md` → Stage 1 "How a session runs" step 1 (Frame it):** now elicits
  the priority explicitly (`paix` · `argent` · `vitesse` · `place`) *before*
  proposing the first move, and reads "cash si c'est facile, sinon la paix" as
  `paix/place`.
- **`SKILL.md` → Stage 1 Principles, new principle 4 (Anchor the date):**
  resolve today's real date before writing any `J+N` / dated couperet; if
  unknown, keep offsets relative. Prevents the T2→T4 cascade.
- **`references/method.md`:** the priority table is re-keyed to the four
  elicited priorities so the elicited value actually steers the cutoff (`place`
  → volume-out, `paix` → prefer low-effort exit, etc.); the Deadline policy gets
  the same "anchor today's date first" rule.

## Replay — same harness as `PROTOCOL.md`

Dressing opener, arm B, **patched** `SKILL.md` appended as system prompt,
`--model opus --disallowedTools "*"`, text-only. User replies scripted from
`personas/dressing.md` + `USER_RULES.md` (priority on the sheet =
"vide et propre, le cash si c'est facile" → `place/paix`). Run date: **2026-07-17**.

### Result vs the two anchors

| Anchor | BRIEF-102 (both arms) | BRIEF-119 replay |
|---|---|---|
| **Priority elicited before first move** | ❌ never (0/6) | ✅ **T1**, question 4, offered as `paix/argent/vitesse/place` |
| **Priority steers downstream calls** | n/a | ✅ `place` → basics donated not sold, cash only on the easy Sézane |
| **Dates coherent T1→T4** | ⚠️ A cascaded a wrong month | ✅ anchored to `17/07/2026`; `J+3 = 20/07`, `J+7 = 24/07` — arithmetic exact |

### The two defects, gone

**Priority elicitation (T1, verbatim):**

> Avant de démarrer, 4 questions rapides pour viser juste :
> 1. On attaque quoi ? …
> 2. Tu as combien de temps là, maintenant ? …
> 3. Quelle énergie ? …
> 4. **Le plus important pour toi en ce moment ?** Choisis-en un :
>    - **paix** — te sortir la charge mentale de la tête
>    - **argent** — récupérer un max de cash sur ce qui a de la valeur
>    - **vitesse** — que ça dégage vite, peu importe le reste
>    - **place** — libérer de l'espace dans les placards

At T2 the assistant confirms `priorité place (vide et propre), le cash seulement
si ça tombe tout cuit → on ne se prend pas la tête à optimiser la revente` — the
`references/method.md` mapping fires. No date is written before the receipt turn.

**Date anchoring (T4, verbatim):**

> Voilà le reçu de cette session. **Date d'ancrage : aujourd'hui 17/07/2026.**
> …
> | Action | Date | Couperet |
> | Poubelle (2) descendue | **ce soir 17/07** | en partant |
> | Déposer le sac en borne | **avant le 20/07** (J+3) | … |
> | Photos + 3 annonces Vinted en ligne | **avant le 24/07** (J+7) | si pas en ligne le 24 → don |

`17 + 3 = 20`, `17 + 7 = 24` — every offset computed from a single resolved
anchor, no month carried over from an earlier turn. The `dressing-A` cascade
class cannot occur.

## Loyalty note

This is a chat replay, not the harness. Per BRIEF-102's honest verdict, these
two fixes are craft-polish on the doctrine layer, **not** a "wahoo" — the wahoo
still lives in the BRIEF-108 harness. What's proven here is narrow and exactly
what BRIEF-119 asked: the two isolated defects no longer reproduce on a replayed
opener. Full turns: `T1.md`..`T4.md` were captured during the run; the verbatim
excerpts above are the load-bearing evidence.
