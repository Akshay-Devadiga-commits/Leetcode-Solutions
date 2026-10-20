var twoSum = function(nums, target) {
    const seen = new Map();
    for (let i = 0; i < nums.length; i++) {
        const comp = target - nums[i];
        if (seen.has(comp)) return [seen.get(comp), i];
        seen.set(nums[i], i);
    }
    return [];
};