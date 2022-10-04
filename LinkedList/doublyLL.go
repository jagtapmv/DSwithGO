package doublyLL

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

//Insert at specific location/index in the DLL
func (head *DLLNode) InsertAtinDLL(index int, term any) {
	if head.DLLLength() > index && index > 0 {
		theNode := head
		i := 1
		for i != index {
			i++
			theNode = theNode.Next
		}
		newNode := &DLLNode{term, theNode, theNode.Next}
		theNode.Next.Prev = newNode
		theNode.Next = newNode
	} else {
		fmt.Println("Provided index is out of bound!")
	}
}

//InsertAtHeadDLL and InsertAtEndDLL will be similar to above with minor changes

//Delete from specific index
func (head *DLLNode) DeleteFromDLL(index int) {
	if head.DLLLength()-1 > index && index > 0 {
		theNode := head
		i := 0
		for i != index {
			i++
			theNode = theNode.Next
		}

		theNode.Prev.Next = theNode.Next
		theNode.Next.Prev = theNode.Prev
		theNode = &DLLNode{nil, nil, nil}
	} else {
		fmt.Println("Provided index is out of bound!")
	}
}
