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

- **Security:** Enforce least privilege, validate input, protect credentials, and isolate honeypots from production systems.
- **Auditability:** Record relevant security and administrative actions without logging secrets.
- **Reliability:** Malformed event input should be rejected safely without crashing the service.
- **Maintainability:** Keep honeypot integrations modular and use a common event structure.
- **Testability:** Provide unit, API, integration, authorization, and detection-rule tests as appropriate.
- **Usability:** Provide clear loading, empty, success, and error states in the web interface.
- **Deployment:** Demonstrate the platform in a controlled environment and document known limitations.

No numerical performance or availability targets have been specified yet. The team should agree on measurable targets if required by the course rubric.

## 7. Acceptance and verification

Each user story has acceptance criteria and a corresponding initial test placeholder in `tests/requirements/`. The placeholders intentionally fail until implemented. They must be replaced with behavioral tests as implementation proceeds.

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
