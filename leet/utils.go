package leet

import (
	"fmt"
)

type LinkedList struct {
	Head *ListNode
}

func (l LinkedList) FromVec(v []int) LinkedList {
	head := &ListNode{Val: v[0]}
	curr := head
	for i := 1; i < len(v); i++ {

		next := ListNode{Val: v[i]}
		curr.Next = &next
		curr = curr.Next
	}
	return LinkedList{
		head,
	}
}

func (l LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d, ", curr.Val)
		curr = curr.Next
	}
}
