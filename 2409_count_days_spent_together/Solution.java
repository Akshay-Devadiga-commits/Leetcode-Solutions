class Solution {
    public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {
        int[] daysInMonths = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        
        int aliceArrive = dateToDayNumber(arriveAlice, daysInMonths);
        int aliceLeave = dateToDayNumber(leaveAlice, daysInMonths);
        int bobArrive = dateToDayNumber(arriveBob, daysInMonths);
        int bobLeave = dateToDayNumber(leaveBob, daysInMonths);
        
        int overlapArrive = Math.max(aliceArrive, bobArrive);
        int overlapLeave = Math.min(aliceLeave, bobLeave);
        
        if (overlapArrive <= overlapLeave) {
            return overlapLeave - overlapArrive + 1;
        }
        return 0;
    }
    
    private int dateToDayNumber(String dateStr, int[] daysInMonths) {
        int month = Integer.parseInt(dateStr.substring(0, 2));
        int day = Integer.parseInt(dateStr.substring(3, 5));
        int totalDays = day;
        for (int i = 0; i < month - 1; i++) {
            totalDays += daysInMonths[i];
        }
        return totalDays;
    }
}