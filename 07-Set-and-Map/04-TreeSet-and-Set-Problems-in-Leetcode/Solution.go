package _4_TreeSet_and_Set_Problems_in_Leetcode

import (
	"strings"
)

func UniqueMorseRepresentations(words []string) int  {
	return uniqueMorseRepresentations(words)
}
// Leetcode 804. Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/description/
func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	//使用BSTSet或者LinkedListSet
	//set := NewBSTSet()
	set := NewLinkedListSet()
	for _,word := range words  {
		var res strings.Builder
		for i:=0;i< len(word);i++  {
			res.WriteString(codes[word[i] - 'a'])
		}
		//转为hashcode
		set.Add( res.String() )
	}
	return set.GetSize()
}

