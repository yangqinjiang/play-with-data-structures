package _3_Add_Elements_in_BST
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

//向二分搜索树中添加新的元素e
func (this *BST)Add(e int)  {
	//root为空,则添加它
	if nil == this.root{

		this.root = NewNode(e)
		this.size ++
	}else{
			this._Add(this.root,e)
	}
}
// 向以node为根的二分搜索树中插入元素e，递归算法
func (this *BST)_Add(node *Node,e int)  {
	if e == node.E{
		//重复添加,直接返回
		return
	}else if e < node.E && node.Left == nil{
		//符合放到左子树且左子树为空,则添加到左子树
		node.Left = NewNode(e)
		this.size ++
		return
	} else if e > node.E && node.Right == nil{
		//符合放到右子树且右子树为空,则添加到右子树
		node.Right = NewNode(e)
		this.size ++
		return
	}

	if e < node.E{
		this._Add(node.Left,e)
	}else{// e> node.E
		this._Add(node.Right,e)
	}
}











