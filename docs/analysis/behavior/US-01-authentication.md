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

## 5. UML Sequence Diagrams

### 5.1 Login (Scenarios 1, 2 and 4)

```mermaid
sequenceDiagram
    actor User
    participant UI as Frontend
    participant API as Authentication API
    participant DB as User Database

    User->>UI: Enter username and password
    UI->>API: Submit credentials
    API->>API: Validate request
    API->>DB: Find user by username
    DB-->>API: Return account or not found
    API->>API: Verify password hash and check account is enabled

    alt Account exists, is enabled, and password is valid
        API->>API: Establish authenticated session or issue credential
        API-->>UI: Authentication successful
        UI-->>User: Show authorized application interface
    else Incorrect credentials or disabled account
        API-->>UI: Reject authentication
        UI-->>User: Show generic login error
    end
```

### 5.2 Protected endpoint access (Scenario 3)

```mermaid
sequenceDiagram
    participant Client
    participant API as Backend API

    Client->>API: Request protected endpoint
    API->>API: Check for valid authentication

    alt No valid authentication
        API-->>Client: HTTP 401 (no protected data returned)
    else Valid authentication
        API-->>Client: Return permitted data
    end
```

## 6. Acceptance Criteria Mapping

- Scenario 1: Successful login, shown in the login sequence diagram (5.1).
- Scenario 2: Invalid credentials, covering an incorrect username (including an unknown account) or password, shown in the login sequence diagram (5.1).
- Scenario 3: Unauthenticated API request, shown in the protected endpoint sequence diagram (5.2).
- Scenario 4: Disabled account, shown in the login sequence diagram (5.1).

The model is consistent with US-01.
