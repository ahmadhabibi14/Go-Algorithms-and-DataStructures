package graph

import (
	"fmt"
	"testing"
)


type Graph struct {
	verticles	int
	adjList		map[int][]int	
}

func NewGraph(verticles int) *Graph {
	return &Graph{
		verticles: verticles,
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(source, dest int) {
	g.adjList[source] = append(g.adjList[source], dest)
	g.adjList[dest] = append(g.adjList[dest], source)
}

func (g *Graph) DFSUtil(vertex int, visited map[int]bool) {
	visited[vertex] = true
	fmt.Printf("%d", vertex)

	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			g.DFSUtil(v, visited)
		}
	}
}

func (g *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)
	g.DFSUtil(startVertex, visited)
}

func TestDepthFirstSearch(t *testing.T) {
	graph := NewGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)

	fmt.Println("Depth-first traversal starting from vertex 0:")
	graph.DFS(0)
}