package graphs

import (
	"fmt"
)

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

type Graph struct {
	vertices int
	adjList  map[int][]int
}

func (g *Graph) NewGraph(numVertices int) *Graph {
	g.vertices = numVertices
	g.adjList = make(map[int][]int)
	return g
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}

func (g *Graph) PrintAdjList() {
	fmt.Print(colorMagenta)
	for vertex, neighbours := range g.adjList {
		fmt.Printf("\tVertex: %d, \t Neighbours: %v\n", vertex, neighbours)
	}
	fmt.Print(colorReset)
}

// Global Visited Array
var visited []bool

// The DFSHelper function performs a depth-first search traversal on a graph starting from a given vertex, marking visited vertices
func DFSHelper(g *Graph, v int, visited []bool) {
	visited[v] = true
	fmt.Printf(colorMagenta+"\t%d", v)
	fmt.Print(colorReset)
	neighbours := g.adjList[v]

	// Recur for all the vertices adjacent to this vertex
	for _, neighbour := range neighbours {
		if !visited[neighbour] {
			DFSHelper(g, neighbour, visited)
		}
	}
}

// Graph Traversals
// The `DFS` function is a method of the `Graph` struct. It performs a Depth First Search traversal starting from a given vertex `v`.
func (g *Graph) DFS(v int) {
	visited = make([]bool, g.vertices)
	DFSHelper(g, v, visited)
}

func (g *Graph) BFS(v int) {
	visited = make([]bool, g.vertices)
	queue := make([]int, 0)

	queue = append(queue, v)
	for len(queue) > 0 {
		visited[queue[0]] = true
		fmt.Printf(colorMagenta+"\t%d", queue[0])
		neighbours := g.adjList[queue[0]]
		for _, neighbour := range neighbours {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
		queue = queue[1:]
	}
	fmt.Print(colorReset)
}
