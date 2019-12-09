package main
import (
	"fmt"
	 "github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/03-Hash-Function-in-Java/Student"
)

func main(){
	fmt.Println(Student.HashCode(1))
	fmt.Println(Student.HashCode(1.0))
	fmt.Println(Student.HashCode("hello world"))

	stu := Student.NewStudent(3,2,"Yang","qinjiang")
	fmt.Println(stu.HashCode())

	another_stu := Student.NewStudent(3,2,"Yang","qinjiang")
	fmt.Println(another_stu.HashCode())
	if stu.Equals(*another_stu){
		fmt.Println("stu equals another_stu")
	}
	//todo: set

	//map
	scores := make(map[*Student.Student]int)
	scores[stu] = 100
	fmt.Println(scores)
}