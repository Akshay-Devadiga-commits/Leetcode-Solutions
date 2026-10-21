function solve(data) {
    if (data.length <= 1) {
        return data;
    }
    
    const mid = Math.floor(data.length / 2);
    const leftHalf = solve(data.slice(0, mid));
    const rightHalf = solve(data.slice(mid));
    
    return merge(leftHalf, rightHalf);
}

function merge(left, right) {
    const sortedList = [];
    let i = 0;
    let j = 0;
    
    while (i < left.length && j < right.length) {
        if (left[i] < right[j]) {
            sortedList.push(left[i]);
            i++;
        } else {
            sortedList.push(right[j]);
            j++;
        }
    }
    
    return sortedList.concat(left.slice(i), right.slice(j));
}