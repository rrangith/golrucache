package doublylinkedlist

import "errors"

type node struct {
	key  interface{}
	val  interface{}
	next *node
	prev *node
}

func makeNode(key, val interface{}, next, prev *node) *node {
	return &node{
		key:  key,
		val:  val,
		next: next,
		prev: prev,
	}
}

func (n *node) getKey() interface{} {
	return n.key
}

func (n *node) getVal() interface{} {
	return n.val
}

func (n *node) setVal(val interface{}) error {
	if val == nil {
		return errors.New("val can't be nil")
	}
	n.val = val
	return nil
}

func (n *node) setKey(key interface{}) error {
	if key == nil {
		return errors.New("key can't be nil")
	}
	n.key = key
	return nil
}

func (n *node) removeNode() {
	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	n.next = nil
	n.prev = nil
}

func (n *node) getNext() *node {
	return n.next
}

func (n *node) setNext(newNode *node) {
	if n.next != nil {
		n.next.prev = newNode
	}

	if newNode != nil {
		newNode.next = n.next
		newNode.prev = n
	}

	n.next = newNode
}

func (n *node) setNextVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	newNode := makeNode(key, val, n.next, n)
	n.setNext(newNode)
	return nil
}

func (n *node) getPrev() *node {
	return n.prev
}

func (n *node) setPrev(newNode *node) {
	if n.prev != nil {
		n.prev.next = newNode
	}

	if newNode != nil {
		newNode.next = n
		newNode.prev = n.prev
	}

	n.prev = newNode
}

func (n *node) setPrevVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	newNode := makeNode(key, val, n, n.prev)
	n.setPrev(newNode)
	return nil
}
