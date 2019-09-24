package _6_implement_stack_in_linkedlist

import "strings"

type LinkedListStack struct {
	list *LinkedList
}

func NewLinkedListStack() *LinkedListStack  {
	return &LinkedListStack{
		list:NewLinkedList(),
	}
}
//实现接口
func (this *LinkedListStack)GetSize() int  {
	return this.list.GetSize()
}
func (this *LinkedListStack)IsEmpty() bool  {
	return this.list.IsEmpty()
}
func (this *LinkedListStack)Push(e int)   {
	this.list.AddFirst(e)
}
func (this *LinkedListStack)Pop() int  {
	return this.list.RemoveFirst()
}
func (this *LinkedListStack)Peek() int  {
	return this.list.GetFirst()
}
func (this *LinkedListStack)ToString() string  {
	var sb strings.Builder
	sb.WriteString("Stack: top ")
	sb.WriteString(this.list.ToString())
	return sb.String()
}
