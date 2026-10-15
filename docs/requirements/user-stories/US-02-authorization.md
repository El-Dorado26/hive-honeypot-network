
# US-02: Role-Based Authorization

## User Story

As a system administrator, I want HIVE to enforce role-based access control so that users can only perform actions permitted by their assigned roles.

## Acceptance Criteria

### Scenario 1: Allow an authorized action
Given an authenticated user has a role permitted to perform an action
When the user requests that action
Then the system allows the request

### Scenario 2: Deny an unauthorized action
Given an authenticated user does not have permission to perform an action
When the user requests that action
Then the system denies the request
And does not perform the restricted action

### Scenario 3: Deny access to protected resources without authentication
Given a client is not authenticated
When the client requests a protected resource
Then the system returns HTTP 401
And does not disclose protected data
