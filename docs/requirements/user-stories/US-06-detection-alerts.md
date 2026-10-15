
# US-06: Detection Rules and Alerts

## User Story

As a security analyst, I want HIVE to evaluate collected events against configured detection rules so that potentially suspicious activity can generate alerts.

## Acceptance Criteria

### Scenario 1: Generate an alert for a matching event
Given a detection rule is enabled
And an accepted event satisfies that rule's defined conditions
When HIVE evaluates the event
Then HIVE generates an alert associated with the matching event
And records the alert details

### Scenario 2: Do not generate an alert for a non-matching event
Given a detection rule is enabled
And an accepted event does not satisfy that rule's conditions
When HIVE evaluates the event
Then HIVE does not generate an alert from that rule for the event

### Scenario 3: Preserve alert information
Given HIVE has generated an alert
When an authorized user retrieves the alert
Then HIVE provides the alert's defined details
And includes a reference to the associated event where available
