
SKIP :=

GINKGO_OPTS :=

test:
	ginkgo -v $(GINKGO_OPTS) -focus "${FOCUS}" -skip "${SKIP}" ./tests

dryrun:
	ginkgo -v -dryRun ./tests -- -runtime=dryrun
