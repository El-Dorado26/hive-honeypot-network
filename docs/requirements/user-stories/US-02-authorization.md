# US-02 — Role-Based Authorization

**As a** platform administrator, **I want** access to be restricted by role, **so that** users can only perform permitted actions.

## Acceptance criteria
- A request from an authenticated user with permission is allowed.
- A request from an authenticated user without permission is denied.
- Authorization is enforced by the backend, not only by hiding frontend controls.
- Denied access attempts are auditable.

**Priority:** Must have  
**Related test:** `tests/requirements/US-02-authorization_test.go`
