package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针解法
func reverseList_1(head *ListNode) *ListNode {
	var cur, pre, temp *ListNode
	cur = head
	pre = nil
	for cur != nil {
		temp = cur.Next //用一个临时变量存放cur.Next, 防止下一步赋值将 cur.Next 改变    (注解: 这里将temp赋值cur是没有用的, 因为之后要用到temp.Next, 即cur.Next, 但是cur.Next已经在下一步被改变)
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}


// 递归解法
func reverseList_2(head *ListNode) *ListNode {
	return reverse(head, nil)
}

func reverse(cur *ListNode, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	temp := cur.Next
	cur.Next = pre
	return reverse(temp, cur)
}