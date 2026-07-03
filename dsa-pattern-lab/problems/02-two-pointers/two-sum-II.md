# Problem Name

## 1. Problem Understanding
Finding the indices of the elements of given array whose sum equals given target number.
The array is sorted & memory to be used is always O(1) means constants space complexity doesn't increase with input.

---

## 2. Brute Force Approach
Same as Two sum
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

Time Complexity: O(n2)
Space Complexity: O(1)

3. Why it is inefficient
Repeated checking the target sum for every iteration is repetitive and unncessary.
This does solve the memory constraint.

4. Pattern Identification

What pattern does this look like?
Two pointers

Why this pattern?
Because its a sorted array so a search space reducing based on where the of the target are nearer.
Elements greater than target will create the sum equal to target.

5. Optimized Solution
```go
func twoSum(numbers []int, target int) []int {
    start,end := 0, len(numbers)-1
    for start < end {
        sum := numbers[start] + numbers[end]
        if target > sum {
            start = start + 1
        } else if target < sum {
            end = end - 1
        } else {
            return []int{start+1, end+1}
        }
    }
    return []int{start+1, end+1}
}
```

6. Key Insight

When things sorted it gives a info that non-sorted array doesn't give.
Reducing the search space


7. Mistake / Learning
What did I miss initially?
That use two pointers might work here with the sorted array

What should I notice next time?
I thought fast and slow pointer might work here.
