# 1. Two Sum

## 1. Problem Understanding
Given an array return the element indices that sum up to the target number
ex: [3, 4, 1, 6, 5] target 7

[0,1]
---

## 2. Brute Force Approach
- Go through the array once
- From the current element in the loop go over the array further
- Add the first loop index of the array with inner loop index elements
- if both sum equals target return the pair and exit
- else continue

```go
func twoSum(nums []int, target int) []int {
    for i,n := range nums {
        for j:= i+1; j < len(nums); j++ {
            if n + nums[j] == target {
                return []int{i,j}
            }
        }
    }
    return []int{}
}
```

Time Complexity: O(n2) n square
Space Complexity: O(n)

3. Why it is inefficient?
The solution won't scale and would break and time to process would increase exponentially with increase in the array size

What is being repeated?
Going through the array multiple times. which is repeated and unnecessary.

What is slow or redundant?
Going through array is slow.


4. Pattern Identification
HashMap

What pattern does this look like?
Saving the complement of the target 

Hash Map / Two Pointers / Sliding Window / Binary Search / Recursion

Why this pattern?
Going the through the array multiple time will be saved. lookup would get faster

5. Optimized Solution
```go
func twoSum(nums []int, target int) []int {
    complements := map[int]int{}
    for i,n := range nums {
       comp := target-n
       if idx, ok := complements[comp]; ok {
            return []int{idx,i}
       }
       complements[n] = i
    }
    return []int{}
}
```

6. Key Insight

What changed in thinking from brute force to optimal?
Structured thinking is needed first implement brute force,
then identify optimization we can do by redudant and repeated things.

7. Mistake / Learning
What did I miss initially?
Always thinking of the optimised solution first

What should I notice next time?
Try to solve brute force first
