# 98. Validate Binary Search Tree

## Problem Understanding

Given the root of a binary tree, determine whether it is a valid **Binary Search Tree (BST)**.

A valid BST satisfies:

- Every node in the **left subtree** has a value **strictly less** than the current node.
- Every node in the **right subtree** has a value **strictly greater** than the current node.
- Both the left and right subtrees must also be valid BSTs.

---

## Complexity

**Time Complexity:** `O(N)`

- Every node is visited exactly once during the traversal.

**Space Complexity:** `O(H)`

Where `H` is the height of the tree.

- The stack stores at most one root-to-leaf path.
- Worst case (skewed tree): `O(N)`
- Balanced tree: `O(log N)`

---

## Optimized Solution

```go
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode

	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			lastIndex := len(stack) - 1
			if lastIndex < 0 {
				break
			}

			node = stack[lastIndex]
			stack = stack[:lastIndex]

			if prev != nil && prev.Val >= node.Val {
				return false
			}

			prev = node
			node = node.Right
		}
	}

	return true
}
```

---

## Key Insight

A valid BST produces a **strictly increasing sequence** when traversed using **in-order traversal** (`Left → Root → Right`).

Instead of validating every subtree independently, we can simply verify that every node visited during the in-order traversal has a value greater than the previously visited node.

Algorithm:

1. Perform an in-order traversal.
2. Keep track of the previously visited node.
3. If `prev.Val >= current.Val`, the BST property is violated.
4. Otherwise, continue traversing.

---

## Why This Works

In-order traversal visits nodes in sorted order **only if the tree is a valid BST**.

The BST property is therefore transformed into a much simpler problem:

> Is the in-order traversal strictly increasing?

This avoids checking the BST property separately for every subtree.

---

## Mistakes / Learnings

- A valid BST is **not** validated by comparing a node only with its immediate children.
- BST validation is about maintaining the **global ordering** of all nodes.
- A valid BST always produces a **strictly increasing** sequence during **in-order traversal**.
- Use `>=` instead of `>` because duplicate values are **not allowed** in a valid BST.
- The iterative solution is identical to recursive in-order traversal—the explicit stack simply simulates the recursive call stack.

---

## Mental Model

> **BST validation is not about checking parent-child relationships.**

It is about verifying that the **entire tree** follows the BST ordering property.

If the **in-order traversal is strictly increasing**, the tree is a valid BST.