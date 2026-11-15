# US-01: User Authentication: Sequence Diagrams

## 1. Purpose

These sequence diagrams describe how HIVE authenticates users, rejects invalid login attempts, and denies unauthenticated access to protected API endpoints.

**Primary actor:** Registered user  
**Supporting components:** Frontend, HIVE backend, user database

**Source:** `docs/analysis/behavior/US-01-authentication-bpd.md`

## 2. Sequence Diagram 1: User Login

This diagram covers successful authentication, invalid credentials, unknown usernames, and disabled accounts. Messages are numbered. Every failed login returns the same generic failure, so the response does not reveal whether a username exists.

```mermaid
sequenceDiagram
    autonumber
    actor User
    participant Frontend
    participant Backend as Authentication API
    participant DB as User Database

    User->>Frontend: Enter username and password
    Frontend->>Backend: Submit login credentials
    Backend->>Backend: Validate login request

    alt Request is invalid
        Backend-->>Frontend: Reject invalid request
        Frontend-->>User: Display error message
    else Request is valid
        Backend->>DB: Find account by username
        DB-->>Backend: Return account or not found
        Backend->>Backend: Check account status and password

        alt Account not found
            Backend-->>Frontend: Return generic login failure
            Frontend-->>User: Display generic error
        else Account is disabled
            Backend-->>Frontend: Return generic login failure
            Frontend-->>User: Display generic error
        else Password is incorrect
            Backend-->>Frontend: Return generic login failure
            Frontend-->>User: Display generic error
        else Credentials are valid
            Backend->>Backend: Establish session or issue authentication credential
            Backend-->>Frontend: Return successful authentication result
            Frontend-->>User: Allow access to permitted features
        end
    end
```

## 3. Sequence Diagram 2: Access to a Protected Endpoint

This diagram covers the unauthenticated API request scenario. It shows that the backend must reject a request when the client does not present valid authentication.

```mermaid
sequenceDiagram
    autonumber
    actor Client
    participant Backend as HIVE Backend API

    Client->>Backend: Request protected endpoint
    Backend->>Backend: Check authentication credentials

    alt Authentication is missing or invalid
        Backend-->>Client: HTTP 401 Unauthorized
        Note over Backend,Client: No protected data is returned
    else Authentication is valid
        Backend->>Backend: Continue processing request
        Backend-->>Client: Return permitted data
    end
```

## 4. Acceptance Criteria Mapping

| US-01 Scenario | Description | Sequence Diagram | Branch |
|---|---|---|---|
| Scenario 1 | Successful login with valid credentials | Diagram 1 | Credentials are valid |
| Scenario 2 | Invalid credentials, including an unknown username or incorrect password | Diagram 1 | Account not found, Password is incorrect |
| Scenario 3 | Unauthenticated request to a protected API endpoint | Diagram 2 | Authentication is missing or invalid |
| Scenario 4 | Disabled account cannot authenticate | Diagram 1 | Account is disabled |

The invalid request branch in Diagram 1 is an extra robustness case. US-01 does not require it.

The diagrams are based on the US-01 behavioral model and its acceptance scenarios.
