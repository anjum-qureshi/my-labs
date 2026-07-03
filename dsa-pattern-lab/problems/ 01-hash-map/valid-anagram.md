# Problem Name

## 1. Problem Understanding
Given two strings s and t, Anagram are words that can be constructed with same characters and same length.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    used := make([]bool, len(t))

    for i := 0; i < len(s); i++ {
        found := false

        for j := 0; j < len(t); j++ {
            if !used[j] && s[i] == t[j] {
                used[j] = true
                found = true
                break
            }
        }

        if !found {
            return false
        }
    }

    return true
}
```

Time Complexity: O(n2)
Space Complexity: O(n) used

3. Why it is inefficient
Repeated the checked for the characters appearing in the inner loop

4. Pattern Identification
Frequency count

What pattern does this look like?
Hash Map

Why this pattern?
We can just store the frequency of a character and whether it used in a map, no need to visit it for every loop again. 

5. Optimized Solution
```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make(map[byte]int,0)

    for i := 0; i < len(s); i++ {
        count[s[i]]++
        count[t[i]]--
    }

    for _, f := range count {
        if f != 0 {
            return false
        }
    }
    return true
}
```

6. Key Insight



7. Mistake / Learning
What did I miss initially?
What should I notice next time?