# Go Backend (Gin + GORM + MySQL) — Hexagonal

Seamless with Docker: `docker-compose up --build` → DB migrates + seed → API ready.

## Quickstart

```bash
cp .env.example .env
docker-compose up --build
# in another terminal (first time only)
make migrate-up
make seed
```

API:
- `POST /auth/login` (JSON: `{ "username":"admin", "password":"secret123" }`)
- `GET /products?page=1&page_size=10&search=shoe&category=men&name=Nike`
- `GET /products/{id}`

Makefile:
- `make migrate-up` `make migrate-down` `make migrate-refresh` `make seed`

## Architecture

- `cmd/api` — HTTP server (Gin)
- `cmd/seed` — Seeder
- `internal/domain` — Entities (core)
- `internal/usecase` — Application services
- `internal/adapter/http` — Handlers (delivery)
- `internal/adapter/repository/mysql` — Repos (persistence)
- `internal/middleware` — JWT auth
- `internal/config` — Config loader
- `pkg/*` — small libs (jwt, password)

Migrations in `db/migrations` using golang-migrate.
