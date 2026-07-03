# 387. First Unique Character in a String

## 1. Problem Understanding
Given a string `s`, return the index of the first character that does not repeat anywhere in the string.
If no such character exists, return `-1`.

The goal is not just to find uniqueness, but to find the **first unique occurrence by order**.

---

## 2. Brute Force Approach
For each character, check whether it appears again elsewhere in the string.

- Pick a character at index `i`
- Scan the entire string to see if the same character appears at any other index
- If no duplicate is found, return `i`

```go
func firstUniqChar(s string) int {
    for i := 0; i < len(s); i++ {
        foundDuplicate := false

        for j := 0; j < len(s); j++ {
            if i != j && s[i] == s[j] {
                foundDuplicate = true
                break
            }
        }

        if !foundDuplicate {
            return i
        }
    }

    return -1
}
```

Time Complexity: O(n2)
Space Complexity: O(1)

## 3. Why it is inefficient
For every character, we rescan the entire string to confirm whether it repeats.

What is being repeated?
The same existence check for each character

What is redundant?
Recomputing whether a character is unique again and again instead of remembering past occurrences

4. Pattern Identification
Frequency counting / character aggregation

Why this pattern?
The problem depends on how many times each character appears, not on pairwise comparisons.

Key idea:

Replace repeated existence checks with a precomputed summary of occurrences.


5. Optimized Solution
```go
func firstUniqChar(s string) int {
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
```

Time Complexity: O(n)
Space Complexity: O(1)

6. Key Insight

Brute force repeatedly answers the question:

“Does this character repeat?”

Optimization changes the question to:

“How many times does each character appear?”

So instead of repeated checking, we compute a global summary once.

7. Mistake / Learning
Initial approach focused on pairwise checking and early stopping.

Missed that:

uniqueness is a global property, not a local comparison problem
precomputing frequency removes repeated work entirely

Key Insight:

When a problem repeatedly asks the same question per element, precompute the answer once globally.
