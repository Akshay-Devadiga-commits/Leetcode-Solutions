package main

import (
	"sort"
)

type IndexedQuery struct {
	queryIndex int
	query      int
}

func countServers(n int, logs [][]int, x int, queries []int) []int {
	ans := make([]int, len(queries))
	count := make([]int, n+1)

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	i := 0
	j := 0
	servers := 0

	indexedQueries := make([]IndexedQuery, len(queries))
	for k, query := range queries {
		indexedQueries[k] = IndexedQuery{queryIndex: k, query: query}
	}
	sort.Slice(indexedQueries, func(i, j int) bool {
		return indexedQueries[i].query < indexedQueries[j].query
	})

	for _, indexedQuery := range indexedQueries {
		queryIndex := indexedQuery.queryIndex
		query := indexedQuery.query

		for j < len(logs) && logs[j][1] <= query {
			count[logs[j][0]]++
			if (count[logs[j][0]] == 1) {
				servers++
			}
			j++
		}
		for i < len(logs) && logs[i][1] < query-x {
			count[logs[i][0]]--
			if (count[logs[i][0]] == 0) {
				servers--
			}
			i++
		}
		ans[queryIndex] = n - servers
	}

	return ans
}
