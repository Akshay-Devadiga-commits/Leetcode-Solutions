# 3947. Maximum Number of Items From Sale II

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=Z4ieLPS38xo)

**[Check out the complete code implementations across Python, Java, C++, JavaScript, and Go in our GitHub repository linked below!](https://www.youtube.com/watch?v=Z4ieLPS38xo)**

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

You are given a 2D integer array  items , where  items[i] = [factor_i, price_i]  represents the  i^th  item. You are also given an integer  budget .


There are unlimited copies of each item available for purchase. You may buy any number of copies of any items such that the total cost of the purchased copies is at most  budget .


After buying items, you may receive free copies according to the following rules:




Each purchased copy of item  i  can give you  at most one  free copy of another item  j .


The free item must satisfy  i != j  and  factor_i  divides  factor_j .


For each ordered pair  (i, j) , you can receive a free copy of item  j  from purchases of item  i   at most once , regardless of how many copies of item  i  you buy.


The same item  j  can be received multiple times for free if it is received from purchases of different item types.




Return the  maximum total number of item copies  you can obtain, including both purchased copies and free copies, while spending at most  budget  on purchased items.

## Approach: Official Approach 1

Official LeetCode solution.

## Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)`
