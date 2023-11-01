package main

import (
	"firstHW/linkedlist"
	"fmt"
)

func main() {
	myList := linkedlist.NewLinkedList(6, 2)
	fmt.Print(myList.At(1))
}
