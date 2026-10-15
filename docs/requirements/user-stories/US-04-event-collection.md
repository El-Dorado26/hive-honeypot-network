
# US-04: Event Collection

## User Story

As a security analyst, I want HIVE to collect security events from configured honeypots so that activity captured by those honeypots can be analyzed centrally.

## Acceptance Criteria

### Scenario 1: Accept a valid event
Given a honeypot is registered with HIVE
When the honeypot submits an event containing all required fields in the expected format
Then HIVE accepts the event
And passes it to the event-processing workflow

### Scenario 2: Reject a malformed event
Given a client submits an event to HIVE
When the event payload is malformed or lacks required fields
Then HIVE rejects the event
And returns an appropriate client error response

### Scenario 3: Reject an unknown event source
Given an event is submitted with an unrecognized honeypot identifier
When HIVE validates the event source
Then HIVE rejects the event
And does not accept it as an event from a registered honeypot
