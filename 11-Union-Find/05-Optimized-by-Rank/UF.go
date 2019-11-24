package _5_Optimized_by_Rank
// 并查集  接口
type UF interface {
	GetSize() int //获取元素数量
	IsConnected( p,  q int) bool //判断p,q两点是否相连
	UnionElements( p, q int) // 合并两点
}