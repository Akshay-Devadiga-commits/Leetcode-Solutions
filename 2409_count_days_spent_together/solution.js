/**
 * @param {string} arriveAlice
 * @param {string} leaveAlice
 * @param {string} arriveBob
 * @param {string} leaveBob
 * @return {number}
 */
var countDaysTogether = function(arriveAlice, leaveAlice, arriveBob, leaveBob) {
    const daysInMonths = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    
    const dateToDayNumber = (dateStr) => {
        const month = parseInt(dateStr.substring(0, 2), 10);
        const day = parseInt(dateStr.substring(3), 10);
        let totalDays = day;
        for (let i = 0; i < month - 1; i++) {
            totalDays += daysInMonths[i];
        }
        return totalDays;
    };
    
    const aliceArrive = dateToDayNumber(arriveAlice);
    const aliceLeave = dateToDayNumber(leaveAlice);
    const bobArrive = dateToDayNumber(arriveBob);
    const bobLeave = dateToDayNumber(leaveBob);
    
    const overlapArrive = Math.max(aliceArrive, bobArrive);
    const overlapLeave = Math.min(aliceLeave, bobLeave);
    
    if (overlapArrive <= overlapLeave) {
        return overlapLeave - overlapArrive + 1;
    }
    return 0;
};