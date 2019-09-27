package _5_Search_in_BST

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







