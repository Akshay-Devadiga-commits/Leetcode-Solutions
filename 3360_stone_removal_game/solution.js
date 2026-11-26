var canAliceWin = function (n) {
    const k = Math.floor((21 - Math.sqrt(441 - 8 * n)) / 2);
    return k % 2 === 1;
};