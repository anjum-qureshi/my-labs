# Problem Name

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
func normalize(s string) string {
	s = strings.ToLower(s)

	var builder strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func isPalindrome(s string) bool {
    s = normalize(s)
    r := ""
    for i:=len(s)-1;i>-1; i-- {
        r = r + string(s[i])
    }
    return r == s
}
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
func isAlphaNumeric(b byte) bool {
    r := rune(b)
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    start, end := 0, len(s)-1
    for  start < end {
        for start < end && !isAlphaNumeric(s[start]) {
            start = start + 1
        }
        for  start < end && !isAlphaNumeric(s[end]) {
            end = end - 1
        }
        if rune(s[start]) != rune(s[end]) {
           return false
        } 
        start, end = start+1, end-1
    }
    return true
}```

6. Key Insight

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
What did I miss initially?
What should I notice next time?
