package main

/**
输入：head = [1,2,3,4]
输出：[2,1,4,3]
输入：head = []
输出：[]
输入：head = [1]
输出：[1]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 添加一个虚拟头节点
	dummyHead := &ListNode{
		Val: 0,
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		// cur -> tmp1 -> tmp2 -> tmp3
		tmp1 := cur.Next
		tmp2 := cur.Next.Next
		tmp3 := cur.Next.Next.Next
		cur.Next = tmp2
		cur.Next.Next = tmp1
		cur.Next.Next.Next = tmp3
		cur = cur.Next.Next
	}
	return dummyHead.Next
}