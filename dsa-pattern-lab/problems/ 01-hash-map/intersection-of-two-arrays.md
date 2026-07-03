# 349. Intersection of Two Arrays

## 1. Problem Understanding
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order. (Intersection of sets)

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

---

## 2. Brute Force Approach

```go
func includes(nums []int, target int) bool {
	for k := 0; k < len(nums); k++ {
		if nums[k] == target {
            return true
		}
	}
    return false
}

func intersection(nums1 []int, nums2 []int) []int {
	intersection := []int{}
	for i := 0; i < len(nums1); i++ {
        if includes(intersection, nums1[i]) {
            continue
        }
        if includes(nums2, nums1[i]) {
            intersection = append(intersection,nums1[i])
        }
	}
    return intersection
}
```

Time Complexity: O (mnk)
Space Complexity: O(min(n, m))

3. Why it is inefficient
What is being repeated?
Repeated is check whether an array is present in an other and also the whether its already visited in the intersection

What is slow or redundant?
Checking the array exists in the another array and intersection element is repeated for every element check.

4. Pattern Identification
Intersection elements can be checked once. Using sets would reduce the repeatation

What pattern does this look like?
Hash Map

Why this pattern?
Storing whether element was already

5. Optimized Solution
```go
func includes(nums []int, target int) bool {
	for k := 0; k < len(nums); k++ {
		if nums[k] == target {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	intersection := map[int]struct{}{}
	nums := map[int]struct{}{}

	for _, ele := range nums1 {
		nums[ele] = struct{}{}
	}

	for _, ele := range nums2 {
		if _,exists := nums[ele]; exists {
			intersection[ele] = struct{}{}
		}
	}

	result := []int{}

	for ele, _ := range intersection {
		result = append(result, ele)
	}
	return result
}
```

6. Key Insight
I'm scared everytime I think of the solution, if the time complexity of O(n2) exponential.

What changed in thinking from brute force to optimal?
Being able to identify what is repeated and reduced.


7. Mistake / Learning
I miss confidence. Nad being okay with a bad solution.
Using map came first than brute force, forcing myself to think in brute force first.
Using empty structs instead of bool to store whether its visited or exists in the above array.
