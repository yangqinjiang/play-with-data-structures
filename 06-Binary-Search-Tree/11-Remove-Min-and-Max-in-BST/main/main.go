package main

import (
	"fmt"
	_11_Remove_Min_and_Max_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/11-Remove-Min-and-Max-in-BST"
)

func main() {
	 bst := _11_Remove_Min_and_Max_in_BST.NewBST()
	 nums := []int{5,3,6,8,4,2,1,10}
	for _,num := range nums{
		bst.Add(num)
	}

	fmt.Print("PreOrder=")
	bst.PreOrder()
	fmt.Println()
	fmt.Print("PreOrderNR=")
	bst.PreOrderNR()
	fmt.Println()
	fmt.Print("LevelOrder=")
	bst.LevelOrder()

	if 1 != bst.Minimun(){
		panic("minValue should be 1")
	}
	if 10 != bst.Maximun(){
		panic("maxValue should be 10")
	}
	//第二次
	//删除最小,最大值
	bst.RemoveMin()
	bst.RemoveMax()
	if 2 != bst.Minimun(){
		panic("minValue should be 2")
	}
	if 8 != bst.Maximun(){
		panic("maxValue should be 8")
	}
	//第二次
	//删除最小,最大值
	bst.RemoveMin()
	bst.RemoveMax()
	if 3 != bst.Minimun(){
		panic("minValue should be 3")
	}
	if 6 != bst.Maximun(){
		panic("maxValue should be 6")
	}
}
