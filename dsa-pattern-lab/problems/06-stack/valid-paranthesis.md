# 20. Valid Parentheses

## 1. Problem Understanding

You are given a string `s` consisting only of the characters:

- `(`
- `)`
- `{`
- `}`
- `[`
- `]`

A string is valid if:

1. Every opening bracket has a corresponding closing bracket of the same type.
2. Brackets are closed in the correct order.
3. No closing bracket appears without a matching opening bracket.

Return `true` if the string is valid, otherwise `false`.

---

# 2. Brute Force

## Intuition

Repeatedly remove adjacent valid pairs:

- `()`
- `{}`
- `[]`

If the string eventually becomes empty, it is valid.

If an entire pass removes nothing, the remaining brackets can never be matched.

### Example

```
([{}])

Pass 1:
([{}])
   ^^

↓

([])

Pass 2:
([])
 ^^

↓

()

Pass 3:
()

↓

""
```

Valid.

Example:

```
([)]
```

No adjacent valid pairs exist.

Nothing can be removed.

Invalid.

### Implementation

```go
func isValid(s string) bool {
	for {
		changed := false

		for i := 0; i < len(s)-1; i++ {
			pair := s[i : i+2]

			if pair == "()" || pair == "{}" || pair == "[]" {
				s = s[:i] + s[i+2:]
				changed = true
				break
			}
		}

		if !changed {
			break
		}
	}

	return len(s) == 0
}
```

### Complexity

- Time: **O(n²)**
- Space: **O(n)** (new strings are created during deletion)

Why?

- There can be up to `n/2` removal passes.
- Each pass scans nearly the entire string.

---

# 3. Key Observation

Instead of repeatedly deleting pairs, think about **what information must be remembered while scanning the string once.**

At any point, the only information we need is:

> **Which opening brackets have not been closed yet?**

Example:

```
([{
```

These brackets are still waiting to be matched.

When we encounter:

```
}
```

Which opening bracket should it match?

Not the first.

Not the second.

It must match the **most recently opened unmatched bracket**.

This is **Last-In, First-Out (LIFO)**.

A **Stack** naturally models this behavior.

---

# 4. Algorithm

1. Create an empty stack.
2. Traverse the string from left to right.
3. If the character is an opening bracket:
   - Push it onto the stack.
4. Otherwise it is a closing bracket:
   - If the stack is empty → Invalid.
   - Check whether the top of the stack matches the corresponding opening bracket.
   - If not → Invalid.
   - Otherwise pop the stack.
5. After processing all characters:
   - If the stack is empty → Valid.
   - Otherwise → Invalid.

---

# 5. Implementation

```go
func isValid(s string) bool {
	stack := []byte{}

	brackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		curr := s[i]

		if curr == '(' || curr == '[' || curr == '{' {
			stack = append(stack, curr)
		} else {
			lastIndex := len(stack) - 1

			// Closing bracket without an opening bracket
			if lastIndex < 0 {
				return false
			}

			// Wrong type of opening bracket
			if brackets[curr] != stack[lastIndex] {
				return false
			}

			// Matched, remove it
			stack = stack[:lastIndex]
		}
	}

	return len(stack) == 0
}
```

---

# 6. Dry Run

Input:

```
([{}])
```

| Character | Stack | Action |
|-----------|-------|--------|
| ( | ( | Push |
| [ | ([ | Push |
| { | ([{ | Push |
| } | ([ | Pop |
| ] | ( | Pop |
| ) | Empty | Pop |

Stack is empty.

Answer = `true`

---

# 7. Edge Cases

### Closing bracket without an opening bracket

```
]
```

Stack is empty.

Return `false`.

---

### Wrong bracket type

```
(]
```

Top of stack:

```
(
```

Expected:

```
[
```

Mismatch.

Return `false`.

---

### Unclosed opening brackets

```
(((
```

Traversal finishes.

Stack is not empty.

Return `false`.

---

# 8. Invariants

During traversal:

1. The stack always contains **only unmatched opening brackets**.
2. Every closing bracket must match the **top of the stack**.
3. A closing bracket can never be processed when the stack is empty.
4. The stack must be empty after processing the entire string.

---

# 9. Complexity

- Time: **O(n)**
- Space: **O(n)**

Each character is:

- pushed at most once
- popped at most once

making the solution linear.

---

# Pattern Recognition

When solving a problem, ask:

### What information do I need to remember?

For this problem:

```
Unmatched opening brackets
```

### Which remembered item do I need?

The **most recently added** unmatched opening bracket.

```
Newest
↓
Stack (LIFO)
```

This thought process naturally leads to choosing a **Stack**, rather than memorizing it as a "stack problem."