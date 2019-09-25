package main

import (
	"fmt"
	_4_LinkedList_and_Recursion "github.com/yangqinjiang/play-with-data-structures/05-Recursion/04-LinkedList-and-Recursion"
)

func main() {
	s := _4_LinkedList_and_Recursion.NewSum()
	ss := s.Sum([]int{1,2,3,4,5,6})
	if ss != 21{
		panic("sum error")
	}
	fmt.Println("sum success eq ",ss)

	nums := []int{1,2,3,4,5,6}
	head := _4_LinkedList_and_Recursion.NewListNodeByArray(nums)
	fmt.Println("origin :",head.ToString())

	res := _4_LinkedList_and_Recursion.RemoveElements(head,1)
	fmt.Println("after delete 1 :",res.ToString())

	res = _4_LinkedList_and_Recursion.RemoveElements2(res,2)
	fmt.Println("after delete 2 :",res.ToString())

	res = _4_LinkedList_and_Recursion.RemoveElements3(res,3)
	fmt.Println("after delete 3 :",res.ToString())

	res = _4_LinkedList_and_Recursion.RemoveElements4(res,4)
	fmt.Println("after delete 4 :",res.ToString())

	res = _4_LinkedList_and_Recursion.RemoveElements5(res,5)
	fmt.Println("after delete 5 :",res.ToString())
	if "6->NULL" != res.ToString(){
		panic("remove elements error")
	}else{
		fmt.Println("remove elements success")
	}
}
