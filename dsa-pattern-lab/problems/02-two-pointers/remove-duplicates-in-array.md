# 26. Remove Duplicates from Sorted Array

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
func removeDuplicates(nums []int) int {
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i ; j++ {
            fmt.Println("i: ",i,"j: ",j,"v: ",nums[:i], nums)
            if nums[i] == nums[j] {
                nums = append(nums[:i],nums[i+1:]...)
                i--
            }
        }
    }
    return len(nums)
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
func removeDuplicates(nums []int) int {
    k := 0
    for i := 1; i < len(nums); i++ {
        if nums[k] != nums[i] {
           k = k + 1
           nums[k] = nums[i]
        }
    }
    return k+1
}
```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
What did I miss initially?
What should I notice next time?
