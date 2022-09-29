package main

import "fmt"

//create a node struct with data and next pointer

type Node struct {
	Data int
	Next *Node
}

func main() {
	//creating 3 nodes with 0 as data and nil as pointer
	head := Node{0, nil}
	first := Node{0, nil}
	second := Node{0, nil}

	//assigning arbitrary value as data and linking the nodes by assining next node address
	//to the pointer of previous node
	head.Data = 1
	head.Next = &first
	first.Data = 2
	first.Next = &second
	second.Data = 3

	//Traversing through the list
	data := head.Data
	next := head.Next
	i := 0
	for {
		fmt.Printf("%dth value is: %d\n", i, data)
		data = next.Data
		next = next.Next
		i++
		if next == nil {
			fmt.Printf("%dth value is: %d\n", i, data)
			break
		}
	}
	fmt.Println("List traversal is completed!")
}
