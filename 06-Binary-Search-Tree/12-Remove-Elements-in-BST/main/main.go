package main

import (
	"fmt"
	_12_Remove_Elements_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/12-Remove-Elements-in-BST"
)

func main() {
	 bst := _12_Remove_Elements_in_BST.NewBST()
	 nums := []int{5,3,6,8,4,2,1,10}
	for _,num := range nums{
		bst.Add(num)
	}
	//fmt.Println(bst.ToString())
	bst.Remove(3)
	bst.Remove(6)
	fmt.Println(bst.ToString())
	if 1 != bst.Minimun(){
		panic("minValue should be 1")
	}
	if 10 != bst.Maximun(){
		panic("maxValue should be 10")
	}
}
