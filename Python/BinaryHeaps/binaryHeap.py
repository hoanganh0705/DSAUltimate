'''

Write a max Heap Class that supports the following:

1. Building a Max heap from an input array
2. Inserting integers in the Heap
3. Removing the Heap’s maximum / root value
4. Peeking at the Heap’s maximum / root value

The Heap is to be represented in the form of an array.
'''

class MaxBinaryHeap:
    def __init__(self):
        self.heap = []

    def build_heap(self, array):
        last_parent = len(array) // 2 - 1
        for i in range(last_parent, -1, -1):
            self._bubble_down(array, i)
        self.heap = array
        return self

    def _bubble_down(self, arr, idx):
        length = len(arr)
        while True:
            left = 2 * idx + 1
            right = 2 * idx + 2
            largest = idx

            if left < length and arr[left] > arr[largest]:
                largest = left

            if right < length and arr[right] > arr[largest]:
                largest = right

            if largest == idx:
                break

            arr[idx], arr[largest] = arr[largest], arr[idx]
            idx = largest

    def extract_max(self):
        if not self.heap:
            return None

        max_val = self.heap[0]
        last = self.heap.pop()

        if self.heap:
            self.heap[0] = last
            self._bubble_down(self.heap, 0)

        return max_val

    def insert(self, value):
        self.heap.append(value)
        self._bubble_up()

    def _bubble_up(self):
        idx = len(self.heap) - 1

        while idx > 0:
            parent = (idx - 1) // 2

            if self.heap[idx] <= self.heap[parent]:
                break

            self.heap[idx], self.heap[parent] = self.heap[parent], self.heap[idx]
            idx = parent

    def peek(self):
        return self.heap[0] if self.heap else None