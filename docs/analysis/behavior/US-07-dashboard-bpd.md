# US-07: Dashboard and Event Search

## 1. Purpose

This behavioral model describes how HIVE shows the dashboard to authorized users, searches stored events, returns empty results when nothing matches, and denies event access to users without permission.

**Primary actor:** Security analyst

**Supporting component:** HIVE frontend, backend and event and alert database

## 2. Preconditions

- The user is authenticated (see US-01).
- Permissions for viewing the dashboard and events are defined (see US-02).
- The supported search criteria are defined.
- Stored events exist for a search to return matches.

## 3. Main Success Scenario

1. The user opens the dashboard.
2. The frontend requests the dashboard data from the backend.
3. The backend verifies that the user has permission to view the dashboard.
4. The backend retrieves the available event and alert summaries.
5. The frontend displays the summaries.
6. The user submits a search using the supported search criteria.
7. The backend verifies that the user has permission to view events.
8. The backend queries the stored events and returns those matching the criteria.
9. The frontend displays the matching events.

## 4. Alternative Scenarios

- No results: a valid search matches no stored events. HIVE displays an empty result and does not present unrelated events as matches.
- Unauthorized event access: a user without permission to view events is denied. HIVE does not disclose the protected event data.

A user without permission to view the dashboard is handled by role-based authorization (US-02).

## 5. Business Process Diagrams

**Notation:** Each box is a swimlane for one participant. Rounded shapes are start events, double circles are end events, rectangles are tasks, diamonds are decisions, and cylinders are data stores. Dotted arrows show data read from or written to a data store.

### 5.1 Viewing the dashboard (Scenario 1)

```mermaid
flowchart TB
    subgraph U["User"]
        U1(["Start: user wants an overview"])
        U2["Open dashboard"]
        U3["Review event and alert summaries"]
        U4["Receive denial"]
        UE1((("End: dashboard viewed")))
        UE2((("End: access denied")))
    end

    subgraph F["Frontend"]
        F1["Request dashboard data"]
        F2["Display available event and alert summaries"]
    end

    subgraph B["Backend API"]
        B1{"User permitted to view the dashboard?"}
        B2["Retrieve event and alert summaries"]
        B3["Return summaries"]
        B4["Deny request, handled by US-02"]
    end

    subgraph D["Event and Alert Database"]
        D1[("Events and alerts")]
    end

    U1 --> U2 --> F1 --> B1
    B1 -->|"Yes"| B2
    B2 -.->|"read"| D1
    D1 -.->|"summaries"| B3
    B3 --> F2 --> U3 --> UE1
    B1 -->|"No"| B4 --> U4 --> UE2
```

### 5.2 Searching events (Scenarios 2, 3 and 4)

```mermaid
flowchart TB
    subgraph U["User"]
        U1(["Start: user investigates activity"])
        U2["Enter search criteria"]
        UE1((("End: matches reviewed")))
        UE2((("End: empty result seen")))
        UE3((("End: access denied")))
    end

    subgraph F["Frontend"]
        F1["Submit search request"]
        F2["Show access denied"]
        F3["Display matching events"]
        F4["Display empty result"]
    end

    subgraph B["Backend API"]
        B1{"User permitted to view events?"}
        B2["Deny request, no event data returned"]
        B3["Query events using search criteria"]
        B4{"Matching events found?"}
        B5["Return matching events"]
        B6["Return empty result"]
    end

    subgraph D["Event Database"]
        D1[("Events")]
    end

    U1 --> U2 --> F1 --> B1
    B1 -->|"No"| B2 --> F2 --> UE3
    B1 -->|"Yes"| B3
    B3 -.->|"query"| D1
    D1 -.->|"results"| B4
    B4 -->|"Yes"| B5 --> F3 --> UE1
    B4 -->|"No"| B6 --> F4 --> UE2
```

## 6. Acceptance Criteria Mapping

- Scenario 1: View the dashboard, shown in the dashboard business process diagram (5.1).
- Scenario 2: Search for events, shown in the search business process diagram (5.2).
- Scenario 3: Return no results, shown in the search business process diagram (5.2).
- Scenario 4: Deny unauthorized event access, shown in the search business process diagram (5.2).

The model is consistent with US-07.
