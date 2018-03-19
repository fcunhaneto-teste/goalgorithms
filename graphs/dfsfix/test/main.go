package main

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfsfix"
)

func main() {
	var s string
	graph := graphs.ReadGraph()
	fmt.Println("Enter the initial node:")
	fmt.Scanf("%s", &s)
	dfsfix.DFS(graph, s)

	for _, n := range graph {
		fmt.Println("Node:", n.N)
		fmt.Println("Color:", n.C)
		fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
		if n.P != nil {
			fmt.Println("Predecessor:", n.P.N)
		} else {
			fmt.Println("Predecessor: null")
		}
		fmt.Println("*********************************")

	}
}
