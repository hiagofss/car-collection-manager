.PHONY: run
run:
	npx nodemon -e html,go --exec go run main.go --signal SIGTERM
.PHONY: dcup
dev:
	docker compose up app
.PHONY: dcdown
dcdown:
	docker compose down