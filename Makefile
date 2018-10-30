
SKIP :=

GINKGO_OPTS :=

test:
	ginkgo -v $(GINKGO_OPTS) -focus "${FOCUS}" -skip "${SKIP}" ./tests/v40

dryrun:
	ginkgo -v -dryRun ./tests/v40 -- -runtime=dryrun
