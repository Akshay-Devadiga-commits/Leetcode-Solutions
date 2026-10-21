def solve(data):
    if len(data) <= 1:
        return data
    
    mid = len(data) // 2
    left_half = solve(data[:mid])
    right_half = solve(data[mid:])
    
    return merge(left_half, right_half)

def merge(left, right):
    sorted_list = []
    i = j = 0
    
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            sorted_list.append(left[i])
            i += 1
        else:
            sorted_list.append(right[j])
            j += 1
            
    sorted_list.extend(left[i:])
    sorted_list.extend(right[j:])
    return sorted_list