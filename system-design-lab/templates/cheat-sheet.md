# System Design Capacity Reference Cheat Sheet

> Goal: Quickly classify numbers as Small / Medium / Large / Huge and decide whether additional architecture is needed.

---

# QPS Calibration

| Requests/sec | Interpretation |
| ------------ | -------------- |
| < 100        | Tiny           |
| 100 - 1K     | Small          |
| 1K - 10K     | Medium         |
| 10K - 100K   | Large          |
| 100K - 1M    | Very Large     |
| > 1M         | Internet Scale |

---

# Application Servers

| Traffic        | First Thought                       |
| -------------- | ----------------------------------- |
| < 1K QPS       | Single server                       |
| 1K - 10K QPS   | Few servers behind Load Balancer    |
| 10K - 100K QPS | Horizontal scaling                  |
| > 100K QPS     | Multiple services, caching required |

### Rule of Thumb

```text
1 server
↓
Few servers
↓
Many servers
↓
Distributed architecture
```

---

# SQL Database Writes

| Writes/sec | Interpretation                 |
| ---------- | ------------------------------ |
| < 100      | Tiny                           |
| 100 - 1K   | Small                          |
| 1K - 10K   | Medium                         |
| 10K - 50K  | Large                          |
| > 50K      | Consider partitioning/sharding |

### Interview Shortcut

```text
< 1K writes/sec
→ Single primary DB usually sufficient

> 10K writes/sec
→ Start discussing scaling strategy
```

---

# SQL Database Reads

| Reads/sec  | Interpretation          |
| ---------- | ----------------------- |
| < 1K       | Tiny                    |
| 1K - 10K   | Small                   |
| 10K - 100K | Medium                  |
| > 100K     | Cache becomes important |

---

# Redis Cache

| Requests/sec | Interpretation           |
| ------------ | ------------------------ |
| < 10K        | Easy                     |
| 10K - 100K   | Comfortable              |
| 100K - 1M    | High but common          |
| > 1M         | Redis Cluster discussion |

### Typical Use Cases

```text
Read-heavy systems
Session storage
Rate limiting
Caching
```

---

# Kafka / Message Queue

| Messages/sec | Interpretation       |
| ------------ | -------------------- |
| < 1K         | Tiny                 |
| 1K - 10K     | Small                |
| 10K - 100K   | Medium               |
| 100K - 1M    | Large                |
| > 1M         | Event platform scale |

### Typical Problems

```text
Producer faster than consumer
Queue backlog
Duplicate messages
Lost messages
```

---

# Storage Capacity

| Total Data     | Interpretation                 |
| -------------- | ------------------------------ |
| < 100 GB       | Tiny                           |
| 100 GB - 1 TB  | Small                          |
| 1 TB - 10 TB   | Medium                         |
| 10 TB - 100 TB | Large                          |
| > 100 TB       | Distributed storage discussion |

---

# Daily Data Growth

| Growth/day        | Interpretation       |
| ----------------- | -------------------- |
| < 1 GB/day        | Ignore               |
| 1 - 100 GB/day    | Normal               |
| 100 GB - 1 TB/day | Significant          |
| > 1 TB/day        | Design around growth |

---

# Network Bandwidth

| Traffic           | Interpretation          |
| ----------------- | ----------------------- |
| < 10 MB/s         | Tiny                    |
| 10 - 100 MB/s     | Small                   |
| 100 MB/s - 1 GB/s | Large                   |
| > 1 GB/s          | CDN / Edge architecture |

---

# User Scale

| Active Users | Interpretation |
| ------------ | -------------- |
| < 10K        | Small          |
| 10K - 1M     | Medium         |
| 1M - 100M    | Large          |
| > 100M       | Internet Scale |

---

# Cache Hit Rate

| Hit Rate | Interpretation |
| -------- | -------------- |
| < 50%    | Poor           |
| 50-80%   | Okay           |
| 80-95%   | Good           |
| > 95%    | Excellent      |

---

# Availability Targets

| SLA     | Downtime / Year |
| ------- | --------------- |
| 99%     | ~3.6 days       |
| 99.9%   | ~8.7 hours      |
| 99.99%  | ~52 minutes     |
| 99.999% | ~5 minutes      |

---

# Failure Mode Checklist

For every component ask:

## Cache

```text
What if cache misses?
What if cache is stale?
What if cache dies?
What if cache is full?
```

## Database

```text
What if DB is slow?
What if DB dies?
What if writes increase 100x?
What if storage grows 100x?
```

## Queue

```text
What if producer is faster than consumer?
What if queue fills up?
What if messages are duplicated?
What if messages are lost?
```

## Load Balancer

```text
What if one server dies?
What if traffic becomes uneven?
```

## ID Generator

```text
Can IDs collide?
Can ID generation become a bottleneck?
```

---

# System Design Thinking Framework

After every design answer:

## Assumptions

```text
What assumptions am I making?
```

## Bottlenecks

```text
What breaks first?
```

## Failure Modes

```text
What if this component fails?
```

## 10x Scale

```text
What changes if traffic increases 10x?
```

## 100x Scale

```text
What changes if traffic increases 100x?
```

---

# Golden Rule

Do NOT ask:

"What technology should I add?"

Ask:

"What breaks next?"

Examples:

Database overloaded
→ Cache

Cache overloaded
→ Sharding

Single machine overloaded
→ Horizontal scaling

Storage growth problem
→ Partitioning / Archival

Single point of failure
→ Replication / Failover

```
```


For every system
1. What are users trying to do?
2. What data do I need to store?
3. What are the main read paths?
4. What are the main write paths?
5. What becomes slow first?
6. How would I fix it?
7. What happens if that component disappears?