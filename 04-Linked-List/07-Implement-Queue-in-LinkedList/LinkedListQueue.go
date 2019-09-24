package _7_Implement_Queue_in_LinkedList

import (
	"strconv"
	"strings"
)
type Node struct {
	E int
	Next *Node
}

func (this *Node)ToString() string  {
	return strconv.Itoa(this.E)//输出Node对象的字符串
}
func NewNode3() *Node {
	return &Node{
		E:    0,
		Next: nil,
	}
}
func NewNode2(e int) *Node {
	return &Node{
		E:    e,
		Next: nil,
	}
}
func NewNode(e int,next *Node) *Node {
	return &Node{
		E:    e,
		Next: next,
	}
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedListQueue() *LinkedListQueue  {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}
//实现接口
func (this *LinkedListQueue)GetSize() int  {
	return this.size
}
func (this *LinkedListQueue)IsEmpty() bool  {
	return this.size == 0
}
func (this *LinkedListQueue)Enqueue(e int)   {
	if nil == this.tail{
		this.tail = NewNode2(e)
		this.head = this.tail
	}else{
		this.tail.Next = NewNode2(e)
		this.tail = this.tail.Next
	}
	this.size ++
}
func (this *LinkedListQueue)Dequeue() int  {
	if this.IsEmpty(){
		panic("Cannot dequeue from an empty queue.")
	}
	retNode := this.head
	this.head = this.head.Next
	retNode.Next = nil
	if nil == this.head{
		this.tail = nil
	}
	this.size --
	return retNode.E
}
func (this *LinkedListQueue)GetFront() int  {
	if this.IsEmpty(){
		panic("Queue is empty.")
	}
	return this.head.E
}
func (this *LinkedListQueue)ToString() string  {
	var sb strings.Builder
	sb.WriteString("Queue: front ")
	cur := this.head
	for (cur != nil){
		sb.WriteString(cur.ToString() + "->")
		cur = cur.Next
	}
	sb.WriteString("NULL tail")
	return sb.String()
}
