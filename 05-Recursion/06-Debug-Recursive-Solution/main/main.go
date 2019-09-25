package main

import (
	"fmt"
	_6_Debug_Recursive_Solution "github.com/yangqinjiang/play-with-data-structures/05-Recursion/06-Debug-Recursive-Solution"
)

func main() {


	nums := []int{1,2,3,4,5,6}
	head := _6_Debug_Recursive_Solution.NewListNodeByArray(nums)
	fmt.Println("origin :",head.ToString())

	res := _6_Debug_Recursive_Solution.RemoveElements(head,6,0)
	fmt.Println("after delete 1 :",res.ToString())

}
