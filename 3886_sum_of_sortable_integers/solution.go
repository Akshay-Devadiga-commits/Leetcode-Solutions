package main

type Solution struct{}

func (s *Solution) SortableIntegers(nums []int) int {
    n := len(nums)
    totalSum := 0
    var divisors []int
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            divisors = append(divisors, i)
        }
    }

    for _, k := range divisors {
        possible := true
        var subarrays [][]int
        for i := 0; i < n; i += k {
            sub := make([]int, k)
            copy(sub, nums[i:i+k])
            drops := 0
            for j := 0; j < k; j++ {
                if sub[j] > sub[(j+1)%k] {
                    drops++
                }
            }
            if drops > 1 {
                possible = false
                break
            }
            var sortedSub []int
            if drops == 1 {
                startIdx := 0
                for j := 0; j < k; j++ {
                    if sub[j] > sub[(j+1)%k] {
                        startIdx = (j + 1) % k
                        break
                    }
                }
                sortedSub = make([]int, k)
                for j := 0; j < k; j++ {
                    sortedSub[j] = sub[(startIdx+j)%k]
                }
            } else {
                sortedSub = sub
            }
            subarrays = append(subarrays, sortedSub)
        }
        if !possible {
            continue
        }
        var flatArray []int
        for _, sub := range subarrays {
            flatArray = append(flatArray, sub...)
        }
        for i := 0; i < n-1; i++ {
            if flatArray[i] > flatArray[i+1] {
                possible = false
                break
            }
        }
        if possible {
            totalSum += k
        }
    }
    return totalSum
}