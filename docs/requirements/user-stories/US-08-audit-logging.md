
# US-08: Audit Logging

## User Story

As a system administrator, I want HIVE to record security-relevant actions so that important system activity can be reviewed and investigated.

## Acceptance Criteria

### Scenario 1: Record a successful security-relevant action
Given an authenticated user performs a defined security-relevant action
When the action completes successfully
Then HIVE creates an audit record containing the defined details
And associates the record with the responsible user where available

### Scenario 2: Record a failed security-relevant action
Given a user attempts a defined security-relevant action
When the action fails
Then HIVE records the failure when it falls within the defined audit policy
And records the available details needed for investigation

### Scenario 3: Protect audit logs from unauthorized access
Given a user does not have permission to review audit logs
When the user requests audit-log data
Then HIVE denies the request
And does not disclose the audit records

### Scenario 4: Avoid storing secrets in audit logs
Given HIVE records a security-relevant action
When the audit record is created
Then the record does not contain passwords, authentication secrets, or other explicitly prohibited secret values

