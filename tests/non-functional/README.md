# HIVE Non-Functional Tests

## 1. Purpose

This directory contains test plans and automated tests for verifying the non-functional requirements of the HIVE Honeypot Network.

The initial tests use a standalone Go HTTP handler prototype to verify event validation, request handling, resilience to invalid input, and preliminary response-time performance.

These prototype tests do not yet verify the complete HIVE application.

## 2. Test Cases and Status

| Test ID | Requirement | Test Objective | Status |
|---|---|---|---|
| NFT-01 | NFR-01 | Verify unauthenticated requests to protected API endpoints are rejected. | Planned: authentication is not implemented. |
| NFT-02 | NFR-02 | Verify malformed JSON and events missing required fields are rejected safely. | Prototype tests implemented; integration with the HIVE backend pending. |
| NFT-03 | NFR-03 | Measure event submission response times under a controlled workload. | Prototype handler performance test implemented; full API and database performance testing pending. |
| NFT-04 | NFR-04 | Verify the handler continues processing valid events after malformed submissions. | Prototype test implemented; full backend resilience testing pending. |
| NFT-05 | NFR-05 | Verify administrative actions are audited and secrets are not logged. | Planned: audit logging is not implemented. |
| NFT-06 | NFR-06 | Verify users without the required permissions cannot perform protected operations. | Planned: role-based authorization is not implemented. |
| NFT-07 | NFR-07 | Verify dashboard loading, empty, and error states. | Planned: frontend implementation is not available. |
| NFT-08 | NFR-08 | Verify supported honeypots produce events conforming to the common event schema. | Planned: final event schema and honeypot integrations are pending. |
| NFT-09 | NFR-10 | Verify the documented development setup works in a clean environment. | Planned: clean-environment deployment testing is pending. |

## 3. Current Prototype Tests

The standalone Go prototype includes the following automated tests:

- `TestMalformedJSONIsRejected`: verifies that malformed JSON receives HTTP 400.
- `TestMissingRequiredFieldIsRejected`: verifies that an event missing required fields receives HTTP 400.
- `TestValidEventIsAccepted`: verifies that a valid event receives HTTP 202.
- `TestNonPostMethodIsRejected`: verifies that unsupported HTTP methods receive HTTP 405.
- `TestHandlerRecoversAfterInvalidRequests`: submits 100 invalid requests and then verifies that a valid request is accepted.
- `TestEventHandlerResponseTime`: measures the response times of 100 in-process handler requests and checks the 95th-percentile time against the preliminary 500 ms target.

These tests exercise the prototype handler, not the complete HIVE backend. The prototype does not currently authenticate users, persist events to PostgreSQL, or generate audit records.

## 4. Test Environment

- **Backend language:** Go
- **Database:** PostgreSQL, planned for integration testing
- **Frontend:** React, planned for usability testing
- **Test data:** Synthetic events and dedicated test accounts
- **Execution environment:** Local development environment or GitHub Actions
- **Dependencies:** Go standard library for the current prototype tests

## 5. Running the Tests

From the repository root, execute:

```bash
cd tests/non-functional
go test -v ./...
```

To run the Go race detector:

```bash
go test -race ./...
```

The tests must compile and execute successfully before their results are reported as passing.

The GitHub Actions workflow at `.github/workflows/non-functional-tests.yml` is configured to run the tests on pushes, pull requests, and manual workflow dispatches.

## 6. Performance Measurement

The preliminary performance test measures the time spent inside the prototype HTTP handler. It does not measure external network latency, database persistence, concurrent production traffic, or end-to-end event ingestion.

The 500 ms threshold is a preliminary target for the prototype. The team must verify the full event ingestion performance target against the implemented API and database in a documented test environment.

Performance results must be recorded with the test environment, workload, response-time measurements, and failure count.

## 7. Test Status and Reporting

Test statuses are defined as follows:

- **Planned:** The test has been designed but has not been implemented.
- **Implemented:** Test code has been written, but successful execution has not yet been confirmed.
- **Passed:** The test has executed successfully against the relevant implementation.
- **Failed:** The test has executed and did not meet its acceptance criteria.
- **Blocked:** The test cannot be completed because a required application component or dependency is unavailable.

For each executed test, record the test ID, execution result, relevant measurements, environment, and any defects found.

Prototype test results must not be presented as verification of the complete HIVE application.

## 8. Next Steps

1. Integrate event validation tests with the actual Go event ingestion API.
2. Add authentication and authorization tests when the security middleware is implemented.
3. Add PostgreSQL integration tests for event persistence and audit logging.
4. Implement frontend tests for dashboard loading, empty, and error states.
5. Run deployment tests in a clean environment.
6. Update each test status based on actual execution results and retain evidence of failures and fixes.
