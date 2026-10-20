func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if idx, ok := seen[target-num]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}