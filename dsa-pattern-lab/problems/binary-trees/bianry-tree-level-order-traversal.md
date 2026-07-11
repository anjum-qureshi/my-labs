# 102. Binary Tree Level Order Traversal

## 1. Problem Understanding
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Reading the elements of a binary tree at level before going to next level.

Ex: Input: root = [3,9,20,null,null,15,7]
    Output: [[3],[9,20],[15,7]]

## 2. Brute Force Approach
Brute force approach would be to find the depth of the tree.
The read the nodes at each level where each level starts from root node.

```go
// brute force code here

Time Complexity: O(n2)
Space Complexity: O(1)

3. Why it is inefficient
We are revisiting the ancestor nodes when reading current level nodes. As we always know the root node.
What is slow or redundant?
Going through the root node to reach the level nodes is redundant.

What pattern does this look like?
Binary Tree Travesal BFS

5. Optimized Solution
```go
func levelOrder(root *TreeNode) [][]int {
	nodes := [][]int{}
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		size := queue.Len()
		levelNodes := []int{}
		for i := 0; i < size; i++ {
			element := queue.Front()
			value := element.Value.(*TreeNode)
			if value != nil {
				if value.Left != nil {
					queue.PushBack(value.Left)
				}
				if value.Right != nil {
					queue.PushBack(value.Right)
				}
				levelNodes = append(levelNodes, value.Val)
			}
			queue.Remove(element)
		}
		nodes = append(nodes, levelNodes)
	}
	return nodes
}
```

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
			element := queue[0]
			if element != nil {
				if element.Left != nil {
					queue = append(queue, element.Left)
				}
				if element.Right != nil {
					queue = append(queue, element.Right)
				}
				levelNodes = append(levelNodes, element.Val)
			}
            queue = queue[1:]
		}
		nodes = append(nodes, levelNodes)
	}
	return nodes
}
```

6. Key Insight
Writing brute force algorithm is harder to think.


What changed in thinking from brute force to optimal?

7. Mistake / Learning
I missed initially when to how to identify the level or separate them.
Think more literally of the solution step by step.
