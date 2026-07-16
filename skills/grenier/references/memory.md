# Reference — persistent inventory memory (file, not a database)

Load this when the person wants to **resume** an inventory, **track statuses**
over time ("what did I still have to sell?", "pick this back up tomorrow"), or
accepts the long-term régime. It is also the storage layer the night operator
writes to ([`night-operator.md`](night-operator.md)).

grenier's long-term memory is a **local folder of plain-text files, not a
database** — offline, free, versionable, and it stays on the user's machine. One
artefact, two régimes: *à l'arrache* (we never open a file) and *long-term* (the
file grows). No new skill, no server, no account, no spend beyond Claude.

## Two régimes — don't auto-create memory

Creating a file on the first object turns a simple conversation into a project.
Don't.

- **À l'arrache (default).** A small one-off lot — under ~10 items, single
  session, no need to track sales/statuses. Stay in conversation. No file.
- **Propose long-term** the moment a signal appears: **≥10 items**, several
  rooms, a "to sell" pile that will carry statuses, an explicit ask to resume
  ("pick up tomorrow", "where was I?"), photos/prices/dates worth keeping, or a
  session producing **more than ~5 decisions** worth persisting.
- **Always an explicit choice before any write.** grenier proposes, the user
  accepts. **Never a surprise file.**

Exact phrase to offer it (cold):

> On peut rester en mode à l'arrache pour cette session, ou ouvrir un inventaire
> long-terme : je crée un dossier local `grenier-inventory/` avec un fichier
> texte `inventory.jsonl` que tu peux lire, modifier, sauvegarder ou supprimer.
> Ça me permettra de reprendre demain sans que tu recolles tout. Tu veux que je
> l'ouvre ?

Short version if already mid-flow:

> Là on dépasse le mode jetable. Je te propose d'ouvrir
> `grenier-inventory/inventory.jsonl` pour garder les décisions et reprendre
> sans recoller le contexte. OK ?

## Platform split — headline it, don't oversell

- **Claude Code / Desktop (default target).** Real local filesystem: grenier
  reads and writes `grenier-inventory/` directly. This is where memory actually
  beats copy-paste.
- **claude.ai (degraded fallback).** No direct local FS. grenier generates
  `inventory.jsonl` as a downloadable file and the user re-uploads it next
  session. Usable, but it is manual re-attach — the very chore memory is meant
  to kill. **Say so plainly; never sell it as seamless.**

## Layout

```text
grenier-inventory/
├── inventory.jsonl     # machine mirror — one item per line, the source of truth
├── inventory.md        # human-facing surface — grouped, readable, for the user
└── archive.jsonl       # optional, later — done/exited items moved out of the way
```

`inventory.jsonl` is the machine mirror; a 15-field single-line JSON object is
the *least* hand-editable format for a non-dev (one bad comma kills the line), so
**grenier owns the writes** and the user reads/edits the friendly `inventory.md`
surface. Regenerate `inventory.md` from the JSONL after each write. The file is
not a "base"; it is a deliberately banal user artefact.

## Item schema (v0)

One item = one JSON line. Flat fields; human detail goes in `notes`.

```json
{"id":"grn-20260704-0001","name":"vélo enfant Decathlon 20 pouces","condition":"good","room":"cave","decision":"sell","channel":"leboncoin","target_price":65,"accept_price":55,"listing_status":"draft","status":"pending","created_at":"2026-07-04","updated_at":"2026-07-04","next_action":"prendre 4 photos","deadline":"2026-07-11","notes":"freins OK, pneus à vérifier"}
```

- `id` — stable, `grn-YYYYMMDD-####` (or a short slug on manual import).
- `name` — human name.
- `condition` — `new|like_new|good|fair|poor|broken|unknown`.
- `room` — room/zone.
- `decision` — `keep|sell|donate|recycle|trash|undecided`.
- `channel` — chosen channel or `none`.
- `target_price` — start price, number or `null`.
- `accept_price` — realistic floor, number or `null`.
- `listing_status` — `none|draft|listed|messaged|sold|expired`.
- `status` — `pending|in_progress|done|blocked`.
- `created_at`, `updated_at` — ISO date.
- `next_action` — next physical action.
- `deadline` — ISO date or `null`.
- `notes` — short text.

Useful later, not required v0: `currency`, `country`, `photo_refs`, `sold_price`,
`completed_at`.

## Mechanics

1. **Detect.** On "resume / inventory / track / sold / list" or the thresholds
   above, offer the long-term phrase. Never write first.
2. **Open (on accept).** Create `grenier-inventory/` with an empty
   `inventory.jsonl` and a one-line `inventory.md` header. Confirm the path.
3. **Read at session start.** Load `inventory.jsonl`, summarise **only the
   active/grouped items** (pending, in_progress, drafts awaiting action), then
   ask for the next batch. Don't dump the whole file into the reply.
4. **Write after each confirmed decision.** Append or update the item's line. On
   Claude Code with FS access, rewrite the file with the updated lines; on
   claude.ai, emit a "add/replace this line" block the user saves. Regenerate
   `inventory.md`.
5. **Close the session.** Show a receipt derived **only** from the file:
   decisions made, realistic cash of `sell` items, dated actions, blocked items.
   No invented numbers outside item lines (same receipt contract as
   [`method.md`](method.md)).
6. **claude.ai fallback.** No local FS → generate `inventory.jsonl` as a
   downloadable file or ask the user to re-upload the existing one. Usable, not
   seamless — say so.

## Anti-accumulation ritual (the "sell it while it's still worth something" ping)

The persistent file is what a naked Claude can't have — use it. When resuming a
long-term inventory, don't just list statuses; **surface decay**:

- **Aging drafts.** `listing_status:draft` older than ~14 days → "still sitting
  as a draft, publish or drop it."
- **Depreciating categories.** Electronics, phones, seasonal gear, kids' items
  outgrown → flag that the realistic price erodes monthly; nudge to act now.
- **Passed deadlines.** `deadline` in the past and not `done` → move to the
  D+7/D+14 exit route from [`method.md`](method.md) (donate/give away), don't
  re-list forever.
- **Seasonal windows.** Ski gear in autumn, garden furniture in spring — a
  one-line ping when the window opens beats a dormant line.

Keep it to a short, dated "worth acting on now" shortlist at the top of the
resume summary — not a nag on every item. The ritual is: **the inventory pings
you, you don't have to remember.**

## Guardrails

- **Privacy first.** The inventory is the user's private belongings — it lives
  only in their local folder, never sent anywhere. It is *not* the anonymised
  studio feedback (see `FEEDBACK.md`); the two never mix.
- **No surprise files, no silent writes.** Every create/write is a confirmed
  action the user asked for.
- **The user owns the file.** They can read, edit, zip, git, or delete it. Treat
  it as theirs, not grenier's.
