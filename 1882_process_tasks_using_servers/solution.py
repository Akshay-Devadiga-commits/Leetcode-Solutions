import heapq

class Solution:
    def assignTasks(self, servers: list[int], tasks: list[int]) -> list[int]:
        n = len(servers)
        m = len(tasks)
        ans = [0] * m

        # Priority queue for available servers: (weight, index)
        # Initially, all servers are available.
        available_servers = []
        for i in range(n):
            heapq.heappush(available_servers, (servers[i], i))

        # Priority queue for busy servers: (free_time, weight, index)
        # This helps us know when servers become free.
        busy_servers = []

        current_time = 0
        task_idx = 0

        while task_idx < m:
            # 1. Free up servers that have finished their tasks by current_time
            while busy_servers and busy_servers[0][0] <= current_time:
                free_time, weight, index = heapq.heappop(busy_servers)
                heapq.heappush(available_servers, (weight, index))

            # 2. Assign tasks if there are available servers
            if available_servers:
                weight, index = heapq.heappop(available_servers)
                ans[task_idx] = index
                
                # Calculate when this server will be free again
                # The task starts processing at current_time
                free_time = current_time + tasks[task_idx]
                heapq.heappush(busy_servers, (free_time, weight, index))
                
                task_idx += 1
            else:
                # 3. If no servers are available, advance time to the next server free-up event.
                # This ensures we don't get stuck if all servers are busy.
                # We must have at least one busy server if available_servers is empty and task_idx < m.
                current_time = busy_servers[0][0]

            # 4. Ensure current_time is at least the arrival time of the next task.
            # The j-th task arrives at second j. If we assigned a task at time t, and the next task arrives at t+1,
            # we need to make sure current_time advances to at least t+1 if it hasn't already.
            # This is crucial for tasks arriving at specific times.
            # If we just assigned a task, current_time might be the same as the previous iteration.
            # We need to consider the arrival time of the *next* task (which is task_idx).
            # If current_time is less than task_idx, it means we are waiting for tasks to arrive.
            # However, the problem states tasks are inserted at second j. So, the next task is available at time task_idx.
            # If we just assigned a task at current_time, and current_time < task_idx, we should advance current_time to task_idx.
            # If current_time >= task_idx, it means the task has already arrived or is arriving now.
            # The logic of advancing time in step 3 handles the case where we wait for servers.
            # If we assigned a task, and current_time is still less than the arrival time of the next task (task_idx), we should advance.
            # The problem states task j is inserted at second j. So, task_idx is the arrival time of the next task.
            # If current_time < task_idx, we must advance current_time to task_idx to consider the arrival of the next task.
            # This is implicitly handled by the fact that if we assign a task, task_idx increments. The next iteration will check against the new task_idx.
            # The critical part is when we advance time in step 3. If we advance to busy_servers[0][0], and that time is less than task_idx,
            # we still need to wait for task_idx to arrive. So, current_time should be max(current_time, task_idx).
            # Let's re-evaluate: if we assign a task, current_time doesn't change in this iteration. The next iteration will check free servers.
            # If available_servers is empty, we advance current_time to busy_servers[0][0]. This is correct.
            # If available_servers is NOT empty, we assign a task. The current_time remains the same for this iteration.
            # The next task to consider is task_idx. Its arrival time is task_idx.
            # So, for the *next* iteration, current_time should be at least task_idx.
            # This means after assigning a task, we should update current_time = max(current_time, task_idx).
            # However, this is only if we assigned a task. If we advanced time in step 3, current_time is already set correctly.
            # Let's simplify: The simulation proceeds second by second implicitly. When we assign a task, it's at `current_time`. The next task is `task_idx`. 
            # If `current_time` is less than `task_idx`, it means the next task hasn't arrived yet. We should wait until `task_idx`.
            # So, if we assigned a task, and `current_time < task_idx`, we should advance `current_time` to `task_idx`.
            # This is equivalent to `current_time = max(current_time, task_idx)` IF a task was assigned.
            # If no task was assigned (because available_servers was empty), we advanced current_time to busy_servers[0][0].
            # In that case, we still need to ensure current_time is at least task_idx.
            # So, a general update `current_time = max(current_time, task_idx)` after step 2 (if task_idx < m) seems correct.
            # Let's test this: If current_time = 0, task_idx = 0. Assign task 0. current_time remains 0. Next iteration, task_idx = 1. current_time = max(0, 1) = 1.
            # This seems to correctly advance time to the arrival of the next task if needed.
            # However, if we advanced time in step 3 to busy_servers[0][0], and busy_servers[0][0] < task_idx, we still need to wait for task_idx.
            # So, the update should be: if task_idx < m, current_time = max(current_time, task_idx).
            # This should happen *after* we potentially advance time in step 3.
            # Let's reconsider the loop structure.
            # The loop should continue as long as there are tasks to assign OR servers that will become free.
            # The current loop `while task_idx < m` is correct for assigning all tasks.
            # The issue is how `current_time` advances.
            # If `available_servers` is NOT empty, we assign a task. `current_time` remains the same for this iteration.
            # The next task is `task_idx`. Its arrival time is `task_idx`.
            # So, for the *next* iteration, `current_time` must be at least `task_idx`.
            # If `current_time` is already `>= task_idx`, it's fine.
            # If `current_time < task_idx`, we must advance `current_time` to `task_idx`.
            # This means `current_time = max(current_time, task_idx)` should be applied if a task was assigned.
            # If `available_servers` IS empty, we advanced `current_time` to `busy_servers[0][0]`. This is the next event time.
            # After this, we still need to ensure `current_time` is at least `task_idx`.
            # So, `current_time = max(current_time, task_idx)` should be applied at the end of the loop if `task_idx < m`.
            # This ensures that if we are waiting for servers, and the next server free-up time is before the next task arrives, we wait for the task.
            # If we assigned a task, and the next task arrives later, we advance time to the next task arrival.
            if task_idx < m:
                current_time = max(current_time, task_idx)

        return ans
