package _6_LinkedListMap

import "strconv"

// 此类要实现Map接口
type LinkedListMap struct {
	dummyHead *Node //虚拟头结点
	size int
}
func NewLinkedListMap() *LinkedListMap  {
	return &LinkedListMap{
		dummyHead: NewNode3(),
		size: 0,
	}
}

//接口:获取链表中的元素个数
func (this *LinkedListMap)GetSize() int  {
	return this.size
}

//接口:返回链表是否为空
func (this *LinkedListMap)IsEmpty() bool  {
	return 0 == this.size
}

//返回以K为键的Node
func (this *LinkedListMap)getNode(K string) *Node{
	cur := this.dummyHead.Next
	for cur != nil {
		if cur.Key == K { //找到,则返回
			return  cur
		}
		cur = cur.Next  //没找到,则下一个
	}
	return nil
}

//接口:是否包含以K为键的元素
func (this *LinkedListMap)Contains(K string) bool{
	return  this.getNode(K) != nil
}
//接口:返回key对应的值
func (this *LinkedListMap)Get(k string) int  {
	node := this.getNode(k)
	if node == nil{
		return 0
	}else{
		return node.Value
	}
}
//接口:往链表头添加元素,存在则更新
func (this *LinkedListMap) Add(key string ,value int ){
	node := this.getNode(key)
	if node == nil{
		this.dummyHead.Next = NewNode(key,value,this.dummyHead.Next)
		this.size ++
	}else{
			node.Value = value
	}
}

//接口:更新以K为键的值,为newValue
func (this *LinkedListMap)Set(key string,newValue int)  {
	node := this.getNode(key)
	if node == nil{
		panic(key + " doesn't exist!")
	}
	//
	node.Value = newValue
}
// 从链表中删除index(0-based)位置的元素, 返回删除的元素
// 在链表中不是一个常用的操作，练习用：）

//接口: 删除key对应的数据,并返回
func (this *LinkedListMap)Remove(key string) int{

	prev := this.dummyHead
	//找到待删除节点的前一个节点prev
	for prev.Next != nil  {
		if prev.Next.Key == key{ //找到,则跳过这一次for
			break
		}
		prev = prev.Next
	}
	if prev.Next != nil{
		//删除元素
		delNode := prev.Next
		prev.Next = delNode.Next
		delNode.Next = nil
		this.size --
		return  delNode.Value
	}
	return 0

}
// 链表结点
type Node struct {
	Key string
	Value int
	Next *Node
}

func NewNode(k string,v int,next *Node) *Node  {
	return &Node{
		Key:k,
		Value:v,
		Next:next,
	}
}
func NewNode2(k string,v int ) *Node  {
	return NewNode(k,v,nil)
}
func NewNode3() *Node  {
	return NewNode("",0,nil)
}
func (this *Node)ToString() string  {
	return this.Key + " : " + strconv.Itoa(this.Value)
}