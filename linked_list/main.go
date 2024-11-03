package linkedlist

import (
	"fmt"
	"log"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

// Insert node to the first
func (ll *LinkedList[T]) Unshift(value T) {
	var newNode = &Node[T]{
		Value: value,
		Next:  nil,
	}
	newNode.Next = ll.Head
	ll.Head = newNode
}

// Insert node with position
func (ll *LinkedList[T]) InsertPos(pos int, value T) {
	if ll.Head == nil {
		log.Fatalf("Failed insert node with position: %v \n", pos)
	}

	var tmp = ll.Head
	var tmpPos = 0
	var newNode = &Node[T]{
		Value: value,
		Next:  nil,
	}
	for tmp.Next != nil {
		tmpPos++
		if tmpPos == pos {
			newNode.Next = tmp.Next
			tmp.Next = newNode
			return
		}
		tmp = tmp.Next
	}

}

// Insert node to the last
func (ll *LinkedList[T]) Push(value T) {
	var newNode = &Node[T]{
		Value: value,
		Next:  nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	var tmp = ll.Head
	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = newNode
}

// Remove the first node
func (ll *LinkedList[T]) Shift() {
	if ll.Head == nil {
		log.Fatalln("Failed to remove the first node")
	}

	ll.Head = ll.Head.Next
}

// Remove node with position
func (ll *LinkedList[T]) RemovePos(pos int) {
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
func (ll *LinkedList[T]) Pop() {
	var tmp = ll.Head
	if ll.Head == nil {
		log.Fatalln("Linked List is empty")
	}

	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp = nil
}

func (ll *LinkedList[T]) Traverse() {
	fmt.Println("======================")
	var curr = ll.Head // pointer variable => & to get value of pointer and * to get value of deference variable
	for curr != nil {
		fmt.Println("Value", (*curr).Value)
		curr = (*curr).Next
	}
}
