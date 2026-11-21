
# US-05: Event Normalization and Storage: Sequence Diagram

## 1. Purpose

This sequence diagram describes how HIVE processes an accepted event, maps its data to a common event schema, preserves the required information, and stores the normalized event in the database.

It also covers storage failures and an additional normalization-failure case.

**Primary actor:** HIVE Event Processing Workflow  
**Supporting components:** Event Collection API, Event Processor, Event Database, Error Handling

**Source:** `docs/analysis/behavior/US-05-event-storage-bpd.md`

## 2. Sequence Diagram: Event Normalization and Storage

This diagram covers the three US-05 acceptance scenarios: normalizing an accepted event, storing the normalized event, and handling a storage failure.

```mermaid
sequenceDiagram
    autonumber
    participant Collection as Event Collection API
    participant Processor as Event Processor
    participant DB as Event Database
    participant Error as Error Handling

    Collection->>Processor: Forward accepted event
    Processor->>Processor: Map event data to common schema

    alt Event cannot be normalized
        Processor->>Error: Report normalization failure
        Error-->>Processor: Confirm failure reported
        Note over Processor,DB: Event is not stored or reported as successfully stored
    else Event is normalized successfully
        Processor->>Processor: Preserve required original event information
        Processor->>DB: Store normalized event

        alt Database stores event successfully
            DB-->>Processor: Confirm successful storage
            Processor->>Processor: Mark event as stored
            Note over Processor,DB: Stored event can later be retrieved through US-07
        else Database storage fails
            DB-->>Processor: Report storage failure
            Processor->>Error: Report or record storage failure
            Error-->>Processor: Confirm failure reported
            Note over Processor,DB: Event is not reported as successfully stored
        end
    end
```

## 3. Acceptance Criteria Mapping

| US-05 Scenario | Description | Sequence Diagram | Branch |
|---|---|---|---|
| Scenario 1 | Normalize an accepted event into the common event schema | Diagram 1 | Event is normalized successfully |
| Scenario 2 | Store a normalized event in the database | Diagram 1 | Database stores event successfully |
| Scenario 3 | Handle a database storage failure | Diagram 1 | Database storage fails |

The normalization-failure branch is an additional robustness case. It is not required by the US-05 acceptance criteria.

## 4. Design Notes

- The diagram begins after the Event Collection API has accepted an event through US-04.
- The Event Processor maps incoming event data to the common event schema.
- Required information from the original event must be preserved during normalization.
- The normalized event is stored in the Event Database.
- A successful database response is required before HIVE reports the event as stored.
- Storage and normalization failures must be reported or recorded through the error-handling mechanism.
- Event retrieval and search are covered separately by US-07; this diagram does not implement or specify the search mechanism.

The diagram is based on the US-05 behavioral model and its three acceptance scenarios.
