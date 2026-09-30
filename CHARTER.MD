# Project Charter: **HIVE Honeypot Network**

## 1. Project Purpose and Objectives

**HIVE (Honeypot Intelligence and Visibility Environment)** is a modular cybersecurity platform that deploys and manages multiple honeypot services, collects security events, analyzes suspicious activity, and visualizes attacker behavior through a centralized web application.

The system will connect controlled honeypots such as **SSH** and **HTTP** services to a central backend, where interactions will be collected, normalized, stored, and analyzed.

The main objective is to provide a practical environment for studying **network attacks, honeypot technology, security monitoring, and attacker behavior** while applying modern software engineering practices.

Advanced capabilities such as **attacker fingerprinting, MITRE ATT&CK mapping, and IOC extraction** may be added depending on time and technical feasibility.

---

## 2. Stakeholders & Roles

- **Product Owner:** Hammad Ali Khan

- **Software Engineering Team:**
  - **Hammad Ali Khan:** Project Manager / Cybersecurity Lead
  - **Numan Tareq Abu Mohammad:** Backend / API Developer
  - **Mohammed Al-Dirbashi:** Honeypot Developer
  - **Samer Asfour:** Frontend Developer
  - **Arda Türkdönmez:** Database / Testing & DevOps Developer

- **GH Repo/Project URL:** *To be filled after charter approval*

---

## 3. Project Scope

### Project Domain

HIVE belongs to the **cybersecurity and network security** domain.

The platform is intended for **security researchers, students, system administrators, and security teams** who need a controlled environment for monitoring and analyzing suspicious network activity.

Potential use cases include:

- Cybersecurity education
- Security research
- Attack-behavior analysis
- Security monitoring
- Honeypot experimentation

### In-Scope

- Centralized management of multiple honeypot services
- Development and integration of selected honeypots, initially focusing on **SSH and HTTP/web services**
- Collection of interaction and security events from honeypots
- Normalization of events into a unified event structure
- Centralized storage of security events
- Detection of suspicious activities such as:
  - Brute-force attempts
  - Service scanning
  - Repeated unauthorized activity
- Alert generation based on configurable detection rules
- Basic correlation of related events and sessions
- Web-based dashboard for monitoring security events and attacker activity
- Authentication and role-based authorization
- Audit logging for security-relevant platform operations
- Automated unit, integration, API, and security-oriented testing where appropriate
- GitHub Issues, Projects, feature branches, pull requests, code reviews, and CI workflows
- Cloud deployment of the platform
- Documentation of the development process and use of generative AI during the software development lifecycle
- Optional advanced security analysis, including:
  - Attacker fingerprinting
  - MITRE ATT&CK mapping
  - IOC extraction
  - Additional honeypot services

### Out-of-Scope

- Replacing a production **SIEM, SOC, or commercial security monitoring platform**
- Providing unrestricted offensive-security infrastructure
- Attacking or scanning systems without authorization
- Providing access from honeypots to real production systems
- Storing real sensitive information in honeypots
- Guaranteeing detection of every possible cyberattack
- Implementing every possible network protocol or honeypot type
- Reproducing the complete functionality of commercial honeypot platforms
- Operating honeypots outside controlled and authorized environments

---

## 4. Technology Stack

### Backend Stack

- **Go**
- **Go HTTP Libraries**
- **RESTful APIs**
- **PostgreSQL**
- **JSON**
- **GitHub Actions**

### Frontend Stack

- **React**
- **JavaScript / TypeScript**
- **HTML**
- **CSS**

### AI Stack

- **Generative AI tools** for development assistance
- AI-assisted requirement decomposition and planning
- AI-assisted code generation and debugging
- AI-assisted test generation
- AI-assisted documentation and code review
- Human review and validation of all AI-generated output

---

## 5. Project Milestones

| Deliverable | Deadline |
|---|---|
| **Project Charter** | 02.10.2026 23:59 |
| **Requirements Definition Report** | 11.10.2026 23:59 |
| **Requirement Analysis Report** | 18.10.2026 23:59 |
| **Software Design Report** | 15.11.2026 23:59 |
| **Software Testing Report** | 22.11.2026 23:59 |
| **Codebase** | 30.11.2026 23:59 |
| **Presentation** | TBA (see sched.) |
