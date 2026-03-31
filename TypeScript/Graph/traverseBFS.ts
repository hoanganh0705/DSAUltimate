const graph: Record<string, string[]> = {
  A: ['B', 'F'],
  B: ['A', 'F', 'C'],
  C: ['B', 'E', 'D'],
  D: ['C', 'E'],
  E: ['D', 'C', 'F'],
  F: ['A', 'B', 'E']
}

function bfs(graph: Record<string, string[]>, start: string): string[] {
  const visited = new Set<string>()
  const queue: string[] = [start]
  const output: string[] = []

  visited.add(start)

  while (queue.length > 0) {
    const current = queue.shift()! // remove first element
    output.push(current)

    for (const neighbor of graph[current]) {
      if (!visited.has(neighbor)) {
        visited.add(neighbor)
        queue.push(neighbor)
      }
    }
  }

  return output
}

console.log(bfs(graph, 'A'))