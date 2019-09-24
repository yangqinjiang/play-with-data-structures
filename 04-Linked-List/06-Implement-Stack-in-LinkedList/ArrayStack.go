package _6_implement_stack_in_linkedlist

import (
	"fmt"
	"strings"
)

//使用数组实现的栈

type ArrayStack struct {
	array *Array
}

func NewArrayStack(capacity int) *ArrayStack  {
	return &ArrayStack{
		array:NewArray(capacity),
	}
}
func NewArrayStackDefault() *ArrayStack  {
	return &ArrayStack{
		array:NewArrayDefault(),
	}
}

//实现接口

//
func (this *ArrayStack)GetSize() int  {
	return  this.array.GetSize()
}

func (this *ArrayStack)IsEmpty() bool  {
	return  this.array.IsEmpty();
}

func (this *ArrayStack)GetCapacity() int  {
	return  this.array.GetCapacity();
}
func (this *ArrayStack)Push(e int)   {
	 this.array.AddLast(e)
}

func (this *ArrayStack)Pop() int   {
	return this.array.RemoveLast()
}

func (this *ArrayStack)Peek() int   {
	return this.array.GetLast()
}

//打印数组字符串
func (this *ArrayStack)ToString() string  {
	var sb strings.Builder
	fmt.Fprintf(&sb,"Stack: ")
	sb.WriteString("[")
	for i := 0;i<this.array.GetSize() ;i++  {
		fmt.Fprintf(&sb,"%d",this.array.Get(i))
		if i != this.array.GetSize() - 1{
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] top")
	return  sb.String()

}













