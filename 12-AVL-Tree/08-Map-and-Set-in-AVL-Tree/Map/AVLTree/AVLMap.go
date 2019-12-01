package Map
import (
	AVLTree "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/AVLTree"
)
//基于AVLTree实现的 映射表
type AVLMap struct{
	avl *AVLTree.AVLTree
}
func NewAVLMap() *AVLMap{
	return &AVLMap{
		avl:AVLTree.NewAVLTree(),
	}
}
//实现接口
func (this *AVLMap)GetSize() int {
	return this.avl.GetSize()
}
//实现接口
func (this *AVLMap)IsEmpty() bool{
	return this.avl.IsEmpty()
}
//实现接口
func (this *AVLMap)Add(key string ,value int ){
	 this.avl.Add(key,value)
}
//实现接口
func (this *AVLMap)Contains(key string) bool{
	return this.avl.Contains(key)
}
//实现接口
func (this *AVLMap)Get(key string) int{
	return this.avl.Get(key)
}
//实现接口
func (this *AVLMap)Set(key string,newValue int){
	 this.avl.Set(key,newValue)
}
//实现接口
func (this *AVLMap)Remove(key string) int{
	return this.avl.Remove(key)
}