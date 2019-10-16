package main

import (
	"container/list"
	"fmt"
)

/*
要求：
非递归

解题思路：
递归.结果集加入左子树递归值,加入根节点,加入又子树递归值,当根节点==nil或左右子树均==nil时结束递归
遍历.通过栈实现,当当前节点不为空时,循环入栈左子树,当前节点指向左子树,直到当前节点为空;若此时栈为空则退出(此步判断在左子树处理之后),随后从栈中取出并加入结果集,再将当前节点指向右子树并继续循环.

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	List *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.List.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	if e := stack.List.Back(); e != nil {
		stack.List.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.List.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}

// 递归
func inorderTraversal1(root *TreeNode) []int { // faster 100% less 33.33%
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 遍历
func inorderTraversal(root *TreeNode) []int { // faster 100% less 33.33%
	for root == nil {
		return nil
	}
	var res []int
	s := NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		if s.Empty() {
			break
		}
		cur = s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func main() {
	t1 := new(TreeNode)
	t1.Val = 1
	t2 := new(TreeNode)
	t2.Val = 2
	t1.Right = t2
	t3 := new(TreeNode)
	t3.Val = 3
	t2.Left = t3
	result := inorderTraversal(t1)
	fmt.Println(result)
}
