'''

Question

Implement a Priority Queue as a min Binary Heap. The Priority Queue class should support the following functions:

1. Enqueue to insert an element
2. Dequeue to extract the element with the highest priority (lowest numerical priority is treated as highest priority)


'''

class Node:
    def __init__(self, value, priority):
        self.value = value
        self.priority = priority


class PriorityQueue:
    def __init__(self):
        self.data = []

    def enqueue(self, value, priority):
        node = Node(value, priority)
        self.data.append(node)
        self._bubble_up()
        return self

    def _bubble_up(self):
        idx = len(self.data) - 1

        while idx > 0:
            parent = (idx - 1) // 2

            if self.data[idx].priority >= self.data[parent].priority:
                break

            self.data[idx], self.data[parent] = self.data[parent], self.data[idx]
            idx = parent

    def dequeue(self):
        if not self.data:
            return None

        min_element = self.data[0]
        last = self.data.pop()

        if self.data:
            self.data[0] = last
            self._bubble_down()

        return min_element

    def _bubble_down(self):
        idx = 0
        length = len(self.data)

        while True:
            left = 2 * idx + 1
            right = 2 * idx + 2
            smallest = idx

            if left < length and self.data[left].priority < self.data[smallest].priority:
                smallest = left

            if right < length and self.data[right].priority < self.data[smallest].priority:
                smallest = right

            if smallest == idx:
                break

            self.data[idx], self.data[smallest] = self.data[smallest], self.data[idx]
            idx = smallest