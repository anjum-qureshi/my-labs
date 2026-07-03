# 27. Remove Element

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
```
// brute force code here

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
func removeElement(nums []int, val int) int {
    k := 0
    for curr := 0; curr < len(nums); curr++ {
        if nums[curr] != val {
            temp := nums[curr]
            nums[curr] = nums[k]
            nums[k] = temp
            k++
        }
    }
    return k
}
```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
What did I miss initially?
What should I notice next time?
