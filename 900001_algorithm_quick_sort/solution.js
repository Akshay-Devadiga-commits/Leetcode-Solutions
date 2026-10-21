function solve(data) {
    if (!data || data.length <= 1) {
        return data;
    }
    const pivot = data[Math.floor(data.length / 2)];
    const left = data.filter(x => x < pivot);
    const middle = data.filter(x => x === pivot);
    const right = data.filter(x => x > pivot);
    return solve(left).concat(middle, solve(right));
}