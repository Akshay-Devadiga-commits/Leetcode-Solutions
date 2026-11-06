class Solution {
    countServers(n, logs, x, queries) {
        const ans = new Array(queries.length).fill(0);
        const count = new Array(n + 1).fill(0);

        logs.sort((a, b) => a[1] - b[1]);

        let i = 0;
        let j = 0;
        let servers = 0;

        const indexedQueries = queries.map((query, index) => ({ queryIndex: index, query: query })).sort((a, b) => a.query - b.query);

        for (const { queryIndex, query } of indexedQueries) {
            while (j < logs.length && logs[j][1] <= query) {
                count[logs[j][0]]++;
                if (count[logs[j][0]] === 1) {
                    servers++;
                }
                j++;
            }
            while (i < logs.length && logs[i][1] < query - x) {
                count[logs[i][0]]--;
                if (count[logs[i][0]] === 0) {
                    servers--;
                }
                i++;
            }
            ans[queryIndex] = n - servers;
        }

        return ans;
    }
}