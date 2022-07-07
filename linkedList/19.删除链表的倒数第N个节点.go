package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 添加虚拟头节点, 就不用考虑头节点的问题了
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	pre := dummyHead
	i := 0 //用来计数
	for cur.Next != nil {
		if i < n { //i < n的时候pre节点不动
			i++
			cur = cur.Next
		} else { //i = n的时候pre节点和cur节点一起向后遍历
			pre = pre.Next
			cur = cur.Next
		}
	}
	//此时pre节点就是要删除节点的前一个节点
	pre.Next = pre.Next.Next
	return dummyHead.Next
}
