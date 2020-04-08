.PHONY: build publish

buildubuntu:
	docker build -t filario/usplay-api-todo:ubuntu -f build/container/ubuntu.Dockerfile .
build:
	docker build -t filario/usplay-api-todo:latest -f build/container/Dockerfile .
publish:
	docker push filario/usplay-api-todo:$(filter-out $@,$(MAKECMDGOALS))
