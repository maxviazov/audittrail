# 🛠️ AuditTrail Roadmap — Development Tasks by File & Feature

This roadmap outlines the recommended order of implementation for core project files based on your current structure.
Tasks are grouped by milestones with associated purpose and target outcomes.

---

## ✅ Milestone 0: Bootstrapping (Completed)

- [x] Initialize GitHub repo and Go module
- [x] Create base folders: `cmd/`, `internal/`, `pkg/`, `api/`, `logs/`
- [x] Implement logger (`pkg/logger/logger.go`)
- [x] Create config handler (`config/config.go`) using Viper
- [x] Setup `Makefile` with version/lint/run/test
- [x] Create `VERSION` file
- [x] Create initial `README.md` with badges and roadmap

---

## 🟢 Milestone 1: Core Infrastructure

- [ ] Define `proto/audit.proto` for gRPC API
- [ ] Create gRPC server scaffolding in `internal/handler/grpc.go`
- [ ] Setup Kafka producer/consumer in `internal/kafka/producer.go`, `consumer.go`
- [ ] Setup PostgreSQL connector in `internal/repository/postgres/postgres.go`
- [ ] Implement Redis caching in `internal/cache/redis.go`
- [ ] Create domain-level `EventService` in `internal/service/event.go`

---

## 🔵 Milestone 2: Runtime & Observability

- [ ] Dockerize app with `Dockerfile` and `docker-compose.yml`
- [ ] Add Prometheus metrics in middleware or gRPC interceptor
- [ ] Add `healthz` and `/metrics` endpoints (optionally via REST or gRPC reflection)
- [ ] Integrate graceful shutdown & signal handling
- [ ] Implement middleware with request ID and contextual logging

---

## 🟣 Milestone 3: Features & Production Readiness

- [ ] Implement event deduplication and basic validation
- [ ] Add querying of audit logs by ID / range
- [ ] Role-based access support (RBAC placeholder)
- [ ] Unit tests + integration tests for logger, service, handlers
- [ ] Add `buf.gen.yaml` and CI formatting checks
- [ ] Release `v0.1.0`

---

## 📘 How to Store & Use This Roadmap

### Option 1: GitHub Project Board

- Use GitHub Projects (https://github.com/maxviazov/audittrail/projects)
- Create columns: Milestone 0, 1, 2, 3
- Create cards from this list and track status

### Option 2: Roadmap.md file in repo

- Save this file as `ROADMAP.md` in the root of your repo
- Link to it from `README.md`
- Update status manually with checkboxes

### Option 3: GitHub Issues with labels

- Break roadmap into issues per file
- Use labels: `milestone-1`, `infra`, `grpc`, `storage`, `redis`

---

## 📫 Suggestion

Create a GitHub Action to verify:

- `VERSION` file matches latest tag
- CI lints + tests
- README badges work

Let me know when you’re back — можем продолжить с proto, gRPC server или Kafka 🧠
