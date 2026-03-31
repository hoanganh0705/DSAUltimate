package main

func dfs(graph map[string][]string, vertex string, visited map[string]bool, output *[]string) {
	*output = append(*output, vertex)
	visited[vertex] = true

	for _, neighbor := range graph[vertex] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited, output)
		}
	}
}
