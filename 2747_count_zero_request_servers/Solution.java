import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

class Solution {
    static class IndexedQuery {
        int queryIndex;
        int query;

        IndexedQuery(int queryIndex, int query) {
            this.queryIndex = queryIndex;
            this.query = query;
        }
    }

    public int[] countServers(int n, int[][] logs, int x, int[] queries) {
        int[] ans = new int[queries.length];
        int[] count = new int[n + 1];

        Arrays.sort(logs, Comparator.comparingInt(a -> a[1]));

        int i = 0;
        int j = 0;
        int servers = 0;

        List<IndexedQuery> indexedQueries = new ArrayList<>();
        for (int k = 0; k < queries.length; k++) {
            indexedQueries.add(new IndexedQuery(k, queries[k]));
        }
        indexedQueries.sort(Comparator.comparingInt(a -> a.query));

        for (IndexedQuery indexedQuery : indexedQueries) {
            int queryIndex = indexedQuery.queryIndex;
            int query = indexedQuery.query;

            while (j < logs.length && logs[j][1] <= query) {
                count[logs[j][0]]++;
                if (count[logs[j][0]] == 1) {
                    servers++;
                }
                j++;
            }
            while (i < logs.length && logs[i][1] < query - x) {
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
}