package main

//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func main() {
//
//	firstListNode1 := ListNode{
//		Val:  2,
//		Next: nil,
//	}
//
//	firstListNode2 := ListNode{
//		Val:  4,
//		Next: nil,
//	}
//
//	firstListNode3 := ListNode{
//		Val:  3,
//		Next: nil,
//	}
//
//	firstListNode1.Next = &firstListNode2
//	firstListNode2.Next = &firstListNode3
//
//	secondListNode1 := ListNode{
//		Val:  5,
//		Next: nil,
//	}
//
//	secondListNode2 := ListNode{
//		Val:  6,
//		Next: nil,
//	}
//
//	secondListNode3 := ListNode{
//		Val:  4,
//		Next: nil,
//	}
//
//	secondListNode1.Next = &secondListNode2
//	secondListNode2.Next = &secondListNode3
//
//	res := addTwoNumbers(&firstListNode1, &secondListNode1)
//
//	fmt.Println(res)
//
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	// Create a dummy head node to simplify the list manipulation
//	dummyHead := &ListNode{}
//	currentNode := dummyHead
//	carry := 0
//
//	// Traverse both input lists until both are exhausted
//	for l1 != nil || l2 != nil {
//		// Get the values of the current nodes (or 0 if the list is exhausted)
//		val1, val2 := 0, 0
//		if l1 != nil {
//			val1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			val2 = l2.Val
//			l2 = l2.Next
//		}
//
//		// Calculate the sum and carry
//		sum := val1 + val2 + carry
//
//		carry = sum / 10
//		sum %= 10
//		// Create a new node with the calculated sum
//		newNode := &ListNode{Val: sum}
//		// Connect the new node to the result list
//		oldC := currentNode
//		currentNode.Next = newNode
//		//fmt.Println(newNode, currentNode, currentNode.Next)
//		currentNode = currentNode.Next
//		newC := currentNode
//		//fmt.Println(newNode, currentNode, currentNode.Next)
//		fmt.Println(oldC, newC)
//		fmt.Println(oldC.Next, newC.Next)
//	}
//
//	// If there's still a carry after traversing both lists, add it as a new node
//	if carry > 0 {
//		currentNode.Next = &ListNode{Val: carry}
//	}
//	//fmt.Println(currentNode)
//	// Return the resulting linked list
//	return dummyHead.Next.Next.Next
//}
