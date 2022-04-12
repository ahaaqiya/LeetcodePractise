//package hotMethod
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func CreatLinklistHead(nums []int) *ListNode {
	var head *ListNode
	head = nil
	for i:=0;i<len(nums);i++{
		tmp := &ListNode{nums[i],nil}
		tmp.Next = head
		head = tmp
	}
	return head
}

func CreatLinklistTail(nums []int) *ListNode {
	cur := &ListNode{}
	head := cur
	for i:=0;i<len(nums);i++{
		tmp := &ListNode{nums[i],nil}
		cur.Next = tmp
		cur = cur.Next
	}
	return head.Next
}

func PrintLinklist(head *ListNode)  {
	for head!=nil{
		fmt.Print(head.Val)
		head = head.Next
		if head!=nil{
			fmt.Print("-")
		}
	}
}


