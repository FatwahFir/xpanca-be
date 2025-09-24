include .env

DB_URL = mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true

.PHONY: migrate-up migrate-down migrate-refresh seed

migrate-up:
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL)" up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL)" down 1

migrate-refresh:
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL)" drop -f
	docker compose run --rm migrate -path=/migrations -database "$(DB_URL)" up

seed:
	docker compose exec app /bin/seed
