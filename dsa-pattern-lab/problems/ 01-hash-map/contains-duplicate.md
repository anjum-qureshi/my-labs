# 217. Contains Duplicate

## 1. Problem Understanding
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true
Explanation: The element 1 occurs at the indices 0 and 3.

Example 2:
Input: nums = [1,2,3,4]
Output: false
Explanation: All elements are distinct.

---

## 2. Brute Force Approach
- Iterate over the array outter loop (i)
- Iterate over the array again from i+1 (inner loop) j
- If value is found internally in the inner loop return true and stop the processing
- continue the check for further elements

```go
func containsDuplicate(nums []int) bool {
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
    }
    return false
}
```

Time Complexity: O(n2)
Space Complexity: O(n)

3. Why it is inefficient
The apperance of the value is calculated multiple times.

What is slow or redundant?
4. Pattern Identification
HashMap

What pattern does this look like?
Frequency identification

Why this pattern?
HashMap would reduced the computation done multiple times

5. Optimized Solution
```go
func containsDuplicate(nums []int) bool {
    visited := make(map[int]bool)
    for _, num := range nums {
        if visited[num] {
            return true
        }
        visited[num] = true
    }
    return false
}
```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?
Easy identification of mapping
I stored the frequencies which wasn't needed

7. Mistake / Learning
Storing the frequency count not needed just seen or not is enough
Exit if even one is seen stops the array processing further