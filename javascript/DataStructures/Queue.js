class Node {
    constructor(value) {
        this.data = value;
        this.next = null;
    }
}

class Queue {
    constructor() {
        this.head = null;
        this.tail = null;
    }

    enqueue(value) {
        let newNode = new Node(value);

        if (this.head === null) {
            this.head = newNode;
            this.tail = newNode;
        } else {
            this.tail.next = newNode;
            this.tail = newNode;
        }
    }

    dequeue() {
        if (this.isEmpty()) {
            throw new Error("Empty queue. Cannot perform pop!")
        }
        if (this.head.next === null) {
            this.head = null;
            this.tail = null;
        } else {
            this.head = this.head.next;
        }
    }

    isEmpty() {
        return this.head === null;
    }

    peek() {
        return this.head.data;
    }
}
