package main
import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/01-Hash-Table-Basics"
)
func main(){
	resIndex := _1_hash_table_basics.FirstUniqChar("leetcode")
	if 0 != resIndex{
		panic("should be 0")
	}

	resIndex = _1_hash_table_basics.FirstUniqChar("loveleetcode")
	if 2 != resIndex{
		panic("should be 2")
	}
	fmt.Println("OK")
}