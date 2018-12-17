.PHONY: container push deploy

REVISION := $(shell git rev-parse HEAD)

container:
	@echo "********** build contaienr ************"
	docker build -t pcodk/masterclass-api -f docker/app/Dockerfile .

push:
	@echo "********** push contaienr ************"
	docker push pcodk/masterclass-api



