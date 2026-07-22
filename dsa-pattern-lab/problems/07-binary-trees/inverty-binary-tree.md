# 226. Invert Binary Tree

## Problem Understanding

Given the root of a binary tree, invert the tree and return its root.

Inverting a binary tree means:

> At every node, swap its left and right subtrees.

Example:

```text
Original

        4
      /   \
     2     7
    / \   / \
   1   3 6   9

↓

Inverted

        4
      /   \
     7     2
    / \   / \
   9   6 3   1
```

---

**Time Complexity:** **O(N)**

Every node is visited exactly once.

**Space Complexity:** **O(H)**

Where **H** is the height of the tree.

- Balanced tree: **O(log N)**
- Skewed tree: **O(N)**

---

## Brute Force Approach

Create a completely new tree.

For every node:

- Create a new node with the same value.
- The new node's left child becomes the inverted right subtree.
- The new node's right child becomes the inverted left subtree.

This works but requires creating an entirely new tree.

**Time Complexity:** **O(N)**

**Space Complexity:** **O(N)** (new tree) + **O(H)** (recursion stack)

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    return &TreeNode{
        Val:   root.Val,
        Left:  invertTree(root.Right),
        Right: invertTree(root.Left),
    }
}
```

---

## Recursive Contract

The recursive function answers one question:

> **Invert the subtree rooted at the current node and return its root.**

### Base Case

If the node is `nil`:

```text
Return nil
```

### Recursive Case

1. Swap the left and right children.
2. Invert the new left subtree.
3. Invert the new right subtree.
4. Return the current node.

---

## Key Observation

The inversion operation is **completely local**.

Every node performs exactly the same operation:

```text
Swap

Left ↔ Right
```

After swapping, each child subtree is simply another smaller instance of the same problem.

This naturally leads to recursion.

---

## Optimized Solution (In-place)

```go
func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        root.Left, root.Right = root.Right, root.Left

        invertTree(root.Left)
        invertTree(root.Right)
    }

    return root
}
```

---

## Why this works

For every node:

1. Swap the left and right children.
2. Recursively invert both subtrees.
3. Return the current node.

Since every subtree is itself a binary tree, recursively applying the same operation eventually inverts the entire tree.

---

## Recursive Flow

```text
invert(node)

↓

Swap left and right children

↓

invert(left subtree)

↓

invert(right subtree)

↓

Return current node
```

---

## Key Insight

Unlike many tree recursion problems, the parent does **not** need any information from its children.

Each node can immediately perform its own work (swap the children), then recursively ask each subtree to do the same.

The recursive function simply returns the root of the inverted subtree.

---

## Mistake / Learning

- The operation is performed **at every node independently**.
- No additional data structures are required.
- There is no need to calculate heights, paths, or maintain global state.
- The recursive contract is simply:

```text
Invert my subtree and return me.
```

- An alternative solution is to construct a completely new inverted tree, but the in-place approach is simpler and uses only the recursion stack.