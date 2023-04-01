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
		-e SERVICEWEAVER_CONFIG=weaver.toml \
		api go run .

weaver-single-test:
	echo "[list products]"
	curl "localhost:8000/products"
	echo "\n"
	echo "[get product]"
	curl "localhost:8000/products/P001"
	echo "\n"
	echo "[purchase product]"
	curl "localhost:8000/products/P001/purchase"

weaver-single-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver single $*

weaver-multi-deploy:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver multi deploy weaver.toml

weaver-multi-test:
	echo "[list products]"
	curl "localhost:8000/products"
	echo "\n"
	echo "[get product]"
	curl "localhost:8000/products/P001"
	echo "\n"
	echo "[purchase product]"
	curl "localhost:8000/products/P001/purchase"

weaver-multi-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver multi $*

weaver-gke-local-deploy:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver gke-local deploy weaver.toml

weaver-gke-local-test:
	# I configured this application by `api/weaver.toml` to associate host name `hello-serviceweaver.example.com` with the `hello-serviceweaver` listener.
	echo "[list products]"
	curl --header 'Host: hello-serviceweaver.example.com' "localhost:8000/products"
	echo "\n"
	echo "[get product]"
	curl --header 'Host: hello-serviceweaver.example.com' "localhost:8000/products/P001"
	echo "\n"
	echo "[purchase product]"
	curl --header 'Host: hello-serviceweaver.example.com' "localhost:8000/products/P001/purchase"

weaver-gke-local-%:
	docker exec -it $(shell docker ps | grep "hello-serviceweaver" | awk '{ print $$1 }') weaver gke-local $*

weaver: weaver---help
weaver-%:
	docker compose -f ./tools/compose.yaml run --rm --service-ports \
		api weaver $*
