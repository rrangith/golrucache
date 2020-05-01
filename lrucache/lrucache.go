package lrucache

import (
	"errors"
	"sync"

	"github.com/rrangith/golrucache/doublylinkedlist"
)

type LRUCache struct {
	cache map[interface{}]*doublylinkedlist.Node
	list  *doublylinkedlist.DoublyLinkedList
	cap   int
	sync.Mutex
}

func MakeLRUCache(cap int) (*LRUCache, error) {
	if cap <= 0 {
		return nil, errors.New("cap must be greater than 0")
	}

	return &LRUCache{
		cache: make(map[interface{}]*doublylinkedlist.Node, cap),
		list:  doublylinkedlist.MakeDoublyLinkedList(),
		cap:   cap,
	}, nil
}

// Need a lock because a write could be taking place, which might alter the size
func (l *LRUCache) GetSize() int {
	l.Lock()
	defer l.Unlock()
	return l.list.GetSize()
}

// Don't need a lock here since cap is fixed
func (l *LRUCache) GetCap() int {
	return l.cap
}

// If the key is found, move the node to the front
func (l *LRUCache) Get(key interface{}) interface{} {
	l.Lock()
	defer l.Unlock()

	n, found := l.cache[key]

	if found {
		l.list.MoveToFront(n)
		return n.GetVal()
	} else {
		return nil
	}
}

func (l *LRUCache) Set(key, val interface{}) error {
	if key == nil || val == nil {
		return errors.New("key and val must not be nil")
	}
	l.Lock()
	defer l.Unlock()

	oldNode, found := l.cache[key]

	if found {
		l.list.MoveToFront(oldNode)
		oldNode.SetVal(val)
	} else {
		newNode := doublylinkedlist.MakeNode(key, val, nil, nil)
		if l.list.GetSize() >= l.GetCap() { // cache is full, need to use list.GetSize since local GetSize locks
			nodeToRemove := l.list.GetTail()
			delete(l.cache, nodeToRemove.GetKey())
			l.list.RemoveBack() // this will update the list's size
		}

		l.cache[key] = newNode
		l.list.InsertFront(newNode) // this will update the list's size
	}
	return nil
}
