package Linked_List_Set

import (
	"fmt"
	"strings"
)

type Node struct {
	E string
	Next *Node
}

func (this *Node)ToString() string  {
	return this.E//输出Node对象的字符串
}
func NewNode3() *Node {
	return &Node{
		E:    "",
		Next: nil,
	}
}
func NewNode2(e string) *Node {
	return &Node{
		E:    e,
		Next: nil,
	}
}
func NewNode(e string,next *Node) *Node {
	return &Node{
		E:    e,
		Next: next,
	}
}


//--------------------------------------------------

type LinkedList struct {
	dummyHead *Node //虚拟头结点
	size int
}

func NewLinkedList() *LinkedList  {
	return &LinkedList{
		dummyHead: NewNode3(),
		size: 0,
	}
}

//获取链表中的元素个数
func (this *LinkedList)GetSize() int  {
	return this.size
}

//返回链表是否为空
func (this *LinkedList)IsEmpty() bool  {
	return 0 == this.size
}

//在链表头添加新的元素e
func (this *LinkedList)AddFirst(e string)  {
	this.Add(0,e)
}

//在链表的index(0-based)位置添加新的元素e
//在链表中不是一个常用的操作,练习用:)
func (this *LinkedList)Add(index int,e string)  {
	//边界
	if index < 0 || index > this.size{
		panic("Add failed. Illegal index.")
	}


		//查找将要插入元素之前的元素prev
		prev := this.dummyHead
		for i:=0;i<index ; i++ {
			prev = prev.Next
		}

		node := NewNode2(e)
		node.Next = prev.Next
		prev.Next = node

		//等价于
		//prev.Next = NewNode(e,prev.Next)

		this.size ++
}

// 在链表末尾添加新的元素e
func (this *LinkedList)AddLast(e string)  {
	this.Add(this.size,e)
}
// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用：）

func (this *LinkedList)Get(index int) string {
	if index <0 || index >= this.size{
		panic("Get failed. Illegal index.")
	}
	cur := this.dummyHead.Next
	for i:=0;i<index ;i++  {
		cur = cur.Next
	}
	return cur.E
}
// 获得链表的第一个元素
func (this *LinkedList)GetFirst() string  {
	return this.Get(0)
}
//// 获得链表的最后一个元素
func (this *LinkedList)GetLast() string  {
	return this.Get(this.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）

func (this *LinkedList)Set(index int,e string)  {
	if index <0 || index >= this.size{
		panic("Get failed. Illegal index.")
	}

	cur := this.dummyHead.Next
	//查找index位置的元素
	for i:=0;i<index ;i++  {
		cur = cur.Next
	}
	cur.E = e
}

// 查找链表中是否有元素e
func (this *LinkedList)Contains(e string) bool  {
	cur := this.dummyHead.Next
	for cur != nil {
		if cur.E == e{
			return true
		}
		cur = cur.Next
	}
	return false
}

// 从链表中删除index(0-based)位置的元素, 返回删除的元素
// 在链表中不是一个常用的操作，练习用：）
func (this *LinkedList)Remove(index int) string{
	if index <0 || index >=this.size{
		panic("Remove failed. Index is illegal.")
	}

	prev := this.dummyHead
	for i:=0;i<index ;i++  {
		prev = prev.Next
	}
	//删除元素
	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	this.size --
	return  retNode.E
}

// 从链表中删除第一个元素, 返回删除的元素
func (this *LinkedList)RemoveFirst() string  {
	return this.Remove(0)
}
// 从链表中删除最后一个元素, 返回删除的元素
func (this *LinkedList)RemoveLast() string  {
	return this.Remove(this.size - 1)
}

//从链表中删除元素e
func (this *LinkedList)RemoveElement(e string)  {
	prev := this.dummyHead
	for(prev.Next != nil){
		if prev.Next.E == e{
			break//找到,中止for
		}
		prev = prev.Next
	}

	//
	if prev.Next != nil{
		delNode := prev.Next
		prev.Next = delNode.Next
		delNode.Next = nil
		this.size --
	}
}

func (this *LinkedList)ToString() string  {
	var sb strings.Builder
	for cur := this.dummyHead.Next;cur != nil;cur = cur.Next{
		fmt.Fprintf(&sb,"%s->",cur.ToString())
	}
	sb.WriteString("NULL")

	return  sb.String()
}