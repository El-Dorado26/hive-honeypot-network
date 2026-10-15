
# US-03: Honeypot Registration and Management

## User Story

As a system administrator, I want to register and manage honeypots so that HIVE can identify and manage its configured event sources.

## Acceptance Criteria

### Scenario 1: Register a honeypot
Given an authenticated administrator has permission to manage honeypots
And the administrator provides valid required honeypot details
When the administrator submits the registration request
Then the system registers the honeypot
And assigns it a unique identifier

### Scenario 2: Reject invalid registration details
Given an authenticated administrator has permission to manage honeypots
When the administrator submits a registration request with missing or invalid required fields
Then the system rejects the request
And reports that the supplied details are invalid

### Scenario 3: Prevent unauthorized registration
Given an authenticated user does not have permission to manage honeypots
When the user submits a honeypot registration request
Then the system denies the request
And does not register the honeypot

