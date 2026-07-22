# 3. Longest Substring Without Repeating Characters

## Problem Understanding

Given a string `s`, return the length of the longest substring without repeating characters.

A substring consists of **contiguous** characters.

Example:

```text
Input:
"abba"

Possible substrings:

"a"
"ab"
"abb"
"abba"
"b"
"bb"
"bba"
"b"
"ba"
"a"

Longest unique substring = "ab"

Answer = 2
```

---

# Brute Force

## Idea

Generate every possible substring.

For each substring:

- Check whether every character is unique.
- Keep track of the maximum length.

---

## Complexity

### Number of Substrings

A string of length `n` has

```text
n + (n-1) + ... + 1

=

n(n+1)/2

=

O(N²)
```

### Validating One Substring

Maintain a hash set while scanning the substring.

Worst case:

```text
O(N)
```

### Overall Complexity

```text
Generate all substrings

↓

O(N²)

Validate each substring

↓

O(N)

Total

↓

O(N³)
```

---

## Why It Is Inefficient

We're repeatedly validating the same characters.

Example:

```text
"a"

"ab"

"abb"

"abba"
```

When validating `"abb"`:

We already know `"ab"` contains unique characters.

Yet we scan

```text
'a'

'b'
```

again.

Then for `"abba"` we scan them again.

The common prefix is recomputed repeatedly.

---

# Better Brute Force

## Observation

When extending a substring by one character,

the previous characters have already been validated.

Instead of rebuilding the hash set,

reuse it while extending the substring.

For each starting index:

- Start with an empty hash set.
- Keep adding characters.
- Stop when a duplicate is found.
- Record the longest length.
- Reset the hash set for the next starting index.

---

## Complexity

Outer loop

```text
O(N)
```

Inner loop

```text
O(N)
```

Overall

```text
O(N²)
```

Space

```text
O(N)
```

---

# Sliding Window

## Observation

When a duplicate is found,

there is no need to restart from the next index.

Instead,

move only the left boundary until the duplicate is removed.

The window always contains unique characters.

---

## Key Invariant

> **The current window `[i...j]` always contains unique characters.**

Maintain:

- Left pointer (`i`)
- Right pointer (`j`)
- Hash set of characters in the current window

---

## Approach

1. Expand the window by moving `j`.
2. If the character is not in the set:
   - Add it.
   - Update the longest length.
3. If a duplicate is found:
   - Remove characters from the left.
   - Move `i` until the duplicate is removed.
4. Continue expanding.

Each character enters and leaves the window at most once.

---

## Complexity

**Time Complexity:** `O(N)`

- Every character is added to the window once.
- Every character is removed from the window once.

**Space Complexity:** `O(N)`

- Hash set stores at most the characters in the current window.

---

# Mistakes / Learnings

- A string of length `n` has `n(n+1)/2` substrings.
- Brute force repeatedly validates the same prefixes.
- Reusing the hash set removes one level of repeated work.
- The biggest optimization is **not restarting** after a duplicate.
- The sliding window maintains a valid window instead of rebuilding it.

---

# Mental Model

Think of the window as a rubber band.

```text
Duplicate?

↓

Shrink from the left.

No duplicate?

↓

Expand to the right.
```

The window is always valid.