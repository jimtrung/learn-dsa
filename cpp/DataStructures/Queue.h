#include <iostream>
using namespace std;

struct Node {
    int data;
    Node* next;

    Node(int value) : data(value), next(NULL) {}
};

class Queue {
private:
    Node* head;
    Node* tail;
public:
    Queue() : head(NULL), tail(NULL) {}

    void enqueue(int value) {
        Node* newNode = new Node(value);

        if (tail == NULL) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            tail = newNode;
        }
    }

    void dequeue() {
        if (this->isEmpty()) {
            cout << "Empty list. Cannot perform popFront" << endl;
            return;
        }
        if (head->next == NULL) {
            delete head;
            head = tail = NULL;
            return;
        } else {
            Node* nodeToDelete = head;
            head = head->next;
            delete nodeToDelete;
        }
    }

    bool isEmpty() {
        return head == NULL && tail == NULL;
    }

    int peek() {
        if (isEmpty()) {
            return -1;
        }

        return head->data;
    }
};
