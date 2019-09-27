package main

import (
	"fmt"
	_9_Non_Recursion_Preorder_Traverse_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/09-Non-Recursion-Preorder-Traverse-in-BST"
)

func main() {
	 bst := _9_Non_Recursion_Preorder_Traverse_in_BST.NewBST()
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
	fmt.Print("PreOrderNR=")
	bst.PreOrderNR()

	//fmt.Println(bst.ToString())

	//
	//PreorderTraversal()
}
