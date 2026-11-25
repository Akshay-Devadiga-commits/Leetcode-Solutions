func smallestUniqueSubarray(nums []int) int {
    n := len(nums)
    
    for length := 1; length <= n; length++ {
        subarrayCounts := make(map[string]int)
        
        for i := 0; i <= n-length; i++ {
            subarray := nums[i : i+length]
            key := make([]byte, 0, length*4)
            for _, num := range subarray {
                key = append(key, byte(num))
            }
            subarrayCounts[string(key)]++
        }
        
        for _, count := range subarrayCounts {
            if count == 1 {
                return length
            }
        }
    }
    
    return n
}