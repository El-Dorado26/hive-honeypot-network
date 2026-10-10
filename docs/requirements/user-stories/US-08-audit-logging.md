# US-08 — Audit Logging

**As a** platform administrator, **I want** security-relevant platform actions to be recorded, **so that** access and administrative activity can be reviewed.

## Acceptance criteria
- Authentication failures and access denials are recorded.
- Honeypot registration and other administrative changes are recorded.
- Audit entries contain an event type, timestamp, and relevant actor/resource identifier when available.
- Passwords, tokens, and secret payloads are not written to audit logs.
- Audit records can be viewed only by authorized users.

**Priority:** Must have  
**Related test:** `tests/requirements/US-08-audit-logging_test.go`
