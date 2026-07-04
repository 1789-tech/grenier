# Evals

`grenier` is prompt-only, so fixtures are lightweight assertions that describe
what a good answer must include or avoid.

`evals/evals.json` is intentionally simple:

- `skill` points to a folder in `skills/`;
- `prompt` is the user ask;
- `delta` states what should improve versus a raw answer without the skill;
- `must_include` names required concepts or strings;
- `must_avoid` catches anti-patterns.

Run the static checks with:

```bash
make test
```

The current target validates JSON syntax, required skill folders and fixture
shape. A model-backed with/without-skill runner can be added later.
