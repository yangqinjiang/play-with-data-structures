package _2_Linked_List_Set

import "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/02-LinkedListSet/BST"

/**
集合类, 实现Set接口的BST
 */
type BSTSet struct {
	bst *bst.BST //这个BST的元素是string类型
}

func NewBSTSet() *BSTSet  {
	return &BSTSet{bst: bst.NewBST()}
}
//添加元素到集合内
func (this *BSTSet)Add(e string)  {
	this.bst.Add(e)
}
//查询集合内,是否存在元素e
func (this *BSTSet)Contains(e string) bool{
	return this.bst.Contains(e)
}
//从集合内,删除元素e
func (this *BSTSet)Remove(e string){
	this.bst.Remove(e)
}
//返回集合的元素个数
func (this *BSTSet)GetSize() int{
	return this.bst.Size()
}
//判断集合是否为空
func (this *BSTSet)IsEmpty() bool{
	return this.bst.IsEmpty()
}