# 🧠 Engineering Thinking Framework

> **Goal:** Don't memorize system designs. Learn how to think through any system.

---

# Phase 0 – Think (Don't Write Yet)

Before opening Notion or Markdown, answer these questions in rough notes.

❓ What problem am I solving?

❓ Who are the users?

❓ What is the main user journey?

❓ Why does this system exist?

Once you can explain it in one paragraph, start documenting.

---

# 1. Problem Statement

Write 2–3 sentences.

Answer:

- What are we building?
- Why are we building it?
- Who will use it?

---

# 2. Functional Requirements

Think:

> What can users do?

Checklist

□ Create

□ Read

□ Update

□ Delete

□ Search

□ Share

□ Upload

□ Download

□ Notify

□ Schedule

□ Authenticate

□ Authorize

□ Other?

Now remove anything not applicable.

---

# 3. Non-Functional Requirements

Think:

> How good should the system be?

Performance

□ Low latency

□ High throughput

Availability

□ Highly available

□ Disaster recovery

Scalability

□ Scale reads

□ Scale writes

□ Multi-region

Consistency

□ Strong consistency

□ Eventual consistency

Reliability

□ No data loss

Durability

□ Persistent storage

Security

□ Authentication

□ Authorization

□ Encryption

Cost

□ Low infrastructure cost

Observability

□ Metrics

□ Logging

□ Tracing

---

# 4. Out of Scope

What are we NOT solving today?

Examples

- Analytics

- Admin Dashboard

- Billing

- Notifications

---

# 5. Assumptions

What assumptions am I making?

Examples

- Read-heavy

- Internet available

- Users authenticated

- URLs are unique

- Mobile-first

---

# 6. Capacity Estimation

Traffic

Reads/day

Writes/day

Peak QPS

Storage

Record Size

Growth/day

Growth/year

Bandwidth

Average request

Average response

Memory

Cache Size

Hot Data

Don't aim for perfect numbers.

Order-of-magnitude is enough.

---

# 7. API Design

For every API write

Endpoint

Purpose

Request

Response

Errors

Questions

□ Is it REST?

□ Is it idempotent?

□ Authentication required?

□ Pagination?

---

# 8. Data Model

Question

What information must survive?

For every entity

Fields

Primary Key

Indexes

Relationships

TTL?

Soft Delete?

---

# 9. Read Flow

Ask

How does a read request travel?

Typical path

Client

↓

Load Balancer

↓

Application

↓

Cache

↓

Database

Questions

□ Cache hit?

□ Cache miss?

□ Latency?

□ Failure?

---

# 10. Write Flow

Ask

How does new data enter the system?

Typical path

Client

↓

Application

↓

Validation

↓

Database

↓

Cache

↓

Response

Questions

□ Validation?

□ Duplicate requests?

□ Transactions?

□ Idempotency?

---

# 11. High-Level Design

Draw boxes.

Don't optimize yet.

Client

↓

Gateway

↓

Services

↓

Storage

↓

External Systems

---

# 12. Scaling

Question

Where is today's bottleneck?

Think

CPU

Memory

Database

Network

Storage

Then ask

How do I remove ONE bottleneck?

Possible solutions

□ Load Balancer

□ Read Replica

□ Partitioning

□ Sharding

□ Cache

□ CDN

□ Queue

□ Async Processing

Never add technology without identifying the bottleneck first.

---

# 13. Failure Modes

This is the most important section.

For EVERY component ask

> What if this component disappears?

Component

↓

What breaks?

↓

Can users continue?

↓

How do we recover?

Checklist

□ Client

□ DNS

□ Load Balancer

□ API Server

□ Cache

□ Database

□ Queue

□ External APIs

□ Object Storage

□ CDN

Examples

Redis down

↓

Fallback to DB

↓

Higher latency

↓

Alert

Database down

↓

Failover

↓

Read Replica

↓

503

Queue down

↓

Retry

↓

DLQ

↓

Alert

---

# 14. Trade-offs

Every major decision must answer

WHY?

Questions

Why PostgreSQL?

Why not MongoDB?

Why Redis?

Why not Memcached?

Why Kafka?

Why RabbitMQ?

Why UUID?

Why Base62?

Why Cache Aside?

Why Write Through?

If you cannot answer WHY,

you probably don't understand the decision yet.

---

# 15. Security

Checklist

□ Authentication

□ Authorization

□ Rate Limiting

□ Encryption

□ Secret Management

□ Input Validation

□ SQL Injection

□ XSS

□ CSRF

□ Abuse Prevention

---

# 16. Monitoring

How do I know production is healthy?

Metrics

Logs

Tracing

Dashboards

Alerts

SLIs

SLOs

---

# 17. Future Improvements

If I had another month,

what would I build?

Examples

Analytics

Multi-region

Premium Features

AI

CDN

Notifications

---

# 18. Reflection (MOST IMPORTANT)

What assumptions did I make?

What trade-offs did I discover?

What mistakes did I make?

What did I learn?

What would I redesign next time?

---

# Engineering Rules

Rule 1

Understand the problem before designing.

---

Rule 2

Separate requirements from solutions.

---

Rule 3

Don't optimize imaginary bottlenecks.

---

Rule 4

Every architecture decision deserves a WHY.

---

Rule 5

Every component can fail.

Always ask

"What happens if this disappears?"

---

Rule 6

Build Version 1 first.

Scale later.

---

Rule 7

Simple beats clever.

---

Rule 8

Think first.

Sketch second.

Document last.