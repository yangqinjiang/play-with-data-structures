package main

import (
	"fmt"
	_7_InOrder_and_PostOrder_Traverse_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/07-InOrder-and-PostOrder-Traverse-in-BST"
)

func main() {
	 bst := _7_InOrder_and_PostOrder_Traverse_in_BST.NewBST()
	 nums := []int{5,3,6,8,4,2}
	for _,num := range nums{
		bst.Add(num)
	}
	/////////////////
	//      5      //
	//    /   \    //
	//   3    6    //
	//  / \    \   //
	// 2  4     8  //
	/////////////////
	fmt.Print("PreOrder=")
	bst.PreOrder()
	fmt.Println()

	fmt.Print("InOrder=")
	bst.InOrder()//中序遍历,按从小到大的顺序排序
	fmt.Println()

	fmt.Print("PostOrder=")
	bst.PostOrder()
	fmt.Println()

	fmt.Println(bst.ToString())
}
