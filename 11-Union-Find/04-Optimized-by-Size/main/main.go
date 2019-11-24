package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/11-Union-Find/04-Optimized-by-Size"
	"math/rand"
	"time"
)

func main() {

	// 情景一:UnionFind1 慢于 UnionFind2
	//size := 100000
	//m := 10000
	// 情景二:UnionFind2 慢于 UnionFind1, 但UnionFind3最快
	size := 100000
	m := 100000
	fmt.Println("------------------test uf1,uf2,uf3 ------------------")
	uf1 :=  NewUnionFind1(size)
	 fmt.Printf("UF1 : %f s ,Array\n",testUF(uf1,m))
	uf2 :=  NewUnionFind2(size)
	fmt.Printf("UF2 : %f s ,Parent\n",testUF(uf2,m))
	uf3 :=  NewUnionFind3(size)
	fmt.Printf("UF3 : %f s ,optimized by Size.\n",testUF(uf3,m))
}

func testUF(uf UF,m int) float64{
	startTime := time.Now()
	size := uf.GetSize()
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<m ;i++  {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.UnionElements(a,b)
	}
	for i:=0;i<m ;i++  {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.IsConnected(a,b)
	}
	elapsed := time.Since(startTime)
	return elapsed.Seconds()
}
