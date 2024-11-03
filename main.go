package main

import (
	"fmt"
	"log"
)

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

// Insert node to the first
func (ll *LinkedList[T]) unshift(node *Node[T]) {
	(*node).Next = ll.Head
	ll.Head = node
}

// Insert node with position
func (ll *LinkedList[T]) insertPos(pos int, node *Node[T]) {
	if ll.Head == nil {
		log.Fatalf("Failed insert node with position: %v \n", pos)
	}

	var tmp = ll.Head
	var tmpPos = 0
	for tmp.Next != nil {
		tmpPos++
		if tmpPos == pos {
			node.Next = tmp.Next
			tmp.Next = node
			return
		}
		tmp = tmp.Next
	}

}

// Insert node to the last
func (ll *LinkedList[T]) push(node *Node[T]) {
	if ll.Head == nil {
		ll.Head = node
		return
	}

	var tmp = ll.Head
	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = node
}

// Remove the first node
func (ll *LinkedList[T]) shift() {
	if ll.Head == nil {
		log.Fatalln("Failed to remove the first node")
	}

	ll.Head = ll.Head.Next
}

// Remove node with position
func (ll *LinkedList[T]) removePos(pos int) {
	if ll.Head == nil {
		log.Fatalln("Failed to remove the first node")
	}

	var tmp = ll.Head
	var tmpPos = 0
	var preNode *Node[T]
	for tmp.Next != nil {
		preNode = tmp
		tmpPos++
		if tmpPos == pos {
			preNode.Next = tmp.Next.Next
			return
		}

		tmp = tmp.Next
	}
}

// Remove the last node
func (ll *LinkedList[T]) pop() {

}

func (ll *LinkedList[T]) traverse() {
	fmt.Println("======================")
	var curr = ll.Head // pointer variable => & to get value of pointer and * to get value of deference variable
	for curr != nil {
		fmt.Println("data", (*curr).Data)
		curr = (*curr).Next
	}
}

func main() {
	var newNode = &Node[string]{
		Data: "10",
		Next: nil,
	}
	var newNode2 = &Node[string]{
		Data: "11",
		Next: nil,
	}
	newNode.Next = newNode2

	var newNode3 = &Node[string]{
		Data: "13",
		Next: nil,
	}

	var newNode4 = &Node[string]{
		Data: "14",
		Next: nil,
	}

	var ll = &LinkedList[string]{}

	// head is the pointer to point the first node in Linked List
	ll.Head = newNode2

	ll.unshift(newNode)
	ll.push(newNode3)
	ll.insertPos(1, newNode4)
	ll.traverse()

	ll.shift()
	ll.removePos(2)
	ll.traverse()

}
