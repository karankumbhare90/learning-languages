class Node<T> {
  public next: Node<T> | null = null;
  public prev: Node<T> | null = null;

  constructor(public value: T) {}
}

export class LinkedList<TElement> {
  private head: Node<TElement> | null = null;
  private tail: Node<TElement> | null = null;
  private length: number = 0;

  public push(element: unknown) {
    const newNode = new Node(element as TElement);

    if (!this.tail) this.head = this.tail = newNode;
    else {
      this.tail.next = newNode;
      newNode.prev = this.tail;
      this.tail = newNode;
    }

    this.length++;
  }

  public pop(): unknown {
    if (!this.tail) return undefined;

    const removed = this.tail;
    this.tail = this.tail.prev;

    if (this.tail) this.tail.next = null;
    else this.head = null;

    this.length--;
    return removed.value;
  }

  public shift(): unknown {
    if (!this.head) return undefined;

    const removed = this.head;
    this.head = this.head.next;

    if (this.head) this.head.prev = null;
    else this.tail = null;

    this.length--;
    return removed.value;
  }

  public unshift(element: unknown) {
    const newNode = new Node(element as TElement);

    if (!this.head) this.head = this.tail = newNode;
    else {
      newNode.next = this.head;
      this.head.prev = newNode;
      this.head = newNode;
    }

    this.length++;
  }

  public delete(element: unknown) {
    let current = this.head;

    while (current) {
      if (current.value === (element as TElement)) {
        if (current.prev) current.prev.next = current.next;
        else this.head = current.next;

        if (current.next) current.next.prev = current.prev;
        else this.tail = current.prev;

        this.length--;
        return true;
      }
      current = current.next;
    }

    return false;
  }

  public count(): unknown {
    return this.length;
  }
}
