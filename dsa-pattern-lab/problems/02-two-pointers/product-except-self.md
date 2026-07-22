# 238. Product of Array Except Self

## Problem Understanding

Given an integer array `nums`, return an array `answer` where:

```text
answer[i]
=
Product of every element except nums[i]
```

The solution must run in **O(N)** time **without using division**.

Example:

```text
Input:
[1,2,3,4]

Output:
[24,12,8,6]
```

---

## Brute Force Approach

### Idea

For every element:

- Multiply every other element.
- Skip the current index.

```go
for i := range nums {
    product := 1

    for j := range nums {
        if i == j {
            continue
        }

        product *= nums[j]
    }
}
```

---

## Complexity

**Time Complexity:** `O(N²)`

- Every element iterates over the entire array.

**Space Complexity:** `O(N)`

- Result array.

---

## Why It Is Inefficient

For every index, the same multiplications are repeated.

Example:

```text
nums

1 2 3 4

Index 0

2×3×4

Index 1

1×3×4

Index 2

1×2×4

...
```

The products of the left and right portions are recalculated many times.

---

## Pattern

**Prefix Product + Suffix Product**

Instead of recomputing products for every index,

precompute:

- Product of all elements to the left.
- Product of all elements to the right.

---

## Key Invariant

> At every index:

```text
Answer

=

Product of Left

×

Product of Right
```

The current element is naturally excluded.

---

## Approach

### First Pass (Left → Right)

Store the product of all elements before the current index.

Example:

```text
nums

1 2 3 4

↓

prefix

1 1 2 6
```

---

### Second Pass (Right → Left)

Maintain a running suffix product.

```text
right = 1

↓

4

↓

12

↓

24
```

Multiply the prefix product with the running suffix product.

---

## Optimized Solution

```go
func productExceptSelf(nums []int) []int {

    answer := make([]int, len(nums))

    answer[0] = 1

    for i := 1; i < len(nums); i++ {
        answer[i] = answer[i-1] * nums[i-1]
    }

    right := 1

    for i := len(nums)-1; i >= 0; i-- {
        answer[i] *= right
        right *= nums[i]
    }

    return answer
}
```

---

## Complexity

**Time Complexity:** `O(N)`

- Two linear passes.

**Space Complexity:** `O(1)` extra space

- The output array is not counted as extra space.

---

## Why This Works

Every answer consists of:

```text
Everything on the left

×

Everything on the right
```

The first pass computes all left products.

The second pass computes all right products on the fly.

Multiplying them produces the required result.

---

## Mistakes / Learnings

- Division is not allowed.
- Do **not** recompute left and right products repeatedly.
- The suffix product is **running**, not simply the next element.
- Build prefix products first.
- Maintain the suffix product while traversing from right to left.

---

## Mental Model

> Every index splits the array into two independent parts.

```text
Left Product

×

Right Product

=

Answer
```

Ignore the current element completely.