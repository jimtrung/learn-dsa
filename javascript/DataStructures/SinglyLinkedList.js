class Node {
    constructor(data) {
        this.data = data;
        this.next = null;
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
            this.tail = newNode;
        }
    }

    push(value, pos) {
        if (pos < 1) {
            throw new Error("Invalid position");
        }
        if (pos === 1) {
            this.pushFront(value);
            return
        }

        let newNode = new Node(value);
        let curr = this.head;

        for (let i = 1; i < pos - 1 && curr; i++) {
            curr = curr.next;
        }

        if (curr === null) {
            throw new Error("Out of bound");
        }

        newNode.next = curr.next;
        curr.next = newNode;
        if (newNode.next === null) {
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
            let curr = this.head;
            while (curr.next.next !== null) {
                curr = curr.next;
            }
            curr.next = null;
            this.tail = curr;
        }
    }

    pop(pos) {
        if (pos < 1) {
            throw new Error("Invalid position");
        }
        if (pos === 1) {
            this.popFront();
            return
        }

        let curr = this.head;

        for (let i = 1; i < pos - 1 && curr; i++) {
            curr = curr.next;
        }
        if (curr === null || curr.next === null) {
            throw new Error("Out of bound");
        }

        if (curr.next === this.tail) {
            this.tail = curr;
        }
        curr.next = curr.next.next;
    }

    display() {
        if (this.head === null) {
            throw new Error("Empty list. Cannot perform display!")
        }

        let curr = this.head;
        let displayList = "";
        while (curr !== null) {
            displayList += `${curr.data}->`;
            curr = curr.next;
        }
        displayList += "NULL";
        console.log(displayList);
    }

    contains(target) {
        let curr = this.head;
        while (curr) {
            if (curr.data === target) {
                return true;
            }
            curr = curr.next;
        }
        return false;
    }
}

// MAIN AREA //
const list = new LinkedList();
list.pushBack(1);
list.pushBack(3);
list.pushBack(4);
list.pushBack(5);
list.pushFront(0);
list.push(2, 3);

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

try {
    list.pop(2);
} catch (err) {
    console.log(err);
}

list.display();
