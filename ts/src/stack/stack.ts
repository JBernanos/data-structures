type Node = {
  value: number;
  next: Node | null;
};

export default class Stack {
  private top: Node | null;
  private count: number;
  private maxSize: number;

  constructor() {
    this.top = null;
    this.count = 0;
    this.maxSize = 127;
  }

  public destroy() {
    this.top = null;
    this.count = 0;
  }

  public push(value: number) {
    if (this.isFull()) {
      console.log(`(ERROR) - Stack max size (${this.maxSize}) reached.`);
      return;
    }

    this.top = { value, next: this.top };
    this.count += 1;
  }

  public pop() {
    if (this.isEmpty()) {
      console.log("(ERROR) - Cannot pop, stack is empty.");
      return;
    }

    this.top = this.top!.next;
    this.count -= 1;
  }

  public peek(): Node | null {
    return this.top;
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
}
