# 206. Reverse Linked List

## 1. Problem Understanding

Given the head of a singly linked list, reverse the list and return the new head.

Example:

1 -> 2 -> 3 -> 4 -> 5

becomes

5 -> 4 -> 3 -> 2 -> 1

---

## 2. Brute Force Thought

Store all nodes in an array.

Reverse the array.

Reconnect each node to the previous one.

Time: O(n)
Space: O(n)

Can we reverse the pointers without storing extra nodes?

Yes.

---

## 3. Key Observation

Each node currently points **forward**.

Example:

1 -> 2 -> 3 -> 4

When reversing,

1 should point to NULL

2 should point to 1

3 should point to 2

4 should point to 3

So for every node we need to:

Current
↓

current.next = previous

But...

Changing current.next immediately loses the rest of the list.

Example:

1 -> 2 -> 3

If we do

1.next = NULL

we lose access to

2 -> 3

unless we save it first.

This means we need another pointer.

---

## 4. The Three Pointer Pattern

We always maintain three pointers.

prev
current
next

Initially

prev = nil

current = head

next = nil

Example

nil <- 1 -> 2 -> 3 -> 4

---

## 5. Algorithm Thinking

At every node we perform exactly four steps.

### Step 1

Save the next node.

next = current.next

Why?

Because we are about to overwrite current.next.

Without saving it,
the remaining list is lost.

---

### Step 2

Reverse the pointer.

current.next = prev

Now current points backward.

Before

prev <- current -> next

After

prev <- current

---

### Step 3

Move prev forward.

prev = current

The reversed list has grown by one node.

---

### Step 4

Move current forward.

current = next

Continue processing the remaining list.

---

## 6. Visualization

Initial

nil <- 1 -> 2 -> 3 -> 4

Iteration 1

Save next

next = 2

Reverse

1 -> nil

Move

prev = 1

current = 2

Result

nil <- 1

Remaining

2 -> 3 -> 4

---

Iteration 2

Save next

next = 3

Reverse

2 -> 1

Move

prev = 2

current = 3

Result

nil <- 1 <- 2

Remaining

3 -> 4

---

Iteration 3

3 -> 2 -> 1

---

Iteration 4

4 -> 3 -> 2 -> 1

current = nil

Loop ends.

Return prev.

---

## 7. Why Return prev?

When current becomes nil,

prev points to

5 -> 4 -> 3 -> 2 -> 1

which is the new head.

---

## 8. Pattern Recognition

Whenever you need to reverse links in a linked list,

remember the three-pointer pattern.

prev

current

next

Save next
↓

Reverse pointer
↓

Move prev
↓

Move current

Repeat.

---

## 9. Complexity

Time

O(n)

Each node visited once.

Space

O(1)

Only three pointers are used.

---

## 10. Interview Thought Process

Interviewer:

"How would you reverse a linked list?"

Thinking:

- Every node should point backward.
- If I reverse immediately, I lose the remaining list.
- Save the next node first.
- Reverse current pointer.
- Move both pointers.
- Repeat until current becomes nil.
- prev will become the new head.

This naturally leads to the three-pointer solution.

---

## Cheat Code

Need to reverse pointers?

Think:

Save

Reverse

Advance

```
next = current.next
current.next = prev
prev = current
current = next
```

Repeat until current == nil.