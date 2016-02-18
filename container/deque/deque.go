package deque

type Node struct {
	prevNode *Node
	nextNode *Node
	value    interface{}
}

func (node *Node) GetValue() interface{} {
	return node.value
}

type Deque struct {
	head, tail *Node
	size       int
}

func NewDeque() Deque {
	return Deque{nil, nil, 0}
}

func (deque *Deque) PushBack(val interface{}) {
	node := &Node{deque.tail, nil, val}
	//First insert
	if deque.size == 0 {
		deque.head = node
	} else {
		deque.tail.nextNode = node
	}

	deque.tail = node
	deque.size++
}

func (deque *Deque) PushFront(val interface{}) {
	//Fist insert
	if deque.size == 0 {
		deque.PushBack(val)
		return
	} else {
		node := &Node{nil, deque.head, val}
		deque.head.prevNode = node
		deque.head = node
	}
	deque.size++
}

func (deque *Deque) PopBack() interface{} {
	temp := deque.tail
	if deque.size == 1 {
		deque.Clear()
	} else if deque.size > 1 {
		deque.tail = temp.prevNode
		deque.size--
	} else {
		return nil
	}
	return temp.GetValue()
}

func (deque *Deque) PopFront() interface{} {
	temp := deque.head
	if deque.size == 1 {
		deque.Clear()
	} else if deque.size > 1 {
		deque.head = temp.nextNode
		deque.size--
	} else {
		return nil
	}
	return temp.GetValue()
}

func (deque *Deque) Clear() {
	deque.head = nil
	deque.tail = nil
	deque.size = 0
}

func (deque *Deque) Front() interface{} {
	return deque.head.GetValue()
}

func (deque *Deque) Back() interface{} {
	return deque.tail.GetValue()
}

func (deque *Deque) Empty() bool {
	if deque.size == 0 {
		return true
	}
	return false
}
