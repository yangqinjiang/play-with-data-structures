package _4_Extract_and_Sift_Down_in_Heap

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
//看堆中的最大元素,即index为0的元素
func (this *MaxHeap)FindMax() int{
	if (0 == this.data.GetSize()){
		panic("Can not findMax when heap is empty.")
	}
	return this.data.Get(0)
}

//取出堆中最大元素
func (this *MaxHeap)ExtractMax() int{
	ret := this.FindMax()
	//将最大元素 与 最后索引的元素  交换
	this.data.Swap(0,this.data.GetSize() - 1)
	this.data.RemoveLast() // 删除最后一个元素,即最大元素
	this.siftDown(0)
	return ret
}
//数据下沉
func (this *MaxHeap)siftDown(k int){
	// for循环的终止条件是: 有左子节点, 否则没有可以siftDown的子节点了
	for this.leftChild(k) < this.data.GetSize() {
		// 在此轮循环中,data[k]和data[j]交换位置
		j := this.leftChild(k)

		//判断左右子节点的元素, 哪个比较大,
		//如果 [存在右子节点]  并且 [右子节点比较大],则 j += 1, 即将要交换位置的索引 更新为 右子节点的
		if j+ 1 < this.data.GetSize() && this.data.Get(j+1) > this.data.Get(j){
			j = this.rightChild(k)
		}
		// data[j] 是 leftChild 和 rightChild 中的最大值

		//如果父节点的值,比子节点的值 更大, 则不需要交换
		if this.data.Get(k) > this.data.Get(j){
			break
		}

		//否则,父节点小于子节点, 交换
		this.data.Swap(k,j)
		k = j
	}
}

