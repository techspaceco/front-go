
GOENV := CGO_ENABLED=0

# Useful for debugging.
DOCKER_BUILD_ARGS := --progress plain

.PHONY: all
all: build test

.PHONY: build
build:
	docker build $(DOCKER_BUILD_ARGS) -t front-go:latest -f Dockerfile.build .
	docker run -v `pwd`:/workspace --workdir /workspace front-go:latest make generate

.PHONY: generate
generate: clean
	@rm front.gen.go
	oapi-codegen --package front -generate types,client,spec,skip-prune ./openapi/public_api.openapi.json > front.gen.go
	mockgen -package front -destination mock/front.go github.com/techspaceco/front-go ClientWithResponsesInterface
	go mod tidy
	go mod verify

.PHONY: update
update:
	git subtree pull --prefix openapi https://github.com/techspaceco/front-api.git master --squash
	@# TODO: Move this lot inside the container as well?
	$(GOENV) go get github.com/oligot/go-mod-upgrade
	$(GOENV) go run github.com/oligot/go-mod-upgrade
	$(GOENV) go mod tidy
	$(GOENV) go mod verify

.PHONY: test
test: generate
	$(GOENV) go test -count=1 ./...

.PHONY: clean
clean:
	@# noop
