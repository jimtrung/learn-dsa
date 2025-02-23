#include <iostream>
using namespace std;

struct Node {
    int data;
    Node *next;

    Node(int value) : data(value), next(NULL) {}
};

class LinkedList {
private:
    Node *tail;

public:
    LinkedList() : tail(NULL) {}

    void pushFront(int value) {
        Node *newNode = new Node(value);

        if (this->isEmpty()) {
            newNode->next = newNode;
            tail = newNode;
        } else {
            newNode->next = tail->next;
            tail->next = newNode;
        }
    }

    void pushBack(int value) {
        Node *newNode = new Node(value);

        if (this->isEmpty()) {
            newNode->next = newNode;
            tail = newNode;
        } else {
            newNode->next = tail->next;
            tail->next = newNode;
            tail = newNode;
        }
    }

    void popFront() {
        if (this->isEmpty()) {
            cout << "Empty list. Cannot perform popFront" << endl;
            return;
        }
        Node* nodeToDelete = tail->next;
        if (tail == tail->next) {
            tail = NULL;
        } else {
            tail->next = tail->next->next;
        }
        delete nodeToDelete;
    }

    void popBack() {
        if (this->isEmpty()) {
            cout << "Empty list. Cannot perform popBack" << endl;
            return;
        }
        if (tail == tail->next) {
            delete tail;
            tail = NULL;
        } else {
            Node *curr = tail->next;

            while (curr->next != tail) {
                curr = curr->next;
            }

            curr->next = tail->next;
            delete tail;
            tail = curr;
        }
    }

    bool isEmpty() { return tail == NULL; }

    void display() {
        if (!this->isEmpty()) {
            Node *curr = tail->next;

            do {
                cout << curr->data << "->";
                curr = curr->next;
            } while (curr != tail->next);
        }
        cout << "NULL" << endl;
    }
};
