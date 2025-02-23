#include <iostream>
#include "SinglyLinkedList.h"

using namespace std;

void testPushFront() {
    LinkedList list;
    list.pushFront(10);
    list.pushFront(20);
    list.pushFront(30);
    cout << "Test pushFront: " << ((list.isEmpty() == false) ? "PASSED" : "FAILED") << endl;
    list.display();
}

void testPushBack() {
    LinkedList list;
    list.pushBack(10);
    list.pushBack(20);
    list.pushBack(30);
    cout << "Test pushBack: " << ((list.isEmpty() == false) ? "PASSED" : "FAILED") << endl;
    list.display();
}

void testPopFront() {
    LinkedList list;
    list.pushBack(10);
    list.pushBack(20);
    list.pushBack(30);
    list.popFront();
    cout << "Test popFront: " << ((list.isEmpty() == false) ? "PASSED" : "FAILED") << endl;
    list.display();
}

void testPopBack() {
    LinkedList list;
    list.pushBack(10);
    list.pushBack(20);
    list.pushBack(30);
    list.popBack();
    cout << "Test popBack: " << ((list.isEmpty() == false) ? "PASSED" : "FAILED") << endl;
    list.display();
}

void testIsEmpty() {
    LinkedList list;
    bool initialEmpty = list.isEmpty();
    list.pushBack(10);
    bool afterPushEmpty = list.isEmpty();
    list.popFront();
    bool afterPopEmpty = list.isEmpty();

    cout << "Test isEmpty (initial): " << (initialEmpty ? "PASSED" : "FAILED") << endl;
    cout << "Test isEmpty (after push): " << (!afterPushEmpty ? "PASSED" : "FAILED") << endl;
    cout << "Test isEmpty (after pop): " << (afterPopEmpty ? "PASSED" : "FAILED") << endl;
}

int main() {
    testPushFront();
    testPushBack();
    testPopFront();
    testPopBack();
    testIsEmpty();
    return 0;
}
