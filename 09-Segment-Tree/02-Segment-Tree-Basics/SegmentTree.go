package _2_Segment_Tree_Basics

/**
线段树
 */
type SegmentTree struct {
   tree []int
   data []int
}

func NewSegmentTree(arr []int) *SegmentTree  {
	data := make([]int,len(arr))
	for i:=0;i<len(arr) ;i++  {
		data[i] = arr[i]
	}
	//tree是容量是data的4倍
	tree := make([]int,4* len(arr))

	return &SegmentTree{
		data:data,
		tree:tree,
	}
}
//读取data的长度
func (this *SegmentTree)GetSize() int  {
	return len(this.data)
}
//读取data的数据
func (this *SegmentTree)Get(index int) int  {
	if index < 0 || index >= len(this.data){
		panic("Index is illegal.")
	}
	return this.data[index]
}
// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (this *SegmentTree)leftChild(index int) int  {
	return  2*index + 1
}
// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (this *SegmentTree)rightChild(index int) int  {
	return  2*index + 2
}