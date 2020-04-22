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
	d := MakeDoublyLinkedListVal("hi")

	if (d.head.GetVal() != "hi") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.head.GetVal(), "hi")
	}

	if (d.tail.GetVal() != "hi") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.tail.GetVal(), "hi")
	}

	if (d.size != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.size, 1)
	}
}

func TestGetSize(t *testing.T) {
	d := MakeDoublyLinkedListVal("hi")

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestGetHead(t *testing.T) {
	d := MakeDoublyLinkedListVal("hi")

	head := d.GetHead()

	if (head.GetVal() != "hi") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", head.GetVal(), "hi")
	}
}

func TestInsertFront(t *testing.T) {
	n1 := node.MakeNode("hi", nil, nil)
	n2 := node.MakeNode(1, nil, nil)

	d := MakeDoublyLinkedList()

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertFront(n2)

	if (d.head != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.tail != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertFront(n1)

	if (d.head != n1) {
		t.Errorf("Head should be n1, but it is not")
	}

	if (d.GetHead().GetNext() != n2) {
		t.Errorf("Head's next should be n2, but it is not")
	}

	if (d.GetSize() != 2) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertFrontVal("2")

	if (d.GetHead().GetVal() != "2") {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().GetVal(), "hi")
	}

	if (d.GetSize() != 3) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}

func TestGetTail(t *testing.T) {
	n1 := node.MakeNode("hi", nil, nil)
	
	d := MakeDoublyLinkedList()
	d.InsertBack(n1)

	tail := d.GetTail()

	if (tail != n1) {
		t.Errorf("Tail should be n1, but it is not")
	}
}

func TestInsertBack(t *testing.T) {
	n1 := node.MakeNode("hi", nil, nil)
	n2 := node.MakeNode(1, nil, nil)

	d := MakeDoublyLinkedList()

	if (d.GetSize() != 0) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertBack(n2)

	if (d.head != n2) {
		t.Errorf("Head should be n2, but it is not")
	}

	if (d.tail != n2) {
		t.Errorf("Tail should be n2, but it is not")
	}

	if (d.GetSize() != 1) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertBack(n1)

	if (d.tail != n1) {
		t.Errorf("Tail should be n1, but it is not")
	}

	if (d.GetTail().GetPrev() != n2) {
		t.Errorf("Tail's next should be n2, but it is not")
	}

	if (d.GetSize() != 2) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertBackVal("2")

	if (d.GetTail().GetVal() != "2") {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().GetVal(), "hi")
	}

	if (d.GetSize() != 3) {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}
