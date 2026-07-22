# 128. Longest Consecutive Sequence

## Problem Understanding

Given an unsorted array of integers, return the length of the longest consecutive sequence.

The solution must run in **O(N)** time.

Example:

```text
Input

[100,4,200,1,3,2]

Output

4

Sequence

1,2,3,4
```

---

# Brute Force

## Idea

For every element:

- Find the next consecutive number.
- Continue searching until the sequence ends.
- Track the longest sequence.

---

## Complexity

Searching for every next number requires scanning the array.

```text
Outer loop

↓

O(N)

Finding next element

↓

O(N)

Sequence length

↓

O(N)

Total

↓

O(N³)
```

---

## Why It Is Inefficient

The same sequence is explored multiple times.

Example:

```text
1

↓

2

↓

3

↓

4
```

Starting from

```text
1
```

explores

```text
1,2,3,4
```

Starting from

```text
2
```

explores

```text
2,3,4
```

Starting from

```text
3
```

explores

```text
3,4
```

The same numbers are counted repeatedly.

---

# Observation

Finding the next number repeatedly is expensive.

Instead,

store every number in a hash set.

Now,

checking whether

```text
x + 1
```

exists becomes

```text
O(1)
```

---

# Key Insight

Only begin counting from the **start** of a sequence.

A number is the start if

```text
x - 1
```

does **not** exist.

Example:

```text
1 2 3 4
```

Start from

```text
1
```

Skip

```text
2

3

4
```

because they already belong to an existing sequence.

---

# Key Invariant

> **Every consecutive sequence is counted exactly once—from its smallest element.**

---

# Approach

1. Build a hash set from the array.
2. Initialize `maxLength`.
3. For every number `x`:
   - If `x-1` exists:
     - Skip.
   - Otherwise:
     - This is the beginning of a sequence.
     - Continue checking:
       - `x+1`
       - `x+2`
       - `x+3`
     - Count the sequence length.
     - Update the maximum length.

---

## Optimized Solution

```go
func longestConsecutive(nums []int) int {

    set := map[int]struct{}{}

    for _, x := range nums {
        set[x] = struct{}{}
    }

    longest := 0

    for x := range set {

        if _, exists := set[x-1]; exists {
            continue
        }

        length := 1

        for {
            if _, exists := set[x+length]; !exists {
                break
            }
            length++
        }

        if length > longest {
            longest = length
        }
    }

    return longest
}
```

---

## Complexity

**Time Complexity:** `O(N)`

- Building the hash set takes `O(N)`.
- Every sequence is traversed exactly once.
- Every element belongs to exactly one sequence traversal.

**Space Complexity:** `O(N)`

- Hash set stores every element.

---

## Why This Works

Instead of starting from every number,

we only start from numbers that have **no predecessor**.

This guarantees that every consecutive sequence is explored exactly once.

---

## Mistakes / Learnings

- The hash set is used for **constant-time existence checks**, not sorting.
- Skip any number whose predecessor (`x-1`) exists.
- Only the smallest number of a sequence should begin counting.
- Every number is visited at most once while extending a sequence.

---

## Mental Model

Imagine every sequence has one **head**.

```text
1 → 2 → 3 → 4
```

Only the head starts counting.

Everyone else is skipped.