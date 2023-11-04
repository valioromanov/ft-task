calculator-api-tests:
	ginkgo -r -race -randomize-all -randomize-suites "./cmd/calcurator-api"

calculator-repo-tests:
	ginkgo -r -race -randomize-all -randomize-suites "./pkg/calculator"

all-tests:
	ginkgo -r -race -randomize-all -randomize-suites .

get-coverage:
	ginkgo -r -v -race --trace --cover --coverprofile=.coverage-report.out --coverpkg=./... ./...