package Set

import (
	LinkedList "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/LinkedList"
)
/**
集合类, 实现Set接口的LinkedList
*/
type LinkedListSet struct {
	linkedList * LinkedList.LinkedList //这个LinkedList的元素是string类型
}

func NewLinkedListSet() *LinkedListSet  {
	return &LinkedListSet{linkedList: LinkedList.NewLinkedList()}
}
//添加元素到集合内
func (this *LinkedListSet)Add(e string)  {
	if !this.linkedList.Contains(e){
		this.linkedList.AddFirst(e) // 添加到链表头部
	}

}
//查询集合内,是否存在元素e
func (this *LinkedListSet)Contains(e string) bool{
	return this.linkedList.Contains(e)
}
//从集合内,删除元素e
func (this *LinkedListSet)Remove(e string){
	this.linkedList.RemoveElement(e)
}
//返回集合的元素个数
func (this *LinkedListSet)GetSize() int{
	return this.linkedList.GetSize()
}
//判断集合是否为空
func (this *LinkedListSet)IsEmpty() bool{
	return this.linkedList.IsEmpty()
}