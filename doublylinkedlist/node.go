package doublylinkedlist

import "errors"

// Node struct has a key and val of any type, with access to the previous and next nodes
type Node struct {
	key  interface{}
	val  interface{}
	next *Node
	prev *Node
}

func makeNode(key, val interface{}, next, prev *Node) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: next,
		prev: prev,
	}
}

func (n *Node) getKey() interface{} {
	return n.key
}

func (n *Node) getVal() interface{} {
	return n.val
}

func (n *Node) setVal(val interface{}) error {
	if val == nil {
		return errors.New("val can't be nil")
	}
	n.val = val
	return nil
}

func (n *Node) setKey(key interface{}) error {
	if key == nil {
		return errors.New("key can't be nil")
	}
	n.key = key
	return nil
}

func (n *Node) removeNode() {
	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	n.next = nil
	n.prev = nil
}

func (n *Node) getNext() *Node {
	return n.next
}

func (n *Node) setNext(newNode *Node) {
	if n.next != nil {
		n.next.prev = newNode
	}

	if newNode != nil {
		newNode.next = n.next
		newNode.prev = n
	}

	n.next = newNode
}

func (n *Node) setNextVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	newNode := makeNode(key, val, n.next, n)
	n.setNext(newNode)
	return nil
}

func (n *Node) getPrev() *Node {
	return n.prev
}

func (n *Node) setPrev(newNode *Node) {
	if n.prev != nil {
		n.prev.next = newNode
	}

	if newNode != nil {
		newNode.next = n
		newNode.prev = n.prev
	}

	n.prev = newNode
}

func (n *Node) setPrevVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	newNode := makeNode(key, val, n, n.prev)
	n.setPrev(newNode)
	return nil
}
