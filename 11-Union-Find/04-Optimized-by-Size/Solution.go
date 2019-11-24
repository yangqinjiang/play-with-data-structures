package _4_Optimized_by_Size

/// Leetcode 547. Friend Circles
/// https://leetcode.com/problems/friend-circles/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的UF类
func findCircleNum(M [][]int) int {
	n := len(M)
	uf := NewUnionFind3(n)
	for i:=0;i<n ;i++  {
		for j:=0;j<i ;j++  {
			if 1 ==M[i][j]{
				uf.UnionElements(i,j)
			}
		}
	}
	set := make(map[int]int)
	for i:=0;i<n ;i++  {
		set[uf.find(i)] = 1  // uf.find 查询父节点
	}

	return len(set)
}

func FindCircleNum(M [][]int) int {
	return findCircleNum(M)
}
