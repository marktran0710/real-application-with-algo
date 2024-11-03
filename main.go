package main

import (
	"fmt"

	"github.com/marktran0710/real-application-with-algo/hashmap"
	linkedlist "github.com/marktran0710/real-application-with-algo/linked_list"
)

func main() {
	var ll = &linkedlist.LinkedList[string]{}

	// head is the pointer to point the first node in Linked List
	ll.Push("10")
	ll.Push("11")
	ll.Push("12")
	ll.Push("13")

	ll.Unshift("14")
	ll.InsertPos(1, "15")
	ll.Traverse()

	ll.Shift()
	ll.RemovePos(2)
	ll.Traverse()

	hm := hashmap.HashMap{}
	hm.Size(10)

	var key = "a"
	hm.Set(key, 20)
	value, ok := hm.Get(key)
	if !ok {
		fmt.Printf("Not found key: %s \n", key)
	}

	fmt.Printf("Key:%s, Value:%v \n", key, value)

	value, ok = hm.Get("b")
	if !ok {
		fmt.Printf("Not found key: %s \n", "b")
	}

}
