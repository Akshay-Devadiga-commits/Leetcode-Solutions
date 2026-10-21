package main

func solve(data []int) []int {
    if len(data) <= 1 {
        return data
    }

    mid := len(data) / 2
    leftHalf := solve(data[:mid])
    rightHalf := solve(data[mid:])

    return merge(leftHalf, rightHalf)
}

func merge(left, right []int) []int {
    sortedList := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            sortedList = append(sortedList, left[i])
            i++
        } else {
            sortedList = append(sortedList, right[j])
            j++
        }
    }

    sortedList = append(sortedList, left[i:]...)
    sortedList = append(sortedList, right[j:]...)
    return sortedList
}