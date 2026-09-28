
HIVE Honeypot Network: Project Charter

1. Project Overview

Project Name: HIVE Honeypot Network

HIVE is a modular cybersecurity platform which allows honeypots to be deployed and managed, gathers security events, analyzes suspicious activity, and uses a central web application to visualize attacker behaviour.

The system will link together a number of controlled honeypot services to a central backend, and the interactions with these services will be recorded as security events and then made available for analysis via a web dashboard.

The project provides a practical environment for studying honeypot technology, network attacks, security monitoring, and attacker behavior while applying modern software engineering practices.

---

2. Problem Statement

Individual honeypots can capture valuable information about suspicious activity, but managing multiple honeypots and analyzing their events independently can make it difficult to obtain a centralized view of attacker behavior.

HIVE aims to address this by providing a single platform through which honeypots can report their activity and security events can be collected, stored, analyzed, and visualized.

The platform will allow the project team to experiment with different honeypot types and progressively add more advanced security-analysis capabilities without requiring the entire system to be redesigned.

---

3. Project Objectives

The main objectives of HIVE are to:

* Develop a centralized platform for managing multiple honeypot services.
* Implement honeypots capable of simulating selected network services.
* Collect and normalize interaction data from different honeypots.
* Store security events in a centralized database.
* Analyze events using configurable detection rules.
* Generate alerts from suspicious activity.
* Correlate related activity across honeypot services and sessions.
* Provide a web dashboard for monitoring events and attacker activity.
* Implement authentication and role-based authorization.
* Maintain audit logs for security-relevant platform operations.
* Apply automated testing and test-driven development where appropriate.
* Use GitHub Issues, Projects, branches, pull requests, and CI throughout development.
* Deploy the platform to a cloud environment.
* Incorporate generative AI into the software development lifecycle and document the team’s AI-assisted development process.

Advanced capabilities such as attacker fingerprinting, MITRE ATT&CK mapping, IOC extraction, and additional honeypot services may be developed incrementally depending on time and technical feasibility.

---

4. Scope

4.1 Core Scope

The initial HIVE platform will consist of:

* A Go backend.
* A RESTful API.
* A React web frontend.
* A PostgreSQL database.
* Honeypot service integrations.
* A centralized event collection mechanism.
* Event normalization and storage.
* Basic detection rules.
* Alert generation.
* Authentication and authorization.
* Audit logging.
* Automated testing.
* CI workflows.
* Cloud deployment.

4.2 Honeypot Scope

HIVE will use a modular architecture so that individual honeypot services can be developed and deployed independently.

The initial implementation is expected to focus on selected high-value services such as:

* SSH.
* HTTP/web.
* Simulated API services.

Additional services such as FTP, SMB, MySQL, and Redis may be added as extensions if project resources and development time permit.

The SSH honeypot may provide a controlled simulated shell environment capable of recording commands and other interactions without providing access to a real operating system.

The HTTP honeypot may simulate vulnerable-looking web resources and record requests, authentication attempts, paths, headers, and other relevant interaction data.

4.3 Security Analysis Scope

The initial analysis functionality will focus on:

* Repeated authentication attempts.
* Brute-force behavior.
* Service scanning.
* Suspicious requests.
* Suspicious commands.
* Repeated unauthorized activity.
* Activity originating from the same source.

Future extensions may include:

* Attacker session correlation.
* Attacker behavioral profiling.
* Tool fingerprinting.
* MITRE ATT&CK technique mapping.
* IOC extraction.
* Session replay.
* Threat-intelligence export.

These advanced features are considered extensible capabilities rather than mandatory requirements for the initial system.

4.4 Out of Scope

The project will not initially attempt to:

* Replace a production SIEM or SOC platform.
* Provide unrestricted offensive-security infrastructure.
* Attack or scan systems without authorization.
* Store real sensitive information in honeypots.
* Provide access to real production systems.
* Guarantee detection of all attacks.
* Implement every possible network protocol.
* Reproduce the complete functionality of commercial honeypot platforms.

All honeypot services and security testing will operate in controlled and authorized environments.

---

5. High-Level Architecture

The initial architecture will follow a modular structure:

                         HIVE PLATFORM
                              |
                    +---------+---------+
                    |                   |
                    v                   v
              Go Backend           React Frontend
                    |
          +---------+---------+
          |         |         |
          v         v         v
        SSH       HTTP       API
     Honeypot   Honeypot   Honeypot
          |         |         |
          +---------+---------+
                    |
                    v
             Event Collector
                    |
                    v
            Event Normalization
                    |
                    v
              PostgreSQL
                    |
                    v
            Detection Engine
                    |
              +-----+-----+
              |           |
              v           v
            Alerts      Analysis
              |           |
              +-----+-----+
                    |
                    v
              Web Dashboard

The architecture is intended to be extensible. Additional honeypot services and analysis modules should be able to connect to the central platform without requiring major changes to existing components.

---

6. Event Model

HIVE will use a unified event structure so that activity from different honeypots can be analyzed consistently.

A security event may contain information such as:

{
  "honeypotId": "ssh-01",
  "serviceType": "SSH",
  "eventType": "LOGIN_ATTEMPT",
  "sourceIp": "example",
  "action": "authentication",
  "timestamp": "example"
}

Depending on the honeypot and event type, additional information may include:

* Session identifier.
* Username.
* Requested resource.
* Command.
* Request method.
* Payload metadata.
* Authentication result.
* File interaction information.
* Detection information.

Sensitive or unnecessary information will not be collected solely for the purpose of increasing data volume.

---

7. Technology Stack

Frontend

* React
* JavaScript/TypeScript as appropriate
* HTML/CSS

Backend

* Go
* Go HTTP libraries
* RESTful APIs

Database

* PostgreSQL

Additional technologies may be introduced when they provide a clear benefit to the project and remain consistent with the course requirements.

---

8. Security Principles

Security will be considered throughout the development lifecycle.

The project will follow principles including:

* Least privilege.
* Authentication before protected operations.
* Role-based authorization.
* Input validation.
* Secure handling of credentials and secrets.
* No hard-coded credentials or sensitive configuration.
* Audit logging.
* Honeypot isolation.
* Controlled security testing.
* Safe handling of captured attacker data.

Honeypots will be designed to prevent interactions from unintentionally providing access to real systems or sensitive information.

---

9. Development Methodology

HIVE will use an iterative and incremental development process.

Development work will generally follow:

GitHub Issue
     |
     v
Feature Branch
     |
     v
Implementation
     |
     v
Tests
     |
     v
Pull Request
     |
     v
CI Checks
     |
     v
Code Review
     |
     v
Merge

The main branch will contain reviewed and integrated changes.

Features should be developed in small, manageable increments so that they can be tested and reviewed independently.

---

10. Testing and Quality Assurance

Testing will be performed throughout development rather than only at the end of the project.

Where appropriate, the team will use test-driven development:

Write Test
    |
    v
Test Fails
    |
    v
Implement Feature
    |
    v
Test Passes
    |
    v
Refactor

Testing may include:

* Unit tests.
* API tests.
* Functional tests.
* Integration tests.
* Authentication and authorization tests.
* Detection-rule tests.
* Database tests.
* Frontend tests.
* Security-oriented tests.

Automated tests will be integrated into GitHub Actions so that important changes can be verified before merging.

---

11. Generative AI Development

Generative AI will be incorporated into the software development lifecycle in accordance with the course requirements.

Potential uses include:

* Requirement decomposition.
* Architecture exploration.
* Implementation planning.
* Code generation.
* Test generation.
* Debugging assistance.
* Refactoring.
* Documentation.
* Code review assistance.
* Security review assistance.

The team will review, test, and modify AI-generated output before incorporating it into the project.

The team will document the AI tools and workflows used during development. The final project presentation will demonstrate how AI was incorporated into the development process.

A possible development pipeline is:

Requirement
     |
     v
AI-Assisted Planning
     |
     v
Human Review
     |
     v
AI-Assisted Implementation
     |
     v
Testing
     |
     v
AI-Assisted Review
     |
     v
Human Review
     |
     v
Pull Request

AI assistance will support development but will not replace human responsibility for the correctness, security, and maintainability of the project.

---

12. Repository and Collaboration Workflow

GitHub will be used as the primary collaboration platform.

The project will use:

* GitHub Issues for requirements and tasks.
* GitHub Projects for project management.
* Feature branches for development.
* Pull requests for proposed changes.
* Code reviews for collaboration and quality control.
* GitHub Actions for automated testing and CI.

Example branch names include:

feature/project-foundation
feature/database-integration
feature/authentication
feature/honeypot-management
feature/event-ingestion
feature/detection-engine
feature/dashboard
feature/ssh-honeypot

Commit messages should clearly describe the purpose of each change.

Examples include:

feat: add event ingestion endpoint
fix: validate honeypot registration
test: add detection rule tests
docs: update architecture documentation
refactor: simplify event processing

---

13. Project Management

The project will be managed using GitHub Projects.

Tasks may progress through:

BACKLOG → TODO → IN PROGRESS → REVIEW → TESTING → DONE

Major features will be represented by GitHub Issues containing requirements and acceptance criteria.

Development will prioritize the core platform before optional advanced functionality.

---

14. Planned Development Phases

The project is expected to develop incrementally.

Phase 1: Platform Foundation

* Go backend.
* Project structure.
* REST API.
* React frontend foundation.
* PostgreSQL integration.
* Authentication and authorization.
* CI and testing setup.

Phase 2: Honeypot Integration

* Honeypot registration.
* SSH honeypot.
* HTTP/web honeypot.
* Event collection.

Phase 3: Event Analysis

* Event normalization.
* Event storage.
* Detection rules.
* Alert generation.
* Basic session correlation.

Phase 4: Dashboard

* Real-time or near-real-time event display.
* Alert monitoring.
* Attacker/source information.
* Event timelines.
* Service statistics.

Phase 5: Advanced Security Analysis

Subject to time and technical feasibility:

* Additional honeypots.
* Attacker fingerprinting.
* Cross-service session correlation.
* MITRE ATT&CK mapping.
* IOC extraction.
* Session replay.
* Threat-intelligence export.

---

15. Deliverables

The expected deliverables include:

1. Go-based HIVE backend.
2. React-based web frontend.
3. PostgreSQL database.
4. Modular honeypot integrations.
5. Central event collection system.
6. Event normalization and storage.
7. Detection and alert functionality.
8. Authentication and authorization.
9. Automated tests.
10. GitHub CI workflow.
11. Cloud deployment.
12. Technical documentation.
13. GitHub Issues and Project management history.
14. Pull requests and code-review history.
15. Documentation of the generative-AI development workflow.

Advanced deliverables such as additional service honeypots, MITRE ATT&CK mapping, attacker profiling, IOC extraction, and session replay may be included depending on project progress.

---

16. Success Criteria

The core HIVE system should be capable of demonstrating the following workflow:

Interaction with Honeypot
          |
          v
     Event Generated
          |
          v
      HIVE Receives
          |
          v
    Event Normalized
          |
          v
      Event Stored
          |
          v
   Detection Evaluation
          |
          v
     Alert Generated
          |
          v
     Dashboard Display

The project should also demonstrate:

* Working authentication and authorization.
* Automated testing.
* CI checks.
* GitHub-based collaborative development.
* Controlled honeypot operation.
* A working cloud deployment.
* Documented use of generative AI.

---

17. Constraints

The project is subject to the time, technical, and resource limitations of the course.

The team will prioritize a functional, secure, and well-tested core platform before implementing advanced features.

The architecture will remain modular so that additional honeypots and analysis capabilities can be added without requiring a complete redesign.

---

18. Charter Maintenance

This charter represents the initial direction and scope of the HIVE Honeypot Network.

As development progresses, the team may refine the scope based on technical feasibility, course requirements, testing results, and available development time.

Major changes to the project’s objectives or scope should be discussed by the team and documented through the project’s GitHub workflow.

Project: HIVE Honeypot Network
Repository: hive-honeypot-network
Charter: docs/CHARTER.MD
