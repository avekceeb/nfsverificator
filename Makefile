
GINKGO_OPTS :=

test:
	ginkgo -v $(GINKGO_OPTS) -focus "${FOCUS}" -skip "Stress" ./tests/v40

dryrun:
	ginkgo -v -dryRun ./tests/v40 -- -runtime=dryrun

local:
	ginkgo -v -skip "Stress" ./tests/v40 -- -config $(PWD)/local.json

stress:
	ginkgo -v -focus "Stress.*Multi" ./tests/v40
	#ginkgo -v -focus "Stress.*Multi" ./tests/v40 -- -config $(PWD)/local.json

test41:
	ginkgo -v ./tests/v41 -- -config $(PWD)/local.json

