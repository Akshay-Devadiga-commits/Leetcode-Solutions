class Solution {
public:
    std::string digitSum(std::string s, int k) {
        while (s.length() > k) {
            std::string next_s = "";
            for (int i = 0; i < s.length(); i += k) {
                int current_sum = 0;
                for (int j = i; j < std::min(s.length(), i + k); ++j) {
                    current_sum += s[j] - '0';
                }
                next_s += std::to_string(current_sum);
            }
            s = next_s;
        }
        return s;
    }
};