# HIVE Dependency Network (PERT-style)



```mermaid
flowchart TD
    A[Requirements definition] --> B[Requirements analysis and design]
    B --> C[Go backend and REST API]
    B --> D[PostgreSQL schema]
    B --> E[React frontend foundation]
    B --> F[Honeypot interface design]
    C --> G[Authentication and authorization]
    C --> H[Honeypot registration API]
    F --> I[SSH honeypot]
    F --> J[HTTP/web honeypot]
    I --> K[Event collection]
    J --> K
    H --> K
    D --> L[Unified event model and storage]
    K --> L
    L --> M[Normalization]
    M --> N[Detection rules]
    N --> O[Alert generation]
    E --> P[Dashboard foundation]
    O --> P
    P --> Q[Full system integration]
    G --> R[Security tests]
    L --> S[Integration tests]
    N --> T[Detection tests]
    Q --> U[Final verification]
    R --> U
    S --> U
    T --> U
    U --> V[Release and documentation]
```
