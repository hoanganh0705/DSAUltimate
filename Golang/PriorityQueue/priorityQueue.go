/*

Question

Implement a Priority Queue as a min Binary Heap. The Priority Queue class should support the following functions:

1. Enqueue to insert an element
2. Dequeue to extract the element with the highest priority (lowest numerical priority is treated as highest priority)

*/

package priorityqueue

type Node struct {
	Value    any
	Priority int
}

type PriorityQueue struct {
	data []Node
}

// constructor
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		data: []Node{},
	}
}

// Enqueue
func (pq *PriorityQueue) Enqueue(value any, priority int) *PriorityQueue {
	node := Node{Value: value, Priority: priority}
	pq.data = append(pq.data, node)
	pq.bubbleUp()
	return pq
}

func (pq *PriorityQueue) bubbleUp() {
	idx := len(pq.data) - 1

	for idx > 0 {
		parent := (idx - 1) / 2

		if pq.data[idx].Priority >= pq.data[parent].Priority {
			break
		}

		pq.data[idx], pq.data[parent] = pq.data[parent], pq.data[idx]
		idx = parent
	}
}

// Dequeue
func (pq *PriorityQueue) Dequeue() *Node {
	if len(pq.data) == 0 {
		return nil
	}

	min := pq.data[0]
	last := pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]

	if len(pq.data) > 0 {
		pq.data[0] = last
		pq.bubbleDown()
	}

	return &min
}

func (pq *PriorityQueue) bubbleDown() {
	idx := 0
	length := len(pq.data)

	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left < length && pq.data[left].Priority < pq.data[smallest].Priority {
			smallest = left
		}

		if right < length && pq.data[right].Priority < pq.data[smallest].Priority {
			smallest = right
		}

		if smallest == idx {
			break
		}

		pq.data[idx], pq.data[smallest] = pq.data[smallest], pq.data[idx]
		idx = smallest
	}
}
