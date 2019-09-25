package _3_Recursion_Basics

type Sum struct {

}

func (this *Sum)Sum(arr []int) int  {
	return this.sum(arr,0)
}
func (this *Sum)sum(arr []int,l int) int  {
	if l == len(arr){
		return 0
	}

	return  arr[l] + this.sum(arr,l+1)
}

func NewSum() *Sum  {
	return &Sum{}
}