package main

import (
	. "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/03-A-Stack-Problem-in-Leetcode"
)

func main() {
	if !IsValid("()[]{}"){
		panic("should be true")
	}
	if IsValid("([)]"){
		panic("should be false")
	}
}
