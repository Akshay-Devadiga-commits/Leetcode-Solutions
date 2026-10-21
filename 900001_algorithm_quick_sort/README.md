# 900001. Quick Sort

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=fuq484zVTuc)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=fuq484zVTuc)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

Explain Quick Sort as a commonly used sorting algorithm. Sort data by partitioning elements around a chosen pivot value such that smaller elements move to the left and larger elements move to the right. Show the core idea, a concrete example, correctness intuition, implementation, and complexity.

## 💡 Approach: Quick Sort

Quick Sort is a divide-and-conquer algorithm. It selects a pivot element from the array and partitions the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively. The Lomuto or Hoare partition schemes can be used to rearrange the array efficiently in place.

## ⏱️ Complexity

- **Time Complexity:** `O(n log n)`
- **Space Complexity:** `O(log n)`
