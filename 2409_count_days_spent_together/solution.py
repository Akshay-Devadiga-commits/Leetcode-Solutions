class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        days_in_months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        
        def date_to_day_number(date_str: str) -> int:
            month = int(date_str[:2])
            day = int(date_str[3:])
            total_days = day
            for i in range(month - 1):
                total_days += days_in_months[i]
            return total_days
        
        alice_arrive = date_to_day_number(arriveAlice)
        alice_leave = date_to_day_number(leaveAlice)
        bob_arrive = date_to_day_number(arriveBob)
        bob_leave = date_to_day_number(leaveBob)
        
        overlap_arrive = max(alice_arrive, bob_arrive)
        overlap_leave = min(alice_leave, bob_leave)
        
        if overlap_arrive <= overlap_leave:
            return overlap_leave - overlap_arrive + 1
        return 0