# case5

Overview
--------
This project is a minimal paste storage application composed with Docker Compose.
Services:
- `db` — PostgreSQL 14 (persistent data stored in `./db_data`)
- `backend` — Go (net/http + pgx) REST API for /pastes
- `frontend` — Static HTML/JS served by Nginx (port 3000). Frontend also contains an Nginx proxy for `/api/`.
- `gateway` — Nginx reverse-proxy that serves the frontend and proxies `/api/` to the backend (port 80).

Prerequisites
-------------
- Docker and Docker Compose installed on the host
- At least a few GB free disk space (build may consume several hundred MB)

Repository layout (relevant)
- `docker-compose.yml`
- `backend/` — Go source and Dockerfile
- `frontend/` — `index.html`, `default.conf`, Dockerfile
- `gateway/` — nginx.conf, Dockerfile
- `db_data/` — host directory used as PostgreSQL data volume (created at runtime)

Ports
-----
- Gateway: http://localhost (host port 80)
- Frontend (direct): http://localhost:3000 (host port 3000)
- Backend (direct): http://localhost:8080 (host port 8080)
- DB: localhost:5432 (Postgres)

Quick start
-----------
1. Build and start everything:
   docker compose build
   docker compose up -d

2. Check running containers:
   docker compose ps

3. Tail logs if needed:
   docker compose logs -f backend
   docker compose logs -f gateway
   docker compose logs -f frontend
   docker compose logs -f db

How the routing works
--------------------
- Gateway listens on port 80 and proxies:
  - `/`  -> frontend service
  - `/api/` -> backend service (requests have `/api/` prefix stripped)
- Frontend service runs on its own container and also includes a local Nginx config so requests made to `http://localhost:3000/api/...` will be proxied to the backend at `backend:8080`.
- You can use either:
  - http://localhost/api/...  (via gateway)
  - http://localhost:3000/api/...  (via frontend's internal proxy)
  - http://localhost:8080/...  (direct to backend — CORS allowed by backend)


Common commands (examples)
--------------------------
- Build all services:
  docker compose build
- Build a single service (e.g. frontend):
  docker compose build frontend
- Start in background:
  docker compose up -d
- Stop and remove containers:
  docker compose down
- Recreate containers (with new builds):
  docker compose up -d --build
- View container status:
  docker compose ps
- Tail logs:
  docker compose logs -f backend
  docker compose logs -f gateway
  docker compose logs -f frontend
- Execute a command inside the frontend container (useful to test internal connectivity):
  docker compose exec frontend sh -c "apk add --no-cache curl >/dev/null 2>&1 || true; curl -sS http://backend:8080/pastes"
