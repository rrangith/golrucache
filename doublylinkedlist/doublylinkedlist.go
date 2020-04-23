package doublylinkedlist

import "errors"

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func MakeDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func MakeDoublyLinkedListVal(key, val interface{}) (*DoublyLinkedList, error) {
	if key == nil || val == nil {
		return nil, errors.New("key and val parameters must not be nil")
	}
	newNode := MakeNode(key, val, nil, nil)
	return &DoublyLinkedList{
		head: newNode,
		tail: newNode,
		size: 1,
	}, nil
}

func (d *DoublyLinkedList) GetSize() int {
	return d.size
}

func (d *DoublyLinkedList) GetHead() *Node {
	return d.head
}

func (d *DoublyLinkedList) RemoveNode(n *Node) error { // must assume that this node is in the linked list
	if n == nil {
		return errors.New("Can not pass in nil")
	}

	if n == d.head {
		d.RemoveFront()
	} else if n == d.tail {
		d.RemoveBack()
	} else {
		d.size--
		n.RemoveNode()
	}

	return nil
}

func (d *DoublyLinkedList) MoveToFront(n *Node) error { // must assume that this node is in the linked list
	if n == nil {
		return errors.New("Can not pass in nil")
	}

	d.RemoveNode(n)
	d.InsertFront(n)

	return nil
}

func (d *DoublyLinkedList) InsertFrontVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	if d.head != nil {
		d.head.SetPrevVal(key, val)
		d.head = d.head.GetPrev()
	} else {
		newNode := MakeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}

	d.size += 1
	return nil
}

func (d *DoublyLinkedList) InsertFront(n *Node) error {
	if n == nil {
		return errors.New("Can not pass in nil")
	}

	if d.head != nil {
		d.head.SetPrev(n)
	} else {
		d.tail = n
	}

	d.head = n
	d.size += 1

	return nil
}

func (d *DoublyLinkedList) GetTail() *Node {
	return d.tail
}

func (d *DoublyLinkedList) InsertBackVal(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val parameters must not be nil")
	}

	if d.tail != nil {
		d.tail.SetNextVal(key, val)
		d.tail = d.tail.GetNext()
	} else {
		newNode := MakeNode(key, val, nil, nil)
		d.head = newNode
		d.tail = newNode
	}

	d.size += 1
	return nil
}

func (d *DoublyLinkedList) InsertBack(n *Node) error {
	if n == nil {
		return errors.New("Can not pass in nil")
	}

	if d.tail != nil {
		d.tail.SetNext(n)
	} else {
		d.head = n
	}

	d.tail = n
	d.size += 1

	return nil
}

func (d *DoublyLinkedList) RemoveFront() error {
	oldHead := d.GetHead()
	if oldHead != nil {
		if oldHead == d.GetTail() {
			d.tail = nil
			d.head = nil
		} else {
			newHead := oldHead.GetNext()
			newHead.SetPrev(nil)
			d.head = newHead
		}

		d.size -= 1
		return nil
	} else {
		return errors.New("Head is nil, so can't remove it")
	}
}

func (d *DoublyLinkedList) RemoveBack() error {
	oldTail := d.GetTail()
	if oldTail != nil {
		if oldTail == d.GetHead() {
			d.tail = nil
			d.head = nil
		} else {
			newTail := oldTail.GetPrev()
			newTail.SetNext(nil)
			d.tail = newTail
		}

		d.size -= 1
		return nil
	} else {
		return errors.New("Tail is nil, so can't remove it")
	}
}
