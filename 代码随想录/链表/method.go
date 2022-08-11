package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val int
	Next *ListNode
}


func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap := make(map[int]int)
	res := 0
	for _,a := range nums1{
		for _,b := range nums2{
			sumMap[a+b]++
		}
	}
	for _,c := range nums3{
		for _,d := range nums4{
			if count,have := sumMap[0-(c+d)];have{
				res = res + count
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int,0)
	for k:=0;k<len(nums)-1;k++{
		if nums[k]>0{
			break
		}
		if k>0 && nums[k]==nums[k-1]{
			continue
		}
		left,right := k+1,len(nums)-1
		for left<right{
			n2,n3 := nums[left],nums[right]
			if nums[k]+nums[left]+nums[right]==0{
				res = append(res,[]int{nums[k],nums[left],nums[right]})
				for left<right && nums[left]==n2{
					left++
				}
				for left<right && nums[right]==n3{
					right--
				}
			}else if nums[k]+nums[left]+nums[right]<0{
				left++
			}else {
				right--
			}
		}
	}
	return res
}

func detectCycle(head *ListNode) *ListNode {
	fast,slow := head,head
	for fast!=nil && fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast==slow{
			break
		}
	}
	if fast==nil || fast.Next==nil{
		return nil
	}
	fast = head
	for fast != slow{
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0,head}
	fast,slow := dummy,dummy
	for i:=0;i<n;i++{
		fast = fast.Next
	}
	for fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return nil
	}
	var pre *ListNode
	newHead := reverseListHelp(pre,head)
	return newHead
}

func reverseListHelp(pre,head *ListNode) *ListNode {
	if head==nil{
		return pre
	}
	tmp := head.Next
	head.Next = pre
	return reverseListHelp(head,tmp)
}

/*func reverseList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return nil
	}
	cur := head
	var pre *ListNode
	for cur != nil{
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A,B := headA,headB
	lenA,lenB := 0,0
	for A!=nil{
		headA = headA.Next
		lenA++
	}
	for B!=nil{
		headB = headB.Next
		lenB++
	}
	var step int
	var fast,slow *ListNode
	if lenA>lenB{
		step = lenA - lenB
		fast,slow = headA,headB
	}else{
		step = lenB - lenA
		fast,slow = headB,headA
	}
	for i:=0;i<step;i++{
		fast = fast.Next
	}
	for fast != slow{
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1,head}
	newHead := dummy
	p,del := newHead,newHead
	for i:=0;i<n;i++{
		p = p.Next
	}
	for p.Next!=nil{
		p = p.Next
		del = del.Next
	}
	del.Next = del.Next.Next
	return newHead.Next
}
*/
func swapPairs(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}
	newHead := &ListNode{-1,head}
	a := newHead.Next
	b := new(ListNode)
	for a.Next!=nil{
		b = a.Next
		a.Next = a.Next.Next
		a = a.Next
		b.Next = a.Next
		a.Next = b
		a = a.Next
	}
	return newHead.Next
}

type MyLinkedList struct {
	Dummy *MyNode
}

type MyNode struct {
	Val int
	Next *MyNode
	Prev *MyNode
}

/*func reverseList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}*/

func main()  {
	nums := []int{1,2,3,4,5}
	list := tailList(nums)
	printList(list)
	newhead := reverseList(list)
	printList(newhead)
}

/*func reverseList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	newHead := &ListNode{head.Val,nil}
	head = head.Next
	for head!=nil{
		tmp := &ListNode{head.Val,nil}
		tmp.Next = newHead
		newHead = tmp
		head = head.Next
	}
	return newHead
}*/


func Constructor() MyLinkedList {
	//虚头节点
	newNode := &MyNode{-1,nil,nil}
	newNode.Prev = newNode
	newNode.Next = newNode
	return MyLinkedList{newNode}
}


func (this *MyLinkedList) Get(index int) int {
	dummy := this.Dummy
	pos := dummy.Next
	i := 0
	if dummy==pos{
		return -1
	}
	for pos.Next!=dummy{
		if i==index{
			break
		}
		pos = pos.Next
		i++
	}
	if i!=index{
		return -1
	}
	return pos.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	dummy := this.Dummy
	tmp := &MyNode{val,dummy.Next,dummy}
	dummy.Next.Prev = tmp
	dummy.Next = tmp
}


func (this *MyLinkedList) AddAtTail(val int)  {
	dummy := this.Dummy
	last := &MyNode{val,dummy,dummy.Prev}
	dummy.Prev.Next = last
	dummy.Prev = last
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index<=0{
		this.AddAtHead(val)
		return
	}
	dummy := this.Dummy
	pos := dummy.Next
	i := 0
	if dummy==pos{
		return
	}
	for pos.Next!=dummy{
		if i==index{
			break
		}
		pos = pos.Next
		i++
	}
	if index==i+1{
		this.AddAtTail(val)
		return
	}else if i!=index{
		return
	}
	newNode := &MyNode{val,pos, pos.Prev}
	pos.Prev.Next = newNode
	pos.Prev = newNode
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	dummy := this.Dummy
	pos := dummy.Next
	i := 0
	if dummy==pos{
		return
	}
	for pos.Next!=dummy{
		if i==index{
			break
		}
		pos = pos.Next
		i++
	}
	if i!=index{
		return
	}
	pos.Next.Prev = pos.Prev
	pos.Prev.Next = pos.Next
}



func removeElements(head *ListNode, val int) *ListNode {
	if head==nil{
		return nil
	}
	tmp := new(ListNode)
	tmp = head
	for tmp!=nil{
		if tmp.Next.Val==val{
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return head
}

func tailList(nums []int) *ListNode {
	if len(nums)==0{
		return nil
	}
	head := &ListNode{nums[0],nil}
	tmp := head
	for i:=1;i<len(nums);i++{
		p := &ListNode{nums[i],nil}
		tmp.Next = p
		tmp = p
	}
	return head
}

func printList(head *ListNode)  {
	if head==nil{
		return
	}
	for head.Next!=nil{
		fmt.Printf("%d-",head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}