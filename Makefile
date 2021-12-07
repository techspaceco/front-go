
GOPRIVATE ?= github.com/techspaceco
GOENV = CGO_ENABLED=0 GOPRIVATE=$(GOPRIVATE)

.PHONY: all
all: generate test

.PHONY: generate
generate: clean
	@rm front.gen.go
	$(GOENV) go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	$(GOENV) go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package front -generate types,client,spec,skip-prune ./openapi/public_api.openapi.json > front.gen.go
	$(GOENV) go install github.com/golang/mock/mockgen@v1.6.0
	$(GOENV) go get -d github.com/golang/mock/mockgen@v1.6.0
	$(GOENV) mockgen -package front -destination mock/front.go github.com/techspaceco/front-go ClientWithResponsesInterface
	$(GOENV) go mod tidy
	$(GOENV) go mod verify

.PHONY: update
update:
	git subtree pull --prefix openapi https://github.com/techspaceco/front-api.git master --squash
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
