# US-03 — Honeypot Registration

**As a** HIVE operator, **I want** to register a configured honeypot, **so that** its events can be associated with the correct service.

## Acceptance criteria
- An authorized operator can register a honeypot with an identifier and service type.
- Required fields are validated.
- Invalid or duplicate registrations are rejected with a clear error.
- Registration changes are recorded in an audit log.

**Priority:** Must have  
**Related test:** `tests/requirements/US-03-honeypot-registration_test.go`
