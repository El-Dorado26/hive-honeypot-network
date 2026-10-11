# HIVE Non-Functional Test Cases

**Project:** HIVE Honeypot Network  
**Status:** Designed, pending implementation  
**Purpose:** Define verifiable tests for the non-functional requirements of HIVE.

## 1. Test environment

- Backend: Go
- Database: PostgreSQL
- Frontend: React
- Test data: Synthetic events and dedicated test accounts
- Execution environment: Isolated development or CI environment

Tests must not use production credentials, real sensitive data, or unrestricted external honeypot services.

## 2. NFT-01: Authentication enforcement

**Requirement:** NFR-01

**Objective:** Verify that protected API endpoints reject requests without valid authentication.

**Preconditions:**
- A protected API endpoint exists.
- Authentication middleware is implemented.

**Test cases:**

1. Send a request without credentials.
2. Send a request with invalid credentials.
3. Send a request with valid credentials.

**Expected results:**
- Requests without valid credentials receive HTTP 401.
- Valid credentials allow access to the authentication layer, subject to authorization rules.
- The response does not expose passwords, tokens, or internal stack traces.

**Verification:** Go API integration tests.

## 3. NFT-02: Malformed event handling

**Requirement:** NFR-02

**Objective:** Verify that invalid event submissions are rejected safely.

**Preconditions:**
- The event ingestion API and event schema are implemented.

**Test cases:**

1. Submit malformed JSON.
2. Submit an event missing a required field.
3. Submit an event containing an invalid field type.
4. Submit a valid event.

**Expected results:**
- Invalid submissions receive an appropriate HTTP 4xx response.
- Invalid events are not persisted as valid events.
- Valid events are processed according to the API contract.
- The backend does not terminate because of invalid input.

**Verification:** Go API and database integration tests.

## 4. NFT-03: Event ingestion performance

**Requirement:** NFR-03

**Objective:** Measure the response time of event ingestion under a controlled workload.

**Preconditions:**
- Event ingestion and database persistence are implemented.
- The test environment and event payload are documented.

**Test procedure:**

1. Prepare 100 valid synthetic event submissions.
2. Submit the events using a documented load profile.
3. Record the response time of every request.
4. Calculate the 95th-percentile response time.
5. Record failed requests and the test environment.

**Acceptance criteria:**
- The 95th-percentile response time is no greater than 500 ms.
- The test report identifies the workload, environment, and failure count.

**Measurement conditions:** Exclude unrelated external network latency where feasible. Document whether the requests are sequential or concurrent and whether database writes are included.

**Verification:** Go benchmark or API performance test.

## 5. NFT-04: Recovery after invalid input

**Requirement:** NFR-04

**Objective:** Verify that the backend continues operating after malformed submissions.

**Preconditions:**
- The event ingestion endpoint is operational.

**Test procedure:**

1. Submit 100 malformed event requests.
2. Confirm that the requests are rejected safely.
3. Submit a valid event.
4. Verify that the valid event is processed and persisted.

**Acceptance criteria:**
- The backend remains operational throughout the test.
- The final valid event is accepted and persisted successfully.
- No invalid event is stored as a valid event.

**Verification:** Go API integration test.

## 6. NFT-05: Audit completeness and secret protection

**Requirement:** NFR-05

**Objective:** Verify that security-relevant administrative actions are audited without exposing secrets.

**Preconditions:**
- Administrative operations and audit logging are implemented.
- A test account with the required administrative permissions exists.

**Test procedure:**

1. Perform a supported administrative action.
2. Retrieve its audit record.
3. Verify the actor, action, timestamp, and outcome.
4. Inspect relevant application and audit logs for passwords, authentication tokens, and other prohibited secret values.

**Acceptance criteria:**
- Every successfully completed administrative action covered by the audit policy has a corresponding audit record.
- Required audit fields are present.
- Prohibited secrets are not recorded in logs.

**Verification:** Go integration tests and log inspection.

## 7. NFT-06: Authorization enforcement

**Requirement:** NFR-06

**Objective:** Verify that users without the required permissions cannot perform protected administrative operations.

**Preconditions:**
- Role-based authorization is implemented.
- Test accounts with different roles exist.

**Test procedure:**

1. Authenticate as a user without administrative permission.
2. Attempt a protected administrative operation through the API.
3. Repeat the request using an authorized test account.

**Acceptance criteria:**
- The unauthorized operation is denied.
- The authorized operation succeeds when all other requirements are satisfied.
- Authorization is enforced by the backend, not only by hiding frontend controls.

**Verification:** Go API authorization tests.

## 8. NFT-07: Dashboard usability states

**Requirement:** NFR-07

**Objective:** Verify that the dashboard communicates its data-loading and data-retrieval states clearly.

**Preconditions:**
- The React dashboard and its data-fetching functionality are implemented.

**Test cases:**

1. Delay the dashboard API response.
2. Return an empty event list.
3. Simulate an API error.
4. Return valid event data.

**Acceptance criteria:**
- A loading indicator is visible while data is being fetched.
- An empty-state message is displayed when no events are available.
- A clear error message is displayed when data retrieval fails.
- Valid data is displayed when the request succeeds.

**Verification:** React component tests using the frontend's chosen testing framework.

## 9. NFT-08: Common event schema conformance

**Requirement:** NFR-08

**Objective:** Verify that each supported honeypot integration produces events conforming to the common HIVE event schema.

**Preconditions:**
- The common event schema is documented.
- At least one honeypot integration is implemented.

**Test procedure:**

1. Generate a synthetic event from each implemented honeypot integration.
2. Pass each event through the normalization component.
3. Validate the normalized event against the common schema.

**Acceptance criteria:**
- Every tested event contains all required fields.
- Field types and permitted values follow the documented schema.
- Invalid events are rejected or handled according to the documented contract.

**Verification:** Go unit and integration tests.

## 10. NFT-09: Reproducible development setup

**Requirement:** NFR-10

**Objective:** Verify that another team member can set up HIVE using the repository documentation.

**Preconditions:**
- Setup instructions and dependency configuration are documented.

**Test procedure:**

1. Use a clean environment or isolated environment without an existing HIVE installation.
2. Follow the documented setup instructions.
3. Configure the required environment variables and test database.
4. Build and start the application.
5. Run the documented test suite.

**Acceptance criteria:**
- The documented setup procedure completes without undocumented manual configuration steps.
- The application starts successfully.
- The documented tests can be executed.

**Verification:** Manual deployment checklist initially; automate suitable steps in CI where feasible.

## 11. Test reporting

For each test, record:

- Test ID and linked requirement.
- Execution date and environment.
- Actual result and measured values where applicable.
- Pass, fail, blocked, or not-run status.
- Defect reference or explanation for failure.

Tests must not be marked as passing until they have been executed against the relevant implementation.
