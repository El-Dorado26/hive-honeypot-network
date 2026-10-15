
# US-01: User Authentication

## User Story

As a registered user, I want to sign in securely so that I can access the HIVE features permitted for my account.

## Acceptance Criteria

### Scenario 1: Successful login

- **Given** the user has an active registered account
- **When** the user submits valid credentials
- **Then** the system authenticates the user and grants access to permitted features

### Scenario 2: Invalid credentials

- **Given** the user is on the login page
- **When** the user submits an incorrect username or password
- **Then** the system denies access and displays a generic error message

### Scenario 3: Unauthenticated API request

- **Given** the client has no valid authentication
- **When** the client requests a protected API endpoint
- **Then** the system returns HTTP 401 without exposing protected data

### Scenario 4: Disabled account

- **Given** the user's account is disabled
- **When** the user submits otherwise valid credentials
- **Then** the system denies authentication

