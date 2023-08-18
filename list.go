package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	// 创建一个链表，包含5个节点
	node1 := &Node{value: 1}
	node2 := &Node{value: 2}
	node3 := &Node{value: 3}
	node4 := &Node{value: 4}
	node5 := &Node{value: 5}

	// 链接5个节点
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	// 插入
	// 在第4个位置插入新的数字6
	// 原来的链表
	// 1 -> 2 -> 3 -> 4 -> 5
	// 新的链表
	// 1 -> 2 -> 3 -> 6 -> 4 -> 5
	node6 := &Node{value: 6}
	node6.next = node4
	node3.next = node6

	// 删除
	// 删除第3个位置的数字3
	// 原来的链表
	// 1 -> 2 -> 3 -> 6 -> 4 -> 5
	// 新的链表
	// 1 -> 2 -> 6 -> 4 -> 5
	node2.next = node6

	// 打印链表
	printList(node1)
}

// 打印链表
func printList(node *Node) {
	for node != nil {
		fmt.Printf("%d -> ", node.value)
		node = node.next
	}
	// 删除最后一个多余的空格
	fmt.Print("\b\b")
	fmt.Println()
}
