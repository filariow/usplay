.PHONY: integration-tests deps

deps:
	go mod download

fname=$(RESULT_DIR)/test-$(shell date +%s).log
integration-tests:
	go test -v ./... | tee $(fname)	