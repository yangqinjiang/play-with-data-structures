package main

import  "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/04-More-about-Leetcode"

func main() {

	if !_4_MORE_ABOUT_LEETCODE.IsValid("()[]{}"){
	    panic("should success")
	}
	if !_4_MORE_ABOUT_LEETCODE.IsValid("([{}])"){
		panic("should success")
	}

	//-------------
	if _4_MORE_ABOUT_LEETCODE.IsValid(""){
		panic("should fail")
	}
	//if _4_MORE_ABOUT_LEETCODE.IsValid("[-"){
	//	panic("should fail")
	//}
	//if _4_MORE_ABOUT_LEETCODE.IsValid("^^&*())(*&^&^&&**[][][]"){
	//	panic("should fail")
	//}
	if _4_MORE_ABOUT_LEETCODE.IsValid("([)]"){
		panic("should fail")
	}
}
