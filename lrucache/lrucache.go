package lrucache

import (
	"golrucache/doublylinkedlist"
	"sync"
)

type LRUCache struct {
	cache map[interface{}]*doublylinkedlist.Node
	list  *doublylinkedlist.DoublyLinkedList
	cap   int
	sync.Mutex
}

func MakeLRUCache(cap int) *LRUCache {
	//if cap is <= 0 then throw error
	return &LRUCache{
		cache: make(map[interface{}]*doublylinkedlist.Node, cap),
		list:  doublylinkedlist.MakeDoublyLinkedList(),
		cap:   cap,
	}
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

	node, found := l.cache[key]

	if found {
		l.list.RemoveNode(node)
		l.list.InsertFront(node)
		return node.GetVal()
	} else {
		return nil
	}
}

func (l *LRUCache) Set(key, val interface{}) bool {
	if key == nil || val == nil {
		return false
	}
	l.Lock()
	defer l.Unlock()

	oldNode, found := l.cache[key]

	if found {
		l.list.RemoveNode(oldNode)
		l.list.InsertFront(oldNode)
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
	return true
}
