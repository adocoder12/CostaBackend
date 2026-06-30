# costaBackend

A backend REST API for managing **touristic apartments** — built in Go, backed by PostgreSQL (via Supabase).

CostaBackend handles the data and business logic behind a short-term rental / holiday-apartment platform: properties, availability, bookings and the people involved. It's designed to be a clean, containerized, production-shaped Go service rather than a throwaway prototype.

> **Status:** Work in progress / personal project. Actively iterating.

---

## Tech stack

| Layer            | Technology                      |
| ---------------- | ------------------------------- |
| Language         | Go                              |
| Database         | PostgreSQL (hosted on Supabase) |
| Containerization | Docker + Docker Compose         |
| Linting          | golangci-lint                   |
| Build / tasks    | Makefile                        |

---

## Features

> _Update this list to match what the API actually does today — keep it honest, only list what's implemented._

- Manage apartments / properties (create, read, update, delete)
- Manage bookings and availability
- Guest / user records
- RESTful JSON API
- Containerized local setup with a single command

---

## Project structure

```
CostaBackend/
├── cmd/
│   └── api/            # Application entry point (main.go)
├── internal/           # Private application code (handlers, models, db, config)
├── services/           # Business logic / service layer
├── docker/             # Docker-related files
├── Dockerfile          # Container image definition
├── docker-compose.yml  # Local multi-container setup (API + database)
├── golangci.yml        # Linter configuration
├── makefile            # Common commands (build, run, lint, test)
├── go.mod / go.sum     # Go module dependencies
└── README.md
```

The layout follows the standard Go convention: `cmd/` for entry points, `internal/` for code that shouldn't be imported by other projects, and a separate service layer for business logic.

---

## Getting started

### Prerequisites

- [Go](https://go.dev/dl/) (version as set in `go.mod`)
- [Docker](https://www.docker.com/) and Docker Compose
- A PostgreSQL database — e.g. a [Supabase](https://supabase.com/) project

### 1. Clone the repository

```bash
git clone https://github.com/adocoder12/CostaBackend.git
cd CostaBackend
```

### 2. Configure environment variables

Create a `.env` file in the project root.

> _Replace these with the actual variable names your code reads (check `internal/` config or wherever you load env vars). These are placeholders._

```env
# Database connection (Supabase / PostgreSQL)
DATABASE_URL=postgresql://USER:PASSWORD@HOST:PORT/DBNAME

# Server
PORT=8080
```

### 3. Run with Docker Compose

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080` (adjust to your configured port).

### 4. Or run locally with the Makefile

```bash
make run
```

> _List your real Makefile targets here — e.g. `make build`, `make run`, `make lint`, `make test`._

---

## API endpoints

> _Fill in your real routes. Example shape below — replace with what's in your handlers._

| Method | Endpoint           | Description            |
| ------ | ------------------ | ---------------------- |
| GET    | `/apartments`      | List all apartments    |
| GET    | `/apartments/{id}` | Get a single apartment |
| POST   | `/apartments`      | Create a new apartment |
| PUT    | `/apartments/{id}` | Update an apartment    |
| DELETE | `/apartments/{id}` | Delete an apartment    |
| GET    | `/bookings`        | List bookings          |
| POST   | `/bookings`        | Create a booking       |

---

## Development

Run the linter (configured in `golangci.yml`):

```bash
golangci-lint run
```

Run tests:

```bash
make test
```

---

## Roadmap

> _Optional but recommended — shows direction. A couple of honest "next steps" is enough._

- [ ] Add automated tests
- [ ] Deploy (e.g. to a cloud provider)
- [ ] Authentication & authorization
- [ ] API documentation (OpenAPI / Swagger)

---

## About

Built by [Adonay D'Agosto](https://github.com/adocoder12) as part of deliberately moving from frontend into full-stack / backend development with Go.

🌐 [adonaydagosto.com](https://www.adonaydagosto.com)
