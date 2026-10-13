# HIVE Backend API Documentation

## 1. Overview

HIVE provides a Go REST API for receiving and retrieving honeypot security events. Events are stored in PostgreSQL.

The API runs locally on port `8080` by default.

**Base URL:** `http://localhost:8080`

## 2. Prerequisites

- Go installed
- PostgreSQL installed and running
- HIVE PostgreSQL database and application role configured
- Required Go dependencies downloaded

## 3. Environment Configuration

The backend reads its database configuration from environment variables.

| Variable | Purpose | Default |
|---|---|---|
| `HIVE_DB_HOST` | PostgreSQL host | `localhost` |
| `HIVE_DB_PORT` | PostgreSQL port | `5432` |
| `HIVE_DB_USER` | Database username | `hive_app` |
| `HIVE_DB_NAME` | Database name | `hive_db` |
| `HIVE_DB_PASSWORD` | Database password | Required |
| `HIVE_DB_SSLMODE` | PostgreSQL SSL mode | `disable` |

Setting the database password in the terminal before starting the API:

```bash
export HIVE_DB_PASSWORD='YOUR_DATABASE_PASSWORD'
```

Replacing the placeholder with my database password.

The default SSL mode is intended for local development. Production deployments require an appropriate secure database configuration.

## 4. Running the Backend

From the repository root, configuring the environment and starting the API:

```bash
export HIVE_DB_PASSWORD='YOUR_DATABASE_PASSWORD'
go run ./cmd/api
```

The server listens on port `8080`.

## 5. Endpoints

### 5.1 Health Check

**Method:** `GET`

**Path:** `/health`

Checks whether the API responds.

Example:

```bash
curl -i http://localhost:8080/health
```

Expected response: `200 OK`

```json
{
  "status": "ok",
  "service": "HIVE API"
}
```

### 5.2 Database Health Check

**Method:** `GET`

**Path:** `/health/db`

Checks connectivity to PostgreSQL.

Example:

```bash
curl -i http://localhost:8080/health/db
```

Expected response when connected: `200 OK`

```json
{
  "status": "ok",
  "database": "connected"
}
```

A database connectivity failure returns an error status.

### 5.3 Create an Event

**Method:** `POST`

**Path:** `/api/v1/events`

**Content-Type:** `application/json`

Stores a valid honeypot event in PostgreSQL.

Example request:

```bash
curl -i -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "source": "ssh-honeypot",
    "event_type": "login_attempt",
    "occurred_at": "2026-10-13T12:00:00+03:00",
    "event_data": {
      "username": "admin",
      "result": "failed",
      "source_ip": "192.0.2.10"
    }
  }'
```

The event must include a non-empty `source`, a non-empty `event_type`, a valid `occurred_at` timestamp, and an `event_data` JSON object.

A successful request returns `201 Created`, including the new event ID and its receipt timestamp.

Invalid requests return an appropriate client error. Unsupported methods return `405 Method Not Allowed`.

### 5.4 List Events

**Method:** `GET`

**Path:** `/api/v1/events`

Retrieves stored events in descending event-time order.

Example:

```bash
curl -i http://localhost:8080/api/v1/events
```

#### Pagination

The endpoint accepts these optional query parameters:

| Parameter | Default | Allowed values |
|---|---:|---|
| `page` | `1` | Integer from `1` to `1000000` |
| `limit` | `20` | Integer from `1` to `100` |

Example:

```bash
curl -i "http://localhost:8080/api/v1/events?page=1&limit=10"
```

The response includes the event list, current page, page size, and whether more results are available.

Example response:

```json
{
  "events": [],
  "page": 1,
  "limit": 10,
  "has_more": false
}
```

The empty list above is illustrative. Actual results depend on the events stored in the database.

Invalid pagination parameters return `400 Bad Request`. Database retrieval failures return a generic server error without exposing internal database details.

## 6. Testing

Run the automated Go tests from the repository root:

```bash
go test ./...
```

Check the health endpoints:

```bash
curl -i http://localhost:8080/health
curl -i http://localhost:8080/health/db
```

Check event retrieval:

```bash
curl -i "http://localhost:8080/api/v1/events?page=1&limit=10"
```

## 7. Current Limitations

- Authentication and role-based access control have not yet been implemented.
- The API is intended for local development and testing until access controls and deployment security are addressed.
- Honeypot integrations, detection rules, and alerting are separate development tasks.
- The health endpoints are operational checks, not a complete security monitoring system.
