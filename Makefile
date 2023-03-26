include .env

build:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api go build -o ./hello-serviceweaver

deps:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api go mod tidy
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api go mod vendor

generate:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api go generate .

weaver-single-deploy:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api go run .

weaver-single-test:
	curl "localhost:8000/products"

weaver-single-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver single $*

weaver-multi-deploy:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver multi deploy weaver.toml

weaver-multi-test:
	curl "localhost:8000/products"

weaver-multi-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver multi $*

weaver-gke-local-deploy:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver gke-local deploy weaver.toml

weaver-gke-local-test:
	# I configured this application by `api/weaver.toml` to associate host name `hello-serviceweaver.example.com` with the `hello-serviceweaver` listener.
	curl --header 'Host: hello-serviceweaver.example.com' "localhost:8000/products"

weaver-gke-local-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver gke-local $*

weaver: weaver---help
weaver-%:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver $*
