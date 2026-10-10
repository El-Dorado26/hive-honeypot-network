# US-04 : Event Collection

**As a** security analyst, **I want** events from connected honeypots to reach the HIVE backend, **so that** activity can be reviewed centrally.

## Acceptance criteria
- The backend accepts a valid event from a registered honeypot.
- Events include at minimum honeypot ID, service type, event type, source IP, action, and timestamp.
- Malformed events are rejected safely.
- Event ingestion does not execute commands or payloads received from a remote source.
- Accepted events can be traced to their source honeypot.

**Priority:** Must have  
**Related test:** `tests/requirements/US-04-event-collection_test.go`
