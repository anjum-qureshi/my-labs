# 11. Container With Most Water

## 1. Problem Understanding

You are given an array `height`.
Each element represents the height of a vertical line.
Choose two lines that together with the x-axis form a container capable of holding the maximum amount of water.
The amount of water is determined by:

```text
Area = min(leftHeight, rightHeight) × width
```

where

```text
width = right - left
```

The container cannot be tilted.

---

## 2. Brute Force Approach

### Idea

Try every possible pair of lines.

For every pair:

1. Calculate the width.
2. Calculate the smaller height.
3. Compute the area.
4. Update the maximum area.

### Algorithm

```text
maxArea = 0

for every i
    for every j > i

        width = j - i
        height = min(height[i], height[j])

        area = width * height

        maxArea = max(maxArea, area)
```

### Implementation

```go
func maxArea(height []int) int {
    maxArea := 0

    for i := 0; i < len(height); i++ {
        for j := i + 1; j < len(height); j++ {

            area := (j - i)

            if height[i] < height[j] {
                area *= height[i]
            } else {
                area *= height[j]
            }

            if area > maxArea {
                maxArea = area
            }
        }
    }

    return maxArea
}
```

### Complexity

- **Time Complexity:** `O(n²)`
- **Space Complexity:** `O(1)`

---

# 3. Pattern Recognition

**Pattern:** Two Pointers (Opposite Direction)

> This problem resembles a shrinking sliding window because the window continuously gets smaller, but it is more accurately classified as a **Two Pointer** problem.

### Why?

- Start with one pointer at each end.
- Compute the answer using the current window.
- Shrink the window by moving exactly one pointer.
- Continue until both pointers meet.

---

# 4. Key Observation

The area depends on two values:

```text
Area = min(height[left], height[right]) × width
```

Notice:

- Width decreases every time we move a pointer.
- Therefore, the only way to potentially increase the area is by increasing the **limiting height**.

The limiting height is always:

```text
min(height[left], height[right])
```

Example:

```text
Left = 3
Right = 7
Width = 10

Area = 3 × 10 = 30
```

Even if the taller line becomes taller:

```text
Left = 3
Right = 100
```

The area is still limited by the shorter line:

```text
Area = 3 × width
```

The taller line does not determine the amount of water.

---

# 5. Why Move the Smaller Pointer?

Suppose:

```text
Left = 3
Right = 8

Area = 3 × width
```

## Move the Taller Pointer

```text
Left = 3
Right = 6
```

- Width decreases.
- Limiting height remains `3`.

The area cannot improve.

---

## Move the Smaller Pointer

Suppose we find:

```text
Left = 9
Right = 8
```

Now:

```text
Area = 8 × smallerWidth
```

Although the width became smaller, the limiting height increased significantly.

This is the **only move that has a chance of producing a larger area**.

---

## Core Insight

```text
Never move the taller line.

Always move the shorter line because only that can increase the limiting height.
```

---

# 6. Optimized Algorithm

1. Place one pointer at the beginning.
2. Place one pointer at the end.
3. Calculate the current area.
4. Update the maximum area.
5. Move the pointer pointing to the shorter line.
6. Repeat until both pointers meet.

---

# 7. Optimized Solution

```go
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {

        width := right - left
        area := 0

        if height[left] < height[right] {
            area = width * height[left]
            left++
        } else {
            area = width * height[right]
            right--
        }

        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}
```

---

# 8. Complexity

| Metric | Complexity |
|---------|------------|
| Time | **O(n)** |
| Space | **O(1)** |

---

# 9. Interview Takeaways

- Area depends on the **minimum** of the two heights.
- Width decreases every time a pointer moves.
- Moving the taller pointer only decreases the width while keeping the limiting height unchanged, so it cannot produce a better result.
- Moving the shorter pointer is the only move that might discover a taller limiting wall, potentially increasing the area despite the reduced width.
- This is a classic **Two Pointer (Opposite Direction)** problem with a greedy decision at each step.
- Whenever a problem asks you to optimize over pairs from both ends of an array and one choice can be safely discarded at each step, consider the **Two Pointer** technique.