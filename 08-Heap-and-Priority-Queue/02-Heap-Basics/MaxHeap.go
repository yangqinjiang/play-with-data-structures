package _2_Heap_Basics

type MaxHeap struct {
	data *Array
}

func NewMaxHeap() *MaxHeap  {
	return &MaxHeap{data: NewArrayDefault()}
}
//返回堆中的元素个数
func (this *MaxHeap)Size() int  {
	return this.data.GetSize()
}
//返回一个bool,表示堆中是否为空
func (this *MaxHeap)IsEmpty() bool  {
	return this.data.IsEmpty()
}
//返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (this *MaxHeap)parent(index int) int{
	if 0 == index{
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}
// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (this *MaxHeap)leftChild(index int) int{
	return index * 2 + 1
}
// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (this *MaxHeap)rightChild(index int) int{
	return index * 2 + 2
}

