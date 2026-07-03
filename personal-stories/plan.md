# September 2026 Job Switch Plan

Target Role:

* Senior Backend Engineer
* Staff Engineer
* Platform Engineer
* Architect-Leaning Engineer

Target Compensation:

* ₹60-65 LPA

---

# Weekly Time Allocation

## DSA (35%)

* 7 hrs/week

## System Design (40%)

* 8 hrs/week

## AI Engineering (20%)

* 5 hrs/week

## Communication (5%)

* 1 hr/week

---

# Daily Schedule

## Morning (1 hour)

DSA

## Evening Session 1 (1 hour)

System Design

## Evening Session 2 (45-60 min)

AI Engineering

## Saturday

Deep Design + Revision

## Sunday

Mock Interview + Review

---

# SYSTEM DESIGN APPROACH

For first 20 systems:

DO NOT write formal design documents.

Create Thinking Notes.

Template:

System Name

1. Requirements

2. Traffic

3. Data

4. Write Flow

5. Read Flow

6. Bottlenecks

7. Failure Modes

8. Scaling Ideas

9. Things Learned

Maximum 1 page.

Goal:
Train architecture thinking before documentation.

---

# MONTH 1

(June 22 - July 19)

Goal:
Build System Design Foundations

---

## Week 1

### DSA

* Two Sum
* Contains Duplicate
* Valid Anagram
* Group Anagrams
* Top K Frequent
* Product Except Self

### System Design

#### URL Shortener

Thinking Notes

Questions:

* How do we generate short codes?
* Read heavy or write heavy?
* What data is stored?
* What breaks first?
* How do we scale?

#### Pastebin

Thinking Notes

Questions:

* How do we store large text?
* How do we expire content?

#### QR Link Service

Thinking Notes

Questions:

* What changes from URL shortener?

### AI

* OpenAI API setup
* Hello World LLM
* CLI Chatbot
* Conversation Memory

Deliverable:
Working CLI chatbot

---

## Week 2

### DSA

Sliding Window

Problems:

* Best Time To Buy Stock
* Longest Substring Without Repeating Characters
* Maximum Average Subarray
* Permutation In String

### System Design

#### Notification Service

#### Email Service

#### SMS Platform

Focus:

* Queues
* Async processing
* Retries

### AI

Build:
Notification Summarizer

Input:
Notifications

Output:
AI summary

---

## Week 3

### DSA

Two Pointers

Problems:

* Two Sum Sorted
* Container With Most Water
* 3 Sum
* Move Zeroes

### System Design

#### File Upload Service

#### Image Hosting Service

#### Google Drive Lite

Focus:

* Metadata
* Object storage
* Upload flow

### AI

Build:
Chat With Documents

---

## Week 4

### DSA

15 Medium Problems

### System Design

#### API Gateway

#### Rate Limiter

#### Kong Architecture

Focus:

* Routing
* Throttling
* Load balancing
* Failure handling

### AI

Build:
Simple RAG

Features:

* Upload file
* Ask questions

---

# MONTH 2

(July 20 - August 16)

Goal:
Design Complete Systems

---

## Week 5

System:
Chat Application

Daily:

* Requirements
* Data Model
* Read Flow
* Write Flow
* Scaling
* Failure Modes

DSA:
Heap

Problems:

* Top K
* K Closest
* Merge K Lists

AI:
Vector Databases

---

## Week 6

System:
Food Delivery

Focus:

* Search
* Matching
* Availability

DSA:
Heap Continued

AI:
Semantic Search

---

## Week 7

System:
Uber

Focus:

* Driver Matching
* Geo Indexing
* Availability

DSA:
Linked List

Problems:

* Reverse List
* Merge Lists
* Detect Cycle

AI:
Search Platform Improvements

---

## Week 8

System:
Payment Gateway

Focus:

* Idempotency
* Retries
* Consistency

DSA:
Stack + Queue

AI:
Improve RAG

---

# MONTH 3

(August 17 - September 13)

Goal:
Architect Thinking

---

## Week 9

Topic:
Caching

Daily:

* Cache Aside
* Write Through
* Stampede
* Eviction
* Hot Keys

Design:
Netflix Feed

---

## Week 10

Topic:
Databases

Daily:

* Replication
* Sharding
* CAP
* Consistency
* Transactions

Design:
Banking System

---

## Week 11

Topic:
Event Driven Systems

Daily:

* Queues
* Kafka
* Ordering
* DLQ
* Consumer Groups

Design:
Order Processing

---

## Week 12

Topic:
Multi Region

Daily:

* Active Active
* Active Passive
* Failover
* Disaster Recovery

Design:
Global URL Shortener

---

# MONTH 4

(September 14 - September 30)

Goal:
Interview Readiness

No New Learning

Only Practice

---

## Week 13

Daily

### DSA

1 Medium

### Design

1 Complete Design

Rotate:

* URL Shortener
* Chat
* Uber
* Payment Gateway
* Netflix
* Notification Service

### Behavioural

Practice Stories

1. Kong Plugin Redesign

2. Dynamic Routing Platform

3. Loan/Card Onboarding Platform

4. Latency Reduction (5s → 5ms)

5. Workday Automation

---

## Week 14

Mock Interview Week

Monday

* DSA Mock

Tuesday

* System Design Mock

Wednesday

* Behavioural Mock

Thursday

* AI Discussion

Friday

* Resume Review

Saturday

* Full Interview Loop

Sunday

* Retrospective

---

# Rules

1. Never spend more than 30 minutes reading theory.

2. Build before reading.

3. Every design must have:

   * Bottlenecks
   * Failure Modes
   * Scaling Ideas

4. Maintain a "Things Learned" section for every design.

5. Focus on reasoning, not memorization.

6. Architects think in trade-offs.

Always ask:

* Why?
* What breaks?
* What happens at 10x traffic?
* What would I change if traffic grows 100x?
