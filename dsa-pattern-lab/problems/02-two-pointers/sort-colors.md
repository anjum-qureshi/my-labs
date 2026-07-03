# 75. Sort Colors

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
func sortColors(nums []int)  {
    freq := map[int]int{0:0,1:0,2:0}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    i := 0
    for color := 0; color < 3; color++ {
        for freq[color] > 0 {
            nums[i] = color
            freq[color]--
            i++
        }
    }
}
```

Time Complexity: O(n)
Space Complexity: O(1)

3. Why it is inefficient
What is being repeated?
What is slow or redundant?
4. Pattern Identification

What pattern does this look like?
Hash Map

Why this pattern?

5. Optimized Solution
// optimized code here

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
This has two passes of the same array
We are rewriting sometimes already ordered elements which is unnecessary
Dutch National Flag
Need to be revisited
