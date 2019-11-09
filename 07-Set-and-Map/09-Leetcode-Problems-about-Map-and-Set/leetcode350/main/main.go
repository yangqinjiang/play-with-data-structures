package main

import "fmt"

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	res := intersect(nums1,nums2)
	fmt.Println(res)
	if len(res) != 2{
		panic("len isnot 2")
	}
	if res[0] != 2{
		panic("index 0 isnot 2")
	}
	if res[1] != 2{
		panic("index 1 isnot 2")
	}
	fmt.Println("-----OK--------")
}
/// Leetcode 350. Intersection of Two Arrays II
/// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
func intersect(nums1 []int, nums2 []int) []int {
	//map 映射表
	_map := map[int]int{}
	//遍历第一个数组
	for _,num := range nums1  {
		//如果不存在于m, 即设置初始值为 1
		if _,exist := _map[num]; !exist{
			_map[num] = 1
		}else{
			_map[num] = _map[num] + 1
		}
	}

	res := []int{}
	for _,num := range  nums2{
		if _,exist := _map[num];exist{
			res = append(res, num) // 有重复的,保存此值
			_map[num] = _map[num] - 1
			//由于map的特点, 当上一句代码 执行后,
			// 如果map.num的值是0,应该del它
			if 0  == _map[num]{
				delete(_map,num)
			}
		}
	}
	return res
}