package doublylinkedlist

import "testing"

func TestMake(t *testing.T) {
	d := MakeDoublyLinkedList()

	if d.head != nil {
		t.Errorf("Head should be nil, but it is not")
	}

	if d.tail != nil {
		t.Errorf("Tail should be nil, but it is not")
	}

	if d.size != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.size, 0)
	}
}

func TestMakeVal(t *testing.T) {
	d, err := MakeDoublyLinkedListVal("key", "hi")

	if err != nil {
		t.Errorf(err.Error())
	}

	if d.head.getKey() != "key" {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.head.getKey(), "key")
	}

	if d.head.getVal() != "hi" {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.head.getVal(), "hi")
	}

	if d.tail.getKey() != "key" {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.tail.getKey(), "key")
	}

	if d.tail.getVal() != "hi" {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.tail.getVal(), "hi")
	}

	if d.size != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.size, 1)
	}
}

func TestMakeValNil(t *testing.T) {
	d, err := MakeDoublyLinkedListVal(nil, "hi")

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d != nil {
		t.Errorf("list should be nil")
	}

	d, err = MakeDoublyLinkedListVal("key", nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d != nil {
		t.Errorf("list should be nil")
	}

	d, err = MakeDoublyLinkedListVal(nil, nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d != nil {
		t.Errorf("list should be nil")
	}
}

func TestGetSize(t *testing.T) {
	d, err := MakeDoublyLinkedListVal("key", "hi")

	if err != nil {
		t.Errorf(err.Error())
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestGetHead(t *testing.T) {
	d, err := MakeDoublyLinkedListVal("key", "hi")

	if err != nil {
		t.Errorf(err.Error())
	}

	head := d.GetHead()

	if head.getKey() != "key" {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", head.getKey(), "key")
	}

	if head.getVal() != "hi" {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", head.getVal(), "hi")
	}
}

func TestRemoveNode(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)
	n2 := makeNode("key2", "hi", nil, nil)
	n3 := makeNode("key3", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)
	d.InsertBack(n2)
	d.InsertBack(n3)

	d.RemoveNode(n2)

	if d.GetSize() != 2 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

}

func TestRemoveNodeHead(t *testing.T) { //TODO MORE TESTS FOR THIS
	n1 := makeNode("key1", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertFront(n1)

	d.RemoveNode(n1)

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveNodeNil(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.RemoveNode(nil)

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveNodeTail(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertBackVal("key2", "h2")

	n1 := makeNode("key1", "hi", nil, nil)
	d.InsertBack(n1)

	d.RemoveNode(n1)

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	if d.GetTail().getKey() != "key2" {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().getVal(), "key2")
	}
}

func TestMoveToFront(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertBackVal("key2", "h2")

	n1 := makeNode("key1", "hi", nil, nil)
	d.InsertBack(n1)

	err := d.MoveToFront(n1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if d.GetSize() != 2 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	if d.GetTail().getKey() != "key2" {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().getKey(), "key2")
	}

	if d.GetTail().getNext() != nil {
		t.Errorf("Tail next should point to nil, but it does not")
	}

	if d.GetTail().getPrev() != n1 {
		t.Errorf("Tail prev should point to n1, but it does not")
	}

	if d.GetHead() != n1 {
		t.Errorf("Head should be n1, but it is not")
	}

	if d.GetHead().getNext().getKey() != "key2" {
		t.Errorf("Head next key was incorrect, got: %s, want: %s.", d.GetHead().getNext().getKey(), "key2")
	}

	if d.GetHead().getPrev() != nil {
		t.Errorf("Head prev should point to nil, but it does not")
	}
}

func TestMoveToFrontNil(t *testing.T) {
	d := MakeDoublyLinkedList()
	err := d.MoveToFront(nil)

	if err == nil {
		t.Errorf("An error should have occurred")
	}
}

func TestInsertFront(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)
	n2 := makeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertFront(n2)

	if d.GetHead() != n2 {
		t.Errorf("Head should be n2, but it is not")
	}

	if d.GetTail() != n2 {
		t.Errorf("Tail should be n2, but it is not")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertFront(n1)

	if d.GetHead() != n1 {
		t.Errorf("Head should be n1, but it is not")
	}

	if d.GetHead().getNext() != n2 {
		t.Errorf("Head's next should be n2, but it is not")
	}

	if d.GetSize() != 2 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertFrontVal("key3", "2")

	if d.GetHead().getKey() != "key3" {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().getKey(), "key3")
	}

	if d.GetHead().getVal() != "2" {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().getVal(), "hi")
	}

	if d.GetSize() != 3 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}

func TestInsertFrontNil(t *testing.T) {
	d := MakeDoublyLinkedList()

	err := d.InsertFront(nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}
}

func TestInsertFrontVal(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertFrontVal("key3", "2")

	if d.GetHead().getKey() != "key3" {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().getKey(), "key3")
	}

	if d.GetHead().getVal() != "2" {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().getVal(), "hi")
	}

	if d.GetTail().getKey() != "key3" {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().getKey(), "key3")
	}

	if d.GetTail().getVal() != "2" {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().getVal(), "hi")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestInsertFrontValNil(t *testing.T) {
	d := MakeDoublyLinkedList()

	err := d.InsertFrontVal(nil, "2")

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	err = d.InsertFrontVal("key", nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	err = d.InsertFrontVal(nil, nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestGetTail(t *testing.T) {
	n1 := makeNode("key", "hi", nil, nil)

	d := MakeDoublyLinkedList()
	d.InsertBack(n1)

	if d.GetTail() != n1 {
		t.Errorf("Tail should be n1, but it is not")
	}
}

func TestInsertBack(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)
	n2 := makeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	d.InsertBack(n2)

	if d.GetHead() != n2 {
		t.Errorf("Head should be n2, but it is not")
	}

	if d.GetTail() != n2 {
		t.Errorf("Tail should be n2, but it is not")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	d.InsertBack(n1)

	if d.GetTail() != n1 {
		t.Errorf("Tail should be n1, but it is not")
	}

	if d.GetTail().getPrev() != n2 {
		t.Errorf("Tail's next should be n2, but it is not")
	}

	if d.GetSize() != 2 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 2)
	}

	d.InsertBackVal("key3", "2")

	if d.GetTail().getKey() != "key3" {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().getKey(), "key3")
	}

	if d.GetTail().getVal() != "2" {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().getVal(), "hi")
	}

	if d.GetSize() != 3 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 3)
	}
}

func TestInsertBackNil(t *testing.T) {
	d := MakeDoublyLinkedList()

	err := d.InsertBack(nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}
}

func TestInsertBackVal(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.InsertBackVal("key3", "2")

	if d.GetHead().getKey() != "key3" {
		t.Errorf("Head key was incorrect, got: %s, want: %s.", d.GetHead().getKey(), "key3")
	}

	if d.GetHead().getVal() != "2" {
		t.Errorf("Head val was incorrect, got: %s, want: %s.", d.GetHead().getVal(), "hi")
	}

	if d.GetTail().getKey() != "key3" {
		t.Errorf("Tail key was incorrect, got: %s, want: %s.", d.GetTail().getKey(), "key3")
	}

	if d.GetTail().getVal() != "2" {
		t.Errorf("Tail val was incorrect, got: %s, want: %s.", d.GetTail().getVal(), "hi")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}
}

func TestInsertBackValNil(t *testing.T) {
	d := MakeDoublyLinkedList()

	err := d.InsertBackVal(nil, "2")

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	err = d.InsertBackVal("key", nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}

	err = d.InsertBackVal(nil, nil)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveFront(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)
	n2 := makeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)
	d.InsertBack(n2)

	d.RemoveFront()

	if d.GetHead() != n2 {
		t.Errorf("Head should be n2, but it is not")
	}

	if d.GetTail() != n2 {
		t.Errorf("Tail should be n2, but it is not")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	if n2.getNext() != nil {
		t.Errorf("n2 next should point to nil")
	}

	if n2.getPrev() != nil {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestRemoveFrontOnlyOne(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)

	d.RemoveFront()

	if d.GetHead() != nil {
		t.Errorf("Head should be nil, but it is not")
	}

	if d.GetTail() != nil {
		t.Errorf("Tail should be nil, but it is not")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveFrontEmpty(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.RemoveFront()

	if d.GetHead() != nil {
		t.Errorf("Head should be nil, but it is not")
	}

	if d.GetTail() != nil {
		t.Errorf("Tail should be nil, but it is not")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveBack(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)
	n2 := makeNode("key2", 1, nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertFront(n1)
	d.InsertFront(n2)

	d.RemoveBack()

	if d.GetHead() != n2 {
		t.Errorf("Head should be n2, but it is not")
	}

	if d.GetTail() != n2 {
		t.Errorf("Tail should be n2, but it is not")
	}

	if d.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 1)
	}

	if n2.getNext() != nil {
		t.Errorf("n2 next should point to nil")
	}

	if n2.getPrev() != nil {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestRemoveBackOnlyOne(t *testing.T) {
	n1 := makeNode("key1", "hi", nil, nil)

	d := MakeDoublyLinkedList()

	d.InsertBack(n1)

	d.RemoveBack()

	if d.GetHead() != nil {
		t.Errorf("Head should be nil, but it is not")
	}

	if d.GetTail() != nil {
		t.Errorf("Tail should be nil, but it is not")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}

func TestRemoveBackEmpty(t *testing.T) {
	d := MakeDoublyLinkedList()

	d.RemoveBack()

	if d.GetHead() != nil {
		t.Errorf("Head should be nil, but it is not")
	}

	if d.GetTail() != nil {
		t.Errorf("Tail should be nil, but it is not")
	}

	if d.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", d.GetSize(), 0)
	}
}
