package main

type MaxBinaryHeap struct {
	heap []int
}

// constructor
func NewMaxBinaryHeap() *MaxBinaryHeap {
	return &MaxBinaryHeap{
		heap: []int{},
	}
}

// BuildHeap
func (h *MaxBinaryHeap) BuildHeap(array []int) *MaxBinaryHeap {
	lastParent := len(array)/2 - 1

	for i := lastParent; i >= 0; i-- {
		h.bubbleDown(array, i)
	}

	h.heap = array
	return h
}

// bubbleDown
func (h *MaxBinaryHeap) bubbleDown(arr []int, idx int) {
	length := len(arr)

	for {
		left := 2*idx + 1
		right := 2*idx + 2
		largest := idx

		if left < length && arr[left] > arr[largest] {
			largest = left
		}

		if right < length && arr[right] > arr[largest] {
			largest = right
		}

		if largest == idx {
			break
		}

		arr[idx], arr[largest] = arr[largest], arr[idx]
		idx = largest
	}
}

// ExtractMax
func (h *MaxBinaryHeap) ExtractMax() *int {
	if len(h.heap) == 0 {
		return nil
	}

	max := h.heap[0]
	last := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	if len(h.heap) > 0 {
		h.heap[0] = last
		h.bubbleDown(h.heap, 0)
	}

	return &max
}

// Insert
func (h *MaxBinaryHeap) Insert(value int) *MaxBinaryHeap {
	h.heap = append(h.heap, value)
	h.bubbleUp()
	return h
}

// bubbleUp
func (h *MaxBinaryHeap) bubbleUp() {
	idx := len(h.heap) - 1

	for idx > 0 {
		parent := (idx - 1) / 2

		if h.heap[idx] <= h.heap[parent] {
			break
		}

		h.heap[idx], h.heap[parent] = h.heap[parent], h.heap[idx]
		idx = parent
	}
}

// Peek
func (h *MaxBinaryHeap) Peek() *int {
	if len(h.heap) == 0 {
		return nil
	}
	return &h.heap[0]
}
