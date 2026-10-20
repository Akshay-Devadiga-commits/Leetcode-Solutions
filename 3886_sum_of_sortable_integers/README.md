# 3886. Sum of Sortable Integers

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution
Watch the complete step-by-step visual explanation and algorithmic intuition on YouTube:
👉 **[Watch on YouTube](https://www.youtube.com/watch?v=v78CWGaatuM)**

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
.


An integer 
k
 is called 
sortable
 if 
k
 
divides
 
n
 and you can sort 
nums
 in 
non-decreasing
 order by sequentially performing the following operations:




Partition 
nums
 into 
consecutive 
subarrays
 of length 
k
.


Cyclically rotate each subarray independently
 any number of times to the left or to the right.




Return an integer denoting the sum of all possible sortable integers 
k
.

## 💡 Approach: Divisor Check and Subarray Rotation Validation

To determine if an integer k is sortable, we first check if k divides n. If it does, we partition nums into n / k subarrays of length k. Each subarray can be independently rotated. For the entire array to be sortable into non-decreasing order: 1) The concatenation of the minimum elements of each subarray in their given order must be sorted. 2) Each subarray itself must be a valid cyclic rotation of a sorted sequence. 3) The maximum element of the i-th subarray must be less than or equal to the minimum element of the (i+1)-th subarray. We check all divisors k of n and sum them up if they meet all conditions.

## ⏱️ Complexity

- **Time Complexity:** `O(n^2)`
- **Space Complexity:** `O(n)`
