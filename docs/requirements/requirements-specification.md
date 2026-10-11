# HIVE Requirements Specification

**Project:** HIVE Honeypot Network  
**Version:** Draft 0.1  
**Status:** For team review  
**Prepared by:** Hammad Ali Khan, Project Manager / Cybersecurity Lead  
**Date:** 10 October 2026

## 1. Purpose

HIVE is a modular cybersecurity platform intended to collect, normalize, store, analyze, and visualize activity recorded by controlled honeypot services through a central web application.

## 2. Problem statement

Honeypots can provide useful information about suspicious interactions, but analyzing separate services independently makes it difficult to obtain a centralized view. HIVE aims to centralize event collection, storage, basic detection, alerting, and visualization.

## 3. Project scope

### In scope
- Go backend and REST API.
- PostgreSQL persistence.
- React web interface.
- Authentication and backend-enforced role-based authorization.
- Initial SSH and HTTP/web honeypot integrations, plus a simulated API honeypot where feasible.
- Common event model, event normalization, and event storage.
- Basic configurable detection rules and alerts.
- Audit logging, automated tests, CI, documentation, and controlled deployment.

### Out of scope
- Replacing a production SIEM/SOC.
- Using real sensitive production data or connecting to production systems.
- Unrestricted offensive infrastructure or unauthorized scanning/attacks.
- Guaranteeing detection of every attack.
- Implementing every possible honeypot protocol in the core release.

### Optional stretch scope
Additional honeypots, attacker fingerprinting, MITRE ATT&CK mapping, IOC extraction, session replay, and threat-intelligence export depend on time and feasibility.

## 4. Stakeholders and roles

- **Hammad Ali Khan:** Product Owner, Project Manager, Cybersecurity Lead.
- **Numan Tareq Abu Mohammad:** Backend/API Developer.
- **Mohammed Al-Dirbashi:** Honeypot Developer.
- **Samer Asfour:** Frontend Developer.
- **Arda Türkdönmez:** Database, Testing, and DevOps responsibilities.
 - **Mohammad Rjaie Rshde Shweiki:** Frontend Developer and edge cases testing 

## 5. Functional requirements

| ID | Requirement | Priority | Story |
|---|---|---|---|
| FR-01 | The platform shall authenticate registered users. | Must | US-01 |
| FR-02 | The backend shall enforce role-based permissions. | Must | US-02 |
| FR-03 | Authorized operators shall be able to register honeypots. | Must | US-03 |
| FR-04 | The backend shall accept valid events from registered honeypots. | Must | US-04 |
| FR-05 | The platform shall normalize and persist supported events. | Must | US-05 |
| FR-06 | The platform shall evaluate basic detection rules and generate alerts. | Must | US-06 |
| FR-07 | Authorized users shall be able to view events and alerts in the dashboard. | Must | US-07 |
| FR-08 | The platform shall record security-relevant administrative and access events. | Must | US-08 |

## 6. Non-functional requirements

The following requirements define measurable quality targets for the initial HIVE release. The targets are intended for a controlled development or demonstration environment and must be verified through automated tests or documented measurements.

| ID | Category | Requirement | Verification |
|---|---|---|---|
| NFR-01 | Security | The backend shall reject unauthenticated requests to protected endpoints with HTTP 401. Requests from authenticated users without the required permissions shall receive HTTP 403. | Automated API authorization tests |
| NFR-02 | Input validation | The event ingestion API shall reject malformed JSON and events missing required fields with an appropriate 4xx response. These requests shall not terminate the backend process. | API integration tests |
| NFR-03 | Performance | Under a controlled test load of 100 valid event submissions, at least 95% of requests shall complete within 500 ms, excluding network latency outside the test environment. | Automated performance test |
| NFR-04 | Reliability | The backend shall remain available to process valid events after 100 consecutive malformed event submissions. | Automated resilience test |
| NFR-05 | Auditability | 100% of successfully completed administrative actions covered by the audit policy shall produce an audit record containing the actor, action, timestamp, and outcome. Secrets and passwords shall not be recorded. | Audit integration tests and log inspection |
| NFR-06 | Access control | Users without the required role or resource permission shall be unable to perform protected administrative operations, including through direct API requests. | Automated authorization tests |
| NFR-07 | Usability | The web interface shall display a loading indicator while dashboard data is being fetched and a clear error or empty state when applicable. | Frontend component tests |
| NFR-08 | Maintainability | Each supported honeypot integration shall produce events that conform to the documented common event schema. | Unit and integration tests |
| NFR-09 | Testability | All implemented Must-priority functional requirements shall have automated tests for their principal success path and relevant failure or authorization cases. | Requirements-to-test traceability review |
| NFR-10 | Deployment | A new development environment shall be deployable by following the documented setup instructions without undocumented manual configuration steps. | Clean-environment deployment test |

### Measurement conditions and limitations

- Performance measurements shall record the test environment, request count, and observed response times.
- Performance targets apply to the documented demonstration environment and are not production capacity guarantees.
- Security tests shall use dedicated test accounts and synthetic data.
- Audit tests shall verify that secrets, passwords, and other explicitly excluded sensitive values are not written to logs.
- Any requirement that cannot be met within the project schedule shall be documented with its status, evidence, and justification.

## 7. Acceptance and verification

Functional requirements shall be verified through user-story acceptance criteria and automated behavioral tests.

Non-functional requirements shall be verified through the measurements and tests defined in Section 6, including security, authorization, performance, reliability, auditability, usability, and deployment verification where applicable.

Each test shall reference the requirement it verifies. Tests shall use controlled test accounts, synthetic event data, and an isolated development or test environment.

Test placeholders shall be replaced with executable tests as implementation proceeds. The team shall document test results, unmet targets, and known limitations before the final demonstration.

## 8. Traceability

- User stories: `docs/requirements/user-stories/`
- Requirement test placeholders: `tests/requirements/`
- Work breakdown structure: `docs/planning/wbs.md`
- Dependency network: `docs/planning/pert-dependencies.md`
- Schedule: `docs/planning/gantt.md`

## 9. Assumptions and constraints

- The initial release focuses on controlled SSH and HTTP/web honeypots and centralized event handling.
- Optional advanced analysis features are not release blockers.
- Honeypots must not provide unrestricted access to the host operating system or production resources.
- Team capacity and exact implementation details will be confirmed during requirement analysis and software design.

## 10. Open questions for team review

1. Which authentication method will be implemented for the first release?
2. What exact roles and permissions are required?
3. What event transport will honeypots use to send events?
4. What is the initial threshold/rule for repeated failed authentication?
5. Which cloud provider or controlled hosting environment is available?
6. Does the course rubric require measurable performance, availability, or usability targets?
