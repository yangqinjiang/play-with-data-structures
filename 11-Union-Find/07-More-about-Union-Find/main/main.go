package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/11-Union-Find/07-More-about-Union-Find"
	"math/rand"
	"time"
)

func main() {

	// 比较 uf3 和uf4, uf5
	size := 10000000
	m := 10000000
	fmt.Printf("------------------test uf3,uf4,tf5,tf6 size:=%d , m:=%d ------------------\n",size,m)
	uf3 :=  NewUnionFind3(size)
	fmt.Printf("UF3 : %f s ,optimized by Size.\n",testUF(uf3,m))
	uf4 :=  NewUnionFind4(size)
	fmt.Printf("UF4 : %f s ,optimized by Rank.\n",testUF(uf4,m))
	uf5 :=  NewUnionFind5(size)
	fmt.Printf("UF5 : %f s ,path compression .\n",testUF(uf5,m))
	uf6 :=  NewUnionFind6(size)
	fmt.Printf("UF6 : %f s ,path compression 2,递归算法 .\n",testUF(uf6,m))
}
//------------------test uf3,uf4,tf5,tf6 size:=10000000 , m:=10000000 ------------------
//UF3 : 6.155090 s ,optimized by Size.
//UF4 : 6.840578 s ,optimized by Rank.
//UF5 : 4.943098 s ,path compression .
//UF6 : 5.563016 s ,path compression 2,递归算法 .




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
