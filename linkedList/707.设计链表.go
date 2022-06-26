package main

import "fmt"

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

type MyLinkedList struct {
	node *node
}

type node struct {
	next *node
	prev *node
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{node: nil}
}

func (this *MyLinkedList) Get(index int) int {
	if this.node == nil || index < 0 { //刚刚初始化的链表
		return -1
	}

	i := 0
	head := this.node
	constHead := head //头节点
	for i < index {
		if head.next != constHead { //下一个不是头节点, 说明没有到尾节点
			i++
			head = head.next
		} else {
			return -1
		}
	}
	return head.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.node == nil { //刚刚初始化的链表
		this.node = &node{
			next: nil,
			prev: nil,
			val:  val,
		}
		this.node.next = this.node
		this.node.prev = this.node
		return
	}

	head := this.node
	preHead := &node{
		next: nil,
		prev: nil,
		val:  val,
	}
	preHead.prev = head.prev //新的头节点的下一个
	preHead.next = head      //新的头节点的下一个
	head.prev.next = preHead //尾节点下一个
	head.prev = preHead      //原来的头节点的上一个

	this.node = preHead //需要更换头节点的位置!!!!
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.node == nil { //刚刚初始化链表
		this.node = &node{
			next: nil,
			prev: nil,
			val:  val,
		}
		this.node.next = this.node
		this.node.prev = this.node
		return
	}

	head := this.node
	tail := head.prev
	nextTail := &node{
		next: nil,
		prev: nil,
		val:  val,
	}
	nextTail.next = head
	nextTail.prev = tail
	head.prev = nextTail
	tail.next = nextTail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.node == nil { //刚刚初始化结构体
		if index != 0 { //添加的位置必须是0
			return
		}
		this.node = &node{
			next: nil,
			prev: nil,
			val:  val,
		}
		this.node.next = this.node
		this.node.prev = this.node
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	i := 0
	head := this.node
	constHead := head
	for i < index-1 { //需要遍历到前一个节点
		if head.next != constHead {
			i++
			head = head.next
		} else { //index - 1节点不存在
			return
		}
	}
	indexNode := &node{
		next: nil,
		prev: nil,
		val:  val,
	}
	indexNode.next = head.next
	indexNode.prev = head
	head.next.prev = indexNode
	head.next = indexNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	if index == 0 { //删除头节点
		head := this.node
		if head.next == head && head.prev == head { //就这一个节点了
			this.node = nil
			return
		}
		this.node = head.next //需要更换头节点位置!!!!
		head.next.prev = head.prev
		head.prev.next = head.next
		return
	}

	i := 0
	head := this.node
	constHead := head
	for i < index { //需要遍历到这个节点
		if head.next != constHead {
			i++
			head = head.next
		} else { //index节点不存在
			return
		}
	}
	head.prev.next = head.next
	head.next.prev = head.prev
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// func main() {
// 	list := Constructor()
// 	list.AddAtHead(1)
// 	printList(&list)
// 	list.AddAtHead(2)
// 	printList(&list)
// 	list.AddAtHead(3)
// 	printList(&list)
// 	list.AddAtTail(9)
// 	printList(&list)
// 	list.AddAtTail(8)
// 	printList(&list)
// 	list.AddAtTail(7)
// 	printList(&list)
// 	list.AddAtIndex(2, 10)
// 	printList(&list)
// 	list.AddAtIndex(7, 11)
// 	printList(&list)
// 	list.DeleteAtIndex(0)
// 	printList(&list)
// 	list.DeleteAtIndex(6)
// 	printList(&list)
// 	list.DeleteAtIndex(2)
// 	printList(&list)
// 	list.DeleteAtIndex(0)
// 	printList(&list)
// 	list.DeleteAtIndex(0)
// 	printList(&list)
// 	list.DeleteAtIndex(0)
// 	printList(&list)
// 	list.DeleteAtIndex(2)
// 	printList(&list)
// 	list.DeleteAtIndex(1)
// 	printList(&list)
// 	list.DeleteAtIndex(0)
// 	printList(&list)
// 	fmt.Println(list.Get(0))
// 	fmt.Println(list.Get(1))
// 	list.AddAtIndex(1, 0)
// 	printList(&list)
// 	list.AddAtIndex(0, 1)
// 	printList(&list)
// 	fmt.Println(list.Get(0))
// }

func printList(list *MyLinkedList) {
	constHead := list.node
	head := list.node

	if head == nil {
		fmt.Println("nil")
		return
	}

	for head.next != constHead {
		fmt.Print(head.val, " ")
		head = head.next
	}
	fmt.Print(head.val, " ") //打印尾节点
	fmt.Println()
}
