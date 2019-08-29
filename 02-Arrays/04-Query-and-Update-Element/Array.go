package _4_Query_and_Update_Element

import (
	"fmt"
	"strings"
)

type Array struct {
	data []int
	size int
}

//构造器,传入数组的容量capacity构造Array
func NewArray(capacity int) *Array  {
	return &Array{
		data: make([]int,capacity),
		size: 0,
	}
}
//无参数的构造函数,默认数组的容量capacity=10
func NewArrayDefault() *Array  {
	return NewArray(10)
}

//获取数组的容量
func (this *Array)GetCapacity() int  {
	return len(this.data)
}
//获取数组中的元素个数
func (this *Array)GetSize() int {
	return  this.size
}
func (this *Array)IsEmpty() bool  {
	return  this.size == 0;
}

//向所有元素后添加一个新元素
func (this *Array)AddLast(e int)  {
	//if this.size == len(this.data){
	//	panic("AddLast failed. Array is full.")
	//}
	//this.data[this.size] = e;
	//this.size ++

	this.Add(this.size,e)
}

//在所有元素前添加一个新元素
func (this *Array)AddFirst(e int)  {
	this.Add(0,e)
}
func (this *Array)Add(index,e int)  {
	if this.size == len(this.data){
		panic("AddLast failed. Array is full.")
	}
	if index <0 || index > this.size{
		panic("AddLast failed. Require index >= 0 and index <= size.")
	}

	//挪动位置
	for i := this.size - 1;i>= index ; i--  {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = e

	this.size ++
}

//获取index索引位置的元素
func (this *Array)Get(index int) int  {
	if index <0 || index >= this.size{
		panic("Get failed. Index is illegal.")
	}

	return  this.data[index]
}

//修改index索引位置的元素为e
func (this *Array)Set(index ,e int) {
	if index <0 || index >= this.size{
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}
//打印数组字符串
func (this *Array)ToString() string  {
	var sb strings.Builder
	fmt.Fprintf(&sb,"Array: size = %d,capacity= %d \n",this.size, len(this.data))
	sb.WriteString("[")
	for i := 0;i<this.size ;i++  {
		fmt.Fprintf(&sb,"%d",this.data[i])
		if i != this.size - 1{
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return  sb.String()

}