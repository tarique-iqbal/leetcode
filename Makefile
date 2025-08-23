.PHONY: help build dev test test-local shell clean

IMAGE_NAME=leetcode-solution
APP_DIR=/home/leetcode

help:
	@echo "Available targets:"
	@echo "  build         Build the Docker image"
	@echo "  dev           Run container with local code mounted (dev shell)"
	@echo "  test          Run tests inside a clean container"
	@echo "  test-local    Run tests on local code in mounted container"
	@echo "  shell         Enter the container without mounting"
	@echo "  clean         Remove Docker image"

build:
	docker build -t $(IMAGE_NAME) .

dev:
	docker run -it --rm \
		-v $(PWD):$(APP_DIR) \
		-w $(APP_DIR) \
		--entrypoint bash \
		$(IMAGE_NAME)

test:
	docker run --rm $(IMAGE_NAME) go test -v ./solution/...

test-local:
	docker run --rm \
		-v $(PWD):$(APP_DIR) \
		-w $(APP_DIR) \
		$(IMAGE_NAME) go test -v ./solution/...

shell:
	docker run -it --rm $(IMAGE_NAME) bash

clean:
	docker rmi $(IMAGE_NAME) || true
