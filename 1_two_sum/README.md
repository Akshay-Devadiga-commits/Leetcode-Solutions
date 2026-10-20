# 1. Two Sum

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution
Watch the complete step-by-step visual explanation and algorithmic intuition on YouTube:
👉 **[Watch on YouTube](https://www.youtube.com/watch?v=mmRijkj8_Y8)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## 💡 Approach: Hash Map

Instead of checking every possible pair using nested loops (which would take O(n^2) time), we can iterate through the array once and use a hash map to keep track of the numbers we have already seen and their indices. For each element, we calculate its complement (target - current_element). If the complement exists in our hash map, we have found our two numbers and can immediately return their indices. Otherwise, we add the current element and its index to the hash map.

## ⏱️ Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(n)`
