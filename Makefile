export GO111MODULE = on
APP_NAME?=sample-service


BUILD_VERSION?=0.0.0-snapshot
.PHONY: test run build build-with-docker docker-build docker-push lint install-build-deps 

GOPATH := $(shell go env GOPATH)
LINTERS := \
	github.com/golang/lint/golint \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck \
	honnef.co/go/tools/cmd/unused

PACKAGES = $(shell go list ./... | grep -v /vendor/)


install-build-deps:
	go install -v $(LINTERS)

lint:
	env GO111MODULE=off go fmt ./...
	env GO111MODULE=on go vet -mod=vendor ./...

test: lint
	mkdir -p builds
	env GO111MODULE=on go test -mod=vendor -race -coverprofile=${UNIT_COVERAGE_OUTPUT} ./...

run:
	env GO111MODULE=on go build -mod=vendor
	env ./sample-service

build:
	env GOOS=linux CGO_ENABLED=0 GO111MODULE=on /usr/local/go/bin/go build -mod=vendor -o builds/sample-service cmd/sample-service/main.go

docker-build:
	docker build --rm -t sample-service .


docker-build-images: docker-build
	docker login -u ${ARTIFACTORY_USER} -p ${ARTIFACTORY_PASSWORD}
	docker tag sample-service snagarju/sample-service:latest
	docker tag sample-service snagarju/sample-service:${BUILD_VERSION}
	docker push snagarju/sample-service:latest
	docker push snagarju/sample-service:${BUILD_VERSION}
