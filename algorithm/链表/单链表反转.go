package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	// 定义一个空节点
	var prev *ListNode
	// 获取当前节点
	current := head
	for current != nil {
		// 先存储下一个节点位置
		next := current.Next
		// 把当前节点指向上一个节点【反转】
		current.Next = prev
		// 将当前节点赋值为上一个节点
		prev = current
		// 将下一个节点赋值为当前节点
		current = next
	}

	return prev
}

func main() {
	// 创建一个单链表：1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}

	// 反转单链表
	newHead := reverseLinkedList(head)

	// 打印反转后的单链表：5 -> 4 -> 3 -> 2 -> 1
	current := newHead
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
