# 15. 3Sum

## 1. Problem Understanding
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
Ex:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

---

## 2. Brute Force Approach
```go
    for i := 0; i <len(arr); arr {
        for j := i+1; i<len(arr); arr {
            for k := ; i<len(arr); arr {
                if sum(nums[i], nums[j], nums[k]) == 0 {
                    push to a github package in 
                }
            }
        }
    }

```
// brute force code here

Time Complexity: O(n3)
Space Complexity: zo(2)

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
What did I miss initially?
What should I notice next time?
