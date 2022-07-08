package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
题意： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。
*/

func detectCycle(head *ListNode) *ListNode {
	// fast指针一次向前前进两个, slow指针向前前进一个, 当二者相遇的时候停止
	fast := head
	slow := head
	hasBegin := false
	for !hasBegin || fast != slow {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil || slow == nil || slow.Next == nil {
			return nil
		}
		hasBegin = true
		fast = fast.Next.Next
		slow = slow.Next
	}
	meetIndex := fast //相遇节点
	//从相遇节点和头节点同时向后找, 相遇的时候就是入口节点
	index1 := meetIndex
	index2 := head
	for index1 != index2 {
		if index1 == nil || index1.Next == nil || index2 == nil || index2.Next == nil {
			return nil
		}
		index1 = index1.Next
		index2 = index2.Next
	}
	return index1
}
