# 2409. Count Days Spent Together

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=wkrlMj3znDw)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=wkrlMj3znDw)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

Alice and Bob are traveling to Rome for separate business meetings.


You are given 4 strings 
arriveAlice
, 
leaveAlice
, 
arriveBob
, and 
leaveBob
. Alice will be in the city from the dates 
arriveAlice
 to 
leaveAlice
 (
inclusive
), while Bob will be in the city from the dates 
arriveBob
 to 
leaveBob
 (
inclusive
). Each will be a 5-character string in the format 
"MM-DD"
, corresponding to the month and day of the date.


Return
 the total number of days that Alice and Bob are in Rome together.


You can assume that all dates occur in the 
same
 calendar year, which is 
not
 a leap year. Note that the number of days per month can be represented as: 
[31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
.

## 💡 Approach: Date to Day Number Conversion and Interval Overlap

To find the number of days Alice and Bob spend together, we can first convert each date string in the format 'MM-DD' into an absolute day number of the year (from 1 to 365). We use the array of days per month in a non-leap year to compute this easily. Once we have converted Alice's arrival/leaving dates and Bob's arrival/leaving dates into day numbers, the overlapping period is simply from the maximum of the two arrival days to the minimum of the two leaving days. If the maximum arrival day is less than or equal to the minimum leaving day, the overlap duration is (min_leave - max_arrive + 1). Otherwise, they do not overlap, and the answer is 0.

## ⏱️ Complexity

- **Time Complexity:** `O(1)`
- **Space Complexity:** `O(1)`
