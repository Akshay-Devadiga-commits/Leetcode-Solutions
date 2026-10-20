package main

import (
	"sort"
)

func sortableIntegers(nums []int) int {
	n := len(nums)
	totalSum := 0

	isSortedRotation := func(sub []int) bool {
		m := len(sub)
		if m <= 1 {
			return true
		}
		drops := 0
		for i := 0; i < m; i++ {
			if sub[i] > sub[(i+1)%m] {
				drops++
			}
		}
		return drops <= 1
	}

	for k := 1; k <= n; k++ {
		if n%k != 0 {
			continue
		}

		valid := true
		var sortedSubarrays [][]int

		for i := 0; i < n; i += k {
			sub := make([]int, k)
			copy(sub, nums[i:i+k])

			if !isSortedRotation(sub) {
				valid = false
				break
			}

			m := len(sub)
			startIdx := 0
			for j := 0; j < m; j++ {
				if sub[j] > sub[(j+1)%m] {
					startIdx = (j + 1) % m
					break
				}
			}

			rotated := make([]int, m)
			for j := 0; j < m; j++ {
				rotated[j] = sub[(startIdx+j)%m]
			}
			sortedSubarrays = append(sortedSubarrays, rotated)
		}

		if !valid {
			continue
		}

		var flat []int
		for _, sub := range sortedSubarrays {
			flat = append(flat, sub...)
		}

		isNonDecreasing := true
		for i := 0; i < n-1; i++ {
			if flat[i] > flat[i+1] {
				isNonDecreasing = false
				break
			}
		}

		if isNonDecreasing {
			totalSum += k
		}
	}

	return totalSum
}