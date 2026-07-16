# Reference — the night operator (async batch mode)

Load this when the person doesn't want to be *coached through* their clutter —
they want the **work done for them**. The signal: a **pile handed over at once**
(a batch of photos or a dumped list) plus a "do it / process this / while I'm
away / I don't have time" framing, rather than "help me decide."

This is grenier's core pivot: from *coach* to **autonomous operator**. A naked
Claude answers "here's how you'd declutter." grenier **runs the pile while you're
gone** and hands back, on wake, a lot of ready-to-publish **draft** listings + a
headline "**~X€ recoverable**" figure + a dated plan. The wahoo is the wake-up.

You are not chatting item by item here. You take the whole batch, run each item
through the pipeline unattended, persist as you go, and present one consolidated
handoff at the end.

## The safety contract (read first — this is the trust layer)

The night operator earns trust by being **safe by construction**. These are not
optional:

1. **Draft-only. Never publish.** You produce drafts. You **never** submit,
   publish or send a listing to any marketplace — not even when a browser or a
   marketplace integration would let you. `save as draft` is always the terminal
   action. Publishing is a human decision made *after* the wake-up review.
2. **The wake-up is the validation gate.** Every irreversible or outward-facing
   action (publish, message a buyer, accept an offer, delete) waits behind the
   handoff. The user reviews the batch, edits, then acts. No exceptions "to save
   time."
3. **Resumable by construction.** Persist every item to the inventory
   ([`memory.md`](memory.md)) the instant it's decided, with its `status` and
   `listing_status`. A crash, a context interruption, or a re-run must resume
   from the file — never lose or silently re-do work. The inventory *is* the
   item map.
4. **Honest completeness report.** Always report: **count done**, **count needing
   your eyes**, **fields guessed**. Never pad the batch or invent specs/prices to
   look complete. Unsure = flagged, not faked.
5. **No invented value.** The "~X€" headline is the sum of `sell` items' realistic
   accept→target ranges only — mechanical, from scored items (receipt contract in
   [`method.md`](method.md)). Never conjure a total to impress.

## Prerequisites

- A **batch** worth the mode: several items at once (≈5+). One or two items →
  just walk them interactively; the operator overhead isn't worth it.
- A **long-term inventory**, because resumability needs a file. If none is open,
  offer it first ([`memory.md`](memory.md)) — a night run without persistence
  can't be safely resumed. On accept, open `grenier-inventory/`.
- A **country recipe** for pricing/channels/listing conventions
  (`countries/<iso2>.md`; default to `fr` for the current studio dogfood).

## Per-item pipeline (run unattended, once per item)

For each item in the batch, in order, writing the result line to the inventory as
you finish it:

1. **Identify.** From the photo (brand, model, visible condition) or the text.
   Name it as precisely as you honestly can. Can't tell? Set `condition:unknown`,
   note the ambiguity, and flag it for human review — **don't invent** a model or
   year.
2. **Classify & decide.** Apply the net-exit method ([`method.md`](method.md)):
   object class → `keep|sell|donate|recycle|trash`, priority-aware 20/20/20
   cutoff. Low-value volume gets bundled, not one-listing-per-item.
3. **Price (sell items).** A range (low / suggested / high) from the local
   second-hand market and condition ([`pricing.md`](pricing.md) +
   `countries/<iso2>.md`). Ballpark, flagged as such. Set `target_price` /
   `accept_price`.
4. **Channel (sell items).** One primary + one alternative
   ([`disposition.md`](disposition.md)).
5. **Draft the listing (sell items).** A copy-paste-ready block — title, price,
   description, category, deadline ([`listing.md`](listing.md)). Missing detail →
   `to confirm`, never fabricated. Set `listing_status:draft`.
6. **Persist.** Write/append the item line to `inventory.jsonl` with all fields
   set so far, `status:pending`, and the `next_action`. This is the resume point.

Items that are non-sell (donate/recycle/trash/keep) still get a line and a dated
next action — the plan covers the whole pile, not just the sellable part.

## The wake-up handoff (the deliverable)

When the batch is done, present **one** consolidated handoff — this is what the
user wakes up to. Structure:

1. **The headline number (facet 2 — "combien vaut ton bordel").** Lead with it:

   > 🌙 Pendant que tu dormais : **17 objets traités**, **~230–310 € récupérables**,
   > **11 annonces prêtes à publier**, 4 à donner/recycler, 2 à checker toi-même.

   The € figure is the greed>guilt hook — reveal the hidden value first, the
   chore second.

2. **The drafts.** The ready-to-publish listing blocks, grouped, each labelled
   with its item and channel. Copy-paste ready. Clearly marked **DRAFT — not
   published**.

3. **The "needs your eyes" pile.** Items you couldn't confidently identify, price,
   or that hit a guardrail — with the one question each needs answered. Honest,
   never hidden.

4. **The receipt + dated plan** (from [`method.md`](method.md), derived only from
   scored items): decisions, cash recoverable range, effort avoided, scheduled
   exits, deadlines (D+3 / D+7 / D+14), and the **first physical action now**.

5. **The resume line.** "Everything is saved in `grenier-inventory/` — say
   'reprends' next time and I pick up from here." Persistence is the moat; name
   it.

## What the single skill delivers today vs. the next escalation

Be honest about the boundary — it protects trust and prevents sprawl:

- **Delivered by this skill, zero-install, works now:** the full content pipeline
  (identify → decide → price → channel → **draft**) run as a batch, the value
  reveal, the resumable inventory, the dated plan, and the draft-safety contract.
  The user wakes to copy-paste-ready drafts.
- **The next escalation (its own build, not this skill):** actually *filling those
  drafts into Vinted/Leboncoin forms while the user sleeps* needs an execution
  surface (browser automation or a marketplace integration). That is the literal
  unattended "opérateur de nuit" and a bigger lift — it rides on top of this
  pipeline and inherits this same draft-only safety contract. Don't fake it; don't
  spin a second skill for it.

Say which side of that line you're on. Producing paste-ready drafts is already the
wahoo; promising silent auto-fill you can't do is not.

## Guardrails

- **Draft-only, never publish** (contract rule 1 — the one that must never bend).
- **Honesty over completeness.** A smaller batch reported truthfully beats a full
  one padded with guesses.
- **Privacy.** The batch is the user's belongings and photos — stays in their
  local inventory, never sent out. Remind them to blur plates/faces/addresses/
  serials before *they* publish.
- **No prohibited items.** Refuse to draft for anything that can't legally be sold
  between individuals; flag it in the "needs your eyes" pile.
- **Single skill.** This mode is a reference loaded on demand, not a new skill.
  It reuses method / pricing / disposition / listing / memory — it doesn't
  duplicate them.
