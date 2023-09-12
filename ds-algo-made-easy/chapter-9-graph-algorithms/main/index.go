package main

import (
	"fmt"

	"example.com/graphs"
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
)

func printDottedLine() {
	fmt.Println(colorBlue + "---------------------------------------------------------------------------------------------------------------------------------------------------------------" + colorReset)
}

func GraphBasics() {
	fmt.Println(colorYellow + "\tCreating a graph" + colorReset)

	// DFS
	fmt.Println(colorCyan + "\tDEPTH FIRST SEARCH" + colorReset)
	g1 := &graphs.Graph{}
	g1.NewGraph(5)
	g1.AddEdge(0, 1)
	g1.AddEdge(0, 2)
	g1.AddEdge(0, 3)
	g1.AddEdge(2, 3)
	g1.AddEdge(2, 4)
	fmt.Println(colorYellow + "\tPrinting the Graph's (G1) Adjaceny List" + colorReset)
	g1.PrintAdjList()
	fmt.Println(colorYellow + "\tDFS Traversal of the graph (G1) starting at vertex 2 is:::" + colorReset)
	g1.DFS(2)
	fmt.Println()
	printDottedLine()

	// BFS
	fmt.Println(colorCyan + "\tBREADTH FIRST SEARCH" + colorReset)
	g2 := &graphs.Graph{}
	g2.NewGraph(4)
	g2.AddEdge(0, 1)
	g2.AddEdge(0, 2)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 0)
	g2.AddEdge(2, 3)
	g2.AddEdge(3, 3)
	fmt.Println(colorYellow + "\tPrinting the Graph's (G2) Adjaceny List" + colorReset)
	g2.PrintAdjList()
	fmt.Println(colorYellow + "\tBFS Traversal of the graph (G2) starting at vertex 1 is:::" + colorReset)
	g2.BFS(2)
	fmt.Println()
	printDottedLine()

	// Topological Sorting
	fmt.Println(colorCyan + "\tTOPOLOGICAL SORTING" + colorReset)
	dg := &graphs.Graph{}
	dg.NewGraph(6)
	dg.AddDirectedEdge(5, 2)
	dg.AddDirectedEdge(5, 0)
	dg.AddDirectedEdge(4, 0)
	dg.AddDirectedEdge(4, 1)
	dg.AddDirectedEdge(2, 3)
	dg.AddDirectedEdge(3, 1)
	fmt.Println(colorYellow + "\tPrinting the Graph's (DG-Directed Graph) Adjaceny List" + colorReset)
	dg.PrintAdjList()
	fmt.Println(colorYellow + "\tTopological Sort of the graph (DG):::" + colorReset)
	dg.TopologicalSort()
	fmt.Println()
	printDottedLine()

	// Shortest Path Algorithm
	fmt.Println(colorCyan + "\tSHORTEST PATH ALGORITHMS" + colorReset)
	g3 := &graphs.Graph{}
	g3.NewGraph(7)
	g3.AddDirectedEdge(0, 1)
	g3.AddDirectedEdge(0, 3)
	g3.AddDirectedEdge(1, 3)
	g3.AddDirectedEdge(1, 4)
	g3.AddDirectedEdge(2, 0)
	g3.AddDirectedEdge(2, 5)
	g3.AddDirectedEdge(3, 5)
	g3.AddDirectedEdge(3, 6)
	g3.AddDirectedEdge(4, 6)
	fmt.Println(colorYellow + "\tPrinting the G3's Adjaceny List" + colorReset)
	g3.PrintAdjList()
	fmt.Println(colorYellow + "\tUnweightedShortestPath (G3) starting from vertex 2 :::" + colorReset)
	g3.UnweightedShortestPath(2)
	printDottedLine()
}

func main() {
	GraphBasics()
}
