package Set
import (
	BSTMap "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map/BST"
)
/**
集合类, 实现Set接口的BST
 */
type BSTSet struct {
	bst *BSTMap.BSTMap //使用BSTMap作为BSTSet的底层实现
}

func NewBSTSet() *BSTSet  {
	return &BSTSet{bst: BSTMap.NewBSTMap()}
}
//添加元素到集合内
func (this *BSTSet)Add(e string)  {
	this.bst.Add(e,0)
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
	return this.bst.GetSize()
}
//判断集合是否为空
func (this *BSTSet)IsEmpty() bool{
	return this.bst.IsEmpty()
}