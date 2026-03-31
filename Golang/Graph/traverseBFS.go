package main

func bfs(graph map[string][]string, start string) []string {
	visited := make(map[string]bool)
	queue := []string{start}
	output := []string{}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // pop front
		output = append(output, current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return output
}
