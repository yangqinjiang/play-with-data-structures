package main

import (
	"fmt"
	_10_Level_Traverse_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/10-Level-Traverse-in-BST"
)

func main() {
	 bst := _10_Level_Traverse_in_BST.NewBST()
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
	fmt.Println()
	fmt.Print("LevelOrder=")
	bst.LevelOrder()

	//fmt.Println(bst.ToString())

	//
	//PreorderTraversal()
}
