/*

Question

Implement a Priority Queue as a min Binary Heap. The Priority Queue class should support the following functions:

1. Enqueue to insert an element
2. Dequeue to extract the element with the highest priority (lowest numerical priority is treated as highest priority)

*/

class BNode<T> {
  value: T;
  priority: number;

  constructor(value: T, priority: number) {
    this.value = value;
    this.priority = priority;
  }
}

class PriorityQueue<T> {
  private data: BNode<T>[];

  constructor() {
    this.data = [];
  }

  enqueue(value: T, priority: number): this {
    const node = new BNode(value, priority);
    this.data.push(node);
    this.bubbleUp();
    return this;
  }

  private bubbleUp(): void {
    let idx = this.data.length - 1;

    while (idx > 0) {
      const parent = Math.floor((idx - 1) / 2);

      if (this.data[idx].priority >= this.data[parent].priority) break;

      [this.data[idx], this.data[parent]] = [
        this.data[parent],
        this.data[idx],
      ];

      idx = parent;
    }
  }

  dequeue(): BNode<T> | null {
    if (this.data.length === 0) return null;

    const min = this.data[0];
    const last = this.data.pop()!;

    if (this.data.length > 0) {
      this.data[0] = last;
      this.bubbleDown();
    }

    return min;
  }

  private bubbleDown(): void {
    let idx = 0;
    const length = this.data.length;

    while (true) {
      let left = 2 * idx + 1;
      let right = 2 * idx + 2;
      let smallest = idx;

      if (
        left < length &&
        this.data[left].priority < this.data[smallest].priority
      ) {
        smallest = left;
      }

      if (
        right < length &&
        this.data[right].priority < this.data[smallest].priority
      ) {
        smallest = right;
      }

      if (smallest === idx) break;

      [this.data[idx], this.data[smallest]] = [
        this.data[smallest],
        this.data[idx],
      ];

      idx = smallest;
    }
  }
}