.PHONY: integration-tests

fname=$(RESULT_DIR)/test-$(shell date +%s).log
integration-tests:
	go test -v ./... -tags=integration | tee $(fname)