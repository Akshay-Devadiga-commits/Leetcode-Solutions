#include <vector>
#include <string>
#include <queue>
#include <tuple>
#include <map>

using namespace std;

class Solution {
public:
    int shortestPath(int n, vector<vector<int>>& edges, string labels, int k) {
        vector<vector<pair<int, int>>> adj(n);
        for (const auto& edge : edges) {
            adj[edge[0]].push_back({edge[1], edge[2]});
        }

        using State = tuple<long long, int, char, int>;
        priority_queue<State, vector<State>, greater<State>> pq;
        pq.push({0, 0, labels[0], 1});

        map<tuple<int, char, int>, long long> dist;

        while (!pq.empty()) {
            auto [w, u, last_char, count] = pq.top();
            pq.pop();

            if (u == n - 1) {
                return w;
            }

            auto state = make_tuple(u, last_char, count);
            if (dist.count(state) && dist[state] < w) {
                continue;
            }
            dist[state] = w;

            for (const auto& [v, weight] : adj[u]) {
                char nxt_char = labels[v];
                int nxt_count = (nxt_char == last_char) ? count + 1 : 1;

                if (nxt_count <= k) {
                    long long nxt_w = w + weight;
                    auto nxt_state = make_tuple(v, nxt_char, nxt_count);
                    if (!dist.count(nxt_state) || nxt_w < dist[nxt_state]) {
                        dist[nxt_state] = nxt_w;
                        pq.push({nxt_w, v, nxt_char, nxt_count});
                    }
                }
            }
        }

        return -1;
    }
};