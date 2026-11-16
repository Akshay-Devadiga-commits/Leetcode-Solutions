/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var digitSum = function(s, k) {
    while (s.length > k) {
        let next_s = '';
        for (let i = 0; i < s.length; i += k) {
            let current_sum = 0;
            for (let j = i; j < Math.min(s.length, i + k); j++) {
                current_sum += parseInt(s[j]);
            }
            next_s += current_sum;
        }
        s = next_s;
    }
    return s;
};