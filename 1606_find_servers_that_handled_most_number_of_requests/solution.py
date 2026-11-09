import heapq

class Solution:
    def busiestServers(self, k: int, arrival: list[int], load: list[int]) -> list[int]:
        server_request_counts = [0] * k
        
        # Min-heap to store available servers (server_id)
        # Initially, all servers are available
        available_servers = list(range(k))
        heapq.heapify(available_servers)
        
        # Min-heap to store busy servers (finish_time, server_id)
        busy_servers = []
        
        for i in range(len(arrival)):
            current_time = arrival[i]
            request_load = load[i]
            
            # 1. Free up servers that have finished their tasks by current_time
            while busy_servers and busy_servers[0][0] <= current_time:
                finish_time, server_id = heapq.heappop(busy_servers)
                heapq.heappush(available_servers, server_id)
            
            # 2. Determine the target server index
            target_server_idx = i % k
            
            # 3. If no servers are available, drop the request
            if not available_servers:
                continue
            
            # 4. Find the smallest available server >= target_server_idx
            # We need to temporarily store servers that are smaller than target_server_idx
            # and re-add them if they are not used.
            temp_servers = []
            assigned_server_id = -1
            
            while available_servers:
                server_id = heapq.heappop(available_servers)
                if server_id >= target_server_idx:
                    assigned_server_id = server_id
                    break
                else:
                    temp_servers.append(server_id)
            
            # Re-add the temporarily stored servers back to available_servers
            for s_id in temp_servers:
                heapq.heappush(available_servers, s_id)
            
            # 5. If no server >= target_server_idx was found, wrap around
            if assigned_server_id == -1:
                # If available_servers is not empty, pick the smallest one (top of heap)
                if available_servers:
                    assigned_server_id = heapq.heappop(available_servers)
                else:
                    # This case should ideally not happen if we checked 'if not available_servers' earlier
                    # but as a safeguard, if somehow it's empty here, drop the request.
                    continue
            
            # 6. Assign the request to the found server
            server_request_counts[assigned_server_id] += 1
            finish_time = current_time + request_load
            heapq.heappush(busy_servers, (finish_time, assigned_server_id))
        
        # Find the maximum request count
        max_requests = 0
        if server_request_counts:
            max_requests = max(server_request_counts)
        
        # Find all servers with the maximum request count
        busiest_servers_list = []
        for server_id, count in enumerate(server_request_counts):
            if count == max_requests:
                busiest_servers_list.append(server_id)
        
        return busiest_servers_list
