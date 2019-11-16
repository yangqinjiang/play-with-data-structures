package _4_Query_in_Segment_Tree

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
		fmt.Println("l==r",l,r," ,treeIndex=",treeIndex, " , get data[",l,"]=",this.data[l]," , return")
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
	fmt.Println("merger: leftTreeIndex +rightTreeIndex =",this.tree[leftTreeIndex] ," + ",this.tree[rightTreeIndex],"=",this.tree[treeIndex])
}
//返回区间[queryL,queryR]的值
func (this *SegmentTree)Query(queryL,queryR int) int {
	//判断索引范围
	if queryL < 0 || queryL >= len(this.data) || queryR < 0  || queryR >= len(this.data) || queryL > queryR {
		panic("Index is illegal.")
	}

	return this.query(0,0,len(this.data) - 1,queryL,queryR);
}
// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (this *SegmentTree)query(treeIndex,l,r,queryL,queryR int) int {
	if l==queryL && r == queryR{
		return this.tree[treeIndex]
	}

	mid := l + (r -l)/2
	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := this.leftChild(treeIndex)
	rightTreeIndex := this.rightChild(treeIndex)
	if queryL >= mid +1{
		return this.query(rightTreeIndex,mid+1,r,queryL,queryR)
	}else if ( queryR <= mid){
		return this.query(leftTreeIndex,l,mid,queryL,queryR)
	}
	leftResult := this.query(leftTreeIndex,l,mid,queryL,mid)
	rightResult := this.query(rightTreeIndex,mid+1,r,mid+1,queryR)

	return this.merger.Merge(leftResult,rightResult)
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
			_,err := fmt.Fprintf(&sb,"%d",this.tree[i])
			if nil != err{
				panic(err)
			}
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