# 📘 AuditTrail — Event Audit Service for Distributed Systems

[![Go Report Card](https://goreportcard.com/badge/github.com/maxviazov/audittrail)](https://goreportcard.com/report/github.com/maxviazov/audittrail)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build](https://github.com/maxviazov/audittrail/actions/workflows/build.yml/badge.svg)](https://github.com/maxviazov/audittrail/actions)

**AuditTrail** is a real-time, centralized event auditing solution for distributed systems and microservices. Built with
transparency, integrity, and traceability in mind, it allows teams to:

- 📌 Record critical events
- 🔐 Maintain secure, immutable logs
- 🔎 Enable traceable and queryable audit history

---

## 🔧 Tech Stack

| Core     | Messaging | Storage    | API      | Observability    |
|----------|-----------|------------|----------|------------------|
| Go 1.22+ | Kafka     | PostgreSQL | gRPC     | Prometheus-ready |
| Redis    | Protobuf  | pgxpool    | TLS/MTLS | Zerolog          |

---

## 🧩 Key Features

- ✅ Centralized event collection via gRPC
- 🔐 Write-only append logs (immutability)
- 📊 Structured logs with metadata (who/when/what)
- ⚡ High-throughput Kafka ingestion
- 🧠 Redis for deduplication and fast access
- 🔍 Query API for downstream usage
- 🏢 Enterprise-grade (SOC2, ISO-27001 oriented)

---

## 🎯 Use Cases

- Audit user activity across microservices
- Log security and permission events
- Capture financial transactions (e.g. trades, payments)
- Integrate with SIEM/observability systems
- Debug production issues with forensic audit trails

---

## 🚀 Planned Features

- 🛂 Role-based access to query API
- 🖥 Web UI for event browsing
- 📤 Export to Elastic, BigQuery, etc.
- 🧪 OpenTelemetry integration
- 🚨 Alerting via Prometheus/Alertmanager

---

## 📌 Roadmap

### 🟢 v0.1.0 — Initial MVP (in progress)

- [x] Repo setup with README and structure
- [ ] gRPC server scaffolding
- [ ] PostgreSQL connector (pgxpool)
- [ ] Kafka producer (basic event ingestion)
- [ ] Redis cache setup
- [ ] Protobuf definitions for `AuditEvent`
- [ ] Basic CLI for local test

### 🔵 v0.2.0 — Query API & Observability

- [ ] Query endpoint (GET event by ID)
- [ ] Prometheus metrics
- [ ] TLS/mTLS support
- [ ] Integration tests

### 🟣 v0.3.0 — Enterprise Features

- [ ] Role-based querying (RBAC)
- [ ] Export to Elastic
- [ ] Web UI prototype (optional)
- [ ] Alerting with Prometheus/Alertmanager

---

## 🧪 Local Dev Setup

```bash
git clone https://github.com/maxviazov/audittrail.git
cd audittrail
make up        # or docker-compose up
```

> Requires: Docker, Go 1.22+, `make`, `protoc`, `buf` (for gRPC/protobufs)

---

## 📁 Project Structure

```bash
.
├── cmd/server           # Entry point
├── internal/
│   ├── api              # gRPC server
│   ├── app              # Application logic
│   ├── config           # Viper config
│   ├── logger           # Zerolog setup
│   ├── postgres         # DB layer (pgxpool)
│   ├── redis            # Redis caching/deduplication
│   └── kafka            # Kafka producer/consumer
├── proto/               # Protobuf definitions
├── scripts/             # Dev helpers
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 📫 Contact

Made with ❤️ by [Макс](https://github.com/maxviazov)  
For inquiries: max.viazov@gmail.com
