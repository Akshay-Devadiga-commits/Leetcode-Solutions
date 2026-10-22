# 3942. Minimum Operations to Sort a Permutation

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=XUdwjrb5eBk)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=XUdwjrb5eBk)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You are given an integer array 
nums
 of length 
n
, where 
nums
 is a 
permutation
 of the integers from 0 to 
n - 1
.


You may perform 
only
 the following operations:




Reverse
 the entire array.


Rotate Left by One
: Move the first element to the end of the array, and rest elements to left by one position.




Return an integer denoting the 
minimum
 number of operations required to sort the array in 
increasing
 order. If it is 
not possible
 to sort the array using only the given operations, return -1.

## 💡 Approach: Breadth-First Search (BFS)

Since the constraints on n are very small (n <= 10), we can model the problem as finding the shortest path in an unweighted graph where each state is a permutation of nums and each transition is one of the two allowed operations: reversing the array or rotating left by one. Breadth-First Search (BFS) is guaranteed to find the minimum number of operations to reach the target sorted state (0, 1, 2, ..., n-1).

## ⏱️ Complexity

- **Time Complexity:** `O(n! * n)`
- **Space Complexity:** `O(n!)`
