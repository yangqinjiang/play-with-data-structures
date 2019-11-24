package _5_Optimized_by_Rank

// 我们的第一版Union-Find
type UnionFind1 struct {
	id []int// 我们的第一版Union-Find本质就是一个数组
}

func NewUnionFind1(size int) *UnionFind1  {
	id := make([]int,size)
	//初始化,每一个id[i]指向自己,没有合并的元素
	for i:=0;i<size ;i++  {
		id[i] = i
	}
	return &UnionFind1{id:id}
}

//实现UF接口
func (this *UnionFind1)GetSize() int  {
	return len(this.id)
}

//查找元素p所对应的集合编号
// O(1)复杂度
func (this *UnionFind1)find(p int) int  {
	if p < 0 || p >= len(this.id){
		panic("p is out of bound.")
	}
	return  this.id[p]
}

//查看元素p和元素q是否所属一个集合
// O(1)复杂度
func (this *UnionFind1)IsConnected(p,q int) bool  {
	return  this.find(p) == this.find(q)
}
//合并元素p和q所属的集合
// O(n)复杂度
func (this *UnionFind1)UnionElements(p,q int )  {
	pID := this.find(p)
	qID := this.find(q)

	if pID == qID{
		return
	}

	//合并过程需要遍历一遍所有元素,将两个元素的所属集合编号合并
	for i:=0;i< len(this.id);i++{
		if this.id[i] == pID{
			this.id[i] = qID
		}
	}
}

