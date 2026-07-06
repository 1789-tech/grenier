.PHONY: test run validate evals

test: validate evals

run: validate

# evals runs the Go harness: unit tests + structural validation of the fixtures
# against the current single-skill layout. Set EVAL_GENERATE_CMD to also score
# with-skill vs without-skill answers (see evals/README.md).
evals:
	go test ./...
	go run ./cmd/grenier-evals

validate:
	@jq -e '.cases | type == "array" and length > 0' evals/evals.json >/dev/null
	@jq -e 'all(.cases[]; (.id|type=="string") and (.skill|type=="string") and (.prompt|type=="string") and (.delta|type=="string") and (.must_include|type=="array") and (.must_avoid|type=="array"))' evals/evals.json >/dev/null
	@test -f skills/grenier/SKILL.md || { echo "missing skill: grenier"; exit 1; }
	@for ref in method disposition pricing listing; do \
		test -f "skills/grenier/references/$$ref.md" || { echo "missing reference: $$ref"; exit 1; }; \
	done
	@for iso in fr us uk; do \
		test -f "skills/grenier/countries/$$iso.md" || { echo "missing country recipe: $$iso"; exit 1; }; \
	done
	@jq -r '.cases[].skill' evals/evals.json | while read -r skill; do \
		test -f "skills/$$skill/SKILL.md" || { echo "eval references missing skill: $$skill"; exit 1; }; \
	done
	@sh -n install.sh
