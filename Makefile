postgres:
	docker run --name teckwatch -p5432:5432 -e POSTGRES_USER=teckwatch -e POSTGRES_PASSWORD=teckwatch -d postgres:13-alpine

createdb:
	docker exec -it teckwatch createdb --username=teckwatch --owner=teckwatch teckwatch_dev

dropdb:
	docker exec -it teckwatch dropdb --username=teckwatch teckwatch_dev

migrateup:
	migrate -path database/migration -database "postgres://teckwatch:teckwatch@localhost:5432/teckwatch_dev?sslmode=disable" --verbose up

migratedown:
	migrate -path database/migration -database "postgres://teckwatch:teckwatch@localhost:5432/teckwatch_dev?sslmode=disable" --verbose down

run:
	make createdb
	make migrateup

.PHONE: postgres createdb dropdb migrateup migratedown run