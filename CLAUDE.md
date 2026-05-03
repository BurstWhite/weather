# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Backend (Go 1.26+)
cd backend && go run ./cmd/server/        # Start API server (default :8080)
cd backend && go build -o weather-server ./cmd/server/

# Frontend (Node 22+)
cd frontend && npm run dev                # Vite dev server on :5173, proxies /api to :8080
cd frontend && npm run build              # Production build to frontend/dist/

# Database
# SQLite file is backend/weather.db — auto-created on first run. Delete to reset.
```

## Environment

Backend requires `backend/.env` with at minimum `QWEATHER_API_KEY`. Optional: `QWEATHER_API_HOST` (custom API proxy host — the QWeather client uses this to construct URLs when set), `PORT` (defaults to 8080).

The seed data (`backend/internal/database/seed.go`) populates 200+ Chinese cities on first database initialization. Skip this behavior by pre-filling the `cities` table before first run.

## Architecture

### Backend — Go + Gin + SQLite

The server is wired in `backend/cmd/server/main.go`: loads config, initializes SQLite, creates the QWeather client, registers 5 routes under `/api` with CORS for `localhost:5173`.

```
handler  →  service  →  qweather.Client (external API)
                ↓
          database.DB (SQLite cache)
```

**QWeather client** (`internal/qweather/client.go`): thin HTTP client for three QWeather endpoints (city search, 7-day forecast, AQI by lat/lon). If `QWEATHER_API_HOST` is set, it hits `https://<host>` instead of the official domains — this is used for the custom proxy endpoint.

**Caching** (`internal/service/weather.go`): weather data and AQI are cached in SQLite with a 1-hour TTL. On cache miss, fresh data is fetched from QWeather and upserted. On API failure, stale cache is returned if available.

**Database** (`internal/database/`): SQLite via `modernc.org/sqlite` (CGo-free). Three tables: `cities`, `weather_cache`, `air_cache`. `DB.SetMaxOpenConns(1)` — SQLite serializes all writes.

**City service** (`internal/service/city.go`): standalone functions (not a struct) operating directly on `database.DB`. `AddCity` uses `INSERT OR IGNORE` — duplicate `location_id` values are silently skipped.

### Frontend — Vue 3 + Vite + Pinia

Single-page app with one route (`/` → `Home.vue`). The `weatherStore` Pinia store (`src/stores/weather.js`) owns all state: cities list, selected city, and current weather data.

The Axios API client (`src/api/index.js`) sets `baseURL: '/api'` and 15s timeout. In dev, Vite proxies `/api` to the backend. `addCityTemp` / `replaceCity` in the store handle the optimistic-add pattern: a temporary city is inserted immediately before the server confirms it.
