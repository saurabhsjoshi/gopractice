package container

type Node struct {
	NextNode *Node
	Value    interface{}
}

func (node *Node) GetNextNode() *Node {
	return node.NextNode
}

func (node *Node) GetValue() interface{} {
	return node.Value
}

type LinkedList struct {
	Root *Node
	size int
}

func NewLinkedList() LinkedList {
	return LinkedList{nil, 0}
}

func (list *LinkedList) GetRootNode() *Node {
	return list.Root
}

func (list *LinkedList) Len() int {
	return list.size
}

func (list *LinkedList) Insert(val interface{}) *Node {
	//Pointer to the root pointer
	tNode := &list.Root
	//Navigate to last nil Node
	for ; *tNode != nil; tNode = &(*tNode).NextNode {
	}
	//Change the pointer to the new Node
	*tNode = &Node{nil, val}
	//Increase length
	list.size++

	return *tNode
}

func (list *LinkedList) RemoveValue(val interface{}) bool {
	tNode := list.getNodeByValue(val)
	if tNode != nil {
		*tNode = (*tNode).NextNode
		list.size--
		return true
	}
	return false
}

/* Gets the first node having the specified value */
func (list *LinkedList) SearchValue(val interface{}) *Node {
	tNode := list.getNodeByValue(val)
	if tNode != nil {
		return *tNode
	}
	return nil
}

/* Function to get a pointer to Node pointer (For internal use) */
func (list *LinkedList) getNodeByValue(val interface{}) **Node {
	tNode := &list.Root
	for ; *tNode != nil; tNode = &((*tNode).NextNode) {
		if (*tNode).Value == val {
			return tNode
		}
	}
	//Nil if value not found
	return nil
}
