class Solution {
    minOperations(s) {
        const isSorted = str => {
            for (let i = 1; i < str.length; i++) {
                if (str[i] < str[i - 1]) return false;
            }
            return true;
        };
        
        if (isSorted(s)) return 0;
        
        const n = s.length;
        const visited = new Set([s]);
        const queue = [[s, 0]];
        let head = 0;
        
        while (head < queue.length) {
            const [curr, ops] = queue[head++];
            
            for (let i = 0; i < n; i++) {
                for (let j = i + 1; j <= n; j++) {
                    if (i === 0 && j === n) continue;
                    
                    const subArr = curr.substring(i, j).split('').sort();
                    const nextStr = curr.substring(0, i) + subArr.join('') + curr.substring(j);
                    
                    if (isSorted(nextStr)) return ops + 1;
                    
                    if (!visited.has(nextStr)) {
                        visited.add(nextStr);
                        queue.push([nextStr, ops + 1]);
                    }
                }
            }
        }
        
        return -1;
    }
}