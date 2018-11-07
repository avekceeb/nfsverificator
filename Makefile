.PHONY: v40 v41

# Defaults:
skip := Stress|Slow
focus :=
config := $(PWD)/config.json

opts := -config $(config)
gopts := --trace -skip "$(skip)" -focus "$(focus)"

# -v

all:
	ginkgo $(gopts) ./tests/... -- $(opts)

v40:
	ginkgo $(gopts) ./tests/v40 -- $(opts)

v41:
	ginkgo $(gopts) ./tests/v41 -- $(opts)

stress:
	ginkgo -v -focus "Stress.*Multi" ./tests/v40 -- $(opts)

