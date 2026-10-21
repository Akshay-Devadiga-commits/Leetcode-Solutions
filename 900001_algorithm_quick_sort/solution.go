func solve(data []int) []int {
    if len(data) <= 1 {
        return data
    }
    pivot := data[len(data)/2]
    var left, middle, right []int
    for _, x := range data {
        if x < pivot {
            left = append(left, x)
        } else if x == pivot {
            middle = append(middle, x)
        } else {
            right = append(right, x)
        }
    }
    res := append([]int{}, solve(left)...)
    res = append(res, middle...)
    res = append(res, solve(right)...)
    return res
}