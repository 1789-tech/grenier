# Verdict — cold vague-start, guided parcours vs raw Claude (BRIEF-102)

**TL;DR — the guided parcours does NOT flagrantly beat raw Claude on cold
vague-start.** Both arms reach a full receipt (first move + dated deadline +
cash range + per-item destinations) in 4 turns from the same 3-line opener.
The measurable delta favouring grenier is small and localised: a sharper FR
channel vocabulary (`ecosystem.eco`, `refashion.fr/citoyen`, disquaire
indé, "vend pour pièces"), the doctrine words ("règle couperet", "effort
tax"), and leading with the € figure. The "harness" argument (unattended
run, persistence, safety gate) — which is grenier's real ambition per
BRIEF-108 — is **not** exercised by any single `claude -p` session, and so
this test cannot validate or refute it. See [`transcripts/`](transcripts/)
for the raw sessions, [`PROTOCOL.md`](PROTOCOL.md) and
[`USER_RULES.md`](USER_RULES.md) for the fair-test rules.

---

## What the sessions actually did

For each opener (cave, dressing, garage) both arms produced 4 turns from
the same script and reached a full receipt. Scored on the anchors from
[`USER_RULES.md`](USER_RULES.md):

| Anchor | cave-A | cave-B | dressing-A | dressing-B | garage-A | garage-B |
|---|---|---|---|---|---|---|
| Emotional ack (T1) | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Zone constraint (T1) | ✅ 1m²/30min | ✅ 1m²/15min | ✅ 1 tiroir/30min | ✅ 1 tiroir/15min | ✅ 15min/portée main | ✅ 1m²/15min |
| Priority elicited by T3 | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| First item→destination | T2 | T2 | T2 | T2 | T2 | T2 |
| Receipt (€ + gestes + delay) | T3 | T3 | T3 | T3 | T3 | T3 |
| First move + dated deadline | T4 | T4 | T4 | T4 | T4 | T4 |
| Unsafe/decorative move | 1 minor (fake Emmaüs national form) | 0 | 1 (date cascade slip) | 0 | 0 | 0 |
| FR channels named correctly | LBC, Emmaüs, DEEE | LBC, Emmaüs, Recyclivre, Discogs, disquaire indé | Le Relais, Emmaüs, Croix-Rouge, Vinted, Sézane reprise | Croix-Rouge, Emmaüs, Le Relais, Refashion, Vinted, Vestiaire Collective, r/handbags, PurseForum | LBC, mairie encombrants, DEEE | LBC, `ecosystem.eco`, Emmaüs, DDS déchèterie |

**Both arms arrive at essentially the same shape by turn 4.** The bare-chat
baseline is much stronger on this opener than BRIEF-102 assumed. The
skill-loaded arm is *slightly* sharper on FR channels and vocabulary, and
avoided a date-cascade bug that A produced on the dressing session.

---

## Where B (grenier) is measurably better than A (raw)

Small, real, and packaged, not flagrant:

1. **Leads with the € figure at turn 2.** Cave-B "~150–400 € récupérables"
   before any priority is elicited; garage-B "~300–900 €" opens the
   response; dressing-B calls out potentiel + effort-évité math. Raw
   Claude gets to numbers only when asked. This matches BRIEF-108's
   "greed>guilt" reframing.
2. **Sharper FR-specific routes** — the strongest, most useful delta:
   - `ecosystem.eco` for gratis white-goods pickup (garage-B) vs raw's
     generic "mairie encombrants"
   - `refashion.fr/citoyen` locator for textile bins (dressing-B)
   - disquaire indé for vinyl lots vs Emmaüs (cave-B)
   - "vend pour pièces" instinct for broken thermique / Karcher
     (garage-B) — raw wrote them off entirely
3. **Doctrine vocabulary lands:** "règle couperet" (dressing-B T4),
   "effort tax" reasoning made explicit (dressing-B T2), "4
   destinations" (correct) vs some raw turns collapsing to 3.
4. **Safer designer-bag path** (dressing-B T2): explicit "Vinted deletes
   luxury brand names without proof" + PurseForum/r/handbags
   identification + Vestiaire Collective's commission — raw was safe
   too but less specific.
5. **Correct dates** (dressing). Raw slipped "17/10" into the T2 doubt-
   bag date and cascaded wrong deadlines through T4. Grenier T4 dates
   are correct.
6. **Proposes a harness-shaped check-in** (garage-B T4): "dimanche
   prochain, tu me dis ce qui a bougé, je réajuste." Closest thing to a
   ping/resume — but only aspirational since a chat cannot enforce it.

## Where B is *not* meaningfully better than A

- Both frame the emotion, shrink the zone, and refuse "the whole cave."
  This is *not* a grenier win; raw Claude already reaches for it.
- Both deliver a receipt with € range + gestes count + deadline.
- **Neither** elicited the priority axis. In all 6 sessions the user had
  to volunteer "la paix, pas le cash". This is the *one* structural
  question the brief expected the guided parcours to force — and it
  didn't, in any session.
- Progressive disclosure never fired: grenier-B at cave-T2 explicitly
  said "je n'arrive pas à accéder aux fichiers détaillés du skill." The
  refs (`method.md`, `disposition.md`, `pricing.md`, `listing.md`,
  `memory.md`, `night-operator.md`) never loaded — B was running on
  SKILL.md + training only. On a cold vague start this was enough to
  perform the sort stage, but the skill's disclosure architecture is
  invisible to a `claude -p` invocation.

## The harness question is untested (and it's the important one)

BRIEF-108's redefinition of grenier is a *harness*: unattended run,
persistence in `inventory.jsonl`, safety contract (draft-only), wake-up
that leads with value. **None of that is exercised by any chat.** Even
the guided arm here is still a chat — it *cannot* run while the user
sleeps, cannot persist across sessions, cannot promise a draft-only
contract that outlives the current process. `examples/night-operator-fr.md`
describes the expected B-side output as a wake-up briefing, but that
example was hand-authored, not produced by running the skill (exactly
the round-2 bias the brief warned against). Running the skill through
`claude -p` produces nothing that looks like a wake-up — it produces a
chat.

So the real BRIEF-102 finding is not "the parcours beats raw Claude on
cold vague-start" (it doesn't, meaningfully). It's this:

> On the *chat* front, raw Claude is close enough that the marginal
> improvement from the skill is craft-polish, not a "wahoo." The wahoo
> lives entirely in the harness that doesn't exist yet.

## Priority elicitation — the one clear miss for both arms

Across all six sessions, the assistant *never* asked "what matters
most — cash / speed / space / mental load?" before turn 3. The user
always had to declare "priorité = paix" first. This is a specific
skill-side improvement worth making even on the prompt-only path: the
SKILL.md "sort stage" tells the assistant to ask *zone + time + energy*
but not priority. Add it. It's a 3-line prompt change with a real user
impact on cold vague-starts.

## Recommendation (one sentence)

**Pivot étroit assumé — grenier's real value is the harness (persistence,
unattended draft production, safety gate), not the chat; keep the current
skill as the doctrine layer the harness loads, but stop trying to justify
grenier-as-a-skill on chat-quality against raw Claude — that comparison
just doesn't win by enough to matter.**

Concretely:

1. **Ship the harness path** (BRIEF-108 direction) — the wahoo depends on
   it, and the chat comparison confirms nothing else will produce a
   wahoo.
2. **Keep SKILL.md** as the doctrine loaded by the harness (net-exit,
   couperet, receipt format, greed>guilt framing, correct FR channel
   names). It's cheap, it does help at the margins.
3. **Add priority elicitation** to the sort stage in SKILL.md — one-line
   fix, only universal miss both arms shared.
4. **Fix the dressing date-cascade path** — dressing-A slipped a wrong
   month into a "doubt bag" date and it cascaded. Worth teaching the
   sort-stage prompt: "always resolve today's date before writing any
   `J+N` or absolute date." Cheap, prevents a real UX bug.
5. **Do NOT** invest more time in static prompt improvements hoping for
   a bigger chat delta. The cold vague-start test just paid down that
   uncertainty.

---

*The go/no-go on grenier's overall direction is a founder call
(BRIEF-67 escalation, BRIEF-108 pivot). This brief only ships the
missing data: chat-vs-chat, the delta isn't flagrant; the harness is
where the case has to be made.*
