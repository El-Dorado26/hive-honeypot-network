
# US-07: Dashboard and Event Search

## User Story

As a security analyst, I want to view a dashboard and search collected events so that I can investigate activity captured by HIVE.

## Acceptance Criteria

### Scenario 1: View the dashboard
Given an authenticated user has permission to view the dashboard
When the user opens the dashboard
Then HIVE displays the available event and alert summaries

### Scenario 2: Search for events
Given an authenticated user has permission to view events
And HIVE contains stored events
When the user submits a search using supported search criteria
Then HIVE returns the events matching those criteria

### Scenario 3: Return no results
Given an authenticated user has permission to view events
When the user submits a valid search that matches no stored events
Then HIVE displays an empty result
And does not present unrelated events as matches

### Scenario 4: Deny unauthorized event access
Given a user does not have permission to view events
When the user requests event data
Then HIVE denies the request
And does not disclose the protected event data

