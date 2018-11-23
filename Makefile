.PHONY: v40 v41 v42

# Defaults:
skip := Stress|Slow
focus :=
config := $(PWD)/config.json

opts := -config $(config)
gopts := -v --trace -keepGoing -skip "$(skip)" -focus "$(focus)"

# -v

all:
	ginkgo $(gopts) ./tests/... -- $(opts)

v40:
	ginkgo $(gopts) ./tests/v40 -- $(opts)

v41:
	ginkgo $(gopts) ./tests/v41 -- $(opts)

v42:
	ginkgo $(gopts) ./tests/v42 -- $(opts)

stress:
	ginkgo -v -focus "Stress.*Multi" ./tests/v40 -- $(opts)

