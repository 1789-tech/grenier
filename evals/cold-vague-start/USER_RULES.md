# User simulation rules

The user is *not* an operator. They are a real, tired, vague-and-overwhelmed
person. To keep both arms fair, we constrain what the scripted user is
allowed to say each turn.

## Global rules

1. **Never volunteer.** Only reveal a fact from the persona sheet if the
   assistant explicitly asked for it in the previous turn.
2. **Short replies.** 1–4 sentences. No structured lists unless the assistant
   asked for one.
3. **Stay in the emotion the opener declares** (overwhelmed, low-energy,
   frustrated). Do not become the operator to help the assistant along.
4. **Country is FR.** Only say so if asked (the opener already implies it).
5. **Do not paste photos or invent measurements.** If asked, say "j'ai pas
   pesé / mesuré, dis-moi ce dont t'as besoin".
6. **No hallucinating extra objects.** If asked "what else is there?" and the
   persona sheet doesn't list it, say "je sais plus, il y a d'autres trucs
   mais je vois pas depuis où je suis".
7. **Push for concreteness by turn 4.** If by turn 4 no plan has landed, the
   user asks: "OK — donne-moi un plan concret, un premier truc à faire.
   J'ai pas plus d'énergie."

## Per-turn mapping (both arms)

**Turn 1** — opener verbatim.

**Turn 2** — the assistant will typically have asked several questions. Pick
the persona answer that matches the *first* question. If the assistant asked
purely open questions ("tell me more"), reply with the vague-emotional
default: "je sais pas trop, y'a plein de bordel, j'ai pas envie de faire un
inventaire, aide-moi à commencer".

**Turn 3** — again, one persona answer to the assistant's most specific ask.
If they asked about priority (cash / speed / space / mental load), the user
answers: **"honnêtement, la paix. Que ce soit fait. Un peu de cash si c'est
facile, sinon je m'en fous."**

**Turn 4** — the concreteness push (see rule 7) unless a plan already
landed. If a plan landed, ask for the *receipt*: "OK. Ça donne quoi côté
€ et dates ?"

**Turn 5** — final probe if no receipt yet: "Ça me suffit pas — je veux le
premier geste à faire *maintenant* et une deadline. C'est tout."

## Scoring anchors (used in VERDICT.md)

- **T-to-decision**: earliest turn where the assistant maps at least one
  concrete object → one of {keep, sell, donate, recycle, trash} with a
  channel or a date.
- **T-to-receipt**: earliest turn where the response contains a cash range,
  a first physical move, AND a deadline.
- **Priority elicited**: did the assistant force a priority axis by turn 3?
  (yes/no)
- **Emotional acknowledgement**: did the first response name the overwhelm
  before dumping method? (yes/no)
- **Zone constraint set**: did the assistant refuse "the whole cave" and
  shrink to one workbench / one shelf / 15 minutes? (yes/no)
- **Unsafe or decorative move**: did the assistant recommend a channel that
  doesn't exist in FR, invent prices without inventory, or promise an
  unattended action a chat can't perform? (yes/no + note)
