migrateup:
	docker run --network host src -path=/migration/ -database "postgresql://postgres:secret@localhost:1234/log_km?sslmode=disable" up
migratedown:
	docker run --network host src -path=/migration/ -database "postgresql://postgres:secret@localhost:1234/log_km?sslmode=disable" down

.PHONY: migrateup migratedown