package main

import "fmt"

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	res := intersection(nums1,nums2)
	fmt.Println(res)
	if len(res) != 1{
		panic("len isnot 1")
	}
	if res[0] != 2{
		panic("index 0 isnot 2")
	}
	fmt.Println("-----OK--------")

	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	res = intersection(nums1,nums2)
	fmt.Println(res)
	if len(res) != 2{
		panic("len isnot 2")
	}
	if res[0] != 9{
		panic("index 0 isnot 9")
	}
	if res[1] != 4{
		panic("index 0 isnot 4")
	}
	fmt.Println("-----OK--------")
}
//给定两个数组，编写一个函数来计算它们的交集。
/// Leetcode 349. Intersection of Two Arrays
/// https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	//以下代码, 使用map 映射表  , 模拟集合 Set
	_map := map[int]int{}
	//遍历第一个数组
	for _,num := range nums1  {
		_map[num] = 1
	}

	res := []int{}
	for _,num := range  nums2{
		if _,exist := _map[num];exist{
			res = append(res, num) // 有重复的,保存此值
			delete(_map,num)  //删除重复的值
		}
	}
	return res
}