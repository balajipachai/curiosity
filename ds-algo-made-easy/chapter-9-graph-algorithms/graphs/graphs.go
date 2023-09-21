package graphs

import (
	"fmt"

	"example.com/heaps"
)

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// The Graph type represents a graph with a given number of vertices and an adjacency list.
// @property {int} vertices - The "vertices" property represents the number of vertices (or nodes) in
// the graph. It indicates the total number of nodes present in the graph.
// @property adjList - The `adjList` property is a map that represents the adjacency list of the graph.
// It maps each vertex to a list of its adjacent vertices.
type Graph struct {
	vertices int
	adjList  map[int][]int
	weight   [][]int
}

// The `NewGraph` function is a method of the `Graph` struct. It is used to create a new instance of
// the `Graph` struct with the specified number of vertices.
func (g *Graph) NewGraph(numVertices int) *Graph {
	g.vertices = numVertices
	g.adjList = make(map[int][]int)
	g.weight = make([][]int, numVertices)
	return g
}

// The `AddEdge` function is a method of the `Graph` struct. It is used to add an undirected edge
// between two vertices in the graph.
func (g *Graph) AddEdge(v1, v2 int) {
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}

func (g *Graph) AddEdgeWithWeight(v1, v2, weight int) {
	g.AddEdge(v1, v2)
	g.weight[v1]
}

// The `AddDirectedEdge` function is a method of the `Graph` struct. It is used to add a directed edge
// from vertex `v1` to vertex `v2` in the graph.
func (g *Graph) AddDirectedEdge(v1, v2 int) {
	g.adjList[v1] = append(g.adjList[v1], v2)
}

// The `PrintAdjList` function is a method of the `Graph` struct. It is used to print the adjacency
// list of the graph.
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

// The `BFS` function is a method of the `Graph` struct. It performs a Breadth First Search traversal
// starting from a given vertex `v`.
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

var stack = make([]int, 0)

// The function performs a depth-first search on a graph and appends the visited vertices to a stack in topological order.
func TopologicalSortHelper(g *Graph, v int, visited []bool) {
	visited[v] = true
	neighbours := g.adjList[v]

	for _, neighbour := range neighbours {
		if !visited[neighbour] {
			TopologicalSortHelper(g, neighbour, visited)
		}
	}
	stack = append(stack, v)
}

// The function performs a topological sort on the graph.
func (g *Graph) TopologicalSort() {
	visited = make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			TopologicalSortHelper(g, i, visited)
		}
	}

	fmt.Println(colorMagenta)
	// Using only array as a stack
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("\t%d", stack[i])
	}
	fmt.Println(colorReset)
}

// The `UnweightedShortestPath` function calculates the shortest path from a given source vertex `s` to
// all other vertices in the graph using the unweighted shortest path algorithm.
func (g *Graph) UnweightedShortestPath(s int) {
	queue := make([]int, 0)
	distance := make([]int, g.vertices)
	path := make([]int, g.vertices)

	queue = append(queue, s) // EnQueue
	for i := 0; i < g.vertices; i++ {
		distance[i] = -1
	}
	distance[s] = 0
	for len(queue) > 0 {
		v := queue[0]
		neighbours := g.adjList[v]
		queue = queue[1:] // DeQueue
		for _, neighbour := range neighbours {
			if distance[neighbour] == -1 {
				distance[neighbour] = distance[v] + 1
				path[neighbour] = v
				queue = append(queue, neighbour) // EnQueue
			}
		}
	}
	fmt.Println(colorMagenta, "\tDistance from vertex ", s, " to:")
	for i, d := range distance {
		fmt.Printf("\t\tVertex = %d\tDistance = %d\n", i, d)
	}
	// fmt.Println("\tPath from vertex s = ", s)
	// for _, p := range distance {
	// 	// fmt.Printf("\t\tVertex = %d\tPath = %d\n", i, p)
	// 	fmt.Printf("\t\tVertex name through which we get shortest distance: %d\n", p)
	// }
}

/*
Dijkstra's Shortest Path Algorithm using PriorityQueue (Heap)
 1. Initialize distances of all vertices as infinite.
 2. Create an empty priority_queue pq.  Every item of pq is a pair (weight, vertex).
    Weight (or distance) is used  as first item  of pair, as first item is by default used to compare two pairs.
 3. Insert source vertex into pq and make its distance as 0.
 4. While either pq doesn’t become empty
    ======> 1. Extract minimum distance vertex from pq.
    Let the extracted vertex be u.
    ======> 2. Loop through all adjacent of u and do the following for every vertex v.
    ======> 3. If there is a shorter path to v through u.
    ===========> 1.Update distance of v, i.e., do dist[v] = dist[u] + weight(u, v)
    ===========> 2. Insert v into the pq (Even if v is already there)
 5. Print distance array dist[] to print all shortest paths.
*/
func (g *Graph) Dijkstra(s int) {
	pq := &heaps.Heap{}
	// Min-Heap
	pq.Create([]int{}, g.vertices, 0)
	distance := make([]int, g.vertices)
	heaps.Insert(pq, s)
	for i := 0; i < g.vertices; i++ {
		distance[i] = -1
	}
	distance[s] = 0
	// While PriorityQueue is not empty
	for !heaps.IsEmpty(pq) 
		u := heaps.Delete(pq)
		neighbours := g.adjList[u]
		for _, v := range neighbours {

		}
	}

}
