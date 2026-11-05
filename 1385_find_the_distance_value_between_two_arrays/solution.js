class Solution {
    findTheDistanceValue(arr1, arr2, d) {
        arr2.sort((a, b) => a - b);
        let distanceValue = 0;

        for (const num1 of arr1) {
            // We need to check if there is any num2 in arr2 such that |num1 - num2| <= d
            // This is equivalent to checking if there is any num2 such that num1 - d <= num2 <= num1 + d

            // Find the index of the first element in arr2 that is >= num1 - d
            let low = 0;
            let high = arr2.length;
            let leftBoundIndex = -1;

            while (low < high) {
                const mid = Math.floor((low + high) / 2);
                if (arr2[mid] >= num1 - d) {
                    leftBoundIndex = mid;
                    high = mid;
                } else {
                    low = mid + 1;
                }
            }

            // Check if this index is valid and if the element at this index satisfies the condition
            // The element arr2[leftBoundIndex] is the smallest element >= num1 - d.
            // We need to check if this element is also <= num1 + d.
            let foundCloseElement = false;
            if (leftBoundIndex !== -1 && arr2[leftBoundIndex] <= num1 + d) {
                foundCloseElement = true;
            }

            // If no element in arr2 is found within the range [num1 - d, num1 + d],
            // then num1 contributes to the distance value.
            if (!foundCloseElement) {
                distanceValue++;
            }
        }

        return distanceValue;
    }
}