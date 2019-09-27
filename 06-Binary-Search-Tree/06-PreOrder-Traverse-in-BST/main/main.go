package main

import (
	"fmt"
	_6_PreOrder_Traverse_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/06-PreOrder-Traverse-in-BST"
	"hash/fnv"
	"strings"
)
/// Leetcode 804. Unique Morse Code Words
/// https://leetcode.com/problems/unique-morse-code-words/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的BST类
func main() {
	words := []string{"gin","zen","gig","msg"}
	if 2 != uniqueMorseRepresentations(words){
		panic("error,should be 2")
	}
	fmt.Println("success")

	bst := _6_PreOrder_Traverse_in_BST.NewBST()
	bst.Add(3)
	bst.Add(2)
	bst.Add(4)
	if !bst.Contains(4){
		panic("error,should Contains 4")
	}
	fmt.Println("found 4")
	if bst.Contains(-1){
		panic("error,should not Contains -1")
	}
}
//
func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	bst := _6_PreOrder_Traverse_in_BST.NewBST()
	for _,word := range words  {
		var res strings.Builder
		for i:=0;i< len(word);i++  {
			res.WriteString(codes[word[i] - 'a'])
		}
		//转为hashcode
		bst.Add( int(hash(res.String())))
	}
	return bst.Size()
}
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
