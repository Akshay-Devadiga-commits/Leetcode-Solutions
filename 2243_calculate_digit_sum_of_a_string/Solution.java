class Solution {
    public String digitSum(String s, int k) {
        while (s.length() > k) {
            StringBuilder next_s = new StringBuilder();
            for (int i = 0; i < s.length(); i += k) {
                int current_sum = 0;
                for (int j = i; j < Math.min(s.length(), i + k); j++) {
                    current_sum += s.charAt(j) - '0';
                }
                next_s.append(current_sum);
            }
            s = next_s.toString();
        }
        return s;
    }
}