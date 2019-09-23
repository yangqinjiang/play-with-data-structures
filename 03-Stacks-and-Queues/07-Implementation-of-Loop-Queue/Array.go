package _7_Implementation_of_Loop_Queue

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

	if index <0 || index > this.size{
		panic("AddLast failed. Require index >= 0 and index <= size.")
	}
	if this.size == len(this.data){
		//调整容量
		this.resize(2* len(this.data))
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

func (this *Array)GetLast() int {
	return this.Get(this.size - 1)
}
func (this *Array)GetFirst() int {
	return this.Get(0)
}
//修改index索引位置的元素为e
func (this *Array)Set(index ,e int) {
	if index <0 || index >= this.size{
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}

//查找数组中是否有元素e
func (this *Array)Contains(e int) bool  {
	for i:=0;i<this.size;i++{
		if e == this.data[i]{
			return  true
		}
	}
	return  false
}

//查找数组中元素e所在的索引,如果不存在元素e,则返回 -1
func (this *Array)Find(e int) int  {
	for i:=0;i<this.size;i++{
		if e == this.data[i]{
			return  i
		}
	}
	return  -1
}

//从数组中删除index位置的元素,返回删除的元素
func (this *Array)Remove(index int) int  {
	if index <0 || index >= this.size{
		panic("Remove failed. Index is illegal.")
	}

	ret := this.data[index]
	//向左移动
	for i:=index +1;i<this.size;i++{
		this.data[i-1] = this.data[i]
	}
	this.size --

	//调整容量,size达到1/4时,则调整为原容量的1/2
	if this.size == len(this.data)/4 && len(this.data) /2 != 0{
		this.resize(len(this.data)/2)
	}
	return  ret
}
//从数组中删除第一个元素,返回删除的元素
func (this *Array)RemoveFirst() int {
	return this.Remove(0)
}

//从数组中删除最后一个元素,返回删除的元素
func (this *Array)RemoveLast() int {
	return this.Remove(this.size -1 )
}
//从数组中删除元素e
func (this *Array)RemoveElement(e int)  {
	index := this.Find(e)
	if -1 != index{
		this.Remove(index)
	}
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

//将数组空间的容量变成newCapacity大小
func (this *Array)resize(newCapacity int)  {
	fmt.Println("--------call resize func---------",this.size,"-->",newCapacity)
	newData := make([]int,newCapacity)
	for i:=0;i<this.size ;i++  {
		newData[i] = this.data[i]
	}
	this.data = newData
}