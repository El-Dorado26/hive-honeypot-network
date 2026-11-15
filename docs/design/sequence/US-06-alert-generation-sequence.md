# US-06: Detection Rules and Alerts: Sequence Diagrams

## 1. Purpose

These sequence diagrams describe how HIVE evaluates accepted events against enabled detection rules, generates and stores alerts for matching events, and allows authorized analysts to retrieve alert information.

**Primary actor:** Security Analyst  
**Supporting components:** Event Processor, Detection Engine, Detection Rule Store, Alert Store, HIVE Backend API

**Source:** `docs/analysis/behavior/US-06-detection-alerts-bpd.md`

## 2. Sequence Diagram 1: Event Evaluation and Alert Generation

This diagram covers the evaluation of accepted events against enabled detection rules. An alert is generated and stored when an event satisfies a rule's conditions. Disabled rules are excluded from evaluation, and non-matching events do not generate alerts.

```mermaid
sequenceDiagram
    autonumber
    participant Processor as Event Processor
    participant Engine as Detection Engine
    participant RuleDB as Detection Rule Store
    participant AlertDB as Alert Store

    Processor->>Engine: Submit accepted event for evaluation
    Engine->>RuleDB: Retrieve enabled detection rules
    RuleDB-->>Engine: Return enabled rules

    loop For each enabled rule
        Engine->>Engine: Evaluate event against rule conditions

        alt Event satisfies rule conditions
            Engine->>Engine: Generate alert for matching event
            Engine->>AlertDB: Store alert with rule and event references
            AlertDB-->>Engine: Confirm alert stored
        else Event does not satisfy rule conditions
            Note over Engine,AlertDB: No alert is generated for this rule
        end
    end

    Engine-->>Processor: Confirm evaluation completed
```

## 3. Sequence Diagram 2: Alert Retrieval

This diagram covers the analyst's request to retrieve alert information. The backend checks authorization before accessing and returning the requested alert details.

```mermaid
sequenceDiagram
    autonumber
    actor Analyst as Security Analyst
    participant Backend as HIVE Backend API
    participant AlertDB as Alert Store

    Analyst->>Backend: Request alert details
    Backend->>Backend: Verify authentication and alert retrieval permission

    alt User is not authorized
        Backend-->>Analyst: Return access denied
        Note over Backend,Analyst: Alert information is not returned
    else User is authorized
        Backend->>AlertDB: Retrieve requested alert
        AlertDB-->>Backend: Return alert details and event reference
        Backend-->>Analyst: Return alert information
        Analyst->>Analyst: Review alert details
    end
```

## 4. Acceptance Criteria Mapping

| US-06 Scenario | Description | Sequence Diagram | Branch |
|---|---|---|---|
| Scenario 1 | Generate an alert when an event matches a detection rule | Diagram 1 | Event satisfies rule conditions |
| Scenario 2 | Do not generate an alert when an event does not match a detection rule | Diagram 1 | Event does not satisfy rule conditions |
| Scenario 3 | Preserve alert information and make it retrievable | Diagram 2 | User is authorized → Alert details retrieved |

Disabled rules are excluded when the Detection Engine retrieves enabled rules in Diagram 1. Therefore, disabled rules are not evaluated and cannot generate alerts.

## 5. Design Notes

- The Detection Engine receives accepted events from the event-processing workflow after US-04 and US-05.
- Only enabled detection rules are retrieved for evaluation.
- Each enabled rule is evaluated independently against the event.
- An alert is generated only when the event satisfies a rule's conditions.
- Each generated alert is associated with its matching detection rule and event.
- Alert details are stored in the Alert Store.
- The backend verifies authorization before returning alert information to an analyst. Role-based authorization is defined separately in US-02.

The diagrams are based on the US-06 behavioral model and its three acceptance scenarios.
