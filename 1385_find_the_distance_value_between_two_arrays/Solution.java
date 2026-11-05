import java.util.Arrays;

class Solution {
    public int findTheDistanceValue(int[] arr1, int[] arr2, int d) {
        Arrays.sort(arr2);
        int distanceValue = 0;

        for (int num1 : arr1) {
            // We need to check if there is any num2 in arr2 such that |num1 - num2| <= d
            // This is equivalent to checking if there is any num2 such that num1 - d <= num2 <= num1 + d

            // Find the index of the first element in arr2 that is >= num1 - d
            int leftBoundIndex = binarySearchLowerBound(arr2, num1 - d);

            // Check if this index is valid and if the element at this index satisfies the condition
            // The element arr2[leftBoundIndex] is the smallest element >= num1 - d.
            // We need to check if this element is also <= num1 + d.
            boolean foundCloseElement = false;
            if (leftBoundIndex < arr2.length) {
                if (arr2[leftBoundIndex] <= num1 + d) {
                    foundCloseElement = true;
                }
            }

            // If no element in arr2 is found within the range [num1 - d, num1 + d],
            // then num1 contributes to the distance value.
            if (!foundCloseElement) {
                distanceValue++;
            }
        }

        return distanceValue;
    }

    // Helper function for binary search to find the lower bound (insertion point)
    private int binarySearchLowerBound(int[] arr, int target) {
        int low = 0;
        int high = arr.length;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] < target) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        return low;
    }
}