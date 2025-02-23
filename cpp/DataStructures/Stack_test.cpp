#include <iostream>
#include <cassert>
#include "Stack.h"  // Include the header file for LinkedList

using namespace std;

void testPush() {
    LinkedList stack;
    assert(stack.isEmpty() == true);

    stack.push(10);
    assert(stack.isEmpty() == false);
    assert(stack.peek() == 10);

    stack.push(20);
    assert(stack.peek() == 20);

    stack.push(30);
    assert(stack.peek() == 30);

    cout << "testPush passed!" << endl;
}

void testPop() {
    LinkedList stack;
    stack.push(10);
    stack.push(20);
    stack.push(30);

    stack.pop();
    assert(stack.peek() == 20);

    stack.pop();
    assert(stack.peek() == 10);

    stack.pop();
    assert(stack.isEmpty() == true);

    cout << "testPop passed!" << endl;
}

void testEdgeCases() {
    LinkedList stack;

    // Popping from an empty stack (should not crash)
    stack.pop();

    // Peeking from an empty stack
    assert(stack.peek() == -1);

    cout << "testEdgeCases passed!" << endl;
}

int main() {
    testPush();
    testPop();
    testEdgeCases();

    cout << "All tests passed successfully!" << endl;
    return 0;
}

