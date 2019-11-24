package _5_Optimized_by_Rank

// 我们的第三版Union-Find
type UnionFind4 struct {
	//我们的第二版Union-Find,使用一个数组,构建一棵指向父节点的树
	// parent[i]表示第一个元素所指向的父节点
	parent []int
	//rank[i]表示以i为根的集合所表示的树的层数
	rank []int
}

func NewUnionFind4(size int) *UnionFind4  {
	parent := make([]int,size)
	rank := make([]int,size)
	//初始化,每一个parent[i]指向自己,表示每一个元素自己自成一个集合
	for i:=0;i<size ;i++  {
		parent[i] = i
		rank[i] = i
	}
	return &UnionFind4{parent:parent,rank:rank}
}

//实现UF接口
func (this *UnionFind4)GetSize() int  {
	return len(this.parent)
}

//查找元素p所对应的集合编号
// O(h)复杂度,h为树的高度
func (this *UnionFind4)find(p int) int  {
	if p < 0 || p >= len(this.parent){
		panic("p is out of bound.")
	}
	//不断去查询自己的父节点,直到到达根节点
	//根节点的特点: parent[p] == p , 即自己指向自己
	for p != this.parent[p]{
		p = this.parent[p]
	}
	return p
}

//查看元素p和元素q是否所属一个集合
// O(h)复杂度,h为树的高度
func (this *UnionFind4)IsConnected(p,q int) bool  {
	return  this.find(p) == this.find(q)
}
//合并元素p和q所属的集合
// O(n)复杂度
func (this *UnionFind4)UnionElements(p,q int )  {
	pRoot := this.find(p)
	qRoot := this.find(q)

	if pRoot == qRoot{
		return
	}

	//根据两个元素所在树的元素个数不同判断合并方向
	//将元素个数少的集合合并到元素个数多的集合上
	if this.rank[pRoot] < this.rank[qRoot]{
		this.parent[pRoot] = qRoot
	}else if this.rank[pRoot] >  this.rank[qRoot]{
		this.parent[qRoot] = pRoot
	}else{
		this.parent[pRoot] = qRoot
		this.rank[qRoot] += 1//此时,维护rank的值
	}

}

