package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l11 := ListNode{Val: 2, Next: nil}
	l12 := ListNode{Val: 4, Next: &l11}
	l16 := ListNode{Val: 3, Next: &l12}

	l21 := ListNode{Val: 5, Next: nil}
	l22 := ListNode{Val: 6, Next: &l21}
	l33 := ListNode{Val: 4, Next: &l22}
	l := addTwoNumbers(&l16, &l33)
	fmt.Println(l)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	xVal, yVal := 0, 0
	tmp := 0
	zList := &ListNode{}
	zList = nil
	for {

		xVal = l1.Val
		yVal = l2.Val
		fmt.Println(xVal, yVal, tmp)
		plus := xVal + yVal + tmp
		if plus >= 10 {
			tmp = 1
			plus -= 10
		} else {
			tmp = 0
		}

		fmt.Println("plus", plus)
		zList = &ListNode{Val: plus, Next: zList}

		if tmp == 0 && nil == l1.Next && nil == l2.Next {
			break
		}

		if nil == l1.Next {
			l1 = &ListNode{0, nil}
		} else {
			l1 = l1.Next
		}
		if nil == l2.Next {
			l2 = &ListNode{0, nil}
		} else {
			l2 = l2.Next
		}
		fmt.Println("----", l1.Next, l2.Next)

	}

	var rtn *ListNode
	for tmp := zList; tmp != nil; tmp = tmp.Next {
		rtn = &ListNode{Val: tmp.Val, Next: rtn}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	return rtn

}
