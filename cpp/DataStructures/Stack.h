#include <iostream>
using namespace std;

struct Node {
    int data;
    Node* next;

    Node(int value) : data(value), next(NULL) {}
};

class LinkedList {
private:
    Node* head;

public:
    LinkedList() : head(NULL) {}

    void push(int value) {
        Node* newNode = new Node(value);

        newNode->next = head;
        head = newNode;
    }

    void pop() {
        if (this->isEmpty()) {
            cout << "Empty stack. Cannot perform pop" << endl;
            return;
        }
        Node *curr = head;
        head = head->next;
        delete curr;
    }

    bool isEmpty() {
        return head == NULL;
    }

    int peek() {
        if (this->isEmpty()) {
            cout << "Empty stack. Cannot perform pop" << endl;
            return -1;
        }

        return head->data;
    }
};
