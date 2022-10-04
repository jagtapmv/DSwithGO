package singlyLL

import (
	"fmt"
	"log"
)

//create a node struct with data and next pointer

type Node struct {
	Data any
	Next *Node
}

//Method to create the new LL
func CreateLL(listData []any) *Node {
	head := &Node{listData[0], nil}
	for i := 1; i < len(listData); i++ {
		head.InsertLastinLL(listData[i])
	}
	return head
}

//Method to traverse the Linked List
func (head *Node) TraverseLL() {
	data := head.Data
	next := head.Next
	i := 0
	for {
		fmt.Printf("%dth value is: %v\n", i, data)
		data = next.Data
		next = next.Next
		i++
		if next == nil {
			fmt.Printf("%dth value is: %v\n", i, data)
			break
		}
	}
}

//Method to search for any specific term. It will output the index if term is present
func (head *Node) SearchLL(searchTerm any) {
	data := head.Data
	next := head.Next
	i := 0
	flag := false
	for {
		if data == searchTerm {
			flag = true
			fmt.Println("Found at index: ", i)
			break
		}
		if next == nil {
			break
		}
		data = next.Data
		next = next.Next
		i++
	}
	if !flag {
		fmt.Println("The given search term could not be found!")
	}
}

//Method to insert the node as head Or at the first position
func (head *Node) InsertFirstinLL(term any) *Node {
	return &Node{term, head}
}

//Method to delete the first/head node of a LL
func (head *Node) DeleteFirstfromLL() *Node {
	newHead := head.Next
	head = &Node{nil, nil}
	return newHead
}

//Insert at last position when head is given
func (head *Node) InsertLastinLL(term any) {
	var newNode Node
	newNode.Data = term
	newNode.Next = nil
	lastNode := head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = &newNode
}

//Get length of a given LL
func (head *Node) Length() int {
	length := 1
	lastNode := head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		length++
	}
	return length
}

//Insert at the given position in the list
func (head *Node) InsertAtinLL(index int, term any) {
	if index < head.Length() && index > 0 {
		theNode := head
		i := 1
		for i != index {
			theNode = theNode.Next
			i++
		}
		theNode.Next = &Node{term, theNode.Next}
	} else {
		log.Println("The index should be greater than 0 and less than", head.Length())
	}
}
