# 543. Diameter of Binary Tree

## Problem Understanding

Given the root of a binary tree, return the **diameter** of the tree.

The **diameter** of a binary tree is the **length (number of edges)** of the **longest path between any two nodes** in the tree.

The path:
- May or may not pass through the root.
- Can start and end at any two nodes.

Example:

```text
        1
       / \
      2   3
     / \
    4   5
```

The longest path is:

```text
4 → 2 → 1 → 3
```

Number of edges = **3**

---

**Time Complexity:** **O(N)**

Every node is visited exactly once.

**Space Complexity:** **O(H)**

Where **H** is the height of the tree.

- Balanced tree: **O(log N)**
- Skewed tree: **O(N)**

---

## Recursive Contract

The recursive function answers **one question**:

> **What is the height of the subtree rooted at the current node?**

The recursive function **does not return the diameter**.

It returns only the subtree height.

At every node we perform two independent tasks:

1. Compute the longest path passing through the current node.
2. Return the subtree height to the parent.

---

## Key Observation

Every node is the **meeting point** between its left and right subtrees.

If we already know:

```text
leftHeight
rightHeight
```

then the longest path passing through the current node is

```text
leftHeight + rightHeight
```

This becomes a **candidate diameter**.

The overall diameter is simply the maximum candidate found across all nodes.

---

## Optimized Solution

```go
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

    var height func(*TreeNode) int

    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := height(node.Left)
        rightHeight := height(node.Right)

        diameter = max(diameter, leftHeight+rightHeight)

        return 1 + max(leftHeight, rightHeight)
    }

    height(root)

    return diameter
}
```

---

## Why this works

For every node:

1. Ask the left child for its height.
2. Ask the right child for its height.
3. Compute the longest path passing through the current node.

```text
candidateDiameter = leftHeight + rightHeight
```

4. Update the global diameter if this candidate is larger.
5. Return the subtree height to the parent.

```text
height = 1 + max(leftHeight, rightHeight)
```

The parent only needs the height because a path extending upward can continue through **only one child**.

---

## Mental Model

Every recursive call has **two jobs**:

### Job 1 — Update the answer

Calculate the longest path passing through the current node.

```text
leftHeight + rightHeight
```

Update the maximum diameter.

---

### Job 2 — Answer the parent

Return the height of the subtree.

```text
1 + max(leftHeight, rightHeight)
```

The parent can only continue along **one** branch, so only the height is returned.

---

## Recursive Flow

```text
height(node)

↓

leftHeight = height(node.Left)

↓

rightHeight = height(node.Right)

↓

diameter = max(diameter, leftHeight + rightHeight)

↓

return 1 + max(leftHeight, rightHeight)
```

---

## Key Insight

The problem asks us to find the **longest path**, not the height.

However, the longest path through a node can only be computed if we know:

- the deepest path in the left subtree
- the deepest path in the right subtree

Therefore, recursion returns **height**, while the diameter is updated as a side effect at every node.

---

## Mistake / Learning

- Diameter is **not** the height of the tree.
- Diameter is the **longest path between any two nodes**.
- The recursive function **returns height**, not diameter.
- Each node computes a **candidate diameter** using:

```text
leftHeight + rightHeight
```

- The answer is the **maximum** candidate across the entire tree.
- The parent only needs the child's height because a path extending upward cannot branch into both subtrees.
- Understanding **what the recursive function should return** is the key to solving recursive tree problems.