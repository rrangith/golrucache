package lrucache

import "testing"

func TestNew(t *testing.T) {
	l, err := New(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if l.cap != 1 {
		t.Errorf("Cap was incorrect, got: %d, want: %d.", l.cap, 1)
	}

	if l.list.GetHead() != nil {
		t.Errorf("List head should be nil, but it is not")
	}

	if l.list.GetTail() != nil {
		t.Errorf("List tail should be nil, but it is not")
	}

	if l.list.GetSize() != 0 {
		t.Errorf("List size was incorrect, got: %d, want: %d.", l.list.GetSize(), 0)
	}
}

func TestNewBadCap(t *testing.T) {
	l, err := New(0)

	if err == nil {
		t.Errorf("Error should have occurred")
	}

	if l != nil {
		t.Errorf("List should be nil")
	}
}

func TestGetSize(t *testing.T) {
	l, err := New(2)

	if err != nil {
		t.Errorf(err.Error())
	}

	if l.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 0)
	}

	l.Set("key1", "hi")

	if l.GetSize() != 1 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 1)
	}

	l.Set("key2", "hi")

	if l.GetSize() != 2 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 2)
	}
}

func TestGetCap(t *testing.T) {
	l, err := New(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if l.GetCap() != 1 {
		t.Errorf("Cap was incorrect, got: %d, want: %d.", l.GetCap(), 1)
	}
}

func TestGet(t *testing.T) {
	l, err := New(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	l.Set("key1", "hi")

	if l.Get("key1") != "hi" {
		t.Errorf("Get was incorrect, got: %s, want: %s.", l.Get("key1"), "hi")
	}
}

func TestGetNotFound(t *testing.T) {
	l, err := New(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if l.Get("key1") != nil {
		t.Errorf("Get was incorrect, got: %s, want: nil.", l.Get("key1"))
	}
}

func TestSet(t *testing.T) {
	l, err := New(10)

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 10; i++ {
		l.Set(i, i)
		if l.GetSize() != i+1 {
			t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), i+1)
		}

		if l.Get(i) != i {
			t.Errorf("Get was incorrect, got: %d, want: %d.", l.Get(i), i)
		}
	}

	for i := 2; i < 12; i++ {
		l.Set(i, i+1)

		if l.GetSize() != 10 {
			t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 10)
		}

		if l.Get(i) != i+1 {
			t.Errorf("Get was incorrect, got: %d, want: %d.", l.Get(i), i+1)
		}
	}

	// these 2 should have been evicted
	if l.Get(0) != nil {
		t.Errorf("Get was incorrect, got: %d, want: nil.", l.Get(0))
	}

	if l.Get(1) != nil {
		t.Errorf("Get was incorrect, got: %d, want: nil.", l.Get(1))
	}
}

func TestSetNil(t *testing.T) {
	l, err := New(10)

	if err != nil {
		t.Errorf(err.Error())
	}

	l.Set(nil, "hi")

	if l.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 0)
	}

	l.Set("hi", nil)

	if l.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 0)
	}

	l.Set(nil, nil)

	if l.GetSize() != 0 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 0)
	}
}

func TestRemove(t *testing.T) {
	l, err := New(10)

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 10; i++ {
		l.Set(i, i)
		if l.GetSize() != i+1 {
			t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), i+1)
		}

		if l.Get(i) != i {
			t.Errorf("Get was incorrect, got: %d, want: %d.", l.Get(i), i)
		}
	}

	err = l.Remove(0)

	if err != nil {
		t.Errorf(err.Error())
	}

	if l.GetSize() != 9 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 9)
	}

	_, found := l.cache[0]

	if found {
		t.Errorf("key was not removed from cache")
	}
}

func TestRemoveNil(t *testing.T) {
	l, err := New(10)

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 10; i++ {
		l.Set(i, i)
		if l.GetSize() != i+1 {
			t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), i+1)
		}

		if l.Get(i) != i {
			t.Errorf("Get was incorrect, got: %d, want: %d.", l.Get(i), i)
		}
	}

	err = l.Remove(nil)

	if err == nil {
		t.Errorf("An error should have occurred")
	}

	if l.GetSize() != 10 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 10)
	}
}

func TestRemoveNotFound(t *testing.T) {
	l, err := New(10)

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < 10; i++ {
		l.Set(i, i)
		if l.GetSize() != i+1 {
			t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), i+1)
		}

		if l.Get(i) != i {
			t.Errorf("Get was incorrect, got: %d, want: %d.", l.Get(i), i)
		}
	}

	err = l.Remove("asd")

	if err == nil {
		t.Errorf("An error should have occurred")
	}

	if l.GetSize() != 10 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", l.GetSize(), 10)
	}
}
