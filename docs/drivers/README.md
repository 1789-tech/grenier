# Marketplace drivers — contributor guide

This is the open-source runner core for grenier's night-operator (DEC-0130 /
BRIEF-118). It fills marketplace **drafts** in a browser and stops there. The
build contract is [`docs/marketplace-drivers-v1.md`](../marketplace-drivers-v1.md);
this guide is how you add a driver against it.

## The one rule everything else serves

> `save as draft` is terminal. No driver may publish, send, message, delete, or
> accept anything without explicit human validation after the wake-up handoff.

This is why the [`Driver`](../../driver/types.go) interface has exactly one write
path, `FillDraft`. `Publish`, `MessageBuyer`, `AcceptOffer`, `DeleteListing`
**do not exist** — you cannot call a method that isn't there. That is the design,
not an omission: the safest publish path is one with no code.

## Layout

```
driver/                    open-source core (interface, runner, receipts, audit, safety)
  types.go                 Driver interface + data shapes + PermissionPosture
  page.go                  browser port (real Chrome or fixture both implement it)
  manifest.go              manifest schema + dependency-free TOML loader
  selectcontrol.go         the terminal-control safety rule (shared, tested once)
  runner.go                run one handoff: gate, resume, retry, dry-run
  receipt.go / audit.go    incremental receipt + audit log
  fixturemarket/           in-memory fake DOM implementing Page, for proving drivers
drivers/<id>/              one package per marketplace
  <id>.go                  the Driver implementation
  manifest.toml            embedded, machine-readable, part of the review surface
  <id>_test.go             the required fixture tests
cmd/grenier-driver/        the local CLI: handoff.json -> receipt.json
```

## Adding a driver

1. **Write the manifest** (`drivers/<id>/manifest.toml`). It is the review
   surface — reviewers read it before the code. It must declare:
   - `permission_posture` and `detection_risk` — **different axes** (see below);
   - `[draft] terminal_action` and `publish_action_wired = false`;
   - `[selectors] draft_action_names` / `publish_action_names` (accessible-name
     matchers), a `confidence` label, and a `last_verified` date.
   `Manifest.Validate()` rejects a manifest with `publish_action_wired = true`,
   no draft stop, or a missing confidence/last-verified — at load time, before
   any browser work.

2. **Implement `Driver`.** Fill fields through the `Page` port, then resolve the
   terminal control with `driver.SelectTerminalControl(controls, manifest)` —
   do **not** hand-roll button selection. That helper is the safety rule:
   - unique draft match → click it (even if a publish button is also present);
   - two draft matches, or a draft/publish overlap → `skipped_ambiguous_submit`;
   - publish present, no draft → `skipped_publish_only` (disqualified);
   - nothing matches → `skipped_selector_drift` (skip, never best-effort).
   Honor `driver.IsDryRun(ctx)`: fill fields, then stop before the terminal click.

3. **Register** in `init()`: `driver.Register("<id>", factory)`.

4. **Write the required tests** (`drivers/<id>/<id>_test.go`), each driving a
   `fixturemarket.Scenario`:
   - happy path draft saved, **with a publish button present, asserted never
     clicked** (`mkt.ClickedID("btn-publish") == false`);
   - ambiguous submit → skipped;
   - publish-only page → disqualified;
   - selector drift → skipped;
   - dry-run fills but does not click.

## Permission posture vs detection risk

They are orthogonal. A driver can be low-detection and ToS-forbidden, or
officially authorized but high-detection. Never collapse them into one "risk"
field.

| Posture | Meaning |
| --- | --- |
| `official_api` | A sanctioned API with a native draft primitive. |
| `authorized` | Written permission / allowlist for external software. |
| `user_browser_tos_hostile` | User's own browser, but the ToS is hostile to automation. |
| `forbidden` | Explicitly disallowed. Do not ship. |
| `unknown` | Not yet determined. Do not run against a real account. |

## Source quality labels

Use these in the manifest `confidence` and in driver docs. Do not write field
reports as platform facts.

- `platform-confirmed` — official docs, public terms, robots.txt, help center.
- `vendor-confirmed` — anti-bot vendor case studies, integration vendors.
- `field-report` — blogs, forums, reseller claims, observed behavior.

## The V1 reference driver: `vinted-fr` and its open posture

`drivers/vintedfr` proves the `FillDraft` path end-to-end. Its posture is
`user_browser_tos_hostile`, and DEC-0130 leaves **acceptance of that posture for
the first driver as a Manfred decision**. Until it is ruled:

- the driver runs **only against the fixture DOM** (`--browser=fixture`, CI,
  dry-run). The exact same code runs against a real user-owned browser once the
  posture is accepted — both satisfy `driver.Page` — but this package never
  presumes acceptance;
- `--browser=real` is **deliberately not wired** and errors with the reason.

Do not remove that gate to "make it work." Wiring a real browser against a live
account is the posture decision, in code form.

## Running

```
go test ./...                                   # unit + fixture tests
go run ./cmd/grenier-driver --handoff h.json    # handoff -> receipt (fixture)
go run ./cmd/grenier-driver --handoff h.json --dry-run
```
