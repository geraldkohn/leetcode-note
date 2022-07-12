package main

/**
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
实现 MyQueue 类：
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false

说明：
你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
*/

type MyQueue struct {
    stackIn []int
    stackOut  []int
}

/** Initialize your data structure here. */
func Constructor_MyQueue() MyQueue {
    return MyQueue{
        stackIn: make([]int, 0),
        stackOut:  make([]int, 0),
    }
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
    for len(this.stackOut) != 0 {
        val := this.stackOut[len(this.stackOut)-1]
        this.stackOut = this.stackOut[:len(this.stackOut)-1]
        this.stackIn = append(this.stackIn, val)
    }
    this.stackIn = append(this.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    for len(this.stackIn) != 0 {
        val := this.stackIn[len(this.stackIn)-1]
        this.stackIn = this.stackIn[:len(this.stackIn)-1]
        this.stackOut = append(this.stackOut, val)
    }
    if len(this.stackOut) == 0 {
        return 0
    }
    val := this.stackOut[len(this.stackOut)-1]
    this.stackOut = this.stackOut[:len(this.stackOut)-1]
    return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    val := this.Pop()
    if val == 0 {
        return 0
    }
    this.stackOut = append(this.stackOut, val)
    return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.stackIn) == 0 && len(this.stackOut) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */