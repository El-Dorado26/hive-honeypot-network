# US-01 — User Authentication

**As a** registered HIVE team member, **I want** to sign in, **so that** protected platform functions are not available to unauthenticated users.

## Acceptance criteria
- A user with valid credentials can authenticate.
- Invalid credentials are rejected without revealing whether the username exists.
- Protected endpoints reject requests without valid authentication.
- Passwords are not stored as plaintext.
- Authentication outcomes are logged without recording passwords or other secrets.

**Priority:** Must have  
**Related test:** `tests/requirements/US-01-authentication_test.go`
