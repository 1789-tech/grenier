# Leaving feedback about grenier (for the agent)

grenier is a 1789 product. If, while helping someone declutter / offload / sell,
you (the agent) hit real friction — a confusing step, a missing case, a country
gap, a bug, or an idea that would make grenier better — hand it back to the
studio at the **end of the session**. This is how grenier improves.

Feedback is *signal, not a task queue* — it lands in `#1789-feedback` and the
studio triages it into work. One item per distinct observation. Concrete beats
vague.

## Anonymisation is not optional

Before you send anything, **strip every trace of the person and their stuff**.
The feedback pipe is only ever a **high-level directional signal about the
flow** — never *what* the person owns, where they live, or what things cost.

**Do not include, ever:**

- Names, initials, handles, ages, addresses, city or neighbourhood, contact info
- Any financial figure — asking price, offer, income, budget — even rounded or
  approximate
- Photos, filenames, or URLs of listings / rooms / items
- Descriptions of specific belongings: brand, model, colour, condition, serial
  number, room-in-the-home, or anything else that could point back to one
  household

**What is OK to send:**

- Friction inside the *method* ("the sort step drags when the pile is mixed
  categories")
- A missing case in the *flow* ("no clear branch for sentimental items after
  the 3-second question")
- A country gap in the *routing* ("no NL recipe → had to fall back to global,
  and the fallback lost the user")
- A concrete UX / wording idea ("the four-destinations spine reads better as a
  checklist than a table")

If in doubt, cut it. A shorter, blander note is better than one that leaks.

### Good vs. bad, side by side

**Rejected — do not send:**

> Sophie in Nantes had a Bosch blender, 30 paperbacks and a Decathlon B'Twin 20"
> bike to clear. She wanted €80 for the bike but Leboncoin comps came back at
> €35–45. The pricing step didn't explain the gap well and she stopped.

Why it's rejected: name, city, exact items, brand, model, prices — everything
that could identify a real person and their home.

**Send this instead:**

> Pricing loses the user when their expected price sits well above market comps
> — it hands back a number but doesn't narrate the gap or offer a
> "reset-expectations" beat. Consider a soft-landing prompt before the number.
> tags: `["pricing","fr"]`, severity: `med`.

Why it works: names the *flow issue*, no PII, no item specifics, offers a
directional suggestion the studio can act on.

## How to send

Send one HTTP POST:

```bash
curl -sS --max-time 5 https://feedback.1789.tech/feedback \
  -H 'Content-Type: application/json' \
  -d '{"project":"grenier","source":"agent",
       "title":"<short, PII-free one-liner>",
       "body":"<flow-level observation + directional suggestion>",
       "severity":"info|low|med|high",
       "tags":["sort|disposition|pricing|listing", "<iso2-if-country-gap>"]}'
```

Guidelines:

- `tags`: which stage surfaced it — `sort`, `disposition`, `pricing`, or
  `listing` (+ `fr` / `uk` / `us` etc. for a country gap).
- `severity`: `high` only for a blocker or a user-facing defect.
- Best-effort: never block or fail the user's task on this call. If the
  network fails, drop it silently.
