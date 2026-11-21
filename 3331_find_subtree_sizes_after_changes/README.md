# 3331. Find Subtree Sizes After Changes

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=qqERM_Mo0z0)

**[Check out the full code implementation and more resources at the link in the description.](https://www.youtube.com/watch?v=qqERM_Mo0z0)**

**YouTube Playlist:** [Medium LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLDwpLRpPSYZI)

## Solutions

| Language | Solution File |
|---|---|
| Python 3 | [`solution.py`](solution.py) |
| Java | [`Solution.java`](Solution.java) |
| C++ | [`solution.cpp`](solution.cpp) |
| JavaScript | [`solution.js`](solution.js) |
| Go | [`solution.go`](solution.go) |

## Problem Statement

You are given a tree rooted at node 0 that consists of  n  nodes numbered from  0  to  n - 1 . The tree is represented by an array  parent  of size  n , where  parent[i]  is the parent of node  i . Since node 0 is the root,  parent[0] == -1 .


You are also given a string  s  of length  n , where  s[i]  is the character assigned to node  i .


We make the following changes on the tree  one  time  simultaneously  for all nodes  x  from  1  to  n - 1 :




Find the  closest  node  y  to node  x  such that  y  is an ancestor of  x , and  s[x] == s[y] .


If node  y  does not exist, do nothing.


Otherwise,  remove  the edge between  x  and its current parent and make node  y  the new parent of  x  by adding an edge between them.




Return an array  answer  of size  n  where  answer[i]  is the  size  of the  subtree  rooted at node  i  in the  final  tree.

## Approach: Array Approach

This solution uses the Array Approach with time complexity O(nh) and auxiliary space O(n). It achieves optimal performance by taking advantage of Array, Depth-First Search, Hash Table, String, Tree.

## Complexity

- **Time Complexity:** `O(nh)`
- **Space Complexity:** `O(n)`
