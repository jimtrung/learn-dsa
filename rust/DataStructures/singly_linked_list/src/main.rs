use std::rc::Rc;
use std::cell::RefCell;

#[derive(Clone)]
struct Node {
    data: i64,
    next: Option<Rc<RefCell<Node>>>
}

impl Node {
    fn new(value: i64) -> Self {
        Node { data: value, next: None }
    }
}

#[derive(Clone)]
struct LinkedList {
    head: Option<Rc<RefCell<Node>>>,
    tail: Option<Rc<RefCell<Node>>>,
}

impl LinkedList {
    fn new() -> LinkedList {
        LinkedList { head: None, tail: None }
    }

    fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    fn push_front(&mut self, value: i64) {
        let new_node = Rc::new(RefCell::new(Node::new(value)));
        if self.is_empty() {
            self.tail = Some(new_node.clone());
            self.head = Some(new_node);
        } else {

        }

    }
}

fn main() {
    let list = LinkedList::new();
}
