package _4_TreeSet_and_Set_Problems_in_Leetcode

import Linked_List_Set "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/04-TreeSet-and-Set-Problems-in-Leetcode/LinkedListSet"

/**
集合类, 实现Set接口的LinkedList
*/
type LinkedListSet struct {
	linkedList * Linked_List_Set.LinkedList //这个LinkedList的元素是string类型
}

func NewLinkedListSet() *LinkedListSet  {
	return &LinkedListSet{linkedList: Linked_List_Set.NewLinkedList()}
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