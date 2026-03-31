package main

import "fmt"

func main() {
	graph := map[string][]string{
		"A": {"B", "F"},
		"B": {"A", "C"},
		"C": {"B", "E", "D"},
		"D": {"C", "E"},
		"E": {"D", "C", "F"},
		"F": {"A", "E"},
	}

	visited := make(map[string]bool)
	output := []string{}

	dfs(graph, "A", visited, &output)

	fmt.Println(output)

	graph1 := map[string][]string{
		"A": {"B", "F"},
		"B": {"A", "F", "C"},
		"C": {"B", "E", "D"},
		"D": {"C", "E"},
		"E": {"D", "C", "F"},
		"F": {"A", "B", "E"},
	}

	fmt.Println(bfs(graph1, "A"))
}
