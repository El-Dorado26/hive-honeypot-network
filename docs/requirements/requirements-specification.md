# HIVE Requirements Specification

**Project:** HIVE Honeypot Network  
**Version:** Draft 0.2  
**Status:** For team review  
**Prepared by:** Hammad Ali Khan, Project Manager / Cybersecurity Lead  
**Date:** 11 October 2026

## 1. Purpose

HIVE is a modular cybersecurity platform designed to collect, normalize, store, analyze, and visualize activity recorded by controlled honeypot services through a centralized web application.

The platform aims to provide a centralized view of suspicious activity and basic security alerts for educational, research, and controlled demonstration environments.

## 2. Problem Statement

Honeypots can provide useful information about suspicious interactions, but analyzing the logs of separate services independently can make it difficult to identify patterns and understand activity across multiple sources.

HIVE addresses this problem by providing centralized event collection, normalization, storage, basic detection, alerting, and visualization through a web application.

The initial release focuses on a manageable set of honeypot integrations and core security monitoring capabilities. It is not intended to replace a production Security Information and Event Management (SIEM) platform or a fully staffed Security Operations Center (SOC).

## 3. Project Scope

### 3.1 In Scope

The initial HIVE release includes:

- A Go backend and REST API.
- PostgreSQL persistence for supported application data and security events.
- A React-based web interface.
- Authentication for registered users.
- Backend-enforced role-based authorization.
- Initial SSH and HTTP/web honeypot integrations.
- A simulated API honeypot where feasible within the project schedule.
- A common event model and event normalization.
- Validation and storage of supported security events.
- Basic configurable detection rules and alert generation.
- A dashboard for viewing events and alerts.
- Audit logging for defined security-relevant administrative and access events.
- Automated tests for implemented requirements.
- Git-based version control and continuous integration.
- Technical documentation, deployment instructions, and controlled deployment verification.

### 3.2 Out of Scope

The following are outside the core release:

- Replacing a production SIEM or SOC.
- Using real sensitive production data or connecting to production systems.
- Unrestricted offensive infrastructure or unauthorized scanning and attacks.
- Guaranteeing the detection of every attack or malicious activity.
- Implementing every possible honeypot protocol.
- Claiming production-grade scalability or availability without appropriate testing.
- Deploying additional integrations that would prevent completion of the core requirements.

### 3.3 Optional Stretch Scope

The following features may be considered if time and resources permit:

- Additional honeypot integrations.
- Attacker fingerprinting.
- MITRE ATT&CK technique mapping.
- Indicator of Compromise (IOC) extraction.
- Session replay.
- Threat-intelligence export.
- Additional detection rules and analysis capabilities.

Optional features shall not delay the delivery of the core requirements.

## 4. Stakeholders and Roles

- **Hammad Ali Khan:** Product Owner, Project Manager, Cybersecurity Lead.
- **Numan Tareq Abu Mohammad:** Backend/API Developer.
- **Mohammed Al-Dirbashi:** Honeypot Developer.
- **Samer Asfour:** Frontend Developer.
- **Arda Türkdönmez:** Database, Testing, and DevOps responsibilities.
- **Mohammad Rjaie Rshde Shweiki:** Frontend Developer and edge-case testing.

Team members may assist with additional tasks when needed. Final task ownership and responsibilities shall be confirmed by the team.

## 5. Functional Requirements

The following functional requirements define the core capabilities expected from HIVE.

| ID | Requirement | Priority | Related User Story |
|---|---|---|---|
| FR-01 | The platform shall authenticate registered users before granting access to protected functionality. | Must | US-01 |
| FR-02 | The backend shall enforce role-based permissions for protected operations and resources. | Must | US-02 |
| FR-03 | Authorized operators shall be able to register and manage supported honeypots. | Must | US-03 |
| FR-04 | The backend shall validate and accept supported events from registered honeypots. | Must | US-04 |
| FR-05 | The platform shall normalize and persist supported security events using the documented common event schema. | Must | US-05 |
| FR-06 | The platform shall evaluate configured basic detection rules and generate alerts when their conditions are met. | Must | US-06 |
| FR-07 | Authorized users shall be able to view and search supported events and view alerts through the dashboard. | Must | US-07 |
| FR-08 | The platform shall record security-relevant administrative and access events covered by the audit policy. | Must | US-08 |

### 5.1 Functional Requirement Clarifications

**FR-01: Authentication**

Only registered users with valid credentials shall be authenticated. Authentication failures shall not grant access to protected resources.

**FR-02: Authorization**

Authentication alone shall not grant permission to perform every operation. The backend shall verify the user's role and applicable permissions before allowing protected operations. Frontend-only restrictions shall not be considered sufficient authorization.

**FR-03: Honeypot Management**

The initial release shall support registration of the honeypot types included in the agreed project scope. Management operations shall be restricted to authorized users.

**FR-04: Event Ingestion**

Incoming events shall be validated before being accepted. Invalid JSON, missing required fields, and unsupported event structures shall be rejected with appropriate client-error responses.

**FR-05: Event Normalization and Storage**

Supported events shall be mapped to a common schema so that events from different integrations can be stored and queried consistently.

**FR-06: Detection and Alerts**

The initial release shall implement basic detection rules agreed upon by the team. Advanced threat detection is outside the core scope.

**FR-07: Dashboard**

The dashboard shall provide the event and alert information available to the authenticated user. The interface shall handle loading, empty, and error states.

**FR-08: Audit Logging**

The platform shall record defined security-relevant actions with sufficient information to support review. Passwords, credentials, cryptographic secrets, and other explicitly excluded sensitive values shall not be written to audit logs.

## 6. Non-Functional Requirements

The following requirements define measurable quality targets for the initial HIVE release. They shall be verified in a controlled development, testing, or demonstration environment as applicable.

| ID | Category | Requirement | Verification Method |
|---|---|---|---|
| NFR-01 | Security | Unauthenticated requests to protected endpoints shall receive HTTP 401. Authenticated users who lack the required permission shall receive HTTP 403. | Automated API authentication and authorization tests |
| NFR-02 | Input Validation | Malformed JSON and events missing required fields shall receive an appropriate HTTP 4xx response. Such requests shall not terminate the backend process. | Automated event-ingestion tests |
| NFR-03 | Performance | Under a controlled test load of 100 valid event submissions, at least 95% of requests shall complete within 500 ms, excluding network latency outside the test environment. | Automated performance test and recorded measurements |
| NFR-04 | Reliability | After 100 consecutive malformed event submissions, the backend shall successfully process a subsequent valid event without requiring a restart. | Automated resilience test |
| NFR-05 | Auditability | 100% of successfully completed administrative actions covered by the audit policy shall produce an audit record containing the actor, action, timestamp, and outcome. Passwords and secrets shall not be recorded. | Audit integration tests and log inspection |
| NFR-06 | Access Control | Users without the required role or resource permission shall be unable to perform protected administrative operations, including by directly calling the backend API. | Automated authorization tests |
| NFR-07 | Usability | The dashboard shall display a loading indicator while data is being fetched and a clear empty or error state when applicable. | Frontend component tests |
| NFR-08 | Maintainability | Every supported honeypot integration shall produce events conforming to the documented common event schema. | Unit and integration tests |
| NFR-09 | Testability | Every implemented Must-priority functional requirement shall have automated tests covering its principal success path and relevant failure or authorization cases. | Requirements-to-test traceability review |
| NFR-10 | Deployment | A new development environment shall be deployable by following the documented setup instructions without undocumented manual configuration steps. | Clean-environment deployment test |

### 6.1 Measurement Conditions and Limitations

- Performance measurements shall record the test environment, number of requests, test method, and observed response times.
- The performance target applies to the documented test environment and is not a guarantee of production capacity.
- Performance measurements shall identify whether requests were sent sequentially or concurrently.
- Security and authorization tests shall use dedicated test accounts and synthetic data.
- Audit tests shall verify both required audit fields and the exclusion of passwords and secrets.
- Reliability tests shall verify that a valid event can be processed after the malformed-request sequence.
- Requirements that cannot be verified because the relevant feature has not yet been implemented shall be marked as pending.
- Any unmet target shall be documented with its status, supporting evidence, known limitations, and justification.

## 7. Acceptance and Verification

### 7.1 Functional Verification

Each functional requirement shall be linked to its corresponding user story and acceptance criteria.

Acceptance criteria shall use the **Given/When/Then** format:

- **Given:** The initial context or precondition.
- **When:** The action or event being performed.
- **Then:** The expected observable result.

For example:

**Scenario: Reject an invalid event**

- **Given** the event-ingestion endpoint is available.
- **When** a client submits malformed JSON.
- **Then** the API returns an appropriate HTTP 4xx response and the backend remains operational.

Each user story shall be maintained as an individual Markdown file in the user-stories directory. Acceptance criteria shall describe observable behavior rather than merely restating the requirement.

### 7.2 Non-Functional Verification

Non-functional requirements shall be verified using the methods specified in Section 6.

The verification process shall include, where applicable:

- Authentication and authorization tests.
- Input-validation tests.
- Performance measurements.
- Resilience testing.
- Audit-log verification.
- Frontend usability tests.
- Common event-schema validation.
- Requirements-to-test traceability checks.
- Clean-environment deployment verification.

### 7.3 Current Prototype Test Status

The Go prototype test suite currently contains six passing tests covering:

1. Rejection of malformed JSON.
2. Rejection of events missing required fields.
3. Acceptance of a valid event.
4. Rejection of unsupported HTTP methods.
5. Recovery of the handler after invalid requests.
6. Preliminary event-handler response-time measurement.

The tests run through the project's GitHub Actions workflow.

These tests verify prototype handler behavior. The response-time measurement does not establish the performance of the complete HIVE application, including the database, frontend, network, and all integrations.

Full-system verification remains pending for features not yet covered by the prototype suite, including authentication, role-based authorization, audit logging, frontend usability, and deployment.

Test status shall be updated as features are implemented and verified. A passing prototype test shall not be treated as proof that an entire functional or non-functional requirement has been satisfied.

### 7.4 Verification Evidence

For each requirement, the team shall record:

- Requirement ID.
- Related user story or implementation component.
- Test case or verification method.
- Test result and date.
- Relevant CI run, test report, or measurement evidence where available.
- Status: planned, implemented, passed, failed, or blocked.
- Known limitations and follow-up actions where applicable.

All implemented Must-priority requirements shall have an identifiable verification path before the final demonstration.

## 8. Traceability

The following documents provide the supporting requirements, testing, and planning information for HIVE.

- **User stories and BDD acceptance criteria:** `docs/requirements/user-stories/`
- **Requirements specification:** `docs/requirements/requirements-specifications.md`
- **Use case diagram:** `docs/requirements/use-case-diagram.puml`
- **Non-functional test suite:** `tests/non-functional/`
- **Non-functional test workflow:** `.github/workflows/non-functional-tests.yml`
- **Work breakdown structure:** `docs/planning/wbs.md`
- **PERT dependency network and calculations:** `docs/planning/pert-dependencies.md`
- **Gantt schedule:** `docs/planning/gantt.md`

The paths above shall be checked against the actual files committed to the project repository. If a document has a different filename or location, this section shall be updated to match it.

### 8.1 Cross-Document Consistency

The requirements, user stories, use case diagram, tests, WBS, PERT network, and Gantt chart shall remain consistent.

In particular:

- Each functional requirement shall have a corresponding user story.
- User-story acceptance criteria shall use BDD formatting.
- Requirements shall be linked to their relevant verification methods.
- The WBS shall decompose the planned project deliverables into identifiable work packages.
- PERT task identifiers, durations, dependencies, slack calculations, and critical path shall agree with the dependency network.
- The Gantt chart shall use consistent task identifiers, durations, dependencies, and schedule assumptions.
- Project milestones and the final deadline shall be consistent across the planning documents.

## 9. Assumptions and Constraints

- The initial release focuses on controlled SSH and HTTP/web honeypots and centralized event handling.
- A simulated API honeypot may be included if feasible within the project schedule.
- Optional advanced analysis features are not release blockers.
- Honeypots shall operate in a controlled environment and shall not provide unrestricted access to the host operating system or production resources.
- Only synthetic or otherwise authorized data shall be used for testing and demonstrations.
- The team has limited time and capacity, so the core requirements take priority over optional features.
- Architecture, event transport, authentication method, hosting environment, and exact implementation details shall be confirmed during requirements analysis and software design.
- The schedule shall be managed against the project deadline of **30 November 2026 at 23:59**, subject to the course's final requirements and approved changes.

## 10. Open Questions for Team Review

1. Which authentication method shall be implemented for the initial release?
2. What exact roles and permissions are required, and which operations are restricted to administrators?
3. What event transport shall the honeypots use to send events to the backend?
4. What common event fields are mandatory across all supported integrations?
5. What initial threshold or rule shall trigger an alert for repeated failed authentication?
6. Which controlled hosting or deployment environment is available?
7. Which administrative and access events must be included in the audit policy?
8. What load-testing method and request concurrency shall be used for NFR-03?
9. Which integration tests are required before each core feature is considered complete?
10. Are there additional measurable performance, availability, security, or usability requirements specified by the course rubric?

Open questions shall be resolved by the team and reflected in the requirements, user stories, implementation, and verification documentation as appropriate.

## 11. Change Management

Changes to requirements shall be reviewed by the team before they are treated as part of the agreed release scope.

When a requirement changes, the team shall consider its effect on:

- Related user stories and acceptance criteria.
- System architecture and interfaces.
- Implementation tasks and work packages.
- Tests and verification evidence.
- PERT dependencies, task estimates, and critical path.
- Gantt schedule, milestones, and final delivery.

The requirements specification shall be updated when approved changes affect the agreed scope or acceptance conditions. The document version and review status shall reflect the actual state of the project.
