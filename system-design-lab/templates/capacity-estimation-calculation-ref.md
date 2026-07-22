# ⚡ Capacity Estimation Log Table (Cheat Code)

## Request Scale

| Daily Requests | Avg RPS | Peak (~3×) |
|--------------:|--------:|-----------:|
| 1M | 12 | 36 |
| 10M | 120 | 360 |
| 100M | 1.2K | 3.6K |
| 1B | 11.6K | 35K |
| 10B | 116K | 350K |

### Mental Shortcut

```text
1M/day    ≈ 10 RPS
10M/day   ≈ 100 RPS
100M/day  ≈ 1K RPS
1B/day    ≈ 10K RPS
10B/day   ≈ 100K RPS
```

---

# Storage Scale

### Assume 100 Bytes / Record

| Records | Storage |
|---------:|--------:|
| 10M | 1 GB |
| 100M | 10 GB |
| 1B | 100 GB |
| 10B | 1 TB |

### Assume 1 KB / Record

| Records | Storage |
|---------:|--------:|
| 1M | 1 GB |
| 10M | 10 GB |
| 100M | 100 GB |
| 1B | 1 TB |
| 10B | 10 TB |

### Assume 10 KB / Record

| Records | Storage |
|---------:|--------:|
| 100K | 1 GB |
| 1M | 10 GB |
| 10M | 100 GB |
| 100M | 1 TB |

---

# Memory Estimation

| Objects | Object Size | RAM |
|---------:|------------:|----:|
| 1M | 100 B | 100 MB |
| 1M | 1 KB | 1 GB |
| 10M | 1 KB | 10 GB |
| 100M | 1 KB | 100 GB |

---

# Network Bandwidth

### Payload = 1 KB

| Peak RPS | Bandwidth |
|----------:|----------:|
| 1K | 1 MB/s |
| 10K | 10 MB/s |
| 100K | 100 MB/s |
| 1M | 1 GB/s |

### Payload = 10 KB

| Peak RPS | Bandwidth |
|----------:|----------:|
| 1K | 10 MB/s |
| 10K | 100 MB/s |
| 100K | 1 GB/s |

### Payload = 100 KB

| Peak RPS | Bandwidth |
|----------:|----------:|
| 100 | 10 MB/s |
| 1K | 100 MB/s |
| 10K | 1 GB/s |

---

# Read / Write Split

| Ratio | Reads | Writes |
|------:|------:|-------:|
| 90:10 | 90% | 10% |
| 80:20 | 80% | 20% |
| 70:30 | 70% | 30% |
| 50:50 | 50% | 50% |

Example

```text
10K RPS

90:10

↓

9K Reads

1K Writes
```

---

# Daily Growth

| Daily | Yearly |
|------:|-------:|
| 1 GB | 365 GB |
| 10 GB | 3.6 TB |
| 100 GB | 36 TB |
| 1 TB | 365 TB |

---

# Availability

| SLA | Downtime |
|-----|----------|
| 99% | 3.6 days |
| 99.9% | 8.7 hours |
| 99.99% | 52 minutes |
| 99.999% | 5 minutes |

---

# Engineering Log Scale

```
10
100
1K
10K
100K
1M
10M
100M
1B
10B
```

---

# Mental Multiplication

### Requests

```text
100M users
×
10 requests/day
=
1B requests/day
≈ 10K RPS
```

### Storage

```text
100M records
×
1 KB
=
100 GB
```

```text
1B records
×
1 KB
=
1 TB
```

### Memory

```text
10M sessions
×
1 KB
=
10 GB RAM
```

### Bandwidth

```text
10K RPS
×
1 KB
=
10 MB/s
```

```text
10K RPS
×
10 KB
=
100 MB/s
```

---

# Replication

```text
Raw Storage × Replication Factor

100 GB × 3 = 300 GB

1 TB × 3 = 3 TB
```

---

# Five Numbers to Memorize

```text
1M/day       ≈ 10 RPS

100M/day     ≈ 1K RPS

1B/day       ≈ 10K RPS

100M × 1KB   ≈ 100 GB

10K RPS × 1KB ≈ 10 MB/s
```

---

# Interview Flow

```text
Users
    ↓
Requests / User
    ↓
Requests / Day
    ↓
Average RPS
    ↓
Peak RPS
    ↓
Read / Write Split
    ↓
Storage
    ↓
Daily Growth
    ↓
Bandwidth
    ↓
Memory
```