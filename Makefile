.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: clean # remove data
clean:
	@rm -rf bin charts coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint # execute golint
lint:
	@golint ./...

.PHONY: test-unit # execute unit test
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh
