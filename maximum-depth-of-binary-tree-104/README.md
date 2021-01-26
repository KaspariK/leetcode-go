# 104. Maximum Depth of Binary Tree

Difficulty: Easy

Given a binary tree, find its *maximum depth*.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.
**Note:** A leaf is a node with no children.

**Example 1:**

![Binary Tree](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

**Example 2:**

```
Input: root = [1,null,2]
Output: 2
```

**Example 3:**

```
Input: root = []
Output: 0
```

**Example 4:**

```
Input: root = [0]
Output: 1
```

**Cosntraints:**

* The number of nodes in the tree is in the range `[0, 10^4]`.
* `-100 <= Node.val <= 100`