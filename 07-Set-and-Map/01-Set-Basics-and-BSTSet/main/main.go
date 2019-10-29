package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
)

func main() {
	filename1, _ := filepath.Abs("07-Set-and-Map/01-Set-Basics-and-BSTSet/a-tale-of-two-cities.txt")
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("words1 size = ",len(words1))


	filename2, _ := filepath.Abs("07-Set-and-Map/01-Set-Basics-and-BSTSet/pride-and-prejudice.txt")
	words2,err2 := FileOperation.ReadFile(filename2)
	if err2 != nil{
		return
	}
	fmt.Println("words2 size = ",len(words2))
}
