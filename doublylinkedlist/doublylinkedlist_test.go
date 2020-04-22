package doublylinkedlist

import (
	"testing"
	"golrucache/node"
)

func TestMake(t *testing.T) {
	d := MakeDoublyLinkedList()

	if (d.head != nil) {
		t.Errorf("Head should be nil, but it is not")
	}

	if (d.tail != nil) {
		t.Errorf("Tail should be nil, but it is not")
	}

	if (d.size != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.size, 0)
	}
}

func TestMakeVal(t *testing.T) {
	d := MakeDoublyLinkedListVal("key", "hi")

	if (d.head.GetKey() != "key") {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.head.GetKey(), "key")
	}

	if (d.head.GetVal() != "hi") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.head.GetVal(), "hi")
	}

	if (d.tail.GetKey() != "key") {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.tail.GetKey(), "key")
	}

	if (d.tail.GetVal() != "hi") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.tail.GetVal(), "hi")
	}

	if (d.size != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.size, 1)
	}
}

func TestGetSize(t *testing.T) {
	d := MakeDoublyLinkedListVal("key", "hi")

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestGetHead(t *testing.T) {
	d := MakeDoublyLinkedListVal("key", "hi")

	head := d.GetHead()

	if (head.GetKey() != "key") {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", head.GetKey(), "key")
	}

	if (head.GetVal() != "hi") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", head.GetVal(), "hi")
	}
}

func TestInsertFront(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)
	n2 := node.MakeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertFront(n2)

	if (d.GetHead() != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.GetTail() != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertFront(n1)

	if (d.GetHead() != n1) {
		t.Errorf("Head should be n1, but it is not")
	}

	if (d.GetHead().GetNext() != n2) {
		t.Errorf("Head's next should be n2, but it is not")
	}

	if (d.GetSize() != 2) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertFrontVal("key3", "2")

	if (d.GetHead().GetKey() != "key3") {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().GetKey(), "key3")
	}

	if (d.GetHead().GetVal() != "2") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().GetVal(), "hi")
	}

	if (d.GetSize() != 3) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}

func TestInsertFrontVal(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertFrontVal("key3", "2")

	if (d.GetHead().GetKey() != "key3") {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().GetKey(), "key3")
	}

	if (d.GetHead().GetVal() != "2") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().GetVal(), "hi")
	}

	if (d.GetTail().GetKey() != "key3") {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().GetKey(), "key3")
	}

	if (d.GetTail().GetVal() != "2") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().GetVal(), "hi")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestGetTail(t *testing.T) {
	n1 := node.MakeNode("key", "hi", nil, nil)
	
	d := MakeDoublyLinkedList()
	d.InsertBack(n1)

	if (d.GetTail() != n1) {
		t.Errorf("Tail should be n1, but it is not")
	}
}

func TestInsertBack(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)
	n2 := node.MakeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertBack(n2)

	if (d.GetHead() != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.GetTail() != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertBack(n1)

	if (d.GetTail() != n1) {
		t.Errorf("Tail should be n1, but it is not")
	}

	if (d.GetTail().GetPrev() != n2) {
		t.Errorf("Tail's next should be n2, but it is not")
	}

	if (d.GetSize() != 2) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertBackVal("key3", "2")

	if (d.GetTail().GetKey() != "key3") {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().GetKey(), "key3")
	}

	if (d.GetTail().GetVal() != "2") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().GetVal(), "hi")
	}

	if (d.GetSize() != 3) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}

func TestInsertBackVal(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertBackVal("key3", "2")

	if (d.GetHead().GetKey() != "key3") {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().GetKey(), "key3")
	}

	if (d.GetHead().GetVal() != "2") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().GetVal(), "hi")
	}

	if (d.GetTail().GetKey() != "key3") {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().GetKey(), "key3")
	}

	if (d.GetTail().GetVal() != "2") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().GetVal(), "hi")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestRemoveFront(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)
	n2 := node.MakeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)
	d.InsertBack(n2)

	d.RemoveFront()

	if (d.GetHead() != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.GetTail() != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	if (n2.GetNext() != nil) {
		t.Errorf("n2 next should point to nil")
	}

	if (n2.GetPrev() != nil) {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestRemoveFrontOnlyOne(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)

	d.RemoveFront()

	if (d.GetHead() != nil) {
		t.Errorf("Head should be nil, but it is not")
	}

	if (d.GetTail() != nil) {
		t.Errorf("Tail should be nil, but it is not")
	}

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveFrontEmpty(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.RemoveFront()

	if (d.GetHead() != nil) {
		t.Errorf("Head should be nil, but it is not")
	}

	if (d.GetTail() != nil) {
		t.Errorf("Tail should be nil, but it is not")
	}

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveBack(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)
	n2 := node.MakeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertFront(n1)
	d.InsertFront(n2)

	d.RemoveBack()

	if (d.GetHead() != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.GetTail() != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	if (n2.GetNext() != nil) {
		t.Errorf("n2 next should point to nil")
	}

	if (n2.GetPrev() != nil) {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestRemoveBackOnlyOne(t *testing.T) {
	n1 := node.MakeNode("key1", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)

	d.RemoveBack()

	if (d.GetHead() != nil) {
		t.Errorf("Head should be nil, but it is not")
	}

	if (d.GetTail() != nil) {
		t.Errorf("Tail should be nil, but it is not")
	}

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveBackEmpty(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.RemoveBack()

	if (d.GetHead() != nil) {
		t.Errorf("Head should be nil, but it is not")
	}

	if (d.GetTail() != nil) {
		t.Errorf("Tail should be nil, but it is not")
	}

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}
