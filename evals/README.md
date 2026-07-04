# Evals

`grenier` is a single, prompt-only skill with progressive disclosure — a core
`skills/grenier/SKILL.md` plus on-demand `references/` and `countries/` recipes.
The fixtures here are lightweight string assertions describing what a good answer
must include or avoid.

## Fixture shape — `evals/evals.json`

```jsonc
{
  "cases": [
    {
      "id": "price-bike-balanced-one-move",   // unique
      "skill": "grenier",                       // folder under skills/
      "prompt": "France. J'ai un velo enfant…", // the user ask
      "delta": "With the skill, the answer must collapse pricing and channel into one balanced move.",
      "must_include": ["Leboncoin", "35", "remise en main propre"],  // required strings (case-insensitive)
      "must_avoid":   ["eBay", "Vinted", "trois options"]            // anti-patterns that must NOT appear
    }
  ]
}
```

- `skill` must point at the current single-skill layout (`grenier`), not the old
  `haggle` / `listing-drafter` / `where-to-sell-*` taxonomy.
- Each case needs at least one `must_include` or `must_avoid` assertion.
- `delta` documents the intended improvement; it is not scored.

## Running

```bash
make evals        # go test ./...  +  structural validation (no model call)
```

`make evals` (or `go run ./cmd/grenier-evals`) parses the fixtures and confirms
every referenced skill and its progressive-disclosure bundle
(`references/` + `countries/`) exists. It runs in CI with zero external
dependencies — no API key, no network. (`make test` also runs the jq-based
static checks in the Makefile.)

### With/without scoring (optional)

Point the runner at any command that reads a prompt on **stdin** and writes the
model's answer to **stdout**. The runner sends each prompt twice — once with the
full skill context (SKILL.md + all references + all countries) prepended, once
raw — and reports the with-vs-without delta per skill:

```bash
EVAL_GENERATE_CMD='your-llm-cli --quiet' make evals
# or directly:
go run ./cmd/grenier-evals -generate-cmd 'your-llm-cli --quiet'
```

Flags: `-evals` (fixture path), `-skills` (skills dir), `-generate-cmd`.

## Adding a case

1. Append an object to `cases` in `evals/evals.json` with a unique `id`.
2. Write `must_include` / `must_avoid` from the *decision* the skill should
   drive (one best move, priority fork, refuse a low-value sale…), not from
   surface wording.
3. Run `make evals`.
