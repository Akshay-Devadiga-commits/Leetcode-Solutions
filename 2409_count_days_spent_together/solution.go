import "strconv"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    daysInMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    
    dateToDayNumber := func(dateStr string) int {
        month, _ := strconv.Atoi(dateStr[:2])
        day, _ := strconv.Atoi(dateStr[3:])
        totalDays := day
        for i := 0; i < month-1; i++ {
            totalDays += daysInMonths[i]
        }
        return totalDays
    }
    
    aliceArrive := dateToDayNumber(arriveAlice)
    aliceLeave := dateToDayNumber(leaveAlice)
    bobArrive := dateToDayNumber(arriveBob)
    bobLeave := dateToDayNumber(leaveBob)
    
    overlapArrive := max(aliceArrive, bobArrive)
    overlapLeave := min(aliceLeave, bobLeave)
    
    if overlapArrive <= overlapLeave {
        return overlapLeave - overlapArrive + 1
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}