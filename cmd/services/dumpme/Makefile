.PHONY: build publish

buildubuntu:
	docker build -t filario/usplay-api-dumpme:ubuntu -f build/container/ubuntu.Dockerfile .
build:
	docker build -t filario/usplay-api-dumpme:latest -f build/container/Dockerfile .
publish:
	docker push filario/usplay-api-dumpme:$(filter-out $@,$(MAKECMDGOALS))
