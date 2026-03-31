adjacency_list = {
    'A': ['B', 'F'],
    'B': ['A', 'C'],
    'C': ['B', 'E', 'D'],
    'D': ['C', 'E'],
    'E': ['D', 'C', 'F'],
    'F': ['A', 'E']
}

def dfs(graph, vertex, visited=None, output=None):
    if visited is None:
        visited = set()
    if output is None:
        output = []

    output.append(vertex)
    visited.add(vertex)

    for neighbor in graph[vertex]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited, output)

    return output


print(dfs(adjacency_list, 'A'))