var maxSelectedElements = function(nums) {
    nums.sort((a, b) => a - b);
    const dp = new Map();
    let maxLength = 0;
    
    for (const x of nums) {
        const dpX = dp.get(x) || 0;
        const dpXMinus1 = dp.get(x - 1) || 0;
        
        const newDpXPlus1 = dpX + 1;
        const newDpX = dpXMinus1 + 1;
        
        dp.set(x + 1, newDpXPlus1);
        dp.set(x, newDpX);
        
        maxLength = Math.max(maxLength, newDpXPlus1, newDpX);
    }
    
    return maxLength;
};