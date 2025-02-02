CONTAINER_NAME=
.PHONY: run
run:
	npx nodemon -e html,go --exec go run main.go --signal SIGTERM
.PHONY: dev
dev:
	docker compose up app
.PHONY: chk
chk:
	docker inspect --format "{{json .State.Health }}" $(CONTAINER_NAME) | jq
.PHONY: dcdown
dcdown:
	docker compose down
.PHONY: gen-docs
gen-docs:
	swag init

