# Simple Go Backend

Seamless with Docker: `docker-compose up --build` → DB migrates + seed → API ready.

## Quickstart

```bash
cp .env.example .env
docker-compose up --build -d
make seed
```

## API

Base URL: http://your-host:8080/api/v1

* `GET /health`

* `POST /auth/login`
  JSON: `{ "username":"cole", "password":"password" }`

* `GET /products?page=1&page_size=10&search=shoe&category=men&name=Nike&sort=price_desc`

* `GET /products/{id}`

* `GET /cart`

* `POST /cart/add` (JSON: `{ "product_id": 1, "qty": 2 }`)

* `POST /cart/{pid}/inc`

* `POST /cart/{pid}/dec`

* `DELETE /cart/{pid}`

## Users (seed)

| Username | Password  |
| -------- | --------- |
| cole     | password  |
| brian    | password  |
| alice    | password  |

## Makefile

* `make migrate-up`
* `make migrate-down`
* `make migrate-refresh`
* `make seed`
