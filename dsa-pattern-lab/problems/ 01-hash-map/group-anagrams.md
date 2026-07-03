# 49. Group Anagrams

## 1. Problem Understanding
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:

There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.


---

## 2. Brute Force Approach
- Store the first element of input as the first group
- Iterate over the array of string (outter loop)
- Now Iterate over the groups (inner loop)
- Take first elment from each group sort it
- Sort the current element in outter loop(x)
- if both match append to current group and break
- Save matched or not for outter loop
- If not matched, Create new group with element x append to groups

Notes: Don't forget 
Save the matched state to make sure you not create groups even when matched
Remember to break the loop to avoid further group checks after matched
Don't iterate over a growing group length




```go

import (
	"sort"
    "strings"
)

func sortSlice(s string) string {
	r := []rune(s) // convert to rune slice (handles Unicode safely)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}


func groupAnagrams(strs []string) [][]string {
    groups := [][]string{}
    for i := 0; i < len(strs); i++ {
        currElement := sortSlice(strs[i])
        matched := false
        ngroups := len(groups)
        for j := 0; j < ngroups; j++ {
            matched = sortSlice(groups[j][0]) == currElement
            if matched {
               groups[j] = append(groups[j], strs[i])
               break
            }
        }
        if !matched {
            groups = append(groups, []string{strs[i]})
        }
    }
    return groups
}
```


Time Complexity: O (n2) + sorting logic O(k log k)
Space Complexity: O(n)

3. Why it is inefficient
Repeated sorting of the already sorting group checks
Won't scale for larger string as sorting is a heavy operation to do O(n2) times

4. Pattern Identification
We need to the pattern of sorted words against the group, this mapping

What pattern does this look like?
HashMap

Why this pattern?
The identity of the group is need to be the key for the group instead of array index

5. Optimized Solution
```go
import (
	"sort"
	"strings"
)

func sortSlice(s string) string {
	r := []rune(s) // convert to rune slice (handles Unicode safely)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		key := sortSlice(word)
		groups[key] = append(groups[key], word)
	}
	result := [][]string{}
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}
```

6. Key Insight
Go mqp functioning so brutally clean to implement


What changed in thinking from brute force to optimal?
How to look optimization and repetition able to indetify the required optimization

7. Mistake / Learning
In correct handling of loop with brute force Added as notes above
Iterating over a growing array unknowinglu
Time tracking is also important