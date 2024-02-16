package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices int
	Edges    map[int]map[int]int
}

func (g *Graph) addEdge(u, v, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[int]map[int]int)
	}
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[int]int)
	}
	g.Edges[u][v] = weight
}

func dijkstra(graph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	for i := 0; i < graph.Vertices; i++ {
		distances[i] = math.MaxInt32
	}

	distances[start] = 0

	for range graph.Edges {
		u := minDistance(distances, visited)
		visited[u] = true

		for v, weight := range graph.Edges[u] {
			if !visited[v] && distances[u] != math.MaxInt32 && distances[u]+weight < distances[v] {
				distances[v] = distances[u] + weight
			}
		}
	}

	return distances
}

func minDistance(distances map[int]int, visited map[int]bool) int {
	min := math.MaxInt32
	var minIndex int

	for v, dist := range distances {
		if !visited[v] && dist <= min {
			min = dist
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	var graph Graph
	var u, v, weight int

	fmt.Scan(&graph.Vertices)

	for i := 0; i < graph.Vertices; i++ {
		fmt.Scan(&u, &v, &weight)
		graph.addEdge(u, v, weight)
	}

	startNode := 0
	distances := dijkstra(&graph, startNode)

	for _, distance := range distances {
		fmt.Printf("%d\n", distance)
	}
}
