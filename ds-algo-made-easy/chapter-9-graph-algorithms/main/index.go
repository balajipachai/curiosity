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
	g := &graphs.Graph{}
	g.NewGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)

	fmt.Println(colorYellow + "\tPrinting the Graph's Adjaceny List" + colorReset)
	g.PrintAdjList()
	fmt.Println(colorYellow + "\tDFS Traversal of the graph starting at vertex 2 is:::" + colorReset)
	g.DFS(2)
	fmt.Println()
	printDottedLine()

	fmt.Println(colorYellow + "\tBFS Traversal of the graph starting at vertex 1 is:::" + colorReset)
	g.BFS(1)
	fmt.Println()
	printDottedLine()
}

func main() {
	GraphBasics()
}
