package main

import (
	"fmt"
	_3_Quick_Union "github.com/yangqinjiang/play-with-data-structures/11-Union-Find/03-Quick-Union"
)

func main() {
	fmt.Println("------------------test uf1 ------------------")
	uf1 := _3_Quick_Union.NewUnionFind1(6)
	uf1.UnionElements(0,1)
	uf1.UnionElements(2,3)
	uf1.UnionElements(4,5)
	if 6 != uf1.GetSize(){
		fmt.Println(uf1.GetSize())
		panic("size should be 6")

	}

	//测试相连的
	if !uf1.IsConnected(0,1){
		panic(" 0 should be connect 1")
	}
	if !uf1.IsConnected(2,3){
		panic(" 2 should be connect 3")
	}
	if !uf1.IsConnected(4,5){
		panic(" 4 should be connect 5")
	}
	//测试不相连的
	if uf1.IsConnected(0,2){
		panic(" 0 should not be connect 2")
	}
	if uf1.IsConnected(4,3){
		panic(" 4 should not be connect 3")
	}
	if uf1.IsConnected(1,5){
		panic(" 1 should not be connect 5")
	}

	fmt.Println("uf1 is OK")

	fmt.Println("------------------test uf2 ------------------")
	uf2 := _3_Quick_Union.NewUnionFind1(6)
	uf2.UnionElements(0,1)
	uf2.UnionElements(2,3)
	uf2.UnionElements(4,5)
	if 6 != uf2.GetSize(){
		fmt.Println(uf2.GetSize())
		panic("size should be 6")

	}

	//测试相连的
	if !uf2.IsConnected(0,1){
		panic(" 0 should be connect 1")
	}
	if !uf2.IsConnected(2,3){
		panic(" 2 should be connect 3")
	}
	if !uf2.IsConnected(4,5){
		panic(" 4 should be connect 5")
	}
	//测试不相连的
	if uf2.IsConnected(0,2){
		panic(" 0 should not be connect 2")
	}
	if uf2.IsConnected(4,3){
		panic(" 4 should not be connect 3")
	}
	if uf2.IsConnected(1,5){
		panic(" 1 should not be connect 5")
	}

	fmt.Println("uf2 is OK")
}
