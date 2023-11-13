package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	size uint
}

func NewLinkedList(size uint) *LinkedList {
	if size == 0 {
		return &LinkedList{}
	}
	list := LinkedList{}
	list.head = &node{int(list.size), nil}
	list.size++
	currNode := &list.head
	for list.size < size {
		(*currNode).next = &node{int(list.size), nil}
		currNode = &(*currNode).next
		list.size++
	}
	return &list
}

func (list *LinkedList) returnNode(pos uint) (*node, error) {
	if pos > list.size {
		return nil, errors.New("incorrect pos value")
	}

	needNode := list.head
	for i := uint(0); i < list.size; i++ {
		if i == pos {
			break
		}
		needNode = (*needNode).next
	}
	return needNode, nil
}

func (list *LinkedList) Add(val int) {
	if list.size == 0 {
		list.head = &node{val, nil}
		list.size++
		return
	}

	lastNode, _ := list.returnNode(list.size - 1)
	(*lastNode).next = &node{val, nil}
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
	resNode, err := list.returnNode(pos)
	if err == nil {
		return resNode.value
	}
	fmt.Println("At operation failed, return -1. Error: ", err)
	return -1
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
		list.head = &node{val, tmp}
		list.size++
	} else if penultimateNode, err := list.returnNode(pos - 1); err == nil {
		tmp := (*penultimateNode).next
		(*penultimateNode).next = &node{val, tmp}
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
		list.head = &node{s[0], nil}
		currNode := list.head
		for _, val := range s[1:] {
			currNode.next = &node{val, nil}
			currNode = currNode.next
		}
	}
	list.size = uint(len(s))
	return &list
}
