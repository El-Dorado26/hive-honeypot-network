# HIVE Requirements Starter Pack

This starter pack follows the instructor's five submission steps. Review every file with your team before committing.

## Suggested branch
`requirements_work`

## Included structure
- `docs/planning/` — WBS, dependency network (PERT-style), and Gantt planning drafts
- `docs/requirements/user-stories/` — individual user stories with acceptance criteria
- `docs/requirements/requirements-specification.md` — report linking the artifacts
- `tests/requirements/` — one temporary red test per user story

## Important about the tests
The Go tests are intentionally written as **red placeholders** that compile and fail with a clear `NOT IMPLEMENTED` message. This makes the initial state visibly failing, as requested, but they are not yet behavioral acceptance tests. As each feature is implemented, replace its placeholder with assertions against the actual API or service behavior. Do not submit tests that fail because of broken setup or compilation errors.

The project charter describes SSH, HTTP/web, and simulated API honeypots as initial targets. FTP, SMB, MySQL, Redis, attacker fingerprinting, MITRE ATT&CK mapping, IOC extraction, session replay, and threat-intelligence export are optional/stretch goals. Keep these out of mandatory acceptance criteria unless the team and instructor explicitly approve them.

## Before committing
1. Confirm the repository's actual Go module path and source layout.
2. Review story IDs, scope, and acceptance criteria as a team.
3. Confirm the course's exact report format and any rubric requirements.
4. Run `go test ./...` from the Go module root and verify that the requirement tests fail for the intended TODO reasons.
5. Commit the plan, stories, tests, and report on `requirements_work`.
