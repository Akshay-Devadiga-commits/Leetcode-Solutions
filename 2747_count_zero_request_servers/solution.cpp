#include <vector>
#include <algorithm>

struct IndexedQuery {
    int queryIndex;
    int query;

    bool operator<(const IndexedQuery& other) const {
        return query < other.query;
    }
};

class Solution {
public:
    std::vector<int> countServers(int n, std::vector<std::vector<int>>& logs, int x, std::vector<int>& queries) {
        std::vector<int> ans(queries.size());
        std::vector<int> count(n + 1, 0);

        std::sort(logs.begin(), logs.end(), [](const std::vector<int>& a, const std::vector<int>& b) {
            return a[1] < b[1];
        });

        int i = 0;
        int j = 0;
        int servers = 0;

        std::vector<IndexedQuery> indexedQueries(queries.size());
        for (int k = 0; k < queries.size(); ++k) {
            indexedQueries[k] = {k, queries[k]};
        }
        std::sort(indexedQueries.begin(), indexedQueries.end());

        for (const auto& indexedQuery : indexedQueries) {
            int queryIndex = indexedQuery.queryIndex;
            int query = indexedQuery.query;

            while (j < logs.size() && logs[j][1] <= query) {
                count[logs[j][0]]++;
                if (count[logs[j][0]] == 1) {
                    servers++;
                }
                j++;
            }
            while (i < logs.size() && logs[i][1] < query - x) {
                count[logs[i][0]]--;
                if (count[logs[i][0]] == 0) {
                    servers--;
                }
                i++;
            }
            ans[queryIndex] = n - servers;
        }

        return ans;
    }
};