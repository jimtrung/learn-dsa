#include <iostream>
using namespace std;

struct PhanSo {
    int tuSo;
    int mauSo;
};

typedef PhanSo Data ;

struct Node {
    Data data;
    Node* next;
};

Node* createNode(Data value) {
    Node* newNode = new Node;
    newNode->data = value;
    newNode->next = NULL;
    return newNode;
}

struct Queue {
    Node* head;
    Node* tail;
};

void newQueue(Queue &q) {
    q.head = q.tail = NULL;
}

void enqueue(Queue &q, Data data) {
    Node* newNode = createNode(data);
    if (q.head == NULL) {
        q.head = newNode;
        q.tail = newNode;
    } else {
        q.tail->next = newNode;
        q.tail = newNode;
    }
}

void dequeue(Queue &q) {
    if (q.head == NULL) {
        cout << "Hang doi rong" << endl;
        return;
    }
    if (q.head->next == NULL) {
        delete q.head;
        q.head = q.tail = NULL;
    } else {
        Node* nodeToDelete = q.head;
        q.head = q.head->next;
        delete nodeToDelete;
    }
}

bool peek(Queue &q, Data &ps) {
    if (q.head == NULL) {
        cout << "Hang doi rong" << endl;
        return false;
    }
    ps = q.head->data;
    return true;
}

bool isEmpty(Queue &q) {
    return q.head == NULL;
}

int main() {
    Queue q;
    newQueue(q);

    int N;
    do {
        cout << "Nhap so luong phan so: ";
        cin >> N;
    } while (N <= 0);

    for (int i = 1; i <= N; i++) {
        Data ps;
        cout << "Nhap thong tin phan so thu " << i << endl;
        cout << " - Nhap tu so: "; cin >> ps.tuSo;
        cout << " - Nhap mau so: "; cin >> ps.mauSo;
        enqueue(q, ps);
    }

    float tong = 0;
    cout << "Cac phan so trong hang doi: " << endl;
    while (!isEmpty(q)) {
        Data ps;
        peek(q, ps);
        cout << ps.tuSo << "/" << ps.mauSo << endl;
        tong += float(ps.tuSo) / ps.mauSo;

        dequeue(q);
    }

    cout << "Tong cua cac phan so = " << tong << endl;
}
