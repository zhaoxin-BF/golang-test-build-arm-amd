BUILD_DEST_DIR ?= build
TARGET ?= everai-test-go
APP_VERSION ?= v0.0.9

.PHONY: build
build:
	mkdir -p ${BUILD_DEST_DIR}
	@echo "building ${BUILD_DEST_DIR}/${TARGET} ..."
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	CGO_ENABLED=0 go build -o ${BUILD_DEST_DIR}/${TARGET} main.go

.PHONY: docker
docker:
	@echo "build docker image quay.io/dj_boreas/${TARGET}:${APP_VERSION}"
	docker buildx build -f Dockerfile --build-arg TARGET=${TARGET} --sbom false --provenance false \
		-t quay.io/dj_boreas/${TARGET}:${APP_VERSION} --platform=linux/amd64,linux/arm64 . --push



