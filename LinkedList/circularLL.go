package main

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

//Main function
func main() {
	first := CLLNode{nil, nil}
	second := CLLNode{nil, nil}
	third := CLLNode{nil, nil}
	fourth := CLLNode{nil, nil}

	first.Data = "first"
	second.Data = "second"
	third.Data = "third"
	fourth.Data = "fourth"
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &first

	//Traverse the circular list
	fourth.TraverseCLL()

	//Insert New node after given node
	fourth.InsertAfter("fifth")
	first.TraverseCLL()
}
