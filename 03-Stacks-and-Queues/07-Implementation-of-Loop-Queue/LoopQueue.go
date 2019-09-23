package _7_Implementation_of_Loop_Queue

import (
	"fmt"
	"strings"
)

//使用数组实现的队列,

type LoopQueue struct {
	data []int
	front,tail int
	size int// 有兴趣的同学，在完成这一章后，可以思考一下：
			// LoopQueue中不声明size，如何完成所有的逻辑？
			// 这个问题可能会比大家想象的要难一点点：）
}

func NewLoopQueue(capacity int) *LoopQueue  {
	return &LoopQueue{
		data: make([]int,capacity+1),//比容量大1
		front:0,
		tail:0,
		size:0,
	}
}
//默认构造函数
func NewLoopQueueDefault() *LoopQueue {
	return NewLoopQueue(10)
}

func (this *LoopQueue)GetCapacity() int  {
	return len(this.data) -1
}

//实现Queue接口
func (this *LoopQueue)GetSize() int {
	return this.size
}

func (this *LoopQueue)IsEmpty() bool {
	return this.front == this.tail//两者相等,则说明为空
}
// 入队
func (this *LoopQueue)Enqueue(e int) {
	if (this.tail + 1) % len(this.data) == this.front{
		this.resize(this.GetCapacity()*2)
	}

	this.data[this.tail] = e
	//修改tail值
	this.tail = (this.tail + 1) % len(this.data)
	this.size ++
}
// 下一小节再做具体实现
func (this *LoopQueue)Dequeue() int {
	if this.IsEmpty(){
		panic("Cannot dequeue from an empty queue.")
	}
	ret := this.data[this.front]
	this.data[this.front] = 0 // nil值
	//更新front值
	this.front = (this.front + 1) % len(this.data)
	this.size --
	if this.size == this.GetCapacity() / 4 && this.GetCapacity() /2 != 0{
		this.resize(this.GetCapacity()/2)
	}

	return ret
}
// 下一小节再做具体实现
func (this *LoopQueue)getFront(e int) int {
	if this.IsEmpty(){
		panic("Cannot dequeue from an empty queue.")
	}
	return this.data[this.front]
}
//打印数组字符串
func (this *LoopQueue)ToString() string  {
	var sb strings.Builder
	fmt.Fprintf(&sb,"Queue: size= %d , capacity = %d\n",this.size,this.GetCapacity())
	sb.WriteString("front [")
	//

	for i := this.front; i!= this.tail;i= (i+1) % len(this.data) {
		fmt.Fprintf(&sb,"%d",this.data[i])
		if (i + 1) % len(this.data) != this.tail{
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] tail")
	return  sb.String()

}
//将数组空间的容量变成newCapacity大小
func (this *LoopQueue)resize(newCapacity int)  {
	fmt.Println("--------call resize func---------",this.size,"-->",newCapacity)
	newData := make([]int,newCapacity)
	for i:=0;i<this.size ;i++  {
		newData[i] = this.data[i]
	}
	this.data = newData
	this.front = 0
	this.tail = this.size

}