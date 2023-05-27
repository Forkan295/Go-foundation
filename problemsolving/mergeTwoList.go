package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	firstListNode1 := ListNode{
		Val:  2,
		Next: nil,
	}

	firstListNode2 := ListNode{
		Val:  4,
		Next: nil,
	}

	firstListNode3 := ListNode{
		Val:  3,
		Next: nil,
	}

	firstListNode1.Next = &firstListNode2
	firstListNode2.Next = &firstListNode3

	secondListNode1 := ListNode{
		Val:  5,
		Next: nil,
	}

	secondListNode2 := ListNode{
		Val:  6,
		Next: nil,
	}

	secondListNode3 := ListNode{
		Val:  0,
		Next: nil,
	}

	secondListNode1.Next = &secondListNode2
	secondListNode2.Next = &secondListNode3

	//res := mergeTwoList(&firstListNode1, &secondListNode1)
	res := mergeTwoList(&ListNode{}, &ListNode{
		Val:  0,
		Next: nil,
	})
	//for res.Next != nil {
	//	fmt.Println(res.Val)
	//	res = res.Next
	//}

	fmt.Println(res)
	//mergeTwoList(&firstListNode1, &secondListNode1)

}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	//arr := []int{}
	//for list1 != nil && list2 != nil {
	//	val1 := 0
	//	val2 := 0
	//
	//	if list1.Val != 0 {
	//		val1 = list1.Val
	//	}
	//
	//	if list2.Val != 0 {
	//		val2 = list2.Val
	//	}
	//
	//	arr = append(arr, val1)
	//	arr = append(arr, val2)
	//
	//	if list1.Next == nil && list2.Next == nil {
	//		break
	//	}
	//
	//	if list1.Next != nil {
	//		list1 = list1.Next
	//	}
	//
	//	if list2.Next != nil {
	//		list2 = list2.Next
	//	}
	//}
	//
	//head := &ListNode{Val: 0}
	//current := head
	//
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//	newNode := &ListNode{Val: arr[i]}
	//	current.Next = newNode
	//	current = newNode
	//}
	//return head.Next

	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		// Compare the values of the nodes in both lists
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append the remaining nodes from either list, if any
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}
