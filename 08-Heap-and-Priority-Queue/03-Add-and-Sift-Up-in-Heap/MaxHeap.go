package _3_add_and_sift_up_in_heap

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

//向堆中添加元素
func (this *MaxHeap)Add(e int)  {
	this.data.AddLast(e)
	this.siftUp(this.data.GetSize() - 1)
}
//堆的siftUp操作
func (this *MaxHeap)siftUp(k int){
	for k > 0 && this.data.Get( this.parent(k)) < this.data.Get(k) {
		//交换两个元素
		this.data.Swap(k,this.parent(k))
		//并且 更新下一个k 值
		k = this.parent(k)
	}
}

