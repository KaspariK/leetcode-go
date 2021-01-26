# 100. Same Tree

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

**Example 1:**

![Example 1](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

**Example 2:**

![Example 1](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)


```
Input: p = [1,2], q = [1,null,2]
Output: false
```

**Example 3:**

![Example 1](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

**Constraints:**

* The number of nodes in both trees is in the range of `[0, 100]`.
* `-10^4 <= Node.val <= 10^4`