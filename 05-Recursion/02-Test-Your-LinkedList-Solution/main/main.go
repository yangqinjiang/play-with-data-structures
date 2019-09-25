package main

import (
	"fmt"
	_2_Test_Your_LinkedList_Solution "github.com/yangqinjiang/play-with-data-structures/05-Recursion/02-Test-Your-LinkedList-Solution"
)

func main() {
	nums := []int{1,2,3,4,5,6}
	head := _2_Test_Your_LinkedList_Solution.NewListNodeByArray(nums)
	fmt.Println("origin :",head.ToString())

	res := _2_Test_Your_LinkedList_Solution.RemoveElements(head,1)
	fmt.Println("after delete 1 :",res.ToString())

	res = _2_Test_Your_LinkedList_Solution.RemoveElements2(res,2)
	fmt.Println("after delete 2 :",res.ToString())

	res = _2_Test_Your_LinkedList_Solution.RemoveElements3(res,3)
	fmt.Println("after delete 3 :",res.ToString())
	if "4->5->6->NULL" != res.ToString(){
		panic("remove elements error")
	}else{
		fmt.Println("remove elements success")
	}
}
