package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type ListNode struct {
	head *Node
}

func (l *ListNode) insert(val int) {
	newNode := &Node{Val: val, Next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}
func (l *ListNode) deleteNode(val int) {
	if l.head.Val == val {
		l.head = l.head.Next
		return
	}

	cur := l.head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}

}

func (l *ListNode) insertAfter(val int, data int) {
	newNode := &Node{Val: data, Next: nil}

	if l.head.Val == val {
		newNode.Next = l.head.Next
		l.head.Next = newNode
		return
	}

	cur := l.head

	for cur.Next != nil {
		if cur.Val == val {
			newNode.Next = cur.Next
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}
	cur.Next = newNode
}

func (l *ListNode) insertAtBeginningAtaValu(val int, data int) {
	newNode := &Node{Val: data, Next: nil}
	if l.head.Val == val {
		newNode.Next = l.head
		l.head = newNode
		return
	}
	cur := l.head
	for cur.Next != nil {
		if cur.Next.Val == val {
			newNode.Next = cur.Next
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}

}

func (l *ListNode) insertAtBeginning(val int) {
	newNode := &Node{Val: val, Next: nil}
	if l.head != nil {
		newNode.Next = l.head
		l.head = newNode
	}
}

func (l *ListNode) display() {
	cur := l.head

	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func main() {
	list := ListNode{}
	list.insert(10)
	list.insert(20)
	list.insert(30)
	list.insert(40)
	list.insert(50)

	// list.insertAtBeginning(200)
	// list.deleteNode(50)
	// list.insertAtBeginningAtaValu(50, 2000)
	list.insertAfter(50, 200)
	list.display()

}
