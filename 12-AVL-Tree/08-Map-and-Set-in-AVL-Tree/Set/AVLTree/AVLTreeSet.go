package Set
import (
	AVLTree "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/AVLTree"
)
//基于AVLTree实现的 映射表
type AVLSet struct{
	avl *AVLTree.AVLTree
}
func NewAVLSet() *AVLSet{
	return &AVLSet{
		avl:AVLTree.NewAVLTree(),
	}
}
//实现接口
func (this *AVLSet)GetSize() int {
	return this.avl.GetSize()
}
//实现接口
func (this *AVLSet)IsEmpty() bool{
	return this.avl.IsEmpty()
}
//实现接口
func (this *AVLSet)Add(key string ){
	 this.avl.Add(key,0)
}
//实现接口
func (this *AVLSet)Contains(key string) bool{
	return this.avl.Contains(key)
}
 
 
//实现接口
func (this *AVLSet)Remove(key string) {
	 this.avl.Remove(key)
}