class Node {
    constructor(value) {
        this.data = value;
        this.next = null;
    }
}

class Stack {
    constructor() {
        this.head = null;
    }

    push(value) {
        let newNode = new Node(value);
        if (this.head === null) {
            this.head = newNode;
        } else {
            newNode.next = this.head;
            this.head = newNode;
        }
    }

    pop() {
        if (this.isEmpty()) {
            throw new Error("Empty stack. Cannot perform pop!")
        }
        if (this.head.next === null) {
            this.head = null;
        } else {
            this.head = this.head.next;
        }
    }

    peek() {
        return this.head.data
    }

    isEmpty() {
        return this.head === null;
    }
}

// MAIN //
const s = new Stack();
s.push(1);
s.push(2);
s.push(3);
s.push(4);

while (!s.isEmpty()) {
    console.log(s.peek());
    s.pop();
}
