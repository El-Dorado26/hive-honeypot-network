
# US-05: Event Normalization and Storage

## User Story

As a security analyst, I want HIVE to normalize and store accepted security events so that events from different honeypots can be searched and analyzed consistently.

## Acceptance Criteria

### Scenario 1: Normalize an accepted event
Given HIVE has accepted a valid event
When the event is processed
Then HIVE maps the available event data to the defined common event schema
And preserves the required information from the original event

### Scenario 2: Store a normalized event
Given an accepted event has been normalized successfully
When HIVE persists the event
Then the event is stored in the configured database
And the stored event can be retrieved using the supported search functionality

### Scenario 3: Handle a storage failure
Given HIVE has normalized an accepted event
When the database cannot store the event
Then HIVE reports or records the processing failure through the defined error-handling mechanism
And does not report the event as successfully stored

