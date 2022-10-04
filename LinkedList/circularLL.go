package circularLL

import "fmt"

type CLLNode struct {
	Data any
	Next *CLLNode
}

//Traversing
func (node *CLLNode) TraverseCLL() {
	nodeaddr := &node
	lastNode := node
	for lastNode.Next != *nodeaddr {
		fmt.Println(lastNode.Data)
		lastNode = lastNode.Next
	}
	fmt.Println(lastNode.Data)
}

//Insert New node after given node
func (node *CLLNode) InsertAfter(term any) {
	newNode := CLLNode{term, node.Next}
	node.Next = &newNode
}

//Here delete, reverse traversal and other methods can come
//They are simple and similar to that of singlyLL
