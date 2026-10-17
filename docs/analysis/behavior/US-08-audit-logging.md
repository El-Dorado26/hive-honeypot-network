# US-08: Audit Logging

## 1. Purpose

This behavioral model describes how HIVE records security-relevant actions, whether they succeed or fail, keeps secrets out of audit records, and protects audit logs from users without permission to review them.

**Primary actor:** Authenticated user performing an action, and administrator reviewing audit logs

**Supporting component:** HIVE backend, audit logger and audit log database

## 2. Preconditions

- The security-relevant actions and the audit policy, including which failures are recorded, are defined.
- The details an audit record must contain are defined.
- Audit log storage is available.
- Permission to review audit logs is defined (see US-02).

## 3. Main Success Scenario

1. An authenticated user performs a defined security-relevant action.
2. The backend processes the action and it completes successfully.
3. The backend reports the action to the audit logger.
4. The audit logger creates an audit record containing the defined details and associates it with the responsible user where available.
5. The audit logger excludes passwords, authentication secrets and other prohibited secret values from the record.
6. The audit logger stores the record in the audit log database.

## 4. Alternative Scenarios

- Failed action: a user attempts a defined security-relevant action and it fails. HIVE records the failure when it falls within the audit policy, together with the available details needed for investigation.
- Unauthorized audit log access: a user without permission to review audit logs requests audit-log data. HIVE denies the request and does not disclose the audit records.

Secrets are never written to audit records, whether the action succeeded or failed.

## 5. Business Process Diagrams

**Notation:** Each box is a swimlane for one participant. Rounded shapes are start events, double circles are end events, rectangles are tasks, diamonds are decisions, and cylinders are data stores. Dotted arrows show data read from or written to a data store.

### 5.1 Recording audit records (Scenarios 1, 2 and 4)

```mermaid
flowchart TB
    subgraph U["User"]
        U1(["Start: user attempts a security-relevant action"])
        U2["Request the action"]
        U3["Receive result of the action"]
        UE((("End: action result received")))
    end

    subgraph B["Backend API"]
        B1["Process the action"]
        B2{"Action completed successfully?"}
        B3["Report success with action details"]
        B4["Report failure with available details"]
        B5{"Failure within the audit policy?"}
        B6["Return result of the action"]
    end

    subgraph L["Audit Logger"]
        L1["Build audit record excluding passwords and secrets"]
        L2["Store audit record, linked to responsible user where available"]
        L3["Create no audit record"]
    end

    subgraph D["Audit Log Database"]
        D1[("Audit records")]
    end

    U1 --> U2 --> B1 --> B2
    B2 -->|"Yes"| B3 --> L1
    B2 -->|"No"| B4 --> B5
    B5 -->|"Yes"| L1
    B5 -->|"No"| L3 --> B6
    L1 --> L2
    L2 -.->|"write"| D1
    L2 --> B6
    B6 --> U3 --> UE
```

### 5.2 Reviewing audit logs (Scenario 3)

```mermaid
flowchart TB
    subgraph U["User"]
        U1(["Start: user wants to review audit logs"])
        U2["Request audit-log data"]
        U3["Review audit records"]
        U4["Receive denial"]
        UE1((("End: records reviewed")))
        UE2((("End: access denied")))
    end

    subgraph B["Backend API"]
        B1{"Authenticated and permitted to review audit logs?"}
        B2["Deny request, no audit records returned"]
        B3["Retrieve audit records"]
        B4["Return audit records"]
    end

    subgraph D["Audit Log Database"]
        D1[("Audit records")]
    end

    U1 --> U2 --> B1
    B1 -->|"No"| B2 --> U4 --> UE2
    B1 -->|"Yes"| B3
    B3 -.->|"read"| D1
    D1 -.->|"records"| B4
    B4 --> U3 --> UE1
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Record a successful security-relevant action, shown in the recording business process diagram (5.1).
- Scenario 2: Record a failed security-relevant action, shown in the recording business process diagram (5.1).
- Scenario 3: Protect audit logs from unauthorized access, shown in the review business process diagram (5.2).
- Scenario 4: Avoid storing secrets in audit logs, shown in the recording business process diagram (5.1) at the step that builds the audit record.

The model is consistent with US-08.
