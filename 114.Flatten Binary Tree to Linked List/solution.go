package main

import (
	"fmt"
)

/*
要求：
in-place

解题思路：
若无需in-place,则可将root转换成list再flatten.
需in-place,可用递归.先找到最左子节点,然后将其插入其父节点与父节点的右子节点之间,随后回到上一父节点继续,如下所示：
     1				  1
    / \				 / \
   2   5     -->    2   5		--> ...
  / \   \            \   \
 3   4   6            3   6
                       \
                        4
非递归情况下,先检测其左子节点是否存在,若存在则将根节点和其右子节点断开,将其左子节点机器后续一起连到原右子节点的位置,把原右子节点连到原左子节点最后面的右子节点之后,如下所示：
     1				  1
    / \				   \
   2   5     -->        2		--> ...
  / \   \              / \
 3   4   6            3   4
                           \
                            5
                             \
                              6

关键点：


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten1(root *TreeNode) { // faster 81.60% less 33.33%
	var list []int
	r(root, &list)
	for k := range list {
		root.Val = list[k]
		root.Left = nil
		if k == len(list)-1 {
			break
		}
		root.Right = new(TreeNode)
		root = root.Right
	}
}

func r(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	r(root.Left, list)
	r(root.Right, list)
}

func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp
}

func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 5
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 6
	flatten(root)
	fmt.Println(root)
}
