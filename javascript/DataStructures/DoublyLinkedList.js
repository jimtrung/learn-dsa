class Node {
    constructor(value) {
        this.data = value;
        this.next = null;
        this.prev = null;
    }
}

class LinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
    }

    pushFront(value) {
        let newNode = new Node(value);
        if (this.head === null) {
            this.head = newNode;
            this.tail = newNode;
        } else {
            newNode.next = this.head;
            this.head.prev = newNode;
            this.head = newNode;
        }
    }

    pushBack(value) {
        let newNode = new Node(value);
        if (this.tail === null) {
            this.head = newNode;
            this.tail = newNode;
        } else {
            this.tail.next = newNode;
            newNode.prev = this.tail;
            this.tail = newNode;
        }
    }

    popFront() {
        if (this.head === null) {
            throw new Error("Empty list. Cannot perform popFront!")
        }
        if (this.head.next === null) {
            this.head = null;
            this.tail = null;
        } else {
            this.head = this.head.next;
            this.prev = null;
        }
    }

    popBack() {
        if (this.head === null) {
            throw new Error("Empty list. Cannot perform popBack!")
        }
        if (this.head.next === null) {
            this.head = null;
            this.tail = null;
        } else {
            this.tail = this.tail.prev;
            this.tail.next = null;
        }
    }

    display() {
        if (this.head === null) {
            throw new Error("Empty list. Cannot perform display!")
        }

        let curr = this.head;
        let displayList = ""
        while (curr) {
            displayList += `${curr.data}->`;
            curr = curr.next;
        }
        displayList += "NULL";
        console.log(displayList);
    }
}

// MAIN AREA //
const list = new LinkedList();
list.pushBack(1);
list.pushBack(3);
list.pushBack(4);
list.pushBack(5);
list.pushFront(0);

try {
    list.popFront();
} catch (err) {
    console.log(err);
}

try {
    list.popBack();
} catch (err) {
    console.log(err);
}

list.display();

