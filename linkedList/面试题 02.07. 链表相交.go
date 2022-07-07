package main

/**
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0
	// 求出链表A和B的长度
	cur := headA
	for cur != nil {
		cur = cur.Next
		lenA++
	}
	cur = headB
	for cur != nil {
		cur = cur.Next
		lenB++
	}

	//长链表移动到剩下的长度与短链表长度相同
	//然后遍历链表查看是否能够对应上
	if lenA > lenB {
		leftLength := lenA
		for cur = headA; cur != nil && leftLength > lenB; cur = cur.Next {
			leftLength--
		}

		//双指针的方法
		retA := cur
		retB := headB
		postNodeA := retA
		postNodeB := retB
		for retA != nil && retB != nil && postNodeA != nil && postNodeB != nil {
			if retA.Val == retB.Val { //找到相等的节点
				if postNodeA.Next != postNodeB.Next { //后续节点不相等, 改变相等的节点位置, 此时节点位置应该在后续节点之后
					retA = postNodeA.Next
					retB = postNodeB.Next
				} else {
					postNodeA = postNodeA.Next
					postNodeB = postNodeB.Next
				}

				continue
			}

			//没有找到相等的节点就继续寻找
			retA = retA.Next
			retB = retB.Next
			postNodeA = retA
			postNodeB = retB
		}
		return retA
	} else if lenA < lenB {
		leftLength := lenB
		for cur = headB; cur != nil && leftLength > lenA; cur = cur.Next {
			leftLength--
		}

		//双指针的方法
		retB := cur
		retA := headA
		postNodeA := retA
		postNodeB := retB
		for retA != nil && retB != nil && postNodeA != nil && postNodeB != nil {
			if retA.Val == retB.Val { //找到相等的节点
				if postNodeA.Next != postNodeB.Next { //后续节点不相等, 改变相等的节点位置, 此时节点位置应该在后续节点之后
					retA = postNodeA.Next
					retB = postNodeB.Next
				} else {
					postNodeA = postNodeA.Next
					postNodeB = postNodeB.Next
				}

				continue
			}

			//没有找到相等的节点就继续寻找
			retA = retA.Next
			retB = retB.Next
			postNodeA = retA
			postNodeB = retB
		}
		return retA
	} else {
		//双指针的方法
		retB := headA
		retA := headA
		postNodeA := retA
		postNodeB := retB
		for retA != nil && retB != nil && postNodeA != nil && postNodeB != nil {
			if retA.Val == retB.Val { //找到相等的节点
				if postNodeA.Next != postNodeB.Next { //后续节点不相等, 改变相等的节点位置, 此时节点位置应该在后续节点之后
					retA = postNodeA.Next
					retB = postNodeB.Next
				} else {
					postNodeA = postNodeA.Next
					postNodeB = postNodeB.Next
				}

				continue
			}

			//没有找到相等的节点就继续寻找
			retA = retA.Next
			retB = retB.Next
			postNodeA = retA
			postNodeB = retB
		}
		return retA
	}	
}
