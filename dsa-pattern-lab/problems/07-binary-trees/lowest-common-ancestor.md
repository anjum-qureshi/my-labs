# 236. Lowest Common Ancestor of a Binary Tree

## Problem Understanding
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

ex:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Example:

        3
      /   \
     5     1
    / \   / \
   6   2 0   8
      / \
     7   4

p = 5
q = 1
LCA = 3
---

#. Complexity
** Time Complexity: ** O(N)

Every node is visited exactly once.

** Space Complexity: ** O(H)

3. Key Recursive Contract

The recursive function answers one question:

"Does this subtree contain either p, q, or their Lowest Common Ancestor?"

The function returns one of three things:

nil → Neither target exists in this subtree.
p or q → One target was found in this subtree.
LCA → This subtree already contains the Lowest Common Ancestor.
---

## Recursive Thinking
##### Base Cases

Stop recursion when:

Current node is nil
Current node is p
Current node is q

Return the current node.

## Recursive Step

Search both subtrees.

left = search(left subtree)

right = search(right subtree)

Each subtree tells us whether it found:

Nothing (nil)
One target
The LCA

## Combining Results
# Case 1
left = nil

right = nil

Neither subtree contains a target.

Return:

nil

# Case 2
left = node

right = nil

Only the left subtree contains a target (or the LCA).

Return:

left

# Case 3
left = nil

right = node

Only the right subtree contains a target (or the LCA).

Return:

right

# Case 4
left = node

right = node

One target was found in each subtree.

Therefore, the current node is the first node where both targets meet.

Return:

current node

Why this pattern?

5. Optimized Solution
```go
func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
	if node == nil || node == p || node == q {
		return node
	}

    left := lowestCommonAncestor(node.Left,p,q)
    right := lowestCommonAncestor(node.Right,p,q)
    if left != nil && right != nil {
        return node
    }

    if left != nil {
        return left
    }

    return right
} 
```

## Key Insight
The recursive function does not return a boolean.

It returns useful information to its parent:

no target found
one target found
the Lowest Common Ancestor

The parent combines the results from the left and right subtrees to determine whether the current node is the LCA.

## Mistake / Learning
- The hardest part is defining what recursion should return before writing any code.
- Think about the contract of the recursive function first.
- The parent only needs information from its left and right children to make its decision.
- If both left and right return non-nil, the current node is the Lowest Common Ancestor.
- If only one side returns a node, simply propagate that result upward.
- The recursion naturally stops when it reaches nil, p, or q.
