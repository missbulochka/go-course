package main

import (
	"firstHW/linkedlist"
	"fmt"
)

func main() {
	s := []int{2, 4, 5, 78, 4}
	myList1 := linkedlist.NewLinkedList(0)
	myList2 := linkedlist.NewFromSlice(s)

	myList1.AddTo(0, 220)
	myList1.Add(56)
	myList1.AddTo(myList1.Size()-1, 220)
	fmt.Println("myList1:")
	myList1.PrintList()

	myList2.Pop()
	myList2.DeleteFrom(0)
	myList2.DeleteFrom(myList2.Size() - 1)
	fmt.Println("myList2")
	myList2.PrintList()

	fmt.Printf("Size of myList1 is %d and size of myList2 is %d\n", myList1.Size(), myList2.Size())
	fmt.Printf("myList1[2] = %d\n", myList1.At(2))

	myList2.UpdateAt(0, 1000000000)
	fmt.Println("myList2")
	myList2.PrintList()
}
