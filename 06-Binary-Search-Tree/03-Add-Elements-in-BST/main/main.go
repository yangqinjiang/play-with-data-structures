package main

import (
	"fmt"
	_3_Add_Elements_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/03-Add-Elements-in-BST"
	"strings"
)
/// Leetcode 804. Unique Morse Code Words
/// https://leetcode.com/problems/unique-morse-code-words/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的BST类
func main() {
	codes := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	words := []string{"gin","zen","gig","msg"}
	bst := _3_Add_Elements_in_BST.NewBST()
	for _,word := range words  {
		var res strings.Builder
		for i:=0;i< len(word);i++  {
			res.WriteString(codes[word[i] - 'a'])
		}
		bst.Add(res.String())
	}
	fmt.Println(bst.Size())
	bst2 := _3_Add_Elements_in_BST.NewBST()
	bst2.Add(1)
	fmt.Println(bst2.Size())
	bst2.Add(3)
	fmt.Println(bst2.Size())
	bst2.Add(2)
	fmt.Println(bst2.Size())
	bst2.Add(5)
	fmt.Println(bst2.Size())
}
