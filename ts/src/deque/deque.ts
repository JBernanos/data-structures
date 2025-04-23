type DequeNode = {
  value: number;
  next: DequeNode | null;
  previous: DequeNode | null;
};

export default class Deque {
  private front: DequeNode | null;
  private back: DequeNode | null;
  private count: number;
  private maxSize: number;

  constructor() {
    this.front = null;
    this.back = null;
    this.count = 0;
    this.maxSize = 127;
  }

  public destroy() {
    this.front = null;
    this.back = null;
    this.count = 0;
  }

  public insertFront(value: number) {
    if (this.isFull()) return;

    const node: DequeNode = { value, next: this.front, previous: null };

    if (this.isEmpty()) {
      this.back = node;
    } else {
      this.front!.previous = node;
    }
    this.front = node;
    this.count += 1;
  }

  public insertBack(value: number) {
    if (this.isFull()) return;

    const node: DequeNode = { value, next: null, previous: this.back };

    if (this.isEmpty()) {
      this.front = node;
    } else {
      this.back!.next = node;
    }
    this.back = node;
    this.count += 1;
  }

  public removeFront() {
    if (this.isEmpty()) return;

    this.front = this.front!.next;
    if (this.front) {
      this.front.previous = null;
    } else {
      this.back = null;
    }
    this.count -= 1;
  }

  public removeBack() {
    if (this.isEmpty()) return;

    this.back = this.back!.previous;
    if (this.back) {
      this.back.next = null;
    } else {
      this.front = null;
    }
    this.count -= 1;
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
