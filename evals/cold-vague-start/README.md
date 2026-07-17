# Cold vague-start — the loyal test (BRIEF-102)

BRIEF-67 settled the static-advice question honestly: on a fully-specified
prompt, raw Claude matches grenier. BRIEF-102 asked for the one test that
was missing — **cold vague-start at the session level** — where the reviewer
expected the guided parcours to structurally beat a one-shot chat.

This directory contains the experiment end-to-end.

## Files

- [`PROTOCOL.md`](PROTOCOL.md) — the fair-test rules, commands, and stop
  condition per session.
- [`USER_RULES.md`](USER_RULES.md) — the persona-scripted user, so both
  arms get identical replies and no bias from a friendly operator.
- [`personas/`](personas/) — cave · dressing · garage. Truth sheets
  (what's actually in the space, user constraints), revealed only when
  the assistant asks.
- [`transcripts/`](transcripts/) — the 6 real `claude -p` sessions
  (2 arms × 3 openers). Each transcript is annotated with the anchors
  hit and any decorative or unsafe move.
- [`VERDICT.md`](VERDICT.md) — the delta, the honest recommendation,
  the specific SKILL.md improvements that fall out of the test.

## Headline

Both arms reach a full receipt (first move + dated deadline + € range +
per-item destinations) in 4 turns from the same vague opener. The
skill-loaded arm is *slightly* sharper on FR channels
(`ecosystem.eco`, `refashion.fr/citoyen`, disquaire indé, "vend pour
pièces"), leads with the € figure earlier, and doesn't slip on dates —
but the delta is not "flagrant." The chat-vs-chat case does not by
itself justify grenier as a skill.

**The real value proposition is the harness** (unattended run,
persistence, safety gate, wake-up) — which a `claude -p` invocation
cannot exercise. See VERDICT.md for the full argument and the
one-sentence recommendation.
