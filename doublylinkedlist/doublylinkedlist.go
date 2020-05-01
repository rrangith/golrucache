package doublylinkedlist

import "errors"

// DoublyLinkedList is the structure that keeps track of its head and tail along with size, made up of nodes with access to prev and next nodes
type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

// MakeDoublyLinkedList is the function to create a new one
func MakeDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// MakeDoublyLinkedListVal creates a single node doubly linkedlist with the key and val passed in
func MakeDoublyLinkedListVal(key, val interface{}) (*DoublyLinkedList, error) {
	if key == nil || val == nil {
		return nil, errors.New("key and val parameters must not be nil")
	}
	newNode := makeNode(key, val, nil, nil)
	return &DoublyLinkedList{
		head: newNode,
		tail: newNode,
		size: 1,
	}, nil
}

// GetSize returns the current size of the doubly linked list
func (d *DoublyLinkedList) GetSize() int {
	return d.size
}

// GetHead returns the current head of the doubly linked list
func (d *DoublyLinkedList) GetHead() *node {
	return d.head
}

// GetTail returns the current tail of the doubly linked list
func (d *DoublyLinkedList) GetTail() *node {
	return d.tail
}

// RemoveNode removes the node passed in from the doubly linked list, must assume this node passed in is in the list
func (d *DoublyLinkedList) RemoveNode(n *node) error {
	if n == nil {
		return errors.New("Can't pass in nil")
	}

	if n == d.head {
		d.RemoveFront()
	} else if n == d.tail {
		d.RemoveBack()
	} else {
		d.size--
		n.removeNode()
	}

	return nil
}

// MoveToFront moves the node passed in to the front of the list, must assume this node passed in is in the list
func (d *DoublyLinkedList) MoveToFront(n *node) error {
	if n == nil {
		return errors.New("Can't pass in nil")
	}

	d.RemoveNode(n)
	d.InsertFront(n)

	return nil
}

// InsertFrontVal will insert a node to the front of the list with the given key and val
func (d *DoublyLinkedList) InsertFrontVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	if d.head != nil {
		d.head.setPrevVal(key, val)
		d.head = d.head.getPrev()
	} else {
		newNode := makeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}

	d.size += 1
	return nil
}

// InsertFront inserts the passed in node to the front of the list
func (d *DoublyLinkedList) InsertFront(n *node) error {
	if n == nil {
		return errors.New("Can't pass in nil")
	}

	if d.head != nil {
		d.head.setPrev(n)
	} else {
		d.tail = n
	}

	d.head = n
	d.size += 1

	return nil
}

// InsertBackVal inserts a node to the back of the list with the given key and val
func (d *DoublyLinkedList) InsertBackVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	if d.tail != nil {
		d.tail.setNextVal(key, val)
		d.tail = d.tail.getNext()
	} else {
		newNode := makeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}

	d.size += 1
	return nil
}

// InsertBack inserts the passed in node to the back of the list
func (d *DoublyLinkedList) InsertBack(n *node) error {
	if n == nil {
		return errors.New("Can't pass in nil")
	}

	if d.tail != nil {
		d.tail.setNext(n)
	} else {
		d.head = n
	}

	d.tail = n
	d.size += 1

	return nil
}

// RemoveFront removes the node that is in the front of the list
func (d *DoublyLinkedList) RemoveFront() error {
	oldHead := d.GetHead()
	if oldHead != nil {
		if oldHead == d.GetTail() {
			d.tail = nil
			d.head = nil
		} else {
			newHead := oldHead.getNext()
			newHead.setPrev(nil)
			d.head = newHead
		}

		d.size -= 1
		return nil
	}

	return errors.New("Head is nil, so can't remove it")
}

// RemoveBack removes the node at the back the list
func (d *DoublyLinkedList) RemoveBack() error {
	oldTail := d.GetTail()
	if oldTail != nil {
		if oldTail == d.GetHead() {
			d.tail = nil
			d.head = nil
		} else {
			newTail := oldTail.getPrev()
			newTail.setNext(nil)
			d.tail = newTail
		}

		d.size -= 1
		return nil
	}

	return errors.New("Tail is nil, so can't remove it")
}
