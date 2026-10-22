# US-01: User Authentication

## 1. Purpose

This behavioral model describes how HIVE should authenticate users, reject invalid login attempts, and deny unauthenticated access to protected API endpoints.

**Primary actor:** Registered user

**Supporting component:** HIVE backend and user database

## 2. Preconditions

- The user has an account, or the submitted username does not correspond to an existing account.
- The authentication service is available.
- User account information is stored securely.

## 3. Main Success Scenario

1. The user submits their credentials through the login interface.
2. The frontend sends the authentication request to the backend.
3. The backend validates the request.
4. The backend retrieves the corresponding account.
5. The backend verifies the password against the stored password hash and checks that the account is enabled.
6. The backend establishes an authenticated session or issues the configured authentication credential.
7. The frontend receives the authentication result and allows access to permitted features.

## 4. Alternative Scenarios

- Invalid credentials (incorrect username or password): authentication is rejected and a generic error message is shown.
- Unknown username: handled as invalid credentials; authentication is rejected with the same generic error message.
- Disabled account: authentication is denied, even if the submitted credentials are otherwise valid.
- Unauthenticated request to a protected endpoint: the backend returns HTTP 401 without exposing protected data.

Additional robustness cases (not required by the US-01 acceptance criteria):

- Invalid request: the backend rejects the request without authenticating the user.
- Authentication service failure: the user receives an appropriate error response.

Failed authentication must not grant access to protected endpoints. Error messages should not unnecessarily reveal whether a particular username exists.

## 5. Business Process Diagrams

**Notation:** Each box is a swimlane for one participant. Rounded shapes are start events, double circles are end events, rectangles are tasks, diamonds are decisions, and cylinders are data stores. Dotted arrows show data read from or written to a data store.

### 5.1 Login (Scenarios 1, 2 and 4)

```mermaid
flowchart TB
    subgraph U["User"]
        U1(["Start: user wants to sign in"])
        U2["Enter username and password"]
        U3["Use authorized application interface"]
        U4["Read generic login error"]
        UE1((("End: signed in")))
        UE2((("End: not signed in")))
    end

    subgraph F["Frontend"]
        F1["Send credentials to backend"]
        F2["Allow access to permitted features"]
        F3["Show generic login error"]
    end

    subgraph B["Authentication API"]
        B1["Validate request"]
        B2["Find user by username"]
        B3{"Account exists and is enabled?"}
        B4{"Password matches stored hash?"}
        B5["Establish session or issue credential"]
        B6["Reject authentication"]
    end

    subgraph D["User Database"]
        D1[("User accounts")]
    end

    U1 --> U2 --> F1 --> B1 --> B2
    B2 -.->|"look up"| D1
    D1 -.->|"account or not found"| B3
    B3 -->|"No"| B6
    B3 -->|"Yes"| B4
    B4 -->|"No"| B6
    B4 -->|"Yes"| B5
    B5 --> F2 --> U3 --> UE1
    B6 --> F3 --> U4 --> UE2
```

### 5.2 Protected endpoint access (Scenario 3)

```mermaid
flowchart TB
    subgraph C["Client"]
        C1(["Start: client needs protected data"])
        C2["Request protected endpoint"]
        C3["Receive HTTP 401"]
        C4["Receive permitted data"]
        CE1((("End: access refused")))
        CE2((("End: data received")))
    end

    subgraph A["Backend API"]
        A1{"Valid authentication presented?"}
        A2["Return HTTP 401 with no protected data"]
        A3["Return permitted data"]
    end

    C1 --> C2 --> A1
    A1 -->|"No"| A2 --> C3 --> CE1
    A1 -->|"Yes"| A3 --> C4 --> CE2
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Successful login, shown in the login business process diagram (5.1).
- Scenario 2: Invalid credentials, covering an incorrect username (including an unknown account) or password, shown in the login business process diagram (5.1).
- Scenario 3: Unauthenticated API request, shown in the protected endpoint business process diagram (5.2).
- Scenario 4: Disabled account, shown in the login business process diagram (5.1).

The model is consistent with US-01.
