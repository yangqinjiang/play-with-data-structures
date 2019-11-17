package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/10-Trie/02-Trie-Basics"
)

func main() {
	fmt.Println("Trie...")
	trie:=_2_Trie_Basics.NewTrie()
	trie.Add("AHELLO中方")
	trie.Add("aHI")
	trie.Add("World")
	trie.Add("World") //重复的单词
	if 3 != trie.GetSize(){
		panic("单词个数不一致")
	}
}
