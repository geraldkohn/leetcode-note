package main

/**
题意：删除链表中等于给定值 val 的所有节点。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//添加虚拟头节点
	preHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	preHead.Next = head
	node := preHead
	for node.Next != nil && node != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next //remove
		} else {
			node = node.Next //skip
		}
	}
	return preHead.Next
}
