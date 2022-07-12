package main

/**
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
 

注意：

你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
*/

type MyStack struct {
    //创建两个队列
    queue1 []int
    queue2 []int
}


func Constructor_MyStack() MyStack {
    return MyStack{	//初始化
        queue1:make([]int,0),
        queue2:make([]int,0),
    }
}


func (this *MyStack) Push(x int)  {
     //先将数据存在queue2中
    this.queue2 = append(this.queue2,x)	
   //将queue1中所有元素移到queue2中，再将两个队列进行交换
    this.Move() 
}


func (this *MyStack) Move(){    
    if len(this.queue1) == 0{
        //交换，queue1置为queue2,queue2置为空
        this.queue1,this.queue2 = this.queue2,this.queue1
    }else{
        //queue1元素从头开始一个一个追加到queue2中
            this.queue2 = append(this.queue2,this.queue1[0])  
            this.queue1 = this.queue1[1:]	//去除第一个元素
            this.Move()     //重复
    }
}

func (this *MyStack) Pop() int {
    val := this.queue1[0]
    this.queue1 = this.queue1[1:]	//去除第一个元素
    return val

}


func (this *MyStack) Top() int {
    return this.queue1[0]	//直接返回
}


func (this *MyStack) Empty() bool {
return len(this.queue1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */