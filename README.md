# Weather App

A weather application with a **Vue 3 + Vite** frontend and a **Go (Gin)** backend, powered by the [QWeather API](https://dev.qweather.com/).

## Project Structure

```
weather-app/
├── backend/                  # Go REST API
│   ├── cmd/server/main.go    # Entry point & routes
│   ├── internal/
│   │   ├── config/           # Environment config
│   │   ├── database/         # SQLite init & migrations
│   │   ├── handler/          # HTTP handlers
│   │   ├── model/            # Data models
│   │   ├── qweather/         # QWeather API client
│   │   └── service/          # Business logic
│   ├── go.mod
│   └── go.sum
└── frontend/                 # Vue 3 SPA
    ├── src/
    │   ├── api/              # Axios API client
    │   ├── components/       # Vue components
    │   ├── router/           # Vue Router
    │   ├── stores/           # Pinia state management
    │   └── views/            # Page views
    ├── index.html
    ├── package.json
    └── vite.config.js
```

## API Routes

| Method | Route              | Description            |
|--------|--------------------|------------------------|
| GET    | `/api/cities`      | List saved cities      |
| POST   | `/api/cities`      | Add a city             |
| DELETE | `/api/cities/:id`  | Remove a city          |
| GET    | `/api/search?q=`   | Search cities by name  |
| GET    | `/api/weather/:id` | Get 7-day forecast & AQI |

## Prerequisites

- **Go** 1.26+
- **Node.js** 22+
- **QWeather API key** — register at [dev.qweather.com](https://dev.qweather.com/)

## Setup

### 1. Configure API Key

Create `backend/.env`:

```env
QWEATHER_API_KEY=your_api_key_here
PORT=8080
```

### 2. Install Dependencies

```bash
# Backend
cd backend
go mod download

# Frontend
cd frontend
npm install
```

### 3. Run

Start the backend:

```bash
cd backend
go run ./cmd/server/
```

Start the frontend in another terminal:

```bash
cd frontend
npm run dev
```

The frontend dev server runs at `http://localhost:5173` and proxies `/api` requests to the backend at `http://localhost:8080`.

### Build

```bash
# Backend
cd backend && go build -o weather-server ./cmd/server/

# Frontend
cd frontend && npm run build
```

## Tech Stack

| Layer    | Technology                     |
|----------|--------------------------------|
| Frontend | Vue 3, Vite, Pinia, Vue Router, Axios |
| Backend  | Go, Gin, SQLite                |
| Weather  | QWeather API                   |
