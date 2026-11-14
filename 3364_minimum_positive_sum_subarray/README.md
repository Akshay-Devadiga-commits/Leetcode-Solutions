# 3364. Minimum Positive Sum Subarray 

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=PkfLDGRog74)

**[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=PkfLDGRog74)**

**YouTube Playlist:** [Easy LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLO_wgoXxiqtY)

## Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| Python 3 | [`solution.py`](solution.py) |
| Java | [`Solution.java`](Solution.java) |
| C++ | [`solution.cpp`](solution.cpp) |
| JavaScript | [`solution.js`](solution.js) |
| Go | [`solution.go`](solution.go) |

## Problem Statement

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

## Approach: Prefix Sums with Nested Loops

The problem asks for the minimum positive sum of a subarray whose length is between l and r (inclusive). Given the small constraint on the array length (nums.length <= 100), an O(N^2) solution is efficient enough.

The core idea is to iterate through all possible subarrays, check their length and sum, and update the minimum positive sum found so far.

To efficiently calculate the sum of any subarray nums[i...j], we can use a prefix sum array. A prefix_sum array is constructed such that prefix_sum[k] stores the sum of the first k elements of nums (i.e., nums[0] + ... + nums[k-1]). With this, the sum of nums[i...j] can be calculated in O(1) time as prefix_sum[j+1] - prefix_sum[i].

Algorithm steps:
1. Initialize min_positive_sum to infinity (e.g., math.inf). This variable will store the smallest positive sum found.
2. Create a prefix_sum array of size n + 1. prefix_sum[0] is 0. For k from 1 to n, prefix_sum[k] = prefix_sum[k-1] + nums[k-1].
3. Use nested loops to iterate through all possible subarrays:
   a. The outer loop iterates i from 0 to n-1 (representing the start index of the subarray).
   b. The inner loop iterates j from i to n-1 (representing the end index of the subarray).
4. For each subarray nums[i...j]:
   a. Calculate its length: current_length = j - i + 1.
   b. Check if l <= current_length <= r. If not, this subarray does not meet the length criteria, so continue to the next subarray.
   c. If the length criteria is met, calculate its sum: current_sum = prefix_sum[j+1] - prefix_sum[i].
   d. Check if current_sum > 0. If it is, update min_positive_sum = min(min_positive_sum, current_sum).
5. After checking all possible subarrays, if min_positive_sum is still infinity, it means no subarray satisfied both the length and positive sum criteria. In this case, return -1.
6. Otherwise, return min_positive_sum.

## Complexity

- **Time Complexity:** `O(N^2)`
- **Space Complexity:** `O(N)`
