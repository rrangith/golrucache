package doublylinkedlist

import "errors"

// Node struct has a key and val of any type, with access to the previous and next nodes
type Node struct {
	key  interface{}
	val  interface{}
	next *Node
	prev *Node
}

// MakeNode creates a node
func MakeNode(key, val interface{}, next, prev *Node) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: next,
		prev: prev,
	}
}

// GetKey will return this node's key
func (n *Node) GetKey() interface{} {
	return n.key
}

// GetVal will return this node's val
func (n *Node) GetVal() interface{} {
	return n.val
}

// SetVal will set this node's val to the one passed in
func (n *Node) SetVal(val interface{}) error {
	if val == nil {
		return errors.New("val can't be nil")
	}
	n.val = val
	return nil
}

// SetKey will set this node's key to the one passed in
func (n *Node) SetKey(key interface{}) error {
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

	newNode := MakeNode(key, val, n.next, n)
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

	newNode := MakeNode(key, val, n, n.prev)
	n.setPrev(newNode)
	return nil
}
