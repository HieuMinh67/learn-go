start_db:
	docker  run --name db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e -d postgres:14-alpine

create_db:
	docker exec -it db createdb --username=root --owner=root app

drop_db:
	docker exec -it db app

import_data:
	cat data.sql | docker exec -i db psql -U root -d app

export_data:
	docker exec -t db pg_dumpall -c -U root > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql

.PHONY: start_db create_db drop_db import_data export_data