#include <iostream>
using namespace std;

struct Node {
  int data;
  Node *left;
  Node *right;
};

class BinarySearchTree {
private:
  Node *root;

public:
  BinarySearchTree() : root(NULL) {}
};

