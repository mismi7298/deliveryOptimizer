package router

import (
	"deliveryOptimzer/pkg/model"
)

func allTopologicalSortsUtil(graph map[model.Location][]model.Location, inDegree map[model.Location]int, result []model.Location, allResults *[][]model.Location) {
	flag := false

	for node := range inDegree {
		if inDegree[node] == 0 {
			flag = true
			result = append(result, node)

			for _, neighbor := range graph[node] {
				inDegree[neighbor]--
			}

			inDegree[node] = -1 // Mark as visited
			allTopologicalSortsUtil(graph, inDegree, result, allResults)

			// Backtrack
			inDegree[node] = 0
			result = result[:len(result)-1]
			for _, neighbor := range graph[node] {
				inDegree[neighbor]++
			}
		}
	}

	if !flag {
		tempResult := make([]model.Location, len(result))
		copy(tempResult, result)
		*allResults = append(*allResults, tempResult)
	}
}

func findAllOrders(de model.DeliveryExecutive, orders []model.Order) [][]model.Location {
	graph := make(map[model.Location][]model.Location)
	inDegree := make(map[model.Location]int)

	// Initialize the graph and in-degree map
	for _, order := range orders {
		u, v := order.Restaurant.Location, order.Customer.Location
		graph[u] = append(graph[u], v)
		inDegree[v]++
		if _, exists := inDegree[u]; !exists {
			inDegree[u] = 0
		}
	}

	var result = []model.Location{de.Location}
	var allResults [][]model.Location
	allTopologicalSortsUtil(graph, inDegree, result, &allResults)

	return allResults
}
