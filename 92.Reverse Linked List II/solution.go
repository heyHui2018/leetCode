package main

import (
	"fmt"
)

/*
要求：
1 ≤ m ≤ n ≤ length of list

解题思路：
定位mPre和nNext,从mNode开始翻转,随后连接上mPre和nNext.在定位nNext后,需在nNode处使head.Next=nil以便翻转

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode { // faster 100% less 12.5%
	if m == n {
		return head
	}
	// 添加headPre可使m>=2,可避免单独考虑m=1的情况
	headPre := &ListNode{}
	headPre.Next = head
	m++
	n++
	// 根据m/n切断链表
	mPre, mNode, nNext := split(headPre, m, n)
	fmt.Println(mPre, mNode, nNext)
	// head已被截断在nNext,故h即为n,e为m
	h, e := reverse(mNode)
	fmt.Println(h, e)
	// 连接链表
	mPre.Next = h
	e.Next = nNext
	return headPre.Next
}

func split(head *ListNode, m, n int) (mPre, mNode, nNext *ListNode) {
	i := 1
	for head != nil {
		if i == m-1 {
			mPre = head
			mNode = head.Next
		}
		if i == n {
			nNext = head.Next
			// 截断head,方便reverse处理
			head.Next = nil
			break
		}
		head = head.Next
		i++
	}
	return
}

// 返回新链条的 head 和 end
func reverse(head *ListNode) (h, e *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	var end *ListNode
	h, end = reverse(head.Next)
	end.Next = head
	e = head
	return
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 3
	l2.Next = l3
	l4 := new(ListNode)
	l4.Val = 4
	l3.Next = l4
	l5 := new(ListNode)
	l5.Val = 5
	l4.Next = l5
	m := 2
	n := 4
	result := reverseBetween(l1, m, n)
	for result != nil {
		fmt.Println(result)
		result = result.Next
	}
}
