const assert = require('assert');

class Node {
    constructor(value) {
        this.data = value;
        this.left = null;
        this.right = null;
    }
}

class BinarySearchTree {
    constructor() {
        this.root = null;
    }
}

function preOrderTraversal(root) {
    if (root === null) {
        return
    }
    protArr.push(root.data)
    preOrderTraversal(root.left);
    preOrderTraversal(root.right);
}

function inOrderTraversal(root) {
    if (root === null) {
        return;
    }
    inOrderTraversal(root.left);
    inotArr.push(root.data)
    inOrderTraversal(root.right);
}

function postOrderTraversal(root) {
    if (root === null) {
        return;
    }
    postOrderTraversal(root.left);
    postOrderTraversal(root.right);
    pootArr.push(root.data)
}

function search(root, target) {
    if (root === null) {
        return null;
    } else if (root.data === target) {
        return root
    } else if (target < root.data) {
        return search(root.left, target)
    } else {
        return search(root.right, target)
    }
}

// MAIN AREA //
const bst = new BinarySearchTree();
const nodeR = new Node('R');
const nodeA = new Node('A');
const nodeC = new Node('C');
const nodeD = new Node('D');
const nodeB = new Node('B');
const nodeE = new Node('E');
const nodeF = new Node('F');
const nodeG = new Node('G');

bst.root = nodeR;
bst.root.left = nodeA;
bst.root.left.left = nodeC;
bst.root.left.right = nodeD;

bst.root.right = nodeB;
bst.root.right.left = nodeE;
bst.root.right.right = nodeF;
bst.root.right.right.left = nodeG;

const protArr = [];
preOrderTraversal(bst.root);
assert.deepStrictEqual(protArr, ['R', 'A', 'C', 'D', 'B', 'E', 'F', 'G']);

const inotArr = [];
inOrderTraversal(bst.root);
assert.deepStrictEqual(inotArr, ['C', 'A', 'D', 'R', 'E', 'B', 'G', 'F']);

const pootArr = [];
postOrderTraversal(bst.root);
assert.deepStrictEqual(pootArr, ['C', 'D', 'A', 'E', 'G', 'F', 'B', 'R']);
console.log("Traversal PASSED")

// Test search function
const bstNum = new BinarySearchTree();
const node13 = new Node(13);
const node7 = new Node(7);
const node3 = new Node(3);
const node8 = new Node(8);
const node15 = new Node(15);
const node14 = new Node(14);
const node19 = new Node(19);
const node18 = new Node(18);

bstNum.root = node13;
bstNum.root.left = node7;
bstNum.root.left.left = node3;
bstNum.root.left.right = node8;

bstNum.root.right = node15;
bstNum.root.right.left = node14;
bstNum.root.right.right = node19;
bstNum.root.right.right.left = node18;

assert.deepStrictEqual(node13, search(bstNum.root, 13));
assert.deepStrictEqual(node7, search(bstNum.root, 7));
assert.deepStrictEqual(node19, search(bstNum.root, 19));
assert.deepStrictEqual(node18, search(bstNum.root, 18));

console.log("Search PASSED");
