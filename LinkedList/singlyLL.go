package main

import "fmt"

//create a node struct with data and next pointer

type Node struct {
	Data any
	Next *Node
}

//Method to traverse the Linked List
func (head *Node) traverseLL() {
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
func (head *Node) searchLL(searchTerm any) {
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

func main() {
	//creating 3 nodes with nil as data and pointer
	head := Node{nil, nil}
	first := Node{nil, nil}
	second := Node{nil, nil}

	//assigning arbitrary value as data and linking the nodes by assining next node address
	//to the pointer of previous node
	head.Data = 1
	head.Next = &first
	first.Data = "b"
	first.Next = &second
	second.Data = 10.3

	//Traversing through the list
	head.traverseLL()
	fmt.Println("List traversal is completed!")

	//Search for specific search term
	head.searchLL(1)

	//Insert element at the start of the LL
	newHead := head.InsertFirstinLL("first")
	newHead.traverseLL()

	//Delete first element(Head) of LL
	newHead = newHead.DeleteFirstfromLL()
	newHead.traverseLL()

	//Insert at last postion
	newHead.InsertLastinLL("last")
	newHead.traverseLL()
}
