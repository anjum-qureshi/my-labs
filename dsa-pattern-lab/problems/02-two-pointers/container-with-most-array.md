# 11. Container With Most Water

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			maxArea = max(min(height[i], height[j])*(j-i), maxArea)
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

Time Complexity:
Space Complexity:

3. Why it is inefficient
What is being repeated?
What is slow or redundant?
4. Pattern Identification

What pattern does this look like?
Hash Map

Why this pattern?

5. Optimized Solution

```go
func maxArea(height []int) int {
	max, l, r := 0,0,len(height)-1
	for l < r {
        width := r - l
        area := min(height[l], height[r]) * width
        if area > max {
            max = area
        }

        if height[l] > height[r] {
            r--
        } else {
            l++
        }
	}
	return max
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
What did I miss initially?
What should I notice next time?
