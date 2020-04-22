package node

type Node struct {
	val interface{}
	next *Node
	prev *Node
}

func MakeNode(val interface{}, next, prev *Node) *Node {
	return &Node {
		val: val,
		next: next,
		prev: prev,
	}
}

func (n *Node) GetVal() interface{} {
	return n.val
}

func (n *Node) RemoveNode() {
	n.next.prev = n.prev
	n.prev.next = n.next
	n.next = nil
	n.prev = nil
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(newNode *Node) {
	if (n.next != nil) {
		n.next.prev = newNode
	}
	newNode.next = n.next
	newNode.prev = n
	n.next = newNode
}

func (n *Node) SetNextVal(val interface{}) {
	newNode := MakeNode(val, n.next, n)
	n.SetNext(newNode)
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) SetPrev(newNode *Node) {
	if (n.prev != nil) {
		n.prev.next = newNode
	}
	newNode.next = n
	newNode.prev = n.prev
	n.prev = newNode
}

func (n *Node) SetPrevVal(val interface{}) {
	newNode := MakeNode(val, n, n.prev)
	n.SetPrev(newNode)
}
