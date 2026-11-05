# 1385. Find the Distance Value Between Two Arrays

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=yMmBswq8CqU)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=yMmBswq8CqU)**

📂 **YouTube Playlist:** [Easy LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLO_wgoXxiqtY)

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

Given two integer arrays 
arr1
 and 
arr2
, and the integer 
d
, 
return the distance value between the two arrays
.


The distance value is defined as the number of elements 
arr1[i]
 such that there is not any element 
arr2[j]
 where 
|arr1[i]-arr2[j]| <= d
.

## 💡 Approach: Sorting and Binary Search

The problem asks us to count elements in `arr1` for which there is no element in `arr2` within a distance `d`. A naive approach would be to iterate through each element of `arr1` and for each element, iterate through all elements of `arr2` to check the distance condition. This would result in an O(n*m) time complexity, where n and m are the lengths of `arr1` and `arr2` respectively. Given the constraints (lengths up to 500), O(n*m) might be acceptable (500*500 = 250,000 operations per test case), but we can optimize it.

To optimize, we can sort `arr2`. Once `arr2` is sorted, for each element `x` in `arr1`, we need to efficiently check if there exists any element `y` in `arr2` such that `|x - y| <= d`. This inequality can be rewritten as `x - d <= y <= x + d`. Since `arr2` is sorted, we can use binary search to find if any element `y` falls within the range `[x - d, x + d]`. Specifically, we can search for the smallest element in `arr2` that is greater than or equal to `x - d`. If such an element exists and is also less than or equal to `x + d`, then the condition `|x - y| <= d` is met for this `x`, and `x` should not be counted towards the distance value. If no such element is found in `arr2` within the range `[x - d, x + d]`, then `x` contributes to the distance value.

Using binary search (like `bisect_left` in Python) on the sorted `arr2` for each element in `arr1` reduces the check for each `arr1` element from O(m) to O(log m). Therefore, the overall time complexity becomes O(m log m) for sorting `arr2` plus O(n log m) for iterating through `arr1` and performing binary search, resulting in O(m log m + n log m).

## ⏱️ Complexity

- **Time Complexity:** `O(m log m + n log m)`
- **Space Complexity:** `O(log m) or O(m) depending on sort implementation`
