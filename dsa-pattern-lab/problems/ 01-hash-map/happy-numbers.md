# 202. Happy Number

## 1. Problem Understanding
We repeatedly transform a number by replacing it with the sum of the squares of its digits.
We need to check if this process eventually reaches 1.

If it does not reach 1, the process will repeat values and enter a cycle forever.

---

## 2. Brute Force Approach
The simplest idea is to simulate the process step by step and store all previously seen numbers.

If we reach:
- `1` → return true
- a repeated number → return false

```go
func sumOfSquareDigits(num int) int {
    sum := 0
    for num > 0 {
        r := num % 10
        sum += r * r
        num /= 10
    }
    return sum
}

func isHappy(n int) bool {
    seen := map[int]struct{}{}

    for {
        if n == 1 {
            return true
        }

        if _, ok := seen[n]; ok {
            return false
        }

        seen[n] = struct{}{}
        n = sumOfSquareDigits(n)
    }
}
```

Time Complexity: O(k)
Space Complexity: O(k)

## 3. Why it is inefficient

We are storing all previously seen numbers only to detect repetition.

This means:

extra memory is used just to remember history
but the process itself already has enough structure to detect repetition without storing everything

So the inefficiency is:

using memory to solve a problem that can be solved by structure

## 4. Pattern Identification

Cycle detection in a deterministic function (functional graph).

Why this pattern?

Each number has exactly one next state
The process forms a chain of states
It either:
reaches terminal state (1), or
enters a cycle

Key idea:

If a state repeats, a cycle is guaranteed.

Two approaches:

HashSet (track history explicitly)
Fast & Slow pointers (detect cycle structurally)

## 5. Optimized Solution

Instead of remembering all values, we detect cycles using two pointers:

slow moves one transformation at a time
fast moves two transformations at a time

If there is a cycle, they will eventually meet.
If there is no cycle, fast reaches 1.

func sumOfSquareDigits(num int) int {
    sum := 0
    for num > 0 {
        r := num % 10
        sum += r * r
        num /= 10
    }
    return sum
}

```go
func isHappy(n int) bool {
    slow := n
    fast := n

    for {
        slow = sumOfSquareDigits(slow)
        fast = sumOfSquareDigits(sumOfSquareDigits(fast))

        if fast == 1 {
            return true
        }

        if slow == fast {
            return false
        }
    }
}
```

Time Complexity: O(k)
Space Complexity: O(1)

## 6. Key Insight

This problem is not about computing digit sums efficiently.

It is about recognizing a structure:

repeated deterministic transformations form a cycle or reach a terminal state

Optimization comes from removing history tracking and relying on structural properties of cycles.

## 7. Mistake / Learning

Initially assumed that reducing numbers repeatedly implies eventual convergence.

Missed that:

deterministic processes can loop
repetition does not mean progress
cycle detection is the real objective, not value reduction

## Key takeaway:

If a process repeatedly transforms a value, think in terms of states and cycles, not just computation.


Initial intuition:
Repeatedly reducing digits felt similar to digital root problems,
where numbers eventually collapse into a stable single digit.

Learning:
A process that reduces magnitude does not necessarily preserve
the correct state information.

In iterative number problems, avoid assuming convergence just because
values become smaller — cycles and hidden state transitions may still exist.

Wrong intuitions are useful because they reveal which assumptions
I unconsciously relied on.


2 = 4
4 -> 16
16 -> 37
37 -> 58
58 -> 89


1 2 3 4 5 6

slow = 0
fast  = 2