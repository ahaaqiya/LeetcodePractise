package main

import (
	"container/heap"
	"strconv"
)

func main()  {
	maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3)
}

func removeElement(nums []int, val int) int {
	slow,fast := 0,0
	for fast < len(nums){
		if nums[fast]!=val{
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _,item := range  nums{
		numMap[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key,value := range numMap{
		heap.Push(h,[2]int{key,value})
		if h.Len()>k{
			heap.Pop(h)
		}
	}
	res := make([]int,k)
	for i:=0;i<k;i++{
		res[k-i-1]  = heap.Pop(h).([2]int)[0]
	}
	return res
}

type IHeap [][2]int

func (h IHeap) Len()int{
	return len(h)
}

func (h IHeap) Less(i,j int) bool{
	return h[i][1]<h[j][1]
}

func (h IHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *IHeap) Push(x interface{})  {
	*h = append(*h,x.([2]int))
}

func (h *IHeap) Pop() interface{}  {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type DouQueue struct {
	queue []int
}

func newDouQueue() *DouQueue {
	return &DouQueue{
		queue: make([]int,0),
	}
}

func (q *DouQueue) Front() int {
	return q.queue[0]
}
func (q *DouQueue) push(val int) {
	for len(q.queue)!=0 && q.queue[len(q.queue)-1]<val{
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue,val)
}

func (q *DouQueue) pop(val int)  {
	if len(q.queue)!=0 && q.Front() == val{
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := newDouQueue()
	res := make([]int,0)
	for i:=0;i<k;i++{
		queue.push(nums[i])
	}
	res = append(res,queue.Front())

	for i:=k;i<len(nums);i++{
		queue.pop(nums[i-k])
		queue.push(nums[i])
		res = append(res,queue.Front())
	}
	return res
}

func evalRPN(tokens []string) int {
	stack := make([]string,0)
	for i:=0;i<len(tokens);i++{
		if tokens[i]!="+" && tokens[i]!="-" && tokens[i]!="*" && tokens[i]!="/"{
			stack = append(stack,tokens[i])
		}else{
			second,_ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			first,_ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			tmp := 0
			if tokens[i]=="+"{
				tmp = first+second
			}else if tokens[i]=="-"{
				tmp = first-second
			}else if tokens[i]=="*"{
				tmp = first * second
			}else{
				tmp = first / second
			}
			tmpStr := strconv.Itoa(tmp)
			stack = append(stack,tmpStr)
		}
	}
	res,_ := strconv.Atoi(stack[0])
	return res
}

func repeatedSubstringPattern(s string) bool {
	newS := s + s
	newS = newS[1:len(newS)-1]
	return kmp(newS,s)
}

func kmp(str string,aim string) bool {
	arr := []byte(str)
	aimArr := []byte(aim)
	next := kmpHelp(aimArr)
	i,j := 0,0
	for i<len(arr) && j<len(aimArr){
		if arr[i]==aimArr[j]{
			i++
			j++
		}else{
			if j==0{
				i = i - next[0]
			}else {
				j = next[j]
			}
		}
	}
	if i<=len(arr) && j==len(aimArr){
		return true
	}
	return false
}

func kmpHelp(arr []byte) []int{
	if len(arr)<=2{
		return []int{-1,0}
	}
	next := make([]int,len(arr))
	next[0],next[1] = -1,0
	index := 0
	for i:=1;i<len(arr)-1;{
		if index==-1{
			next[i+1]=0
			index = 0
			i++
		}else if arr[index]==arr[i]{
			next[i+1] = index+1
			index++
			i++
		}else {
			index = next[index]
		}
	}
	return next
}

func removeDuplicates(s string) string {
	if len(s)<2{
		return s
	}
	arr := []byte(s)
	sStack := make([]byte,0)
	sStack = append(sStack,arr[0])
	for i:=1;i<len(s);i++{
		if len(sStack)>0 && arr[i] == sStack[len(sStack)-1]{
			sStack = sStack[:len(sStack)-1]
		}else {
			sStack = append(sStack,arr[i])
		}
	}
	return string(sStack)
}

func isValid(s string) bool {
	arr := []byte(s)
	sStack := make([]byte,0)
	strMap := map[byte]byte{'{':'}','[':']','(':')'}
	for i:=0;i<len(arr);i++{
		if arr[i]=='(' || arr[i]=='{' || arr[i]=='['{
			sStack = append(sStack,strMap[arr[i]])
		}else if arr[i]==')' || arr[i]=='}' || arr[i]==']'{
			if len(sStack)==0 || arr[i] != sStack[len(sStack)-1]{
				return false
			}
			sStack = sStack[:len(sStack)-1]
		}
	}
	if len(sStack)>0{
		return false
	}
	return true
}

func maxProduct(nums []int) int {
	dp := make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		dp[i] = make([]int,2)
	}
	dp[0][0],dp[0][1] = nums[0],nums[0]
	max := dp[0][0]
	for i:=0;i<len(nums);i++{
		dp[i][0] = findMax(findMax(dp[i][0]*nums[i],dp[i][1]*nums[i]),nums[i])
		dp[i][1] = findMin(findMin(dp[i][0]*nums[i],dp[i][1]*nums[i]),nums[i])
		if dp[i][0]>max{
			max = dp[i][0]
		}
	}
	return max
}

func findMax(a,b int) int {
	if a>b{
		return a
	}else {
		return b
	}
}
func findMin(a,b int) int {
	if a>b{
		return b
	}else {
		return a
	}
}
type stack struct {
	nums []int
}

func (this *stack) pushToTop(num int)  {
	this.nums = append(this.nums,num)
}

func (this *stack) peekFromTop() int {
	return this.nums[len(this.nums)-1]
}
func (this *stack) popFromTop(){
	this.nums = this.nums[:len(this.nums)-1]
}

type MyQueue struct {
	InStack *stack
	OutStack *stack
}

func Constructor() MyQueue {
	return MyQueue{
		InStack: &stack{nums:[]int{}},
		OutStack: &stack{nums:[]int{}},
	}
}

func (this *MyQueue) Push(x int)  {
	for len(this.OutStack.nums)!=0{
		val := this.OutStack.peekFromTop()
		this.InStack.pushToTop(val)
		this.OutStack.popFromTop()
	}
	this.InStack.pushToTop(x)
}

func (this *MyQueue) Pop() int {
	for len(this.InStack.nums)!=0{
		val := this.InStack.peekFromTop()
		this.OutStack.pushToTop(val)
		this.InStack.popFromTop()
	}
	if len(this.OutStack.nums)==0{
		return 0
	}
	val := this.OutStack.peekFromTop()
	this.OutStack.popFromTop()
	return val
}


func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val==0{
		return 0
	}
	this.OutStack.pushToTop(val)
	return val
}


func (this *MyQueue) Empty() bool {
	if len(this.InStack.nums)==0 && len(this.OutStack.nums)==0{
		return true
	}
	return false
}
