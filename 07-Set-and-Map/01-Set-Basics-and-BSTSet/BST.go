package _1_Set_Basics_and_BSTSet

import (
	"container/list"
	"fmt"
	"strings"
)
//TODO:Node的E属性,应该实现泛型
// 树结点
type Node struct {
	E string
	Left,Right *Node
}

func NewNode(e string) *Node  {
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
func (this *BST)Add(e string)  {
	this.root = this._Add(this.root,e)
}
//优化
// 向以node为根的二分搜索树中插入元素e，递归算法
// 返回插入新节点后二分搜索树的根
func (this *BST)_Add(node *Node,e string) *Node  {
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
func (this *BST)Contains(e string) bool {
	return this._contains(this.root,e)
}
// 看以node为根的二分搜索树中是否包含元素e, 递归算法
func (this *BST)_contains(node *Node,e string) bool {

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
// 二分搜索树的非递归前序遍历
func (this *BST)PreOrderNR()  {
	if nil == this.root{
		return
	}
	//使用list模拟Stack
	stack := list.New()
	stack.PushBack(this.root)
	for stack.Len() > 0{
		cur := stack.Remove(stack.Back()).(*Node)
		fmt.Println(cur.E)
		//注意,先压入右子树
		if nil != cur.Right{
			stack.PushBack(cur.Right)
		}
		//再压入右子树
		if nil != cur.Left{
			stack.PushBack(cur.Left)
		}
	}
}

//二分搜索树的中序遍历
func (this *BST)InOrder()  {
	this._inOrder(this.root)
}

//中序遍历以node为根的二分搜索树,递归算法
func (this *BST)_inOrder(node *Node)  {
	if nil == node{
		return
	}
	this._inOrder(node.Left)
	fmt.Println(node.E)
	this._inOrder(node.Right)
}
//二分搜索树的后序遍历
func (this *BST)PostOrder()  {
	this._postOrder(this.root)
}

//后序遍历以node为根的二分搜索树,递归算法
func (this *BST)_postOrder(node *Node)  {
	if nil == node{
		return
	}
	this._postOrder(node.Left)
	this._postOrder(node.Right)
	fmt.Println(node.E)
}

// 二分搜索树的层序遍历
func (this *BST)LevelOrder()  {
	if nil == this.root{
		return
	}
	queue := list.New()
	queue.PushFront(this.root)
	for queue.Len() > 0{
		cur := queue.Remove(queue.Back()).(*Node)//出队列
		fmt.Println(cur.E)

		//左子树
		if nil != cur.Left{
			queue.PushFront(cur.Left)//入队列
		}
		//右子树
		if nil != cur.Right{
			queue.PushFront(cur.Right)//入队列
		}
	}
}

// 寻找二分搜索树的最小元素
func (this *BST)Minimun() string  {
	if 0 == this.size{
		panic("BST is empty")
	}
	minNode := this._minimun(this.root)
	return minNode.E
}
// 返回以node为根的二分搜索树的最小值所在的节点
func (this *BST)_minimun(node *Node) *Node  {
	//如果以node为根的左子树为空,则说明node是最小值
	if nil == node.Left{
		return node
	}
	//递归查询
	return this._minimun(node.Left)
}

// 寻找二分搜索树的最大元素
func (this *BST)Maximun() string  {
	if 0 == this.size{
		panic("BST is empty")
	}
	maxNode := this._maximun(this.root)
	return maxNode.E
}
// 返回以node为根的二分搜索树的最大值所在的节点
func (this *BST)_maximun(node *Node) *Node  {
	//如果以node为根的右子树为空,则说明node是最大值
	if nil == node.Right{
		return node
	}
	//递归查询
	return this._maximun(node.Right)
}
// 从二分搜索树中删除最小值所在节点, 返回最小值
func (this *BST)RemoveMin() string  {
	ret := this.Minimun()
	this.root = this._removeMin(this.root)
	return ret
}
// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *BST)_removeMin(node *Node) *Node  {
	//判断左子树
	if nil == node.Left{
		rightNode := node.Right
		node.Right = nil
		this.size --
		return rightNode
	}
	node.Left = this._removeMin(node.Left)
	return node
}
// 从二分搜索树中删除最大值所在节点
func (this *BST)RemoveMax() string  {
	ret := this.Maximun()
	this.root = this._removeMax(this.root)
	return ret
}
// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (this *BST)_removeMax(node *Node) *Node  {
	//判断右子树
	if nil == node.Right{
		leftNode := node.Left
		node.Left = nil
		this.size --
		return leftNode
	}
	node.Right = this._removeMax(node.Right)
	return node
}

//使用后继节点的方式,删除任意元素
//从二分搜索树中删除元素为e的节点
func (this *BST)Remove(e string)  {
	this.root = this._remove(this.root,e)
}
//删除掉以node为根的二分搜索树中,值为e的节点,递归算法
// 返回删除节点后,新的二分搜索树的根
func (this *BST)_remove(node *Node,e string) *Node  {
	if nil == node{
		return nil
	}
	if e < node.E{ //从左子树查询,删除
		node.Left = this._remove(node.Left,e)
		return node
	}else if e > node.E{//从右子树查询,删除
		node.Right = this._remove(node.Right,e)
		return node
	}else{ // e == node.E
		//待删除节点左子树为空的情况
		if nil == node.Left{
			rightNode := node.Right
			node.Right = nil //删除右子树的引用
			this.size --
			return rightNode
		}
		//待删除节点右子树为空的情况
		if nil == node.Right{
			leftNode := node.Left
			node.Left = nil //删除左子树的引用
			this.size --
			return leftNode
		}
		//查找后继节点,顶替 [删除的元素]
		//待删除节点 左右子树均不为空的情况
		//找到比待删除节点 大的最小节点,即待删除节点右子树的最小节点
		//用这个节点顶替待删除节点的位置
		successor := this._minimun(node.Right)
		successor.Right = this._removeMin(node.Right)
		successor.Left = node.Left

		//删除 [待删除节点]的左右子树引用
		node.Left = nil
		node.Right = nil
		return  successor //返回后继节点
	}
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
	res.WriteString(generateDepthString(depth) + node.E + "\n")
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






