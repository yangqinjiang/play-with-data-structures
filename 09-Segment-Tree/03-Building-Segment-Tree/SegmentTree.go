package _3_Building_Segment_Tree

import (
	"fmt"
	"strings"
)

/**
线段树
 */
type SegmentTree struct {
   tree []int
   data []int
   merger Merger
}

func NewSegmentTree(arr []int,merger Merger) *SegmentTree  {
	data := make([]int,len(arr))
	for i:=0;i<len(arr) ;i++  {
		data[i] = arr[i]
	}
	//tree是容量是data的4倍
	tree := make([]int,4* len(arr))
	fmt.Println(len(arr),len(tree))

	st := &SegmentTree{
		data:data,
		tree:tree,
		merger:merger,
	}
	fmt.Println(" ,len(tree) = ", len(st.tree),st.tree, " ,st.data=",st.data)
	//构建线段树
	st.buildSegmentTree(0,0,len(arr) -1)

	return st
}
//在treeIndex的位置创建区间[l...r]的线段树
func (this *SegmentTree)buildSegmentTree(treeIndex,l,r int)  {
	//递归终止条件
	if l==r{
		fmt.Println("l==r",l,r," ,len(tree) = ", len(this.tree))
		this.tree[treeIndex] = this.data[l]
		return
	}
	fmt.Println("treeIndex=",treeIndex," ,l=",l ," ,r=",r)
	leftTreeIndex := this.leftChild(treeIndex)
	rightTreeIndex := this.rightChild(treeIndex)
	fmt.Println("leftTreeIndex=",leftTreeIndex , " rightTreeIndex=",rightTreeIndex)

	mid := l +(r-l)/2
	fmt.Println("left..... mid=",mid)
	this.buildSegmentTree(leftTreeIndex,l,mid)
	fmt.Println("right.....mid=",mid)
	this.buildSegmentTree(rightTreeIndex,mid+1,r)

	this.tree[treeIndex] = this.merger.Merge(this.tree[leftTreeIndex],this.tree[rightTreeIndex])
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

//打印数组字符串
func (this *SegmentTree)ToString() string  {
	var sb strings.Builder
	//fmt.Fprintf(&sb,"[")
	sb.WriteString("[")
	for i := 0;i< len(this.tree);i++  {
		if 0 != this.tree[i]{
			fmt.Fprintf(&sb,"%d",this.tree[i])
		}else {
			sb.WriteString("nil")
		}

		if i != len(this.tree)- 1{
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return  sb.String()

}