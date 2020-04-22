package doublylinkedlist

type Node struct {
	key  interface{}
	val  interface{}
	next *Node
	prev *Node
}

func MakeNode(key, val interface{}, next, prev *Node) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: next,
		prev: prev,
	}
}

func (n *Node) GetKey() interface{} {
	return n.key
}

func (n *Node) GetVal() interface{} {
	return n.val
}

func (n *Node) SetVal(val interface{}) {
	n.val = val
}

func (n *Node) SetKey(key interface{}) {
	n.key = key
}

func (n *Node) RemoveNode() {
	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	n.next = nil
	n.prev = nil
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(newNode *Node) {
	if n.next != nil {
		n.next.prev = newNode
	}

	if newNode != nil {
		newNode.next = n.next
		newNode.prev = n
	}

	n.next = newNode
}

func (n *Node) SetNextVal(key, val interface{}) {
	newNode := MakeNode(key, val, n.next, n)
	n.SetNext(newNode)
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) SetPrev(newNode *Node) {
	if n.prev != nil {
		n.prev.next = newNode
	}

	if newNode != nil {
		newNode.next = n
		newNode.prev = n.prev
	}

	n.prev = newNode
}

func (n *Node) SetPrevVal(key, val interface{}) {
	newNode := MakeNode(key, val, n, n.prev)
	n.SetPrev(newNode)
}
