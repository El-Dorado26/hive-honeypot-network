# US-06 : Basic Detection and Alerts

**As a** security analyst, **I want** configured rules to identify suspicious activity and generate alerts, **so that** repeated or suspicious interactions are easier to notice.

## Acceptance criteria
- A configured rule can match a defined suspicious pattern, such as repeated failed authentication attempts.
- A matching event or event sequence produces an alert.
- A non-matching event does not produce that alert.
- An alert includes a timestamp, rule/reason, and related event or source reference.
- Detection errors are logged without crashing event ingestion.

**Priority:** Must have  
**Related test:** `tests/requirements/US-06-detection-alerts_test.go`
