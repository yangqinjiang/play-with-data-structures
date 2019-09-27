package main

import (
	"fmt"
	_6_PreOrder_Traverse_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/06-PreOrder-Traverse-in-BST"
)

func main() {
	 bst := _6_PreOrder_Traverse_in_BST.NewBST()
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
	bst.PreOrder()
	fmt.Println()
	fmt.Println(bst.ToString())
}
