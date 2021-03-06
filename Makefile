.PHONY: v40 v41 v42

# Defaults:
skip := Stress|Slow|Problematic|Expiration|Reboot
focus :=
config := $(PWD)/config.json
server :=
export :=
trace :=

opts :=
ifneq ($(config),)
    opts = -config $(config)
endif
ifneq ($(server),)
    opts += -server $(server)
endif
ifneq ($(export),)
    opts += -export $(export)
endif
ifneq ($(trace),)
    opts += -trace $(trace)
endif

gopts := -v --trace -keepGoing -skip "$(skip)" -focus "$(focus)"

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

