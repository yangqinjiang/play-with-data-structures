package RBTree


import (
"strconv"
"strings"
)
//颜色的标记
const RED bool = true
const BLACK  bool= false
//TODO:Node的E属性,应该实现泛型
//这个RBTree的元素是string类型
// 树结点
type Node struct {
	Key string
	Value int
	Left,Right *Node
	Color bool
}

func NewNode(key string,value int) *Node  {
	return &Node{
		Key:key,
		Value:value,
		Left:nil,
		Right:nil,
		Color:RED, //default red
	}
}

// 红黑树
type RBTree struct {
	root *Node
	size int
}

func NewRBTree() *RBTree{
	return &RBTree{
		root:nil,
		size:0,
	}
}

func (this *RBTree)GetSize() int {
	return this.size
}
func (this *RBTree)IsEmpty() bool {
	return 0 == this.size
}
//判断节点node的颜色
func (this *RBTree)isRed(node *Node) bool  {
	if nil == node{
		return BLACK
	}
	return node.Color
}
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (this *RBTree)leftRotate(node *Node) *Node  {
	x := node.Right

	//左旋转
	node.Right = x.Left
	x.Left = node

	//颜色
	x.Color = node.Color
	node.Color = RED

	return x
}
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (this *RBTree)rightRotate(node *Node) *Node  {
	x := node.Left

	//右旋转
	node.Left = x.Right
	x.Right = node

	//颜色
	x.Color = node.Color
	node.Color = RED

	return x
}

//颜色翻转
func (this *RBTree)flipColors(node *Node)  {
	node.Color = RED
	node.Left.Color = BLACK
	node.Left.Color = BLACK
}
// 向红黑树中添加新的元素(key,value)
func (this *RBTree)Add(key string,value int)  {
	this.root = this._Add(this.root,key,value)
	this.root.Color = BLACK //最终根节点为黑色节点
}
//优化
// 向以node为根的红黑树中插入元素e，递归算法
// 返回插入新节点后红黑树的根
func (this *RBTree)_Add(node *Node,key string,value int) *Node  {
	if nil == node{
		this.size ++
		return &Node{Key:key,Value:value}
	}

	if key < node.Key{
		node.Left = this._Add(node.Left,key,value)
	}else if key > node.Key{
		node.Right = this._Add(node.Right,key,value)
	}else{//相等,更新value值
		node.Value = value
	}
	//右节点为红色, 左节点不是红色
	if this.isRed(node.Right) && !this.isRed(node.Left){
		node = this.leftRotate(node)
	}
	//左节点为红色, 左左节点也是红色
	if this.isRed(node.Left) && this.isRed(node.Left.Left){
		node = this.rightRotate(node)
	}
	//左右节点都是红色
	if this.isRed(node.Left) && this.isRed(node.Right){
		this.flipColors(node)
	}

	return  node
}
//返回以node为根节点的红黑树中,key所在的节点, 递归算法
func (this *RBTree)getNode(node *Node,key string) *Node {
	if node == nil{
		return nil
	}
	if key == node.Key{//找到
		return node
	}else if key < node.Key{//,查找左子树
		return this.getNode(node.Left,key)
	}else{ // key > node.key E,查找右子树
		return this.getNode(node.Right,key)
	}
}
//接口:看二分搜树中是否包含元素e
func (this *RBTree)Contains(key string) bool {
	return this.getNode(this.root,key) != nil
}
//接口:获取k-v,不存在则返回 0
func (this *RBTree)Get(key string) int  {
	node := this.getNode(this.root,key)
	if node == nil{
		return 0
	}else{
		return node.Value
	}
}
//接口:更新k-v
func (this *RBTree)Set(key string,newValue int)  {
	node := this.getNode(this.root,key)
	if node == nil{
		panic(key + "  doesn't exist!")
	}
	node.Value = newValue
}

// 返回以node为根的红黑树的最小值所在的节点
func (this *RBTree)_minimun(node *Node) *Node  {
	//如果以node为根的左子树为空,则说明node是最小值
	if nil == node.Left{
		return node
	}
	//递归查询
	return this._minimun(node.Left)
}

// 寻找红黑树的最大元素
func (this *RBTree)Maximun() string  {
	if 0 == this.size{
		panic("BSTMap is empty")
	}
	maxNode := this._maximun(this.root)
	return maxNode.Key
}
// 返回以node为根的红黑树的最大值所在的节点
func (this *RBTree)_maximun(node *Node) *Node  {
	//如果以node为根的右子树为空,则说明node是最大值
	if nil == node.Right{
		return node
	}
	//递归查询
	return this._maximun(node.Right)
}
// 从红黑树中删除最小值所在节点, 返回最小值
func (this *RBTree)RemoveMin() string  {
	ret := this.Minimun()
	this.root = this._removeMin(this.root)
	return ret
}
// 寻找红黑树的最小元素
func (this *RBTree)Minimun() string  {
	if 0 == this.size{
		panic("RBTree is empty")
	}
	minNode := this._minimun(this.root)
	return minNode.Key
}
// 删除掉以node为根的红黑树中的最小节点
// 返回删除节点后新的红黑树的根
func (this *RBTree)_removeMin(node *Node) *Node  {
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
// 从红黑树中删除最大值所在节点
func (this *RBTree)RemoveMax() string  {
	ret := this.Maximun()
	this.root = this._removeMax(this.root)
	return ret
}
// 删除掉以node为根的红黑树中的最大节点
// 返回删除节点后新的红黑树的根
func (this *RBTree)_removeMax(node *Node) *Node  {
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
//接口:从红黑树中删除元素为e的节点
func (this *RBTree)Remove(key string) int {
	panic("No remove in RBTree!")
}
func (this *RBTree)ToString() string  {
	var res strings.Builder
	this.generateRBTreeString(this.root,0,&res)
	return res.String()
}
// 生成以node为根节点，深度为depth的描述RBTree的字符串
func (this *RBTree)generateRBTreeString(node *Node,depth int,res *strings.Builder)  {
	if node == nil{
		res.WriteString(generateDepthString(depth) + "null\n")
		return
	}
	res.WriteString(generateDepthString(depth) + node.Key + " : "  + strconv.Itoa(node.Value) + "\n")
	this.generateRBTreeString(node.Left,depth+1,res)
	this.generateRBTreeString(node.Right,depth+1,res)
}
//创建字符串
func generateDepthString(depth int) string{
	var sb strings.Builder

	for i := 0;i<depth;i++{
		sb.WriteString("--")
	}
	return sb.String()
}





