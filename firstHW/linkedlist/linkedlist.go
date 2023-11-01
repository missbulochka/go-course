package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size uint
}

func NewLinkedList(size uint) *LinkedList {
	list := LinkedList{}
	if size > 0 {
		list.head = &Node{int(list.size), nil}
		list.size++
		currNode := &list.head
		for list.size < size {
			(*currNode).next = &Node{int(list.size), nil}
			currNode = &(*currNode).next
			list.size++
		}
	}
	return &list
}

func (list *LinkedList) returnNode(pos uint) (*Node, error) {
	if pos > list.size {
		return nil, errors.New("incorrect pos value")
	}

	needNode := list.head
	for i := uint(0); i < list.size; i++ {
		if i != pos {
			needNode = (*needNode).next
		} else {
			break
		}
	}
	return needNode, nil
}

func (list *LinkedList) Add(val int) {
	if list.size == 0 {
		list.head = &Node{val, nil}
		list.size++
		return
	}

	lastNode, _ := list.returnNode(list.size - 1)
	(*lastNode).next = &Node{val, nil}
	list.size++
}

func (list *LinkedList) Pop() {
	if list.size == 0 {
		fmt.Printf("Pop operation failed. Linked list is empty")
	}

	penultimateNode, _ := list.returnNode(list.size - 2)
	(*penultimateNode).next = nil
	list.size--
}

func (list *LinkedList) At(pos uint) int {
	if resNode, err := list.returnNode(pos); err == nil {
		return resNode.value
	} else {
		fmt.Println("At operation failed, return -1. Error: ", err)
		return -1
	}
}

func (list *LinkedList) Size() uint {
	return list.size
}

func (list *LinkedList) DeleteFrom(pos uint) {
	if list.size == 0 || pos > list.size {
		fmt.Println("DeleteFrom operation failed, incorrect pos value")
		return
	}

	if pos == 0 {
		list.head = (*list.head).next
	} else {
		penultimateNode, _ := list.returnNode(pos - 1)
		(*penultimateNode).next = (*penultimateNode).next.next
	}
	list.size--
}

func (list *LinkedList) AddTo(pos uint, val int) {
	if pos == 0 {
		tmp := list.head
		list.head = &Node{val, tmp}
		list.size++
	} else if penultimateNode, err := list.returnNode(pos - 1); err == nil {
		tmp := (*penultimateNode).next
		(*penultimateNode).next = &Node{val, tmp}
		list.size++
	} else {
		fmt.Print("Operation failed. Error: ", err)
	}
}

func (list *LinkedList) UpdateAt(pos uint, val int) {
	if needNode, err := list.returnNode(pos); err == nil {
		(*needNode).value = val
	} else {
		fmt.Print("Operation failed. Error: ", err)
	}
}

func (list *LinkedList) PrintList() {
	currNode := list.head
	for currNode != nil {
		fmt.Println(currNode.value)
		currNode = currNode.next
	}
}

func NewFromSlice(s []int) *LinkedList {
	list := LinkedList{}

	if len(s) > 0 {
		list.head = &Node{s[0], nil}
		currNode := list.head
		for _, val := range s[1:] {
			currNode.next = &Node{val, nil}
			currNode = currNode.next
		}
	}
	list.size = uint(len(s))
	return &list
}
