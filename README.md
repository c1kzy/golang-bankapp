# Bank App

A small Go pet project implementing a simple banking backend: user management and money transactions (deposit, withdraw, transfer), backed by PostgreSQL and exposed over HTTP.

## Features

- **Users** — create, update, delete users
- **Transactions** — deposit, withdraw, and transfer money between users
- **History** — track and query past transactions
- Structured logging
- Graceful shutdown on `SIGINT` / `SIGTERM`
- PostgreSQL connection pooling (via `pgx`)

## Project Structure

The app is organized by feature, each with its own repository / service / transport (HTTP handler) layers:

```
.
├── cmd/
│   └── main.go
├── migrations/
├── docker-compose.yml
├── Makefile
├── .env
├── core/
│   ├── logger/            # logging setup
│   ├── pgx_pool/          # Postgres connection pool
│   ├── http_server/       # HTTP server + config
│   ├── middleware/        # HTTP middleware (e.g. request logging)
│   ├── user/
│   │   ├── repository/    # user data access
│   │   ├── service/       # user business logic
│   │   └── transport/     # user HTTP handlers
│   ├── transactions/
│   │   ├── repository/
│   │   ├── service/
│   │   └── transport/
│   └── history/
│       ├── repository/
│       ├── service/
│       └── transport/
```

## Requirements

- Go 1.2x+
- Docker & Docker Compose
- `make`

## Configuration

Configuration is loaded from a `.env` file at the project root (the Makefile `include`s and exports it automatically for every target):

| Variable | Description |
|---|---|
| `HTTP_ADDR` | Address/port the HTTP server listens on |
| `HTTP_SHUTDOWN_TIMEOUT` | Graceful shutdown timeout for the HTTP server (e.g. `30s`) |
| `POSTGRES_HOST` | PostgreSQL host |
| `POSTGRES_USER` | PostgreSQL user |
| `POSTGRES_PASSWORD` | PostgreSQL password |
| `POSTGRES_DB` | PostgreSQL database name |
| `POSTGRES_TIMEOUT` | PostgreSQL connection timeout (e.g. `10s`) |
| `LOGGER_LEVEL` | Logger verbosity (e.g. `DEBUG`, `INFO`, `WARN`, `ERROR`) |

> `LOGGER_FOLDER` is set automatically by `make bankapp-run` (to `out/logs`) — no need to set it yourself.

Copy the template below into a `.env` file at the project root and fill in the blanks:

```dotenv
HTTP_ADDR=
HTTP_SHUTDOWN_TIMEOUT=30s

POSTGRES_HOST=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG
```

## Getting Started

1. Clone the repo:
   ```bash
   git clone github.com/c1kzy/golang-bankapp
   cd bank-app
   ```

2. Create your `.env` file with the required variables (see [Configuration](#configuration)).

3. Start PostgreSQL:
   ```bash
   make env-up
   ```

4. Run database migrations:
   ```bash
   make migrate-up
   ```

5. Run the app:
   ```bash
   make bankapp-run
   ```

The server will start, connect to PostgreSQL, and register the user, transaction, and history routes.

## Makefile Commands

| Command | Description |
|---|---|
| `make bankapp-run` | Tidy modules and run the app (`cmd/main.go`) |
| `make env-up` | Start the `bankapp-postgres` container in the background |
| `make env-down` | Stop the `bankapp-postgres` container |
| `make env-cleanup` | Stop Postgres and delete its volume data (`out/pgdata`), with confirmation |
| `make logs-cleanup` | Delete local log files (`out/logs`), with confirmation |
| `make migrate-create seq=<name>` | Create a new migration file pair (e.g. `make migrate-create seq=init`) |
| `make migrate-up` | Apply all pending migrations |
| `make migrate-down` | Roll back migrations |
| `make migrate-action action=<up\|down>` | Run a migration action directly against the DB |

Migrations live in `migrations/` and are run via a `bankapp-postgres-migrate` service (defined in `docker-compose.yml`), using [golang-migrate](https://github.com/golang-migrate/migrate) under the hood.

## Docker

`docker-compose.yml` defines (at least) two services:

- **`bankapp-postgres`** — the PostgreSQL database used by the app
- **`bankapp-postgres-migrate`** — a one-off migration runner used by the `make migrate-*` targets

## API

### Users
- `POST /users` — create a user
- `GET /users` — list users
- `GET /users/{id}` — get a user by ID
- `PATCH /users/{id}` — update a user
- `DELETE /users/{id}` — delete a user

### Transactions
- `POST /transactions/deposit` — deposit money
- `POST /transactions/withdrawal` — withdraw money
- `POST /transactions/transfer` — transfer money between users

### History
- `GET /transactions/all/{user_id}` — list all operations for a user
- `GET /transactions/transfer/{user_id}` — list transfer operations for a user

## Graceful Shutdown

The app listens for `SIGINT`/`SIGTERM` and shuts down cleanly via `signal.NotifyContext`, allowing in-flight requests to complete before exiting.