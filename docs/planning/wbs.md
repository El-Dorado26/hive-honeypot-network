# HIVE Work Breakdown Structure (WBS)

The Work Breakdown Structure (WBS) decomposes the HIVE Honeypot Network project into manageable deliverables and work packages. To keep the diagrams readable, the WBS is shown as one overview diagram followed by one diagram for each major work area.

## WBS Overview

```mermaid
flowchart LR
    root(["HIVE Honeypot Network"]):::root
    n1_0["1.0 Project Management"]:::area
    root --> n1_0
    n2_0["2.0 Requirements and System Design"]:::area
    root --> n2_0
    n3_0["3.0 Platform Foundation"]:::area
    root --> n3_0
    n4_0["4.0 Honeypot Integration"]:::area
    root --> n4_0
    n5_0["5.0 Event Processing and Security Analysis"]:::area
    root --> n5_0
    n6_0["6.0 Web Dashboard"]:::area
    root --> n6_0
    n7_0["7.0 Testing and Quality Assurance"]:::area
    root --> n7_0
    n8_0["8.0 Deployment and Final Delivery"]:::area
    root --> n8_0
    n9_0["9.0 Optional Advanced Features"]:::opt
    root -.-> n9_0
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

Sections 1.0 to 8.0 are the planned core project work. Section 9.0 (dashed) contains optional advanced features that may be considered after the required deliverables are completed. Optional features must not delay the core deliverables.

## 1.0 Project Management

```mermaid
flowchart LR
    n1_0["1.0 Project Management"]:::area
    n1_1["1.1 Project coordination"]:::sub
    n1_0 --> n1_1
    n1_1_1["1.1.1 Maintain project charter and scope"]:::pkg
    n1_1 --> n1_1_1
    n1_1_2["1.1.2 Assign team responsibilities"]:::pkg
    n1_1 --> n1_1_2
    n1_1_3["1.1.3 Track milestones, risks, and dependencies"]:::pkg
    n1_1 --> n1_1_3
    n1_2["1.2 Project documentation"]:::sub
    n1_0 --> n1_2
    n1_2_1["1.2.1 Maintain requirements documentation"]:::pkg
    n1_2 --> n1_2_1
    n1_2_2["1.2.2 Maintain design and planning documents"]:::pkg
    n1_2 --> n1_2_2
    n1_2_3["1.2.3 Prepare final project documentation"]:::pkg
    n1_2 --> n1_2_3
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 2.0 Requirements and System Design

```mermaid
flowchart LR
    n2_0["2.0 Requirements and System Design"]:::area
    n2_1["2.1 Requirements analysis"]:::sub
    n2_0 --> n2_1
    n2_1_1["2.1.1 Review functional requirements"]:::pkg
    n2_1 --> n2_1_1
    n2_1_2["2.1.2 Define measurable non-functional requirements"]:::pkg
    n2_1 --> n2_1_2
    n2_1_3["2.1.3 Define acceptance criteria and use cases"]:::pkg
    n2_1 --> n2_1_3
    n2_2["2.2 System design"]:::sub
    n2_0 --> n2_2
    n2_2_1["2.2.1 Define system architecture and component boundaries"]:::pkg
    n2_2 --> n2_2_1
    n2_2_2["2.2.2 Design REST API contracts"]:::pkg
    n2_2 --> n2_2_2
    n2_2_3["2.2.3 Design the common honeypot event schema"]:::pkg
    n2_2 --> n2_2_3
    n2_2_4["2.2.4 Design the database schema"]:::pkg
    n2_2 --> n2_2_4
    n2_2_5["2.2.5 Document security and authorization decisions"]:::pkg
    n2_2 --> n2_2_5
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 3.0 Platform Foundation

```mermaid
flowchart LR
    n3_0["3.0 Platform Foundation"]:::area
    n3_1["3.1 Backend and API"]:::sub
    n3_0 --> n3_1
    n3_1_1["3.1.1 Set up the Go backend"]:::pkg
    n3_1 --> n3_1_1
    n3_1_2["3.1.2 Implement REST API routing and request handling"]:::pkg
    n3_1 --> n3_1_2
    n3_1_3["3.1.3 Configure application settings and error handling"]:::pkg
    n3_1 --> n3_1_3
    n3_2["3.2 Database"]:::sub
    n3_0 --> n3_2
    n3_2_1["3.2.1 Configure PostgreSQL"]:::pkg
    n3_2 --> n3_2_1
    n3_2_2["3.2.2 Create tables and database relationships"]:::pkg
    n3_2 --> n3_2_2
    n3_2_3["3.2.3 Implement event persistence"]:::pkg
    n3_2 --> n3_2_3
    n3_3["3.3 Application security"]:::sub
    n3_0 --> n3_3
    n3_3_1["3.3.1 Implement authentication"]:::pkg
    n3_3 --> n3_3_1
    n3_3_2["3.3.2 Implement role-based authorization"]:::pkg
    n3_3 --> n3_3_2
    n3_3_3["3.3.3 Implement administrative audit logging"]:::pkg
    n3_3 --> n3_3_3
    n3_4["3.4 Development infrastructure"]:::sub
    n3_0 --> n3_4
    n3_4_1["3.4.1 Configure version control and branching"]:::pkg
    n3_4 --> n3_4_1
    n3_4_2["3.4.2 Configure automated test execution"]:::pkg
    n3_4 --> n3_4_2
    n3_4_3["3.4.3 Configure continuous integration"]:::pkg
    n3_4 --> n3_4_3
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 4.0 Honeypot Integration

```mermaid
flowchart LR
    n4_0["4.0 Honeypot Integration"]:::area
    n4_1["4.1 SSH honeypot"]:::sub
    n4_0 --> n4_1
    n4_1_1["4.1.1 Configure SSH honeypot"]:::pkg
    n4_1 --> n4_1_1
    n4_1_2["4.1.2 Capture and forward SSH events"]:::pkg
    n4_1 --> n4_1_2
    n4_2["4.2 HTTP/web honeypot"]:::sub
    n4_0 --> n4_2
    n4_2_1["4.2.1 Configure HTTP/web honeypot"]:::pkg
    n4_2 --> n4_2_1
    n4_2_2["4.2.2 Capture and forward HTTP events"]:::pkg
    n4_2 --> n4_2_2
    n4_3["4.3 Event ingestion"]:::sub
    n4_0 --> n4_3
    n4_3_1["4.3.1 Receive honeypot event submissions"]:::pkg
    n4_3 --> n4_3_1
    n4_3_2["4.3.2 Validate incoming event data"]:::pkg
    n4_3 --> n4_3_2
    n4_3_3["4.3.3 Handle invalid submissions safely"]:::pkg
    n4_3 --> n4_3_3
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 5.0 Event Processing and Security Analysis

```mermaid
flowchart LR
    n5_0["5.0 Event Processing and Security Analysis"]:::area
    n5_1["5.1 Event normalization"]:::sub
    n5_0 --> n5_1
    n5_1_1["5.1.1 Map incoming data to the common event schema"]:::pkg
    n5_1 --> n5_1_1
    n5_1_2["5.1.2 Validate normalized event fields"]:::pkg
    n5_1 --> n5_1_2
    n5_2["5.2 Event storage"]:::sub
    n5_0 --> n5_2
    n5_2_1["5.2.1 Store normalized events in PostgreSQL"]:::pkg
    n5_2 --> n5_2_1
    n5_2_2["5.2.2 Retrieve stored events through the API"]:::pkg
    n5_2 --> n5_2_2
    n5_3["5.3 Detection and alerting"]:::sub
    n5_0 --> n5_3
    n5_3_1["5.3.1 Implement initial detection rules"]:::pkg
    n5_3 --> n5_3_1
    n5_3_2["5.3.2 Generate alerts from matching events"]:::pkg
    n5_3 --> n5_3_2
    n5_3_3["5.3.3 Display relevant event and alert information"]:::pkg
    n5_3 --> n5_3_3
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 6.0 Web Dashboard

```mermaid
flowchart LR
    n6_0["6.0 Web Dashboard"]:::area
    n6_1["6.1 Frontend foundation"]:::sub
    n6_0 --> n6_1
    n6_1_1["6.1.1 Set up the React application"]:::pkg
    n6_1 --> n6_1_1
    n6_1_2["6.1.2 Implement API communication"]:::pkg
    n6_1 --> n6_1_2
    n6_2["6.2 Dashboard features"]:::sub
    n6_0 --> n6_2
    n6_2_1["6.2.1 Display event summaries"]:::pkg
    n6_2 --> n6_2_1
    n6_2_2["6.2.2 Display events and alerts"]:::pkg
    n6_2 --> n6_2_2
    n6_2_3["6.2.3 Implement loading, empty, and error states"]:::pkg
    n6_2 --> n6_2_3
    n6_3["6.3 Access control in the interface"]:::sub
    n6_0 --> n6_3
    n6_3_1["6.3.1 Display features according to user permissions"]:::pkg
    n6_3 --> n6_3_1
    n6_3_2["6.3.2 Handle unauthorized API responses"]:::pkg
    n6_3 --> n6_3_2
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 7.0 Testing and Quality Assurance

```mermaid
flowchart LR
    n7_0["7.0 Testing and Quality Assurance"]:::area
    n7_1["7.1 Backend testing"]:::sub
    n7_0 --> n7_1
    n7_1_1["7.1.1 Test event validation"]:::pkg
    n7_1 --> n7_1_1
    n7_1_2["7.1.2 Test authentication and authorization"]:::pkg
    n7_1 --> n7_1_2
    n7_1_3["7.1.3 Test event processing and persistence"]:::pkg
    n7_1 --> n7_1_3
    n7_2["7.2 Non-functional testing"]:::sub
    n7_0 --> n7_2
    n7_2_1["7.2.1 Test response-time targets"]:::pkg
    n7_2 --> n7_2_1
    n7_2_2["7.2.2 Test resilience to invalid input"]:::pkg
    n7_2 --> n7_2_2
    n7_2_3["7.2.3 Test audit logging and security controls"]:::pkg
    n7_2 --> n7_2_3
    n7_3["7.3 Integration and regression testing"]:::sub
    n7_0 --> n7_3
    n7_3_1["7.3.1 Test backend, database, and honeypot integration"]:::pkg
    n7_3 --> n7_3_1
    n7_3_2["7.3.2 Test dashboard integration"]:::pkg
    n7_3 --> n7_3_2
    n7_3_3["7.3.3 Fix defects and rerun relevant tests"]:::pkg
    n7_3 --> n7_3_3
    n7_4["7.4 Test reporting"]:::sub
    n7_0 --> n7_4
    n7_4_1["7.4.1 Record results and defects"]:::pkg
    n7_4 --> n7_4_1
    n7_4_2["7.4.2 Prepare the software testing deliverable"]:::pkg
    n7_4 --> n7_4_2
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 8.0 Deployment and Final Delivery

```mermaid
flowchart LR
    n8_0["8.0 Deployment and Final Delivery"]:::area
    n8_1["8.1 Deployment preparation"]:::sub
    n8_0 --> n8_1
    n8_1_1["8.1.1 Document environment configuration"]:::pkg
    n8_1 --> n8_1_1
    n8_1_2["8.1.2 Prepare deployment instructions"]:::pkg
    n8_1 --> n8_1_2
    n8_2["8.2 Deployment verification"]:::sub
    n8_0 --> n8_2
    n8_2_1["8.2.1 Verify application startup"]:::pkg
    n8_2 --> n8_2_1
    n8_2_2["8.2.2 Verify configuration and connectivity"]:::pkg
    n8_2 --> n8_2_2
    n8_3["8.3 Final delivery"]:::sub
    n8_0 --> n8_3
    n8_3_1["8.3.1 Integrate and review the final codebase"]:::pkg
    n8_3 --> n8_3_1
    n8_3_2["8.3.2 Update the README and project documentation"]:::pkg
    n8_3 --> n8_3_2
    n8_3_3["8.3.3 Prepare the final project submission"]:::pkg
    n8_3 --> n8_3_3
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## 9.0 Optional Advanced Features

```mermaid
flowchart LR
    n9_0["9.0 Optional Advanced Features"]:::opt
    n9_1["9.1 Extended detection"]:::sub
    n9_0 --> n9_1
    n9_1_1["9.1.1 Explore additional detection rules"]:::pkg
    n9_1 --> n9_1_1
    n9_1_2["9.1.2 Explore event correlation"]:::pkg
    n9_1 --> n9_1_2
    n9_2["9.2 Additional integrations"]:::sub
    n9_0 --> n9_2
    n9_2_1["9.2.1 Evaluate further honeypot types"]:::pkg
    n9_2 --> n9_2_1
    n9_2_2["9.2.2 Document feasibility and future improvements"]:::pkg
    n9_2 --> n9_2_2
    classDef root fill:#1f3a5f,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef area fill:#2e6da4,stroke:#1f3a5f,color:#fff,font-weight:bold
    classDef sub fill:#dbe9f6,stroke:#2e6da4,color:#000
    classDef pkg fill:#fff,stroke:#8a8a8a,color:#000
    classDef opt fill:#f3f3f3,stroke:#8a8a8a,stroke-dasharray:5 3,color:#000

```

## WBS Interpretation

- Level 1: The complete HIVE Honeypot Network project.
- Level 2: Major project work areas, numbered 1.0 through 9.0.
- Level 3: Subdivisions of each major work area (for example 3.1).
- Level 4: Individual work packages describing specific activities or deliverables (for example 3.1.1).

The WBS describes the scope and decomposition of project work. Task durations, dependencies, slack, and the critical path are documented separately in the PERT planning document. The project schedule is documented in the Gantt chart.

## Scope and Planning Notes

The WBS identifies planned work and does not, by itself, confirm that any task has been completed. Implementation status and testing evidence should be recorded separately. The PERT and Gantt documents should use the same task identifiers (for example 3.1.1) so that work packages can be traced across the planning documents.
