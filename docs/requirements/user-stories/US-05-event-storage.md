# US-05 — Event Normalization and Storage

**As a** security analyst, **I want** events normalized into a common structure and stored, **so that** events from different honeypots can be searched consistently.

## Acceptance criteria
- Supported honeypot event formats are mapped to the common event model.
- A valid normalized event is persisted in PostgreSQL.
- Stored events can be retrieved by authorized users.
- Invalid records do not silently appear as successfully stored.
- Event timestamps and source identifiers are retained.

**Priority:** Must have  
**Related test:** `tests/requirements/US-05-event-storage_test.go`
