#include <iostream>
using namespace std;

struct Node {
    int data;
    Node* next;
    Node *prev;

    Node(int value) : data(value), next(NULL), prev(NULL) {}
};

class LinkedList {
private:
    Node* head;
    Node* tail;
public:
    LinkedList() : head(NULL), tail(NULL) {}

    void pushFront(int value) {
        Node* newNode = new Node(value);

        if (this->isEmpty()) {
            head = tail = newNode;
        } else {
            newNode->next = head;
            head->prev = newNode;
            head = newNode;
        }
    }

    void pushBack(int value) {
        Node* newNode = new Node(value);

        if (this->isEmpty()) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            newNode->prev = tail;
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
        } else {
            Node* nodeToDelete = head;
            head = head->next;
            head->prev = NULL;
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
        } else {
            Node* nodeToDelete = tail;
            tail = tail->prev;
            tail->next = NULL;
            delete tail;
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
