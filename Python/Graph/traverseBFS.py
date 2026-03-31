from collections import deque

adjacency_list = {
    'A': ['B', 'F'],
    'B': ['A', 'F', 'C'],
    'C': ['B', 'E', 'D'],
    'D': ['C', 'E'],
    'E': ['D', 'C', 'F'],
    'F': ['A', 'B', 'E']
}

def bfs(graph, start):
    visited = set()
    queue = deque([start])
    output = []

    visited.add(start)

    while queue:
        current = queue.popleft()
        output.append(current)

        for neighbor in graph[current]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)

    return output


print(bfs(adjacency_list, 'A'))