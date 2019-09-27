package _6_PreOrder_Traverse_in_BST

import (
	"fmt"
	"strconv"
	"strings"
)

// 树结点
type Node struct {
	E int
	Left,Right *Node
}

func NewNode(e int) *Node  {
	return &Node{
		E:e,
		Left:nil,
		Right:nil,
	}
}

// 二分搜索树
type BST struct {
	root *Node
	size int
}

func NewBST() *BST{
	return &BST{
		root:nil,
		size:0,
	}
}

func (this *BST)Size() int {
	return this.size
}

func (this *BST)IsEmpty() bool {
	return 0 == this.size
}
//优化
//向二分搜索树中添加新的元素e
func (this *BST)Add(e int)  {
	this.root = this._Add(this.root,e)
}
//优化
// 向以node为根的二分搜索树中插入元素e，递归算法
// 返回插入新节点后二分搜索树的根
func (this *BST)_Add(node *Node,e int) *Node  {
	if nil == node{
		this.size ++
		return &Node{E:e}
	}

	if e < node.E{
		node.Left = this._Add(node.Left,e)
	}else if e > node.E{
		node.Right = this._Add(node.Right,e)
	}
	//相等,不处理

	return  node
}
// 看二分搜树中是否包含元素e
func (this *BST)Contains(e int) bool {
	return this._contains(this.root,e)
}
// 看以node为根的二分搜索树中是否包含元素e, 递归算法
func (this *BST)_contains(node *Node,e int) bool {

	if nil == node{//查找不到
		return false
	}
	if e == node.E{//已经查找
		return true
	}else if e < node.E{//,查找左子树
		return this._contains(node.Left,e)
	}else { // e > node.E,查找右子树
		return this._contains(node.Right,e)
	}
}
//二分搜索树的前序遍历
func (this *BST)PreOrder()  {
	this._preOrder(this.root)
}

//前序遍历以node为根的二分搜索树,递归算法
func (this *BST)_preOrder(node *Node)  {
	if nil == node{
		return
	}
	fmt.Println(node.E)
	this._preOrder(node.Left)
	this._preOrder(node.Right)
}

func (this *BST)ToString() string  {
	var res strings.Builder
	this.generateBSTString(this.root,0,&res)
	return res.String()
}
// 生成以node为根节点，深度为depth的描述二叉树的字符串
func (this *BST)generateBSTString(node *Node,depth int,res *strings.Builder)  {
	if node == nil{
		res.WriteString(generateDepthString(depth) + "null\n")
		return
	}
	res.WriteString(generateDepthString(depth) + strconv.Itoa(node.E) + "\n")
	this.generateBSTString(node.Left,depth+1,res)
	this.generateBSTString(node.Right,depth+1,res)
}
//创建字符串
func generateDepthString(depth int) string{
	var sb strings.Builder

	for i := 0;i<depth;i++{
		sb.WriteString("--")
	}
	return sb.String()
}






