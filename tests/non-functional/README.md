# HIVE Non-Functional Tests

## Purpose

These tests verify the quality attributes of HIVE, including security, performance, reliability, auditability, usability, and deployment.

## Test cases

| Test ID | Requirement | Test objective | Status |
|---|---|---|---|
| NFT-01 | NFR-01 | Verify unauthorized API requests are rejected. | Planned |
| NFT-02 | NFR-02 | Verify malformed events are rejected safely. | Planned |
| NFT-03 | NFR-03 | Measure event submission response times under a controlled load. | Planned |
| NFT-04 | NFR-04 | Verify the backend continues processing valid events after malformed submissions. | Planned |
| NFT-05 | NFR-05 | Verify administrative actions are audited and secrets are not logged. | Planned |
| NFT-06 | NFR-06 | Verify users without the required permissions cannot perform protected operations. | Planned |
| NFT-07 | NFR-07 | Verify dashboard loading, empty, and error states. | Planned |
| NFT-08 | NFR-08 | Verify supported honeypots produce events that follow the common event schema. | Planned |
| NFT-09 | NFR-10 | Verify the documented development setup works in a clean environment. | Planned |

## Test environment

- Backend: Go
- Database: PostgreSQL
- Frontend: React
- Data: Synthetic test events and dedicated test accounts
- Environment: Isolated development or test environment

## Test execution and reporting

Each test must reference the requirement it verifies. The team shall record the execution result, relevant measurements, and any defects found.

A test shall be marked as passing only after it has been executed successfully against the relevant implementation.

Tests that are not yet implemented or executed must remain marked as planned or not run.
