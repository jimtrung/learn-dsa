class Node {
    constructor(value) {
        this.data = value;
        this.next = null;
    }
}

class LinkedList {
    constructor() {
        this.tail = null;
    }

    pushFront(value) {
        let newNode = new Node(value);
        if (this.tail === null) {
            newNode.next = newNode;
            this.tail = newNode;
        } else {
            newNode.next = this.tail.next;
            this.tail.next = newNode;
        }
    }

    pushBack(value) {
        let newNode = new Node(value);
        if (this.tail === null) {
            newNode.next = newNode;
            this.tail = newNode;
        } else {
            newNode.next = this.tail.next;
            this.tail.next = newNode;
            this.tail = newNode;
        }
    }

    popFront() {
        if (this.tail === null) {
            throw new Error("Empty list. Cannot perform popFront!")
        }
        if (this.tail.next === this.tail) {
            this.tail = null;
        } else {
            this.tail.next = this.tail.next.next;
        }
    }

    popBack() {
        if (this.tail === null) {
            throw new Error("Empty list. Cannot perform popBack!")
        }
        if (this.tail.next === this.tail) {
            this.tail = null;
        } else {
            let curr = this.tail;
            while (curr.next !== this.tail) {
                curr = curr.next;
            }
            curr.next = curr.next.next;
            this.tail = curr;
        }
    }

    display() {
        if (this.tail === null) {
            throw new Error("Empty list. Cannot perform display!")
        }

        let curr = this.tail;
        let displayList = "";
        let lastNode = `${curr.data}->`;
        curr = curr.next;

        while (curr !== this.tail) {
            displayList += `${curr.data}->`;
            curr = curr.next;
        }
        displayList += lastNode;
        displayList += "this.head";
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

