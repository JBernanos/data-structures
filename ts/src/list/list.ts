type ListNode = {
  value: number;
  next: ListNode | null;
  previous: ListNode | null;
};

export default class List {
  private head: ListNode | null;
  private tail: ListNode | null;
  private count: number;
  private maxSize: number;

  constructor() {
    this.head = null;
    this.tail = null;
    this.count = 0;
    this.maxSize = 127;
  }

  public destroy() {
    this.head = null;
    this.tail = null;
    this.count = 0;
  }

  public insertAtHead(value: number): void {
    if (this.isFull()) return;

    if (this.isEmpty()) {
      const node: ListNode = { value, next: null, previous: null };
      node.next = node;
      node.previous = node;
      this.head = node;
      this.tail = node;
    } else {
      const node: ListNode = { value, next: this.head, previous: this.tail };
      this.head!.previous = node;
      this.tail!.next = node;
      this.head = node;
    }

    this.count += 1;
  }

  public insertAtTail(value: number): void {
    if (this.isFull()) return;

    if (this.isEmpty()) {
      const node: ListNode = { value, next: null, previous: null };
      node.next = node;
      node.previous = node;
      this.head = node;
      this.tail = node;
    } else {
      const node: ListNode = { value, next: this.head, previous: this.tail };
      this.head!.previous = node;
      this.tail!.next = node;
      this.tail = node;
    }

    this.count += 1;
  }

  public removeHead(): void {
    if (this.isEmpty()) return;

    if (this.head == this.tail) {
      this.head = null;
      this.tail = null;
    }

    this.head = this.head!.next;
    this.head!.previous = this.tail;
    this.tail!.next = this.head;

    this.count -= 1;
  }

  public removeTail(): void {
    if (this.isEmpty()) return;

    if (this.head == this.tail) {
      this.head = null;
      this.tail = null;
    }

    this.tail = this.tail!.previous;
    this.head!.previous = this.tail;
    this.tail!.next = this.head;

    this.count -= 1;
  }

  public peekHead(): number | null {
    if (this.isEmpty()) return null;

    return this.head!.value;
  }

  public peekTail(): number | null {
    if (this.isEmpty()) return null;

    return this.tail!.value;
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
    let currentNode = this.head;
    while (currentNode != null) {
      if (currentNode.value === value) return true;
      currentNode = currentNode.next;
    }
    return false;
  }
}
