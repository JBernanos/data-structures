type QueueNode = {
  value: number;
  next: QueueNode | null;
};

export default class Queue {
  private front: QueueNode | null;
  private count: number;
  private maxSize: number;

  constructor() {
    this.front = null;
    this.count = 0;
    this.maxSize = 127;
  }

  public destroy() {
    this.front = null;
    this.count = 0;
  }

  public enqueue(value: number) {
    if (this.isFull()) return;

    if (this.isEmpty()) {
      this.front = { value, next: null };
      this.count += 1;
      return;
    }

    let currentNode = this.front;
    while (currentNode?.next != null) {
      currentNode = currentNode.next;
    }

    if (currentNode !== null) {
      currentNode.next = { value, next: null };
      this.count += 1;
    }
  }

  public dequeue() {
    if (this.front === null) {
      console.log("(ERRO) - Cannot dequeue, queue is empty.");
      return;
    }

    this.front = this.front.next;
    this.count -= 1;
  }

  public peek(): number | null {
    if (this.front === null) return null;

    return this.front.value;
  }

  public size(): number {
    return this.maxSize;
  }

  public isEmpty(): boolean {
    return this.count === 0;
  }

  public isFull(): boolean {
    return this.count === this.maxSize;
  }

  public contains(value: number): boolean {
    let currentNode = this.front;
    while (currentNode != null) {
      if (currentNode.value === value) return true;
      currentNode = currentNode.next;
    }

    return false;
  }
}
