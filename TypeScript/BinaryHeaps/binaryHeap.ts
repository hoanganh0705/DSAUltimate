/*

Write a max Heap Class that supports the following:

1. Building a Max heap from an input array
2. Inserting integers in the Heap
3. Removing the Heap’s maximum / root value
4. Peeking at the Heap’s maximum / root value

The Heap is to be represented in the form of an array.
*/

class MaxBinaryHeap {
  heap: number[];

  constructor() {
    this.heap = [];
  }

  buildHeap(array: number[]): this {
    const lastParent = Math.floor(array.length / 2) - 1;

    for (let i = lastParent; i >= 0; i--) {
      this.bubbleDown(array, i);
    }

    this.heap = array;
    return this;
  }

  private bubbleDown(arr: number[], idx: number): void {
    const length = arr.length;

    while (true) {
      let left = 2 * idx + 1;
      let right = 2 * idx + 2;
      let largest = idx;

      if (left < length && arr[left] > arr[largest]) {
        largest = left;
      }

      if (right < length && arr[right] > arr[largest]) {
        largest = right;
      }

      if (largest === idx) break;

      [arr[idx], arr[largest]] = [arr[largest], arr[idx]];
      idx = largest;
    }
  }

  extractMax(): number | null {
    if (this.heap.length === 0) return null;

    const max = this.heap[0];
    const last = this.heap.pop()!;

    if (this.heap.length > 0) {
      this.heap[0] = last;
      this.bubbleDown(this.heap, 0);
    }

    return max;
  }

  insert(value: number): this {
    this.heap.push(value);
    this.bubbleUp();
    return this;
  }

  private bubbleUp(): void {
    let idx = this.heap.length - 1;

    while (idx > 0) {
      const parent = Math.floor((idx - 1) / 2);

      if (this.heap[idx] <= this.heap[parent]) break;

      [this.heap[idx], this.heap[parent]] = [
        this.heap[parent],
        this.heap[idx],
      ];

      idx = parent;
    }
  }

  peek(): number | null {
    return this.heap.length ? this.heap[0] : null;
  }
}