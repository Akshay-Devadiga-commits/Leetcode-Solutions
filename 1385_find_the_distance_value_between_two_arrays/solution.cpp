#include <vector>
#include <algorithm>

class Solution {
public:
    int findTheDistanceValue(std::vector<int>& arr1, std::vector<int>& arr2, int d) {
        std::sort(arr2.begin(), arr2.end());
        int distance_value = 0;

        for (int num1 : arr1) {
            // We need to check if there is any num2 in arr2 such that |num1 - num2| <= d
            // This is equivalent to checking if there is any num2 such that num1 - d <= num2 <= num1 + d

            // Find the iterator to the first element in arr2 that is >= num1 - d
            auto it = std::lower_bound(arr2.begin(), arr2.end(), num1 - d);

            // Check if this iterator is valid and if the element it points to satisfies the condition
            // The element *it is the smallest element >= num1 - d.
            // We need to check if this element is also <= num1 + d.
            bool found_close_element = false;
            if (it != arr2.end()) {
                if (*it <= num1 + d) {
                    found_close_element = true;
                }
            }

            // If no element in arr2 is found within the range [num1 - d, num1 + d],
            // then num1 contributes to the distance value.
            if (!found_close_element) {
                distance_value++;
            }
        }

        return distance_value;
    }
};