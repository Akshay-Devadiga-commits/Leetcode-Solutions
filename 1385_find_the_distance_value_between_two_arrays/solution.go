import (
    "sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    sort.Ints(arr2)
    distanceValue := 0

    for _, num1 := range arr1 {
        // We need to check if there is any num2 in arr2 such that |num1 - num2| <= d
        // This is equivalent to checking if there is any num2 such that num1 - d <= num2 <= num1 + d

        // Find the index of the first element in arr2 that is >= num1 - d
        // sort.SearchInts returns the index where target would be inserted to maintain order.
        // If target is present, it returns the index of the first occurrence.
        leftBoundIndex := sort.SearchInts(arr2, num1-d)

        // Check if this index is valid and if the element at this index satisfies the condition
        // The element arr2[leftBoundIndex] is the smallest element >= num1 - d.
        // We need to check if this element is also <= num1 + d.
        foundCloseElement := false
        if leftBoundIndex < len(arr2) {
            if arr2[leftBoundIndex] <= num1+d {
                foundCloseElement = true
            }
        }

        // If no element in arr2 is found within the range [num1 - d, num1 + d],
        // then num1 contributes to the distance value.
        if !foundCloseElement {
            distanceValue++
        }
    }

    return distanceValue
}