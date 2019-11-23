package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/10-Trie/06-Trie-and-Map"
)
/**
输入: insert("apple", 3), 输出: Null
输入: sum("ap"), 输出: 3
输入: insert("app", 2), 输出: Null
输入: sum("ap"), 输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/map-sum-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {
	fmt.Println("Trie...")
	trie:= Constructor()
	trie.Insert("apple",3)
	if 3 != trie.Sum("ap"){
		panic("应该是3")
	}
	trie.Insert("app",2)
	if 5 != trie.Sum("ap"){
		panic("应该是5")
	}
	fmt.Println("OK")
}
