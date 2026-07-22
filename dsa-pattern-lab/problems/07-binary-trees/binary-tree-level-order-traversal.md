# 102. Binary Tree Level Order Traversal

## Problem Understanding

Given the root of a binary tree, return the **level order traversal** of its nodes' values.

Traverse the tree **level by level**, from **left to right**.

Example:

```text
Input:
        3
       / \
      9  20
         / \
        15  7

Output:
[[3], [9,20], [15,7]]
```

---

## Brute Force Approach

### Idea

1. Compute the height of the tree.
2. For every level `0...height`, traverse the tree again and collect nodes at that level.

### Complexity

**Time Complexity:** `O(N²)`

- Each level requires traversing a large portion of the tree again.

**Space Complexity:** `O(H)`

- Recursive call stack.

---

## Why It Is Inefficient

The same ancestor nodes are revisited for every level.

Example:

```text
Root

↓

Level 1

↓

Level 2

↓

Level 3
```

To reach nodes at Level 3, we repeatedly traverse through the root and intermediate nodes.

This repeated traversal makes the solution inefficient.

---

## Pattern

**Breadth-First Search (BFS)**

BFS naturally visits nodes level by level using a queue.

---

## Key Invariant

> **At the beginning of each iteration, every node currently in the queue belongs to exactly one tree level.**

Therefore:

1. Record the queue length.
2. Process exactly that many nodes.
3. Enqueue their children.
4. The newly added nodes automatically become the next level.

---

## Approach

- Initialize the queue with the root.
- While the queue is not empty:
  - Record the current queue size.
  - Process exactly `size` nodes.
  - Store their values.
  - Add their children to the queue.
- Append the collected values for that level to the result.

---

## Optimized Solution

```go
func levelOrder(root *TreeNode) [][]int {
	nodes := [][]int{}
	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		levelNodes := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		nodes = append(nodes, levelNodes)
	}

	return nodes
}
```

---

## Complexity

**Time Complexity:** `O(N)`

- Every node is visited exactly once.
- Each node is enqueued and dequeued exactly once.

**Space Complexity:** `O(W)`

Where `W` is the maximum width of the tree.

- Worst case: `O(N)`
- Balanced tree: up to approximately `O(N/2)` at the last level.

---

## Why This Works

The queue preserves the order in which nodes are discovered.

By recording the queue length **before** processing a level, we separate the current level from the next level.

Even though child nodes are added while processing, they are **not processed immediately**, because the loop only processes the number of nodes that were already present in the queue.

---

## Mistakes / Learnings

- **The queue length at the start of each iteration represents exactly one level of the tree.**
- Always record the queue length **before** processing the level.
- Child nodes added during processing belong to the **next level**.
- `O(N)` does **not** mean there can only be one loop. Nested loops are acceptable if every node is processed only once overall.
- BFS is the natural choice whenever a tree needs to be traversed **level by level**.

---

## Mental Model

> **A queue naturally separates the tree into levels.**

At the start of each iteration:

```
Queue

↓

Exactly one level
```

Process only those nodes.

Everything added afterward belongs to the next level.