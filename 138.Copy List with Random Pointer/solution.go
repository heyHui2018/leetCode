package main

import (
	"fmt"
)

/*
要求：


解题思路：


关键点：


*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node { // faster 98.63% less 100%

}

func main() {
	head := new(Node)
	result := copyRandomList(head)
	fmt.Println(result)
}
