#include <vector>
#include <string>
#include <queue>
#include <tuple>
#include <map>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int shortestPath(int n, vector<vector<int>>& edges, string labels, int k) {
        vector<vector<pair<int, int>>> adj(n);
        for (const auto& edge : edges) {
            adj[edge[0]].push_back({edge[1], edge[2]});
        }

        char init_char = labels[0];
        if (1 > k) return -1;

        using State = tuple<int, int, char, int>;
        priority_queue<State, vector<State>, greater<State>> pq;
        pq.push({0, 0, init_char, 1});

        map<tuple<int, char, int>, int> dist;
        dist[{0, init_char, 1}] = 0;

        while (!pq.empty()) {
            auto [d, u, last_char, count] = pq.top();
            pq.pop();

            if (u == n - 1) return d;

            auto state_key = make_tuple(u, last_char, count);
            if (dist.count(state_key) && d > dist[state_key]) continue;

            for (const auto& neighbor : adj[u]) {
                int v = neighbor.first;
                int weight = neighbor.second;
                char v_char = labels[v];
                int new_count = (v_char == last_char) ? count + 1 : 1;

                if (new_count <= k) {
                    int new_dist = d + weight;
                    auto next_state = make_tuple(v, v_char, new_count);
                    if (!dist.count(next_state) || new_dist < dist[next_state]) {
                        dist[next_state] = new_dist;
                        pq.push({new_dist, v, v_char, new_count});
                    }
                }
            }
        }

        return -1;
    }
};