# 3364. Minimum Positive Sum Subarray 

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=RrJlv9grQmE)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=RrJlv9grQmE)**

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

You are given an integer array 
nums
 and 
two
 integers 
l
 and 
r
. Your task is to find the 
minimum
 sum of a 
subarray
 whose size is between 
l
 and 
r
 (inclusive) and whose sum is greater than 0.


Return the 
minimum
 sum of such a subarray. If no such subarray exists, return -1.


A 
subarray
 is a contiguous 
non-empty
 sequence of elements within an array.

## 💡 Approach: Brute Force with Sliding Window Optimization

The problem asks for the minimum positive sum of a subarray with a length between `l` and `r`. Since the constraints on `nums.length` are small (up to 100), a brute-force approach of checking all possible subarrays is feasible. We can iterate through all possible start indices `i` and end indices `j` of a subarray. For each subarray `nums[i:j+1]`, we calculate its sum and check if its length is within the range `[l, r]` and if its sum is greater than 0. We keep track of the minimum such sum found. To optimize the sum calculation, we can use a sliding window approach. For a fixed start index `i`, as we extend the end index `j`, we can update the sum incrementally instead of recalculating it from scratch. This avoids redundant computations. The constraints `1 <= l <= r <= nums.length` and `1000 <= nums[i] <= 1000` are important. The constraint on `nums[i]` being positive simplifies things slightly, as any subarray of length at least 1 will have a positive sum. However, the problem statement allows negative numbers in `nums` (as seen in example 1), so we must handle that. The constraint `1000 <= nums[i] <= 1000` seems to be a typo and should likely be `-1000 <= nums[i] <= 1000` or similar, given the examples. Assuming the examples are correct and `nums[i]` can be negative, the logic remains the same: find the minimum positive sum.

## ⏱️ Complexity

- **Time Complexity:** `O(n^2)`
- **Space Complexity:** `O(1)`
