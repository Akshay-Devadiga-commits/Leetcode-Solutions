#include <string>
#include <vector>
#include <algorithm>

class Solution {
public:
    int countDaysTogether(std::string arriveAlice, std::string leaveAlice, std::string arriveBob, std::string leaveBob) {
        std::vector<int> daysInMonths = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        
        auto dateToDayNumber = [&](const std::string& dateStr) -> int {
            int month = std::stoi(dateStr.substr(0, 2));
            int day = std::stoi(dateStr.substr(3, 2));
            int totalDays = day;
            for (int i = 0; i < month - 1; ++i) {
                totalDays += daysInMonths[i];
            }
            return totalDays;
        };
        
        int aliceArrive = dateToDayNumber(arriveAlice);
        int aliceLeave = dateToDayNumber(leaveAlice);
        int bobArrive = dateToDayNumber(arriveBob);
        int bobLeave = dateToDayNumber(leaveBob);
        
        int overlapArrive = std::max(aliceArrive, bobArrive);
        int overlapLeave = std::min(aliceLeave, bobLeave);
        
        if (overlapArrive <= overlapLeave) {
            return overlapLeave - overlapArrive + 1;
        }
        return 0;
    }
};