package main

import "fmt"

type DLLNode struct {
	Data any
	Prev *DLLNode
	Next *DLLNode
}

//Method to traverse from head in forward direction
func (head *DLLNode) TraverseDLL() {
	lastNode := head
	i := 0
	for lastNode.Next != nil {
		i++
		fmt.Printf("%dth element is: %v\n", i, lastNode.Data)
		lastNode = lastNode.Next
	}
	fmt.Printf("last element is: %v\n", lastNode.Data)

}

//Method to traverse in reverse direction
func (lastNode *DLLNode) TraverseReverseDLL() {
	head := lastNode
	i := lastNode.DLLLengthV2()
	for head.Prev != nil {
		fmt.Printf("%dth element is: %v\n", i, head.Data)
		head = head.Prev
		i--
	}
	fmt.Printf("First element is: %v\n", head.Data)
}

//Get length of the DLL
func (head *DLLNode) DLLLength() int {
	lastNode := head
	length := 1
	for lastNode.Next != nil {
		length++
		lastNode = lastNode.Next
	}
	return length
}

//Get length of the DLL
func (lastNode *DLLNode) DLLLengthV2() int {
	head := lastNode
	length := 1
	for head.Prev != nil {
		length++
		head = head.Prev
	}
	return length
}

func main() {
	head := DLLNode{nil, nil, nil}
	first := DLLNode{nil, nil, nil}
	second := DLLNode{nil, nil, nil}

	head.Data = 2
	head.Next = &first
	first.Data = 3
	first.Next = &second
	first.Prev = &head
	second.Data = 5
	second.Prev = &first

	//Traverse the Linked list
	head.TraverseDLL()
	second.TraverseReverseDLL()

}
