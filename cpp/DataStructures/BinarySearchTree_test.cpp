#include <iostream>
#include "BinarySearchTree.h"

void insert(Node* root, int value) {
}

void preOrderTraversal(Node *root) {
  if (root == NULL) {
    return;
  }
  preOrderTraversal(root->left);
  cout << root->data << " ";
  preOrderTraversal(root->right);
}

void inOrderTraversal(Node *root) {
  if (root == NULL) {
    return;
  }
  cout << root->data << " ";
  inOrderTraversal(root->left);
  inOrderTraversal(root->right);
}

void postOrderTraversal(Node *root) {
  if (root == NULL) {
    return;
  }
  postOrderTraversal(root->left);
  postOrderTraversal(root->right);
  cout << root->data << " ";
}

int main() {
    BinarySearchTree* bts = new BinarySearchTree;
}
