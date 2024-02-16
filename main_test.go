package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := Graph{
		Vertices: 5,
		Edges: map[int]map[int]int{
			0: {1: 2, 2: 4},
			1: {2: 1, 3: 7},
			2: {4: 3},
			3: {4: 1},
			4: {},
		},
	}

	startNode := 0
	distances := dijkstra(&graph, startNode)

	expectedDistances := map[int]int{
		0: 0,
		1: 2,
		2: 3,
		3: 9,
		4: 6,
	}

	if !reflect.DeepEqual(distances, expectedDistances) {
		t.Errorf("Expected distances: %v, but got: %v", expectedDistances, distances)
	}
}
