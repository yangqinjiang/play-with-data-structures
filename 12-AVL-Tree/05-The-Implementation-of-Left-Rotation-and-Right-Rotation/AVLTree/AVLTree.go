package AVLTree

import (
	//"fmt"
	"math"
	"strconv"
	"strings"
)



type Node struct {
	Key string
	Value int
	Left,Right *Node
	height int // 节点高度
}

func NewNode(key string,value int) *Node  {
	return &Node{
		Key:key,
		Value:value,
		Left:nil,
		Right:nil,
		height:1,
	}
}


type AVLTree struct {
	root *Node
	size int
}

func NewAVLTree() *AVLTree{
	return &AVLTree{
		root:nil,
		size:0,
	}
}
//接口:
func (this *AVLTree)GetSize() int {
	return this.size
}
//接口:
func (this *AVLTree)IsEmpty() bool {
	return 0 == this.size
}
//判断该二叉树是否是一棵二分搜索树
func (this *AVLTree)IsBST() bool {
	keys := []string{}
	this.inOrder(this.root,keys)
	for i:=1;i< len(keys);i++{
		if keys[i-1] > keys[i]{
			return false
		}
	}
	return true
}
//中序排序
func (this *AVLTree)inOrder(node *Node,keys []string)  {
	if nil == node{
		return
	}
	this.inOrder(node.Left,keys)
	keys = append(keys, node.Key)
	this.inOrder(node.Right,keys)

}
// 判断该二叉树是否是一棵平衡二叉树
func (this *AVLTree)IsBalanced() bool {
	return  this._isBalanced(this.root)
}
// 判断以Node为根的二叉树是否是一棵平衡二叉树，递归算法
func (this *AVLTree)_isBalanced(node *Node) bool {
	if nil == node{
		return  true
	}
	balanceFactor := this.getBalanceFactor(node)

	//大于 1,说明不平衡
	if math.Abs(float64(balanceFactor)) > 1.0 {
		return false
	}
	return  this._isBalanced(node.Left) && this._isBalanced(node.Right)
}
//返回节点的高度
func (this *AVLTree)getHeight(node *Node) int {
	if nil == node{
		return 0
	}
	return node.height
}

//获得节点node的平衡因子
func (this *AVLTree)getBalanceFactor(node *Node) int {
	if nil == node{
		return 0
	}
	//左节点的height 与右节点的height 之差
	return this.getHeight(node.Left) - this.getHeight(node.Right)
}
    // 对节点y进行向右旋转操作，返回旋转后新的根节点x
    //        y                              x
    //       / \                           /   \
    //      x   T4     向右旋转 (y)        z     y
    //     / \       - - - - - - - ->    / \   / \
    //    z   T3                       T1  T2 T3 T4
    //   / \
	// T1   T2
	func (this *AVLTree)rightRotate(y *Node) *Node{
		x := y.Left
		T3 := x.Right

		//向右旋转过程
		x.Right = y
		y.Left = T3

		//更新height
		y.height = 1 + this.max(this.getHeight(y.Left),this.getHeight(y.Right))
		x.height = 1 + this.max(this.getHeight(x.Left),this.getHeight(x.Right))

		return x // 返回以x为根的子树
	}
    // 对节点y进行向左旋转操作，返回旋转后新的根节点x
    //    y                             x
    //  /  \                          /   \
    // T1   x      向左旋转 (y)       y     z
    //     / \   - - - - - - - ->   / \   / \
    //   T2  z                     T1 T2 T3 T4
    //      / \
	//     T3 T4
	func (this *AVLTree)leftRotate(y *Node) *Node{
		x := y.Right
		T2 := x.Left

		//向左旋转过程
		x.Left = y
		y.Right = T2

		//更新height
		y.height = 1 + this.max(this.getHeight(y.Left),this.getHeight(y.Right))
		x.height = 1 + this.max(this.getHeight(x.Left),this.getHeight(x.Right))

		return x // 返回以x为根的子树
	}
//优化
//接口:向二分搜索树中添加新的元素e
func (this *AVLTree)Add(key string,value int)  {
	this.root = this._Add(this.root,key,value)
}
//优化
// 向以node为根的二分搜索树中插入元素e，递归算法
// 返回插入新节点后二分搜索树的根
func (this *AVLTree)_Add(node *Node,key string,value int) *Node  {
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

	//更新节点的height
	node.height = 1 + this.max(this.getHeight(node.Left),this.getHeight(node.Right))

	//计算平衡因子
	balanceFactor := this.getBalanceFactor(node)
	// if math.Abs(float64(balanceFactor))  > 1{
	// 	fmt.Printf("unbalanced : %d\n",balanceFactor)
	// }
	//平衡维护
	if balanceFactor > 1 && this.getBalanceFactor(node.Left) >= 0{
		//fmt.Println("rightRotate")
		return this.rightRotate(node) //RR
	}
	if balanceFactor < -1 && this.getBalanceFactor(node.Right) <= 0{
		//fmt.Println("leftRotate")
		return this.leftRotate(node) // LL
	}

	return  node
}
//求最大值
func (this *AVLTree)max(x,y int) int  {
	if x > y{
		return  x
	}
	return  y
}
//返回以node为根节点的二分搜索树中,key所在的节点, 递归算法
func (this *AVLTree)getNode(node *Node,key string) *Node {
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
func (this *AVLTree)Contains(key string) bool {
	return this.getNode(this.root,key) != nil
}
//接口:获取k-v,不存在则返回 0
func (this *AVLTree)Get(key string) int  {
	node := this.getNode(this.root,key)
	if node == nil{
		return 0
	}else{
		return node.Value
	}
}
//接口:更新k-v
func (this *AVLTree)Set(key string,newValue int)  {
	node := this.getNode(this.root,key)
	if node == nil{
		panic(key + "  doesn't exist!")
	}
	node.Value = newValue
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (this *AVLTree)_minimun(node *Node) *Node  {
	//如果以node为根的左子树为空,则说明node是最小值
	if nil == node.Left{
		return node
	}
	//递归查询
	return this._minimun(node.Left)
}

// 寻找二分搜索树的最大元素
func (this *AVLTree)Maximun() string  {
	if 0 == this.size{
		panic("BST is empty")
	}
	maxNode := this._maximun(this.root)
	return maxNode.Key
}
// 返回以node为根的二分搜索树的最大值所在的节点
func (this *AVLTree)_maximun(node *Node) *Node  {
	//如果以node为根的右子树为空,则说明node是最大值
	if nil == node.Right{
		return node
	}
	//递归查询
	return this._maximun(node.Right)
}
// 从二分搜索树中删除最小值所在节点, 返回最小值
func (this *AVLTree)RemoveMin() string  {
	ret := this.Minimun()
	this.root = this._removeMin(this.root)
	return ret
}
// 寻找二分搜索树的最小元素
func (this *AVLTree)Minimun() string  {
	if 0 == this.size{
		panic("BST is empty")
	}
	minNode := this._minimun(this.root)
	return minNode.Key
}
// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *AVLTree)_removeMin(node *Node) *Node  {
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
func (this *AVLTree)RemoveMax() string  {
	ret := this.Maximun()
	this.root = this._removeMax(this.root)
	return ret
}
// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (this *AVLTree)_removeMax(node *Node) *Node  {
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
//接口:从二分搜索树中删除元素为e的节点
func (this *AVLTree)Remove(key string) int {
	node := this.getNode(this.root,key)
	if node != nil{
		this.root = this._remove(this.root,key)
		return node.Value
	}
	return 0
}
//删除掉以node为根的二分搜索树中,值为e的节点,递归算法
// 返回删除节点后,新的二分搜索树的根
func (this *AVLTree)_remove(node *Node,key string) *Node  {
	if nil == node{
		return nil
	}
	if key < node.Key{ //从左子树查询,删除
		node.Left = this._remove(node.Left,key)
		return node
	}else if key > node.Key{//从右子树查询,删除
		node.Right = this._remove(node.Right,key)
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
func (this *AVLTree)ToString() string  {
	var res strings.Builder
	this.generateBSTString(this.root,0,&res)
	return res.String()
}
// 生成以node为根节点，深度为depth的描述二叉树的字符串
func (this *AVLTree)generateBSTString(node *Node,depth int,res *strings.Builder)  {
	if node == nil{
		res.WriteString(generateDepthString(depth) + "null\n")
		return
	}
	res.WriteString(generateDepthString(depth) + node.Key + " : "  + strconv.Itoa(node.Value) + "\n")
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
