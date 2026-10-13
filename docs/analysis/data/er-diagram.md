
# HIVE Honeypot Network: Entity-Relationship Diagram

## 1. Introduction

The current implementation includes the `events` table in PostgreSQL. The other entities in this diagram represent proposed tables for future development. 
## 2. Proposed E-R Diagram

```mermaid
erDiagram
    USERS ||--o{ USER_ROLES : has
    ROLES ||--o{ USER_ROLES : assigned_through

    HONEYPOTS ||--o{ EVENTS : generates

    DETECTION_RULES ||--o{ ALERTS : generates
    EVENTS ||--o{ ALERTS : relates_to

    USERS ||--o{ AUDIT_LOGS : performs

    USERS {
        bigint id PK
        varchar username UK
        varchar password_hash
        boolean enabled
    }

    ROLES {
        bigint id PK
        varchar name UK
        varchar description
    }

    USER_ROLES {
        bigint user_id PK, FK
        bigint role_id PK, FK
    }

    HONEYPOTS {
        bigint id PK
        varchar name
        varchar honeypot_type
        varchar status
    }

    EVENTS {
        bigint id PK
        varchar source
        varchar event_type
        timestamptz occurred_at
        jsonb event_data
        timestamptz received_at
    }

    DETECTION_RULES {
        bigint id PK
        varchar name
        text description
        boolean enabled
    }

    ALERTS {
        bigint id PK
        bigint rule_id FK
        bigint event_id FK
        varchar severity
        varchar status
        timestamptz created_at
    }

    AUDIT_LOGS {
        bigint id PK
        bigint user_id FK
        varchar action
        timestamptz created_at
    }
```

## 3. Entity Descriptions

| Entity | Purpose |
|---|---|
| `users` | Stores application accounts and password hashes. |
| `roles` | Defines application roles, such as Admin and Analyst. |
| `user_roles` | Implements the many-to-many relationship between users and roles. |
| `honeypots` | Stores information about registered honeypot services. |
| `events` | Stores normalized security events received by the backend. |
| `detection_rules` | Defines conditions for identifying potentially suspicious activity. |
| `alerts` | Stores alerts associated with a detection rule and an event. |
| `audit_logs` | Records important actions performed within the application. |

## 4. Relationship Explanations

- A user can have multiple roles, and each role can be assigned to multiple users. The `user_roles` table connects them.
- A honeypot can generate multiple events. Each event is associated with one registered honeypot in the proposed design.
- A detection rule can generate multiple alerts.
- Each alert references one detection rule and one event in this initial design.
- A user can perform multiple actions that are recorded in the audit logs. A system-generated log entry may have no associated user.

## 5. Implementation Status and Design Notes

The existing PostgreSQL `events` table contains the fields shown in the diagram. The remaining tables are proposed additions and will be implemented before the corresponding relationships can be enforced by PostgreSQL.

The current `events.source` column is a text field, not a foreign key. To enforce the proposed relationship between `honeypots` and `events`, a future database migration will introduce a `honeypot_id` foreign key or an equivalent source-registration design.

The proposed `alerts` table initially references one event per alert. If a future alert needs to reference multiple events, an additional junction table such as `alert_events` may be required.
