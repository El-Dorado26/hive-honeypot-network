# HIVE Requirements Starter Pack

## Overview

This starter pack contains the initial planning and requirements artifacts for the HIVE Honeypot Network project. It provides a structured foundation for defining the project scope, documenting requirements, planning development tasks, and preparing tests that will later verify implemented functionality.

All artifacts should be reviewed and approved by the team before the requirements phase is finalized.

## Suggested Branch

`requirements_work`

Keep the requirements work on this branch rather than committing directly to `main`.

## Included Structure

```text
docs/
├── planning/
│   ├── WBS.md
│   ├── dependency-network.md
│   └── gantt-chart.md
└── requirements/
    ├── user-stories/
    │   ├── US-001.md
    │   ├── US-002.md
    │   └── ...
    └── requirements-specification.md

tests/
└── requirements/
    ├── US-001_test.go
    ├── US-002_test.go
    └── ...
```

*Note: The filenames above illustrate the intended organization. Use the actual filenames provided in the repository.*

### 1. Planning Documents: `docs/planning/`

This directory contains the initial project planning artifacts:

- **Work Breakdown Structure (WBS):** Breaks the project into manageable tasks and deliverables.
- **Dependency Network (PERT-style):** Shows task dependencies and the order in which activities should be performed.
- **Gantt Chart:** Presents the planned development schedule, task durations, and timeline.

### 2. User Stories: `docs/requirements/user-stories/`

Each user story describes a specific capability from the perspective of a user or stakeholder. Stories should include:

- A unique story ID.
- A clear description of the requirement.
- Acceptance criteria defining when the requirement is satisfied.
- Any relevant dependencies or constraints.

Review all story IDs, descriptions, and acceptance criteria with the team before implementation.

### 3. Requirements Specification: `docs/requirements/requirements-specification.md`

This document consolidates the requirements and connects them to the project's planning artifacts and user stories.

It should clearly communicate the project scope, functional and non-functional requirements, acceptance criteria, and relevant constraints. Ensure that the final document follows the instructor's required report format and assessment rubric.

### 4. Requirements Tests: `tests/requirements/`

This directory contains temporary Go test placeholders corresponding to the user stories.

These tests establish an initial testing structure. They are intended to be replaced or expanded with meaningful behavioral acceptance tests as the corresponding features are implemented.

## Project Scope

The project charter identifies the following initial honeypot targets:

- SSH honeypot.
- HTTP/web honeypot.
- Simulated API honeypot.

The following capabilities are optional stretch goals unless the team and instructor explicitly approve them as mandatory requirements:

- FTP, SMB, MySQL, and Redis honeypots.
- Attacker fingerprinting.
- MITRE ATT&CK mapping.
- Indicator of Compromise (IOC) extraction.
- Session replay.
- Threat intelligence export.

Mandatory acceptance criteria should remain aligned with the approved project scope. Stretch goals should not be treated as required deliverables without explicit approval.

## Important: Understanding the Placeholder Tests

The initial Go tests are intentionally written as **red placeholders**. They are designed to compile and fail with a clear `NOT IMPLEMENTED` message, indicating that the corresponding functionality has not yet been implemented.

These placeholders are not behavioral acceptance tests. They do not demonstrate that the application meets its requirements.

As development progresses:

1. Replace each placeholder with assertions against the relevant API, service, or other implemented component.
2. Verify both expected behavior and relevant error conditions.
3. Ensure tests compile and run successfully before evaluating their results.
4. Confirm that implemented functionality satisfies the corresponding user story's acceptance criteria.

Do not submit tests that fail because of compilation errors, missing dependencies, or broken test configuration.

## Submission Checklist

Before committing the requirements artifacts, complete the following checks:

- [ ] Confirm the repository's actual Go module path and source code layout.
- [ ] Review user story IDs, project scope, and acceptance criteria with the team.
- [ ] Verify the requirements report against the course's required format and rubric.
- [ ] Run `go test ./...` from the Go module root.
- [ ] Verify that the placeholder tests compile and fail for the intended `NOT IMPLEMENTED` reasons.
- [ ] Confirm that planning documents, user stories, requirements tests, and the requirements specification are included.
- [ ] Commit the reviewed artifacts to the `requirements_work` branch, not directly to `main`.
- [ ] Check the branch and commit contents before creating a pull request or submitting the work.

## Next Steps

After the requirements artifacts have been reviewed and committed:

1. Obtain team agreement on the requirements and project scope.
2. Finalize the planning documents and requirements specification.
3. Begin implementing the highest priority user stories.
4. Replace placeholder tests with behavioral acceptance tests as features become available.
5. Track implementation progress against the approved plan and acceptance criteria.

The goal is to maintain traceability between the project plan, user stories, implementation, and tests throughout the development lifecycle.
