package _1_Linked_List_Basics

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
type LinkedList struct {

}