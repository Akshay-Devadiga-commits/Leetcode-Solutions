# 3886. Sum of Sortable Integers

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution
Watch the complete step-by-step visual explanation and algorithmic intuition on YouTube:
👉 **[Watch on YouTube](https://www.youtube.com/watch?v=SRjLq07ObBM)**

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

To solve this problem, we first find all divisors k of n (the length of the array). For each divisor k, we partition the array nums into n/k consecutive subarrays of length k. Since we can cyclically rotate each subarray independently, a subarray of length k can achieve a sorted state if and only if its elements can form a contiguous sorted sequence when rotated. More simply, each subarray when rotated must become non-decreasing. Furthermore, the maximum element of the i-th subarray must be less than or equal to the minimum element of the (i+1)-th subarray so that the whole array ends up sorted globally. We iterate through all divisors k, check if these conditions hold for all subarrays, and sum up all valid k values.

## ⏱️ Complexity

- **Time Complexity:** `O(n * sigma(n)) where sigma(n) is the sum of divisors of n, which is very small for n <= 10.`
- **Space Complexity:** `O(n)`
