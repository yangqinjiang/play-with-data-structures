package main

import (
	"fmt"
	_3_Recursion_Basics "github.com/yangqinjiang/play-with-data-structures/05-Recursion/03-Recursion-Basics"
)

func main() {
	s := _3_Recursion_Basics.NewSum()
	ss := s.Sum([]int{1,2,3,4,5,6})
	if ss != 21{
		panic("sum error")
	}
	fmt.Println("sum success eq ",ss)

	nums := []int{1,2,3,4,5,6}
	head := _3_Recursion_Basics.NewListNodeByArray(nums)
	fmt.Println("origin :",head.ToString())

	res := _3_Recursion_Basics.RemoveElements(head,1)
	fmt.Println("after delete 1 :",res.ToString())

	res = _3_Recursion_Basics.RemoveElements2(res,2)
	fmt.Println("after delete 2 :",res.ToString())

	res = _3_Recursion_Basics.RemoveElements3(res,3)
	fmt.Println("after delete 3 :",res.ToString())
	if "4->5->6->NULL" != res.ToString(){
		panic("remove elements error")
	}else{
		fmt.Println("remove elements success")
	}

}
