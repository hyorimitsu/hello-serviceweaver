include .env

build:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api go build -o ./hello-serviceweaver

deps:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api go mod tidy
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api go mod vendor

generate:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api go generate .

weaver-single-deploy:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api go run .

weaver-single-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver single $*

weaver-multi-deploy:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api weaver multi deploy weaver.toml

weaver-multi-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver multi $*

weaver-gke-local-deploy:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api weaver gke-local deploy weaver.toml

weaver-gke-local-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver gke-local $*

weaver-gke-deploy:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api weaver gke deploy weaver.toml

weaver-gke-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver gke $*

weaver: weaver---help
weaver-%:
	docker compose -f ./tools/compose.yaml run --rm \
		-v "$(PWD)/api":"/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		-w "/go/src/github.com/hyorimitsu/hello-serviceweaver/api" \
		api weaver $*
