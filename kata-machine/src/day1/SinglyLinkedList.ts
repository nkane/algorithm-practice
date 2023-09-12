class Node<T> {
    public data: T;

    public next: Node<T> | null = null;
    public prev: Node<T> | null = null;

    constructor(data: T) {
        this.data = data;
    }
}

export default class SinglyLinkedList<T> {
    public length: number;
    public head: Node<T> | null = null;
    public tail: Node<T> | null = null;

    constructor() {
        this.length = 0;
        this.head = null;
        this.tail = null;
    }

    prepend(item: T): void {
        this.insertAt(item, 0);
        this.head = item as any;
    }

    append(item: T): void {
        this.insertAt(item, this.length);
        this.tail = item as any;
    }

    insertAt(item: T, idx: number): void {
        let new_node = new Node(item);
        if (this.length === 0) {
            // list is empty
            this.head = new_node;
            this.tail = this.head;
        } else {
            // make sure index is in range
            if (idx > this.length) {
                this.append(item);
            } else {
                // walk the list until a particular index
                let i = 0;
                let current_node = this.head as any;
                while (idx > i) {
                    if (current_node.next) {
                        current_node = current_node.next;
                    } else {
                        break;
                    }
                    i++;
                }
                new_node.next = current_node.next;
                new_node.prev = current_node;
                current_node.next = new_node;
            }
        }
        this.length++;
    }

    remove(item: T): T | undefined {
        let current_node = this.head as any;
        // locate item
        for (let i = 0; i < this.length; i++) {
            if (current_node.data === item) {
                // found node
                break;
            }
            current_node = current_node.next;
        }
        // remove
        current_node.prev.next = current_node.next;
        current_node.next.prev = current_node.prev;
        this.length--;
        return;
    }

    get(idx: number): T | undefined {
        if (idx > this.length) {
            return undefined;
        }
        let current_node = this.head as any; 
        for (let i = 0; i <= idx; i++) {
            current_node = current_node.next;
        }
        return current_node;
    }

    removeAt(idx: number): T | undefined {
        let current_node = this.get(idx) as any;
        if (current_node) {
            current_node.prev.next = current_node.next;
            current_node.next.prev = current_node.prev;
            this.length--;
        }
        return undefined;
    }
}
