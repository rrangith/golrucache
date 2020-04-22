package doublylinkedlist

import "golrucache/node"

type DoublyLinkedList struct {
	head *node.Node
	tail *node.Node
	size int
}

func MakeDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList {
		head: nil,
		tail: nil,
		size: 0,
	}
}

func MakeDoublyLinkedListVal(key, val interface{}) *DoublyLinkedList {
	newNode := node.MakeNode(key, val, nil, nil)
	return &DoublyLinkedList {
		head: newNode,
		tail: newNode,
		size: 1,
	}
}

func (d *DoublyLinkedList) GetSize() int {
	return d.size
}

func (d *DoublyLinkedList) GetHead() *node.Node {
	return d.head
}

func (d *DoublyLinkedList) InsertFrontVal(key, val interface{}) {
	if (d.head != nil) {
		d.head.SetPrevVal(key, val)
		d.head = d.head.GetPrev()
	} else {
		newNode := node.MakeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}

	d.size += 1
}

func (d *DoublyLinkedList) InsertFront(node *node.Node) {
	if (d.head != nil) {
		d.head.SetPrev(node)
	} else {
		d.tail = node
	}

	d.head = node
	d.size += 1
}

func (d *DoublyLinkedList) GetTail() *node.Node {
	return d.tail
}

func (d *DoublyLinkedList) InsertBackVal(key, val interface{}) {
	if (d.tail != nil) {
		d.tail.SetNextVal(key, val)
		d.tail = d.tail.GetNext()
	} else {
		newNode := node.MakeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}
	d.size += 1
}

func (d *DoublyLinkedList) InsertBack(node *node.Node) {
	if (d.tail != nil) {
		d.tail.SetNext(node)
	} else {
		d.head = node
	}
	
	d.tail = node
	d.size += 1
}

func (d* DoublyLinkedList) RemoveFront() {
	oldHead := d.GetHead()
	if (oldHead != nil) {
		if (oldHead == d.GetTail()) {
			d.tail = nil
			d.head = nil
		} else {
			newHead := oldHead.GetNext()
			newHead.SetPrev(nil)
			d.head = newHead
		}

		d.size -= 1
	}
}

func (d* DoublyLinkedList) RemoveBack() {
	oldTail := d.GetTail()
	if (oldTail != nil) {
		if (oldTail == d.GetHead()) {
			d.tail = nil
			d.head = nil
		} else {
			newTail := oldTail.GetPrev()
			newTail.SetNext(nil)
			d.tail = newTail
		}

		d.size -= 1
	}
}
