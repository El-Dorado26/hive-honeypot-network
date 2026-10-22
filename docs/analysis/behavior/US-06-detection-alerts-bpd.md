# US-06: Detection Rules and Alerts

## 1. Purpose

This behavioral model describes how HIVE evaluates accepted events against enabled detection rules, generates alerts for matching events, and lets authorized users retrieve alert information.

**Primary actor:** Security analyst

**Supporting component:** HIVE detection engine, detection rules, event store and alert store

## 2. Preconditions

- At least one detection rule is configured and enabled, with defined conditions.
- An accepted event is available for evaluation (see US-04 and US-05).
- Alert storage is available.

## 3. Main Success Scenario

1. An accepted event is submitted to the detection engine for evaluation.
2. The detection engine loads the enabled detection rules.
3. The engine evaluates the event against the conditions of each enabled rule.
4. For a rule whose conditions the event satisfies, the engine generates an alert associated with the matching event.
5. The engine records the alert details.
6. An authorized analyst later requests the alert.
7. HIVE verifies that the user is authorized and returns the alert's defined details, including a reference to the associated event where available.

## 4. Alternative Scenarios

- Non-matching event: the event does not satisfy a rule's conditions. HIVE does not generate an alert from that rule for the event.
- Disabled rule: a disabled rule is not evaluated and generates no alerts.

Users without permission to retrieve alerts are handled by role-based authorization (US-02).

## 5. Business Process Diagrams

**Notation:** Each box is a swimlane for one participant. Rounded shapes are start events, double circles are end events, rectangles are tasks, diamonds are decisions, and cylinders are data stores. Dotted arrows show data read from or written to a data store.

### 5.1 Event evaluation and alert generation (Scenarios 1 and 2)

```mermaid
flowchart TB
    subgraph P["Event Processor"]
        P1(["Start: accepted event available"])
        P2["Submit event for evaluation"]
    end

    subgraph E["Detection Engine"]
        E1["Load enabled detection rules"]
        E2{"More enabled rules to evaluate?"}
        E3["Evaluate event against rule conditions"]
        E4{"Event satisfies the conditions?"}
        E5["Record alert linked to the matching event"]
        E6["Generate no alert from this rule"]
        EE((("End: evaluation complete")))
    end

    subgraph D["Rule and Alert Database"]
        D1[("Detection rules")]
        D2[("Alerts")]
    end

    P1 --> P2 --> E1
    E1 -.->|"read"| D1
    D1 -.->|"enabled rules"| E2
    E2 -->|"Yes"| E3 --> E4
    E4 -->|"Yes"| E5 --> E2
    E4 -->|"No"| E6 --> E2
    E5 -.->|"write alert"| D2
    E2 -->|"No"| EE
```

### 5.2 Alert retrieval (Scenario 3)

```mermaid
flowchart TB
    subgraph A["Analyst"]
        A1(["Start: analyst needs alert details"])
        A2["Request alert"]
        A3["Review alert details"]
        A4["Receive denial"]
        AE1((("End: alert reviewed")))
        AE2((("End: access denied")))
    end

    subgraph B["Backend API"]
        B1{"User authorized to retrieve alerts?"}
        B2["Retrieve alert and associated event reference"]
        B3["Return alert details with event reference where available"]
        B4["Deny request, handled by US-02"]
    end

    subgraph D["Rule and Alert Database"]
        D1[("Alerts")]
    end

    A1 --> A2 --> B1
    B1 -->|"Yes"| B2
    B2 -.->|"read"| D1
    D1 -.->|"alert details"| B3
    B3 --> A3 --> AE1
    B1 -->|"No"| B4 --> A4 --> AE2
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Generate an alert for a matching event, shown in the evaluation business process diagram (5.1).
- Scenario 2: Do not generate an alert for a non-matching event, shown in the evaluation business process diagram (5.1).
- Scenario 3: Preserve alert information, shown in the alert retrieval business process diagram (5.2).

The model is consistent with US-06.
