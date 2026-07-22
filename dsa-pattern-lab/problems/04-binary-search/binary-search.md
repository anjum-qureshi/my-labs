# Binary Search - High-Level Notes

## When Can I Use Binary Search?

Use Binary Search when:

- The search space is **sorted**, or
- A monotonic condition exists (once true, always true).

---

## Core Idea

Instead of searching every element,

compare with the middle element and eliminate **half of the remaining search space**.

Every comparison reduces the search range by half.

---

## Key Invariant

> **If the target exists, it always lies within the current search range `[left...right]`.**

Every pointer update must preserve this invariant.

---

## Pointer Updates

```
target == nums[mid]
```

Return `mid`.

---

```
target < nums[mid]
```

Discard the right half.

```
right = mid - 1
```

---

```
target > nums[mid]
```

Discard the left half.

```
left = mid + 1
```

---

## Loop Condition

```go
for left <= right
```

Use `<=` so the last remaining element is also checked.

---

## Mid Calculation

Preferred:

```go
mid := left + (right-left)/2
```

Avoids integer overflow.

---

## Complexity

**Time:** `O(log N)`

Each iteration removes half of the remaining search space.

```
N

↓

N/2

↓

N/4

↓

N/8

↓

...
```

**Space:** `O(1)`

---

## Mental Model

Imagine searching for a word in a dictionary.

Open the middle page.

- If your word comes before it, discard the right half.
- If your word comes after it, discard the left half.

Repeat until the search space is empty or the word is found.

---

## Common Mistakes

- Using Binary Search on an unsorted search space.
- Updating the wrong pointer.
- Using `left < right` when the last element still needs to be checked.
- Forgetting to return `-1` when the target is not found.

---

## Template

```go
left, right := 0, len(nums)-1

for left <= right {

    mid := left + (right-left)/2

    if nums[mid] == target {
        return mid
    }

    if target < nums[mid] {
        right = mid - 1
    } else {
        left = mid + 1
    }
}

return -1
```

---

## One Sentence to Remember

> **Every comparison eliminates half of the remaining search space while maintaining the invariant that the target, if it exists, is still inside `[left...right]`.**