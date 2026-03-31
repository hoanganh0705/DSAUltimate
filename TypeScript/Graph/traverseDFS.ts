const graph1: Record<string, string[]> = {
  A: ['B', 'F'],
  B: ['A', 'C'],
  C: ['B', 'E', 'D'],
  D: ['C', 'E'],
  E: ['D', 'C', 'F'],
  F: ['A', 'E']
}

function dfs(
  graph: Record<string, string[]>,
  vertex: string,
  visited: Set<string> = new Set(),
  output: string[] = []
): string[] {
  output.push(vertex)
  visited.add(vertex)

  for (const neighbor of graph[vertex]) {
    if (!visited.has(neighbor)) {
      dfs(graph, neighbor, visited, output)
    }
  }

  return output
}

console.log(dfs(graph1, 'A'))