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

func NewLinkedList(size uint, val int) LinkedList {
	list := LinkedList{}
	if size == 0 {
		return list
	}

	currNode := &list.head
	for list.size < size {
		if *currNode == nil {
			*currNode = &Node{
				val, nil,
			}
			list.size++

			if size == list.size {
				break
			}
		}

		(*currNode).next = &Node{
			int(list.size) + val,
			nil,
		}
		currNode = &(*currNode).next
		list.size++
	}
	return list
}

func (list *LinkedList) returnNode(pos uint) (*Node, error) {
	needNode := list.head
	for i := uint(0); i < list.size; i++ {
		if i+1 != pos {
			needNode = (*needNode).next
		} else {
			return needNode, nil
		}
	}
	return nil, errors.New("incorrect pos value")
}

func (list *LinkedList) Add(val int) {
	lastNode, _ := list.returnNode(list.size)
	(*lastNode).next = &Node{val, nil}
	list.size++
}

func (list *LinkedList) Pop() {
	penultimateNode, _ := list.returnNode(list.size - 1)
	(*penultimateNode).next = nil
	list.size--
}

func (list *LinkedList) At(pos uint) (int, error) {
	if resNode, err := list.returnNode(pos); err == nil {
		return resNode.value, nil
	} else {
		return 0, nil
	}
}

func (list *LinkedList) Size() uint {
	return list.size
}

func (list *LinkedList) DeleteFrom(pos uint) {
	if pos == 1 {
		list.head = (*list.head).next
		list.size--
	} else {
		penultimateNode, err := list.returnNode(pos - 1)
		if err == nil {
			(*penultimateNode).next = (*penultimateNode).next.next
			list.size--
		} else {
			fmt.Print("Operation failed. Error: ", err)
		}
	}
}

func (list *LinkedList) AddTo(pos uint, val int) {
	var penultimateNode *Node
	var err error
	if pos == 1 {
		tmp := list.head
		list.head = &Node{val, tmp}
		list.size++
	}

	penultimateNode, err = list.returnNode(pos - 1)
	if err == nil {
		tmp := (*penultimateNode).next
		(*penultimateNode).next = &Node{val, tmp}
		list.size++
	} else {
		fmt.Print("Operation failed. Error: ", err)
	}
}

func (list *LinkedList) UpdateAt(pos uint, val int) {
	needNode, err := list.returnNode(pos)
	if err == nil {
		(*needNode).value = val
	} else {
		fmt.Print("Operation failed. Error: ", err)
	}
}
