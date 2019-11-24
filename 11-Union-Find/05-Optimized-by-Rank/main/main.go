package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/11-Union-Find/05-Optimized-by-Rank"
	"math/rand"
	"time"
)

func main() {

	// 比较 uf3 和uf4
	size := 10000000
	m := 10000000
	fmt.Println("------------------test uf3,uf4 ------------------")
	uf3 :=  NewUnionFind3(size)
	fmt.Printf("UF3 : %f s ,optimized by Size.\n",testUF(uf3,m))
	uf4 :=  NewUnionFind4(size)
	fmt.Printf("UF4 : %f s ,optimized by Rank.\n",testUF(uf4,m))
}
//------------------test uf3,uf4 ------------------
//size := 10000000
//m := 10000000
//UF3 : 6.048099 s ,optimized by Size.
//UF4 : 6.658091 s ,optimized by Rank.


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
