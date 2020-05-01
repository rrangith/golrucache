package lrucache

import (
	"errors"
	"sync"

	"github.com/rrangith/golrucache/doublylinkedlist"
)

// LRUCache is a struct that evicts the least recently used item if the cache is full, it has a max size and can set and get in O(1), this structure is thread-safe
type LRUCache struct {
	cache map[interface{}]*doublylinkedlist.Node
	list  *doublylinkedlist.DoublyLinkedList
	cap   int
	sync.Mutex
}

// New will create a cache with the given capacity, the capacity must be greater than 0
func New(cap int) (*LRUCache, error) {
	if cap <= 0 {
		return nil, errors.New("cap must be greater than 0")
	}

	return &LRUCache{
		cache: make(map[interface{}]*doublylinkedlist.Node, cap),
		list:  doublylinkedlist.MakeDoublyLinkedList(),
		cap:   cap,
	}, nil
}

// GetSize will return the current size of the lru cache, need a lock because a write could be taking place, which might alter the size
func (l *LRUCache) GetSize() int {
	l.Lock()
	defer l.Unlock()
	return l.list.GetSize()
}

// GetCap returns the capacity of the lru cache, don't need a lock here since cap is static
func (l *LRUCache) GetCap() int {
	return l.cap
}

// Get the value with the key passed in, if the key is found, move the node to the front
func (l *LRUCache) Get(key interface{}) interface{} {
	l.Lock()
	defer l.Unlock()

	n, found := l.cache[key]

	if found {
		l.list.MoveToFront(n)
		return n.GetVal()
	}

	return nil
}

// Set the key value pair in the cache
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

// Remove deletes the node with the key that was passed in, returns an error if it wasn't found
func (l *LRUCache) Remove(key interface{}) error {
	if key == nil {
		return errors.New("key must not be nil")
	}
	l.Lock()
	defer l.Unlock()

	oldNode, found := l.cache[key]

	if found {
		l.list.RemoveNode(oldNode)
		delete(l.cache, key)
		return nil
	}

	return errors.New("key not found")
}
