package _2_Add_Elements_in_LinkedList

import "strconv"

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


//--------------------------------------------------

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList  {
	return &LinkedList{
		head: nil,
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
func (this *LinkedList)AddFirst(e int)  {
	node := NewNode2(e)
	node.Next = this.head
	this.head = node
	//等价于
	//this.head = NewNode(e,this.head)

	this.size ++
}

//在链表的index(0-based)位置添加新的元素e
//在链表中不是一个常用的操作,练习用:)
func (this *LinkedList)Add(index int,e int)  {
	//边界
	if index < 0 || index > this.size{
		panic("Add failed. Illegal index.")
	}

	if index == 0{
		this.AddFirst(e)//如果是头元素
	}else{
		//查找将要插入元素之前的元素prev
		prev := this.head
		for i:=0;i<index - 1 ; i++ {
			prev = prev.Next
		}

		node := NewNode2(e)
		node.Next = prev.Next
		prev.Next = node

		//等价于
		//prev.Next = NewNode(e,prev.Next)

		this.size ++
	}
}

// 在链表末尾添加新的元素e
func (this *LinkedList)AddLast(e int)  {
	this.Add(this.size,e)
}