package _7_more_about_resizing_in_hash_table
import (
	"github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/07-More-About-Resizing-in-Hash-Table/HashTable"
	"strconv"
)
/// Leetcode 350. Intersection of Two Arrays II
/// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的HashTable类
func Intersect(nums1 []int, nums2 []int) []int{
	return intersect(nums1,nums2)
}
func intersect(nums1 []int, nums2 []int) []int {
	//map 映射表
	
	_map := HashTable.NewHashTable()
	//遍历第一个数组
	for _,num := range nums1  {
		num_key := strconv.Itoa(num)
		//如果不存在于m, 即设置初始值为 1
		if !_map.Contains(num_key){
			_map.Add(num_key,1)
		}else{
			_map.Set(num_key,_map.Get(num_key) + 1)
		}
	}

	res := []int{}
	for _,num := range  nums2{
		num2_key := strconv.Itoa(num)
		if _map.Contains(num2_key){
			res = append(res, num) // 有重复的,保存此值
			_map.Set(num2_key,_map.Get(num2_key) - 1)
			
			//由于map的特点, 当上一句代码 执行后,
			// 如果map.num的值是0,应该del它
			if 0  == _map.Get(num2_key){
				_map.Remove(num2_key)
			}
		}
	}
	return res
}

