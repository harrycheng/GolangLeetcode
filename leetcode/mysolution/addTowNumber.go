package mysolution

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumbers() {
	a := []int{3,4,2,9}
	b := []int{4,6,5,9}
	l1 := constructIntList(a)
	l2 := constructIntList(b)

	printList(l1)
	printList(l2)

	re := addTwoNumbers(l1, l2)
	printList(re)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var re *ListNode = nil
	var trail *ListNode = nil

	flag := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		temp := v1 + v2 + flag
		if temp > 9 {
			temp = temp % 10
			flag = 1
		} else {
			flag = 0
		}

		if(re == nil){
			node := ListNode{Val: temp, Next: trail}
			re = &node
			trail = &node
		} else {
			node := ListNode{Val: temp, Next: nil}
			trail.Next = &node
			trail = &node
		}

	}
	if flag == 1 {
		node:=ListNode{1, nil}
		trail.Next=&node
		trail = &node
	}
	return re
}

func constructIntList(intarray []int) *ListNode {
	var re *ListNode = nil
	for _, v := range intarray {
		node := ListNode{Val: v, Next: re}
		re = &node
	}
	return re
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println(" ")
}
