package graph

import (
	"fmt"
	"testing"
)

type TGraph struct {
	edges map[int][]int
	vertices []int
}

func NewTGraph(vertices []int) *TGraph {
	return &TGraph{
		edges: make(map[int][]int),
		vertices: vertices,
	}
}

func (g *TGraph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *TGraph) topologicalSortUtil(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true

	for _, u := range g.edges[v] {
		if !visited[u] {
			g.topologicalSortUtil(u, visited, stack)
		}
	}

	*stack = append([]int{v}, *stack...)
}

func (g *TGraph) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack)
		}
	}

	return stack
}

func TestTopologicalSorting(t *testing.T) {
	gp := NewTGraph([]int{5, 4, 2, 3, 1, 0})

	gp.addEdge(5, 2)
	gp.addEdge(5, 0)
	gp.addEdge(4, 0)
	gp.addEdge(4, 1)
	gp.addEdge(2, 3)
	gp.addEdge(3, 1)

	result := gp.topologicalSort()

	for _, v := range result {
		fmt.Println(v)
	}
}