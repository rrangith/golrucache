package node

import "testing"

func TestMake(t *testing.T) {
	n := MakeNode(1, 1, nil, nil)

	if (n.key != 1) {
		t.Errorf("Key was incorrect, got: %d, want: %d.", n.key, 1)
	}

	if (n.val != 1) {
		t.Errorf("Val was incorrect, got: %d, want: %d.", n.val, 1)
	}

	if (n.next != nil) {
		t.Errorf("Next should be nil, but it is not")
	}

	if (n.prev != nil) {
		t.Errorf("Prev should be nil, but it is not")
	}
}

func TestMakeString(t *testing.T) {
	n := MakeNode("key", "hi", nil, nil)

	if (n.key != "key") {
		t.Errorf("Key was incorrect, got: %s, want: %s.", n.key, "key")
	}

	if (n.val != "hi") {
		t.Errorf("Val was incorrect, got: %s, want: %s.", n.val, "hi")
	}
}

func TestGetKey(t *testing.T) {
	n := MakeNode("key", "hi", nil, nil)

	if (n.GetKey() != "key") {
		t.Errorf("Key was incorrect, got: %s, want: %s.", n.GetKey(), "key")
	}
}

func TestGetVal(t *testing.T) {
	n := MakeNode("key", "hi", nil, nil)

	if (n.GetVal() != "hi") {
		t.Errorf("Val was incorrect, got: %s, want: %s.", n.GetVal(), "hi")
	}
}

func TestSetVal(t *testing.T) {
	n := MakeNode("key", "hi", nil, nil)

	n.SetVal(1)

	if (n.GetVal() != 1) {
		t.Errorf("Val was incorrect, got: %d, want: %d.", n.GetVal(), 1)
	}
}

func TestSetKey(t *testing.T) {
	n := MakeNode("key", "hi", nil, nil)

	n.SetKey("Newkey")

	if (n.GetKey() != "Newkey") {
		t.Errorf("Key was incorrect, got: %s, want: %s.", n.GetKey(), "Newkey")
	}
}

func TestRemove(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)
	n3 := MakeNode("key3", 3, nil, nil)

	n1.SetNext(n2)
	n2.SetNext(n3)

	n2.RemoveNode()

	if (n1.next != n3) {
		t.Errorf("n1 next should point to n3")
	}

	if (n3.prev != n1) {
		t.Errorf("n3 prev should point to n1")
	}

	if (n2.next != nil) {
		t.Errorf("n2 next should point to nil")
	}

	if (n2.prev != nil) {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestRemoveLast(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)
	n2.RemoveNode()

	if (n1.next != nil) {
		t.Errorf("n1 next should point to nil")
	}
}

func TestRemoveFirst(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)
	n1.RemoveNode()

	if (n2.prev != nil) {
		t.Errorf("n2 prev should point to nil")
	}
}

func TestSetNext(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)
	n2.SetNextVal("key3", 3)

	n3 := n2.next

	if (n1.next != n2) {
		t.Errorf("n1 next should point to n2")
	}

	if (n1.prev != nil) {
		t.Errorf("n1 prev should point to nil")
	}

	if (n2.next != n3) {
		t.Errorf("n2 next should point to n3")
	}

	if (n2.prev != n1) {
		t.Errorf("n2 prev should point to n1")
	}

	if (n3.next != nil) {
		t.Errorf("n3 next should point to nil")
	}

	if (n3.prev != n2) {
		t.Errorf("n3 prev should point to n2")
	}
}

func TestSetNextMiddle(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)
	n2.SetNextVal("key3", 3)
	n3 := n2.next

	n2.SetNextVal("keymiddle", "middle")

	nMid := n2.next

	if (n2.next != nMid) {
		t.Errorf("n2 next should point to nMid")
	}

	if (nMid.prev != n2) {
		t.Errorf("nMid prev should point to n2")
	}

	if (nMid.next != n3) {
		t.Errorf("nMid next should point to n3")
	}

	if (n3.prev != nMid) {
		t.Errorf("n3 prev should point to nMid")
	}
}

func TestSetNextNil(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n1.SetNext(nil)

	if (n1.next != nil) {
		t.Errorf("n1 next should point to nil")
	}

	if (n1.prev != nil) {
		t.Errorf("n1 prev should point to nil")
	}
}

func TestGetNext(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)

	if (n1.GetNext() != n2) {
		t.Errorf("n1 next should point to n2")
	}
}

func TestSetPrev(t *testing.T) {
	n3 := MakeNode("key3", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n3.SetPrev(n2)
	n2.SetPrevVal("key1", 3)

	n1 := n2.prev

	if (n1.next != n2) {
		t.Errorf("n1 next should point to n2")
	}

	if (n1.prev != nil) {
		t.Errorf("n1 prev should point to nil")
	}

	if (n2.next != n3) {
		t.Errorf("n2 next should point to n3")
	}

	if (n2.prev != n1) {
		t.Errorf("n2 prev should point to n1")
	}

	if (n3.next != nil) {
		t.Errorf("n3 next should point to n3")
	}

	if (n3.prev != n2) {
		t.Errorf("n3 prev should point to n2")
	}
}

func TestSetPrevMiddle(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)
	n2.SetNextVal("key3", 3)
	n3 := n2.next

	n3.SetPrevVal("keymiddle", "middle")

	nMid := n3.prev

	if (n2.next != nMid) {
		t.Errorf("n2 next should point to nMid")
	}

	if (nMid.prev != n2) {
		t.Errorf("nMid prev should point to n2")
	}

	if (nMid.next != n3) {
		t.Errorf("nMid next should point to n3")
	}

	if (n3.prev != nMid) {
		t.Errorf("n3 prev should point to nMid")
	}
}

func TestSetPrevNil(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n1.SetPrev(nil)

	if (n1.next != nil) {
		t.Errorf("n1 next should point to nil")
	}

	if (n1.prev != nil) {
		t.Errorf("n1 prev should point to nil")
	}
}

func TestGetPrev(t *testing.T) {
	n1 := MakeNode("key1", "hi", nil, nil)
	n2 := MakeNode("key2", 1, nil, nil)

	n1.SetNext(n2)

	if (n2.GetPrev() != n1) {
		t.Errorf("n2 prev should point to n1")
	}
}