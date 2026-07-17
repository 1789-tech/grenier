# Marketplace driver V1

BRIEF-110 / DEC-0130 turns the night-operator boundary into a build target:
grenier should be able to fill marketplace drafts in a real browser, without
folding marketplace execution back into the Claude Skill and without depending
on one proprietary API.

The V1 goal is deliberately narrow: an open-source driver core plus one French
marketplace reference driver that proves `fill_draft` end-to-end. More drivers
come later.

## Product contract

The non-negotiable rule from the Skill remains the product contract here:

> `save as draft` is terminal. No driver may publish, send, message, delete, or
> accept anything without explicit human validation after the wake-up handoff.

If a marketplace does not expose a stable draft stop, its driver is not eligible
for V1. It can exist as a research note, not as a shipped driver.

## Architecture

Keep the Skill and execution surface separate:

```text
grenier Skill
  -> handoff.json          # listings prepared by the night operator
  -> driver runner         # local CLI / daemon, user-owned browser session
  -> marketplace driver    # one implementation per marketplace
  -> receipt.json          # filled / skipped / needs-human-review
  -> grenier inventory     # read by the Skill on resume
```

The Skill produces copy-ready draft data and consumes receipts. It does not drive
the browser directly. The runner owns browser lifecycle, persistence, retries,
and audit logs. Drivers own marketplace-specific navigation and selectors.

## Driver interface

Use a small capability interface, not a marketplace-specific API wrapper:

```go
type Driver interface {
    Marketplace() Marketplace
    Capabilities(ctx context.Context) (Capabilities, error)
    FillDraft(ctx context.Context, item DraftInput) (DraftResult, error)
    ListDrafts(ctx context.Context) ([]DraftSummary, error)
    LoadState(ctx context.Context) (DriverState, error)
    SaveState(ctx context.Context, state DriverState) error
}
```

`FillDraft` is the only V1 write path. `Publish`, `MessageBuyer`,
`AcceptOffer`, and `DeleteListing` must not exist in the V1 interface.

Minimal data shapes:

```go
type Marketplace struct {
    ID      string // "vinted-fr", "leboncoin-fr"
    Country string // "FR"
    Name    string
}

type Capabilities struct {
    FillDraft       bool
    ListDrafts      bool
    UploadPhotos    bool
    DraftStop       DraftStop
    Permission      PermissionPosture
    DetectionRisk   RiskLevel
    MaxItemsPerRun  int
}

type DraftInput struct {
    ItemID      string
    Title       string
    Description string
    PriceCents  int
    Currency    string
    Category    string
    Condition   string
    PhotoPaths  []string
    Attributes  map[string]string
}

type DraftResult struct {
    ItemID       string
    Marketplace  string
    Status       DraftStatus
    DraftURL     string
    SkippedReason string
    Warnings     []string
}
```

`PermissionPosture` is required, because permission and detection are different
axes:

```go
type PermissionPosture string

const (
    PermissionOfficialAPI PermissionPosture = "official_api"
    PermissionAuthorized  PermissionPosture = "authorized"
    PermissionUserBrowser PermissionPosture = "user_browser_tos_hostile"
    PermissionForbidden   PermissionPosture = "forbidden"
    PermissionUnknown     PermissionPosture = "unknown"
)
```

## Safety model

V1 safety rests on driver behavior, tests, and receipts. Do not claim network
body interception as the core guarantee; browser extension and automation APIs
are not a reliable semantic firewall for JSON request bodies.

Required controls:

- The driver may only select controls whose accessible name / label matches a
  marketplace-specific draft action.
- If both draft and publish controls are visible and the draft control is not
  uniquely identified, `FillDraft` must return `skipped_ambiguous_submit`.
- Selectors must be declared in a driver manifest with a confidence note and a
  last-verified date.
- Tests must include a fixture where a publish button is present and assert that
  it is never clicked.
- The runner must write a receipt after every item, including skipped items, so
  the Skill can resume without pretending the night succeeded.
- All drivers must support a dry-run mode that fills fields but stops before the
  terminal draft action.

Defense-in-depth controls, useful but not sufficient alone:

- Rate limits and jitter per driver.
- Screenshot or accessibility-tree capture before the terminal draft click.
- Optional request logging that flags unexpected publish-like endpoints.

## Driver manifest

Every driver ships a machine-readable manifest next to the code:

```toml
id = "vinted-fr"
country = "FR"
display_name = "Vinted France"
surface = "full_browser"
permission_posture = "user_browser_tos_hostile"
detection_risk = "medium"
max_items_per_run = 6
last_verified = "2026-07-17"

[draft]
supported = true
terminal_action = "save_draft"
publish_action_wired = false

[sources]
terms = "https://www.vinted.com/terms_and_conditions"
official_api = "https://pro-docs.svc.vinted.com/"
```

The manifest is part of the review surface. If the driver is ToS-hostile, say so
in the manifest and docs. The core stays neutral; individual drivers carry their
own posture.

## V1 reference driver choice

Use one FR full-browser driver first, not two.

| Candidate | Draft primitive | Permission posture | Detection risk | V1 call |
| --- | --- | --- | --- | --- |
| Vinted FR consumer browser | Consumer drafts exist; Pro API also has native `is_draft` | ToS-hostile unless Vinted authorizes external software | Medium | Best reference driver if Manfred accepts per-driver risk |
| Leboncoin FR browser | Public page says drafts resume across devices | ToS/robots-hostile without prior permission | High | Park for V1 despite draft primitive |
| Vinted Pro API | Native `is_draft` in official API | Official but allowlisted to select businesses | Low | Keep as later authorized driver, not V1 full-browser proof |

DEC-0130 chooses full-browser drivers for V1, so the Vinted Pro API should not be
the foundation. Still, its `is_draft` support is the clean target for a later
authorized driver if the allowlist path opens.

## Minimal V1 slice

1. `grenier-driver` local CLI:
   - reads `handoff.json`;
   - launches or attaches to a user-owned browser profile;
   - loads exactly one driver by manifest ID;
   - writes `receipt.json` incrementally.
2. `Driver` interface plus manifest schema.
3. One FR reference driver proving:
   - login/session reuse;
   - photo upload from local paths;
   - title / description / price / category / condition fill;
   - terminal save-as-draft only;
   - list drafts for receipt reconciliation.
4. Fixtures and tests:
   - happy path draft;
   - ambiguous submit -> skipped;
   - publish-only page -> disqualified;
   - selector drift -> skipped, not best-effort publish.
5. Documentation for contributors:
   - how to add a driver;
   - how to label permission posture and source quality;
   - why publish actions are not part of the interface.

## Handoff files

`handoff.json` should be generated by the Skill or exported from the local
inventory:

```json
{
  "run_id": "grn-run-20260717-001",
  "country": "FR",
  "marketplace": "vinted-fr",
  "items": [
    {
      "item_id": "grn-20260717-0001",
      "title": "Veste denim Levi's taille M",
      "description": "Bon etat, coupe droite, a retirer a Paris.",
      "price_cents": 2800,
      "currency": "EUR",
      "category": "mode",
      "condition": "good",
      "photo_paths": ["photos/grn-20260717-0001-1.jpg"],
      "attributes": {"brand": "Levi's", "size": "M"}
    }
  ]
}
```

`receipt.json` is append-safe and resume-friendly:

```json
{
  "run_id": "grn-run-20260717-001",
  "marketplace": "vinted-fr",
  "results": [
    {
      "item_id": "grn-20260717-0001",
      "status": "draft_saved",
      "draft_url": "https://example.invalid/drafts/123",
      "warnings": []
    }
  ]
}
```

## Source quality labels

Use three labels in driver docs and review:

- `platform-confirmed`: official docs, public terms, robots.txt, marketplace help.
- `vendor-confirmed`: anti-bot vendor case studies, integration vendors.
- `field-report`: blogs, forums, reseller claims, observed behavior.

Do not write field reports as platform facts.

## Current source anchors

- Vinted Pro Integrations documents an allowlist for select businesses and an
  item API with `is_draft`: https://pro-docs.svc.vinted.com/
- Vinted's public terms are the source to check for external-software
  restrictions before recommending a consumer driver:
  https://www.vinted.com/terms_and_conditions
- Leboncoin's robots.txt says automated access needs special permission:
  https://www.leboncoin.fr/robots.txt
- Leboncoin's CGU restrict extraction/indexing by robots or similar tools without
  prior written authorization: https://www.leboncoin.fr/dc/cgu
- Leboncoin's public nouveautes page says drafts can be saved and resumed across
  devices, so "no draft primitive" is not a valid reason to disqualify it:
  https://www.leboncoin.fr/evenement/nouveautes
- DataDome's Leboncoin case study reports 9.5M malicious requests blocked per
  day on average, 30M peak, 30 sensitive endpoints, and a 0.0018% false-positive
  rate:
  https://datadome.co/customers-stories/how-leboncoin-blocks-millions-of-malicious-requests-every-day/

## Open decisions before code merge

- Which marketplace is the first reference driver under DEC-0130?
- Is `user_browser_tos_hostile` acceptable for the first driver, or must V1 wait
  for an authorized driver?
- Is the runner a Go CLI first, or a browser-extension host with native
  messaging? The driver interface above works for both; Go CLI is the smaller
  open-source core.
- What is the maximum run size for the first driver? Default to a conservative
  manifest cap and make it visible in receipts.
