# 链表

## 链表的理论基础
* 链表的类型: 单链表, 双链表, 循环链表
* 链表的操作: 删除节点, 添加节点, 查询节点

> 注意链表的结构体写法!

```go
// 单链表
type ListNode struct {
    Val int
    Next *ListNode
}
```