.PHONY: test run validate

test: validate

run: validate

validate:
	@jq -e '.cases | type == "array" and length > 0' evals/evals.json >/dev/null
	@jq -e 'all(.cases[]; (.id|type=="string") and (.skill|type=="string") and (.prompt|type=="string") and (.delta|type=="string") and (.must_include|type=="array") and (.must_avoid|type=="array"))' evals/evals.json >/dev/null
	@for skill in declutter offload price-and-decide sell; do \
		test -f "skills/$$skill/SKILL.md" || { echo "missing skill: $$skill"; exit 1; }; \
	done
	@jq -r '.cases[].skill' evals/evals.json | while read -r skill; do \
		test -f "skills/$$skill/SKILL.md" || { echo "eval references missing skill: $$skill"; exit 1; }; \
	done
	@sh -n install.sh
