package binarysearchtree

import "fmt"

type Node struct {
	leftNode  *Node
	rightNode *Node
	value     int
}

type BinarySearchTree struct {
	rootNode *Node
	size     int
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{nil, 0}
}

func (tree *BinarySearchTree) Len() int {
	return tree.size
}

func (tree *BinarySearchTree) GetRootNode() *Node {
	return tree.rootNode
}

func (tree *BinarySearchTree) InsertValue(val int) {
	if tree.size == 0 {
		tree.rootNode = &Node{nil, nil, val}
	} else {
		tree.insertRecursively(tree.rootNode, val)
	}
	tree.size++
}

func (tree *BinarySearchTree) insertRecursively(node *Node, val int) {
	if val > node.value && node.rightNode != nil {
		tree.insertRecursively(node.rightNode, val)
	} else if val > node.value && node.rightNode == nil {
		node.rightNode = &Node{nil, nil, val}
	} else if val < node.value && node.leftNode != nil {
		tree.insertRecursively(node.leftNode, val)
	} else if val < node.value && node.leftNode == nil {
		node.leftNode = &Node{nil, nil, val}
	}
}

func (tree *BinarySearchTree) PrintInOrder(node *Node) {

	if node.leftNode != nil {
		tree.PrintInOrder(node.leftNode)
	}

	fmt.Println(node.value)

	if node.rightNode != nil {
		tree.PrintInOrder(node.rightNode)
	}

}

func (tree *BinarySearchTree) GetPreOrder(node *Node) []int{
	var a []int
	tree.fillPreOrderArray(node, a)
	return a
}

func (tree *BinarySearchTree) fillPreOrderArray(node *Node, result []int){

}

func (tree *BinarySearchTree) SearchValue(node *Node, val int) bool {
	if node == nil {
		return false
	} else if node.value == val {
		return true
	} else if val > node.value {
		return tree.SearchValue(node.rightNode, val)
	} else if val < node.value {
		return tree.SearchValue(node.leftNode, val)
	} else {
		return false
	}
}
