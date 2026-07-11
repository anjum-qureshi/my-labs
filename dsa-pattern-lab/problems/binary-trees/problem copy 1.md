# Problem Name

## 1. Problem Understanding
Write the problem in your own words.

---

## 2. Brute Force Approach
What is the simplest way to solve this?

```go
// brute force code here

Time Complexity: O(N)
Space Complexity: O(H)

3. Why it is inefficient
What is being repeated?
What is slow or redundant?
4. Pattern Identification

What pattern does this look like?
Hash Map

Why this pattern?

5. Optimized Solution
// optimized code here

6. Key Insight

Time Complexity: O(N), where N is the number of nodes in the tree.

Each node is visited exactly once during the DFS traversal. At each node, we perform a constant amount of work—checking the base cases, making two recursive calls, and combining the results from the left and right subtrees. Since no node is processed more than once, the total time complexity is O(N).

One or two lines:

What changed in thinking from brute force to optimal?

7. Mistake / Learning
What did I miss initially?
What should I notice next time?
