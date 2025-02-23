#include <iostream>
#include "Queue.h" // Assuming the queue implementation is in queue.h
using namespace std;

void testQueue() {
    Queue q;

    // Test: Enqueue elements
    q.enqueue(10);
    q.enqueue(20);
    q.enqueue(30);
    cout << "Peek after enqueues: " << q.peek() << " (Expected: 10)" << endl;

    // Test: Dequeue elements
    q.dequeue();
    cout << "Peek after first dequeue: " << q.peek() << " (Expected: 20)" << endl;

    q.dequeue();
    cout << "Peek after second dequeue: " << q.peek() << " (Expected: 30)" << endl;

    q.dequeue();
    cout << "Is queue empty after all dequeues? " << (q.isEmpty() ? "Yes" : "No") << " (Expected: Yes)" << endl;

    // Test: Dequeue from empty queue
    q.dequeue(); // Should print error message

    // Test: Peek from empty queue
    cout << "Peek from empty queue: " << q.peek() << " (Expected: -1)" << endl;
}

int main() {
    testQueue();
    return 0;
}
