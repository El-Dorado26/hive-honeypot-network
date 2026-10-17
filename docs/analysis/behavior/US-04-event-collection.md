# US-04: Event Collection

## 1. Purpose

This behavioral model describes how HIVE accepts security events from registered honeypots, and rejects malformed events and events from unknown sources.

**Primary actor:** Registered honeypot

**Supporting component:** HIVE event collection API, honeypot registry and event-processing workflow

## 2. Preconditions

- For an event to be accepted, the submitting honeypot is registered with HIVE (see US-03).
- The expected event format and required fields are defined.
- The event collection endpoint is available.

## 3. Main Success Scenario

1. The honeypot submits an event to the event collection endpoint.
2. The backend validates that the payload is well formed and contains all required fields.
3. The backend checks the honeypot identifier against the registered honeypots.
4. The backend accepts the event.
5. The backend passes the accepted event to the event-processing workflow (US-05).
6. The backend acknowledges the submission to the honeypot.

## 4. Alternative Scenarios

- Malformed event: the payload is malformed or lacks required fields. The backend rejects the event and returns an appropriate client error response.
- Unknown event source: the honeypot identifier is not recognized. The backend rejects the event and does not accept it as an event from a registered honeypot.

Source validation depends on the honeypot registry created through US-03.

## 5. Business Process Diagram

**Notation:** Each box is a swimlane for one participant. Rounded shapes are start events, double circles are end events, rectangles are tasks, diamonds are decisions, and cylinders are data stores. Dotted arrows show data read from or written to a data store.

### 5.1 Event submission (Scenarios 1, 2 and 3)

```mermaid
flowchart TB
    subgraph H["Honeypot or Client"]
        H1(["Start: event needs to be reported"])
        H2["Submit event"]
        H3["Receive client error response"]
        H4["Receive rejection"]
        H5["Receive acknowledgement"]
        HE1((("End: event rejected")))
        HE2((("End: source rejected")))
        HE3((("End: event accepted")))
    end

    subgraph A["Event Collection API"]
        A1{"Well formed with all required fields?"}
        A2["Reject event with client error response"]
        A3["Look up honeypot identifier"]
        A4{"Honeypot registered?"}
        A5["Reject event, not accepted from a registered honeypot"]
        A6["Accept event"]
        A7["Acknowledge submission"]
    end

    subgraph R["Honeypot Registry"]
        R1[("Registered honeypots")]
    end

    subgraph P["Event Processing Workflow"]
        P1["Receive accepted event for normalization and storage"]
    end

    H1 --> H2 --> A1
    A1 -->|"No"| A2 --> H3 --> HE1
    A1 -->|"Yes"| A3
    A3 -.->|"look up"| R1
    R1 -.->|"registered or not found"| A4
    A4 -->|"No"| A5 --> H4 --> HE2
    A4 -->|"Yes"| A6 --> P1
    A6 --> A7 --> H5 --> HE3
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Accept a valid event, shown in the event submission business process diagram (5.1).
- Scenario 2: Reject a malformed event, shown in the event submission business process diagram (5.1).
- Scenario 3: Reject an unknown event source, shown in the event submission business process diagram (5.1).

The model is consistent with US-04.
