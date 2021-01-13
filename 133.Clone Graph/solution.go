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
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node { // faster 100% less 21.51%
	nodeMap := make(map[int]*Node)
	return DFSClone(node, nodeMap)
}

func DFSClone(node *Node, nodeMap map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if getNode, ok := nodeMap[node.Val]; ok {
		return getNode
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	nodeMap[newNode.Val] = newNode
	for index := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, DFSClone(node.Neighbors[index], nodeMap))
	}
	return newNode
}

func main() {
	n := new(Node)
	result := cloneGraph(n)
	fmt.Println(result)
}
