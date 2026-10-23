import "sort"

func sortPermutation(nums []int) int {
    n := len(nums)
    sorted := true
    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            sorted = false
            break
        }
    }
    if sorted {
        return 0
    }

    maxVal := 0
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }

    for k := maxVal; k >= 0; k-- {
        parent := make([]int, n)
        for i := 0; i < n; i++ {
            parent[i] = i
        }

        var find func(int) int
        find = func(i int) int {
            if parent[i] == i {
                return i
            }
            parent[i] = find(parent[i])
            return parent[i]
        }

        union := func(i, j int) {
            rootI := find(i)
            rootJ := find(j)
            if rootI != rootJ {
                parent[rootI] = rootJ
            }
        }

        for i := 0; i < n; i++ {
            for j := i + 1; j < n; j++ {
                if (nums[i] & nums[j]) == k {
                    union(i, j)
                }
            }
        }

        components := make(map[int][]int)
        for i := 0; i < n; i++ {
            root := find(i)
            components[root] = append(components[root], i)
        }

        possible := true
        for _, indices := range components {
            currentValues := make([]int, len(indices))
            for idx, i := range indices {
                currentValues[idx] = nums[i]
            }
            targetValues := make([]int, len(indices))
            copy(targetValues, indices)

            sort.Ints(currentValues)
            sort.Ints(targetValues)

            match := true
            for i := range currentValues {
                if currentValues[i] != targetValues[i] {
                    match = false
                    break
                }
            }
            if !match {
                possible = false
                break
            }
        }

        if possible {
            return k
        }
    }

    return 0
}