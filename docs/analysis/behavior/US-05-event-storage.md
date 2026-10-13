# US-05: Event Normalization and Storage

## 1. Purpose

This behavioral model describes how HIVE normalizes accepted events into a common schema, stores them, and handles storage failures, so that events from different honeypots can be searched and analyzed consistently.

**Primary actor:** HIVE event-processing workflow

**Supporting component:** HIVE database and error-handling mechanism

## 2. Preconditions

- HIVE has accepted a valid event (see US-04).
- The common event schema is defined.
- The configured database is available, except in the storage failure scenario.

## 3. Main Success Scenario

1. The event-processing workflow receives an accepted event.
2. The workflow maps the available event data to the common event schema.
3. The workflow preserves the required information from the original event.
4. The workflow stores the normalized event in the configured database.
5. The database confirms that the event was stored.
6. The stored event can be retrieved using the supported search functionality (see US-07).

## 4. Alternative Scenarios

- Storage failure: the database cannot store the normalized event. HIVE reports or records the failure through the defined error-handling mechanism and does not report the event as successfully stored.

Additional robustness case (not required by the US-05 acceptance criteria):

- Normalization failure: if an accepted event cannot be mapped to the common schema, the failure is handled through the same error-handling mechanism and the event is not reported as stored.

## 5. UML Sequence Diagram

### 5.1 Normalization and storage (Scenarios 1, 2 and 3)

```mermaid
sequenceDiagram
    participant Collector as Event Collection
    participant Proc as Event Processor
    participant DB as Event Database
    participant Err as Error Handling

    Collector->>Proc: Pass accepted event
    Proc->>Proc: Map event data to common event schema
    Proc->>Proc: Preserve required information from original event
    Proc->>DB: Store normalized event

    alt Event stored successfully
        DB-->>Proc: Confirm storage
        Proc-->>Collector: Report event stored
    else Database cannot store the event
        DB-->>Proc: Storage failure
        Proc->>Err: Report or record processing failure
        Proc-->>Collector: Report failure (event not reported as stored)
    end
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Normalize an accepted event, shown in the normalization and storage sequence diagram (5.1).
- Scenario 2: Store a normalized event, shown in the normalization and storage sequence diagram (5.1). Retrieval through search is covered by US-07.
- Scenario 3: Handle a storage failure, shown in the normalization and storage sequence diagram (5.1).

The model is consistent with US-05.
