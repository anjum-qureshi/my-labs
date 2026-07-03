# Problem Name

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
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
func sortedSquares(nums []int) []int {
    squares := make([]int, len(nums))
    l,r,pos := 0, len(nums)-1,len(nums)-1
    
    for l <= r {
        l2 := nums[l]*nums[l]
        r2 := nums[r]*nums[r]
        if l2 > r2 {
            squares[pos] = l2
            l++
            pos--
        } else {
            squares[pos] = r2
            r--
            pos--
        }
    }
    return squares
}
```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
Read whole preoblem and understand do not jump with assumptions. becoming impatient.
