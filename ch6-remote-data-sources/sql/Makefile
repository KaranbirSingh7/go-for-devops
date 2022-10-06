POSTGRES_USER=postgres
POSTGRES_PWD=postgres
POSTGRES_DB=mydb

.PHONY: db-restart
db-restart: db-stop db-start

.PHONY: db-start
db-start: 
	docker run --name postgres \
		-e POSTGRES_PASSWORD=${POSTGRES_PWD} \
		-e POSTGRES_USER=${POSTGRES_USER} \
		-e POSTGRES_DB=${POSTGRES_DB} \
		-p 5432:5432 \
		-d postgres:14.5-alpine

.PHONY: db-stop
db-stop:
	docker rm -f postgres

