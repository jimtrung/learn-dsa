#include <bits/stdc++.h>
using namespace std;

struct Node {
    int data;
    Node* next;

    Node(int value) : data(value), next(NULL) {}
};

class LinkedList {
private:
    Node* head;
    Node* tail;

public:
    LinkedList() : head(NULL), tail(NULL) {}

    void pushFront(int value) {
        Node* newNode = new Node(value);

        if (head == NULL) {
            head = tail = newNode;
        } else {
            newNode->next = head;
            head = newNode;
        }
    }

    void pushBack(int value) {
        Node* newNode = new Node(value);

        if (tail == NULL) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            tail = newNode;
        }
    }

    void popFront() {
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

    void popBack() {
        if (this->isEmpty()) {
            cout << "Empty list. Cannot perform popBack" << endl;
            return;
        }
        if (head->next == NULL) {
            delete head;
            head = tail = NULL;
            return;
        } else {
            Node* curr = head;

            while (curr->next->next) {
                curr = curr->next;
            }
            delete curr->next;
            curr->next = NULL;
        }
    }

    bool isEmpty() {
        return head == NULL && tail == NULL;
    }

    void display() {
        Node* curr = head;

        while (curr) {
            cout << curr->data << "->";
            curr = curr->next;
        }
        cout << "NULL" << endl;
    }
};
