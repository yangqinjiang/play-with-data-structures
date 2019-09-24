package  _7_Implement_Queue_in_LinkedList

import (
	"fmt"
	"strings"
)

//使用数组实现的队列,
//TODO: 不支持泛型

type ArrayQueue struct {
	array *Array
}

func NewArrayQueue(capacity int) *ArrayQueue  {
	return &ArrayQueue{
		array:NewArray(capacity),
	}
}
func NewArrayQueueDefault() *ArrayQueue  {
	return &ArrayQueue{
		array:NewArrayDefault(),
	}
}

//实现接口

//
func (this *ArrayQueue)GetSize() int  {
	return  this.array.GetSize()
}

func (this *ArrayQueue)IsEmpty() bool  {
	return  this.array.IsEmpty();
}

func (this *ArrayQueue)GetCapacity() int  {
	return  this.array.GetCapacity();
}
func (this *ArrayQueue)Enqueue(e int)   {
	 this.array.AddLast(e)
}

func (this *ArrayQueue)Dequeue() int   {
	return this.array.RemoveFirst()
}

func (this *ArrayQueue)GetFront() int   {
	return this.array.GetFirst()
}

//打印数组字符串
func (this *ArrayQueue)ToString() string  {
	var sb strings.Builder
	fmt.Fprintf(&sb,"Queue: ")
	sb.WriteString("front [")
	for i := 0;i<this.array.GetSize() ;i++  {
		fmt.Fprintf(&sb,"%d",this.array.Get(i))
		if i != this.array.GetSize() - 1{
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] tail")
	return  sb.String()

}













