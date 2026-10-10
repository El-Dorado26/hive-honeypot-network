# US-07 : Event and Alert Dashboard

**As a** security analyst, **I want** to view collected events and generated alerts in the web dashboard, **so that** I can understand honeypot activity from one place.

## Acceptance criteria
- An authorized user can view a list of stored events.
- An authorized user can view generated alerts.
- Event and alert details show relevant timestamps, service type, and source information.
- Loading, empty, and error states are presented clearly.
- Unauthorized users cannot retrieve protected dashboard data through direct API requests.

**Priority:** Must have  
**Related test:** `tests/requirements/US-07-dashboard_test.go`
