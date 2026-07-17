# Cold vague-start — protocol (BRIEF-102)

BRIEF-67 settled the static-advice question honestly: raw Claude matches
grenier when the prompt is fully specified. BRIEF-102 asks for the missing
test — **cold vague-start at the session level** — where the reviewer expected
the guided parcours to structurally beat a one-shot chat.

This directory captures that experiment end-to-end so anyone can rerun it.

## What we're testing

Not "who writes better advice". The static-advice comparison already exists
(`examples/method-delta-fr.md`). Here the question is narrower:

> Given the SAME 3-line vague opener, does the grenier-guided session reach a
> **dated action plan + net-value receipt** more cleanly than a bare Claude
> chat — measured by turns to decision, elicitation quality, emotional
> handling, and whether we actually arrive at receipt + first move?

## Fair-test rules

- **Same opener, same persona, same country, same energy** — for both arms.
- The user is a scripted persona sheet. Facts are only revealed *when the
  assistant asks*. No volunteering. This models "vague, overwhelmed" honestly.
- **No hand-authored grenier output.** The B arm is a real `claude -p`
  invocation with the skill's `SKILL.md` appended as system prompt. Whatever
  the model actually produces is what we score.
- **No hand-authored user replies.** Each user turn is picked from the persona
  fact sheet based on what the assistant asked in the previous turn, using
  the rules in `USER_RULES.md`.
- Both arms have the same tool budget (**text-only, no bash/read**), because
  the naked-chat baseline is a text conversation.
- Both arms use `--model opus` and empty working directories with no
  `CLAUDE.md`.

## Commands

```bash
# Arm A — bare Claude, no skill.
cd /tmp/brief102/<opener>-A
claude -p --model opus --disallowedTools "*" "<opener>"
claude -p --model opus --disallowedTools "*" -c "<user reply 2>"
# ... continue until receipt or 5 turns

# Arm B — grenier skill loaded as system prompt.
cd /tmp/brief102/<opener>-B
claude -p --model opus --disallowedTools "*" \
  --append-system-prompt-file /path/to/skills/grenier/SKILL.md "<opener>"
claude -p --model opus --disallowedTools "*" \
  --append-system-prompt-file /path/to/skills/grenier/SKILL.md -c "<user reply 2>"
# ... continue until receipt or 5 turns
```

Progressive disclosure (references/*.md) is **not** loaded automatically —
that's a `claude` runtime detail this test doesn't grant. If B needs deeper
knowledge and asks for it, we note that as a limitation. On a cold vague
start we expect the sort-stage content (in `SKILL.md` itself) to be enough.

## Stop condition per session

- The assistant delivers a receipt (decisions + cash range + first move + a
  dated deadline), OR
- 5 user turns elapse without one — recorded as a miss.

## What we compare

Per session (see `VERDICT.md`):
- Turns to first structured decision (destinations named, not just method).
- Turns to a receipt with a first physical move and a date.
- Elicitation quality: did the assistant force a priority choice
  (cash / speed / space / mental load) and a zone constraint?
- Emotional handling: does the assistant frame low-energy / overwhelm before
  drowning the user in options?
- Any decorative or unsafe move (e.g. suggesting sale of a "designer" bag
  without proof; naming platforms that don't operate in France; etc.).
