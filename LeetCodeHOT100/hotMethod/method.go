//package hotMethod
package main

import (
	"fmt"
	"math"
	"sort"
)

//两数之和
func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++ {
			if nums[i]+nums[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil && l2==nil{
		return nil
	}
	cur := &ListNode{}
	head := cur
	flag := 0
	for l1!=nil || l2!=nil || flag != 0{
		var sum int
		if l1!=nil && l2!=nil{
			sum = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}else if l1!=nil {
			sum = l1.Val
			l1 = l1.Next
		}else if l2 != nil{
			sum = l2.Val
			l2 = l2.Next
		}
		x := (sum%10 + flag)%10
		flag = (sum + flag)/10
		tmp := &ListNode{x,nil}
		cur.Next = tmp
		cur = tmp
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil && l2==nil{
		return nil
	}
	cur := &ListNode{}
	head := cur
	flag := 0
	for l1!=nil || l2!=nil{
		var sum int
		var a int
		var b int
		if l1==nil{
			a = 0
		}else {
			a = l1.Val
			l1 = l1.Next
		}
		if l2==nil{
			b = 0
		}else {
			b = l2.Val
			l2 = l2.Next
		}
		sum = a + b + flag
		x := sum%10
		flag = sum/10
		cur.Next = &ListNode{Val: x}
		cur = cur.Next
	}
	if flag != 0{
		cur.Next = &ListNode{Val: flag}
	}
	return head.Next
}

//KMP
//求每位上的最大一致前后缀
func KMPHelp(arr []byte)  []int{
	if len(arr)<2{
		return []int{-1,0}
	}
	next := make([]int,len(arr))
	index,i := 0,1
	next[0] = -1
	next[1] = 0
	for i<len(arr)-1{
		if index == -1{
			next[i+1] = 0
			i++
			index = 0
		}else if arr[index] == arr[i]{
			next[i+1] = index+1
			index++
			i++
		}else {
			index = next[index]
		}
	}
	fmt.Println(next)
	return next
}

func KMP(s string,aim string) int {
	sArr := []byte(s)
	aimArr := []byte(aim)
	next := KMPHelp(aimArr)
	i,j := 0,0
	for i<len(sArr) && j<len(aimArr){
		if sArr[i] == aimArr[j]{
			i++
			j++
		}else {
			if j==0{
				i = i-next[j]
			}else {
				j = next[j]
			}
		}
	}
	if j==len(aimArr) && i<=len(sArr){
		return i-j
	}
	return -1
}
//无重复字符的最长字串
func lengthOfLongestSubstring(s string) int {
	arr := []byte(s)
	start,end := 0,0
	max := 0
	keyMap := make(map[byte]int)
	for i:=0;i<len(arr);i++{
		if _,have := keyMap[arr[i]]; !have{
			keyMap[arr[i]] = i
		}else {
			if keyMap[arr[i]]>=start{
				start = keyMap[arr[i]]+1
				keyMap[arr[i]] = i
			}else {
				keyMap[arr[i]] = i
			}
		}
		end = i
		if end - start + 1 > max{
			max = end - start + 1
		}

	}
	return max
}

//寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count,end,need := 0,0,0
	length := len(nums1) + len(nums2)
	if length%2==0{
		need = 2
	}else {
		need = 1
	}
	end = length/2
	res := make([]int,length)
	i,j := 0,0
	var flag bool
	for count<=end{
		flag = findMIN(i,j,nums1,nums2)
		if flag{
			res[count] = nums1[i]
			i++
		}else {
			res[count] = nums2[j]
			j++
		}
		count++
	}
	if need==1{
		return float64(res[count-1])
	}else {
		return float64(res[count-1]+res[count-2])/2
	}
}

func findMIN(i,j int, nums1,nums2 []int) bool {
	if i>=len(nums1){
		return false
	}
	if j>=len(nums2){
		return true
	}
	if nums1[i]<=nums2[j]{
		return true
	}else {
		return false
	}
}

//最长回文字串
func longestPalindrome(s string) string {
	index := 1
	arr := []byte(s)
	res,flag := make([]int,len(s)),make([]bool,len(s))
	res[0],flag[0] = 1,true
	max,maxIndex := 0,0
	lengthEven,lengthOdd := 1,1
	left,right := 0,0
	for index<len(s){
		//考虑偶数
		left,right = index-1,index
		for left>=0 && right<len(s) && arr[left] == arr[right]{
			left--
			right++
		}
		lengthEven = (right-left)-1
		//考虑奇数
		left,right = index-1,index+1
		for left>=0 && right<len(s) && arr[left] == arr[right]{
			left--
			right++
		}
		lengthOdd = (right-left)-1
		//判断采用哪一个
		if lengthOdd>=lengthEven{
			res[index] = lengthOdd
			flag[index] = true
		}else {
			res[index] = lengthEven
			flag[index] = false
		}
		//记录最长的
		if max<res[index]{
			max = res[index]
			maxIndex = index
		}
		index++
	}
	if flag[maxIndex]{
		return s[maxIndex-max/2:maxIndex+max/2+1]
	}else {
		return s[maxIndex-max/2:maxIndex+max/2]
	}
}

//正则表达式匹配
func isMatch(s string, p string) bool {
	dp := make([][]bool,len(s)+1)
	for i:=0;i<=len(s);i++{
		dp[i] = make([]bool,len(p)+1)
	}
	dp[0][0] = true
	for i:=1;i<=len(p);i++{
		if p[i-1] == '*'{
			dp[0][i] = dp[0][i-2]
		}

	}
	for i:=0;i<len(s);i++{
		for j:=0;j<len(p);j++{
			if s[i] == p[j] || p[j] == '.'{
				dp[i+1][j+1] = dp[i][j]
			}else if p[j] == '*'{
				if dp[i+1][j-1] == true{
					dp[i+1][j+1] = true
				}else if s[i] == p[j-1] || p[j-1]=='.'{
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

//盛最多水的容器
func maxArea(height []int) int {
	left,right := 0,len(height)-1
	max := math.MinInt
	for left<right{
		if height[left]<height[right]{
			vol := height[left]*(right-left)
			if max<vol{
				max = vol
			}
			left++
		}else {
			vol := height[right]*(right-left)
			if max<vol{
				max = vol
			}
			right--
		}
	}
	return max
}

//三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int,0)
	for i:=0;i<len(nums);i++{
		left,right := i+1,len(nums)-1
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		if nums[i]>0{
			break
		}
		for left<right{
			if nums[i] + nums[left] + nums[right] < 0{
				left++
			}else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			}else {
				list := []int{nums[i],nums[left],nums[right]}
				res = append(res,list)
				for left<right && nums[left]==nums[left+1]{
					left++
				}
				for left<right && nums[right]==nums[right-1]{
					right--
				}
				right--
				left++
			}
		}
	}
	return res
}

/*func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index + 1, combination + string(letters[i]))
		}
	}
}*/

//电话号码的字母组合
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits)==0{
		return nil
	}
	combinations = []string{}
	traceBack(0,digits,"")
	return combinations
}

func traceBack(index int,digits string,comb string)  {
	if index == len(digits){
		combinations = append(combinations,comb)
	}else {
		digit := string(digits[index])
		str1 := phoneMap[digit]
		for i:=0;i<len(str1);i++{
			a := comb
			a  = a + string(str1[i])
			traceBack(index+1,digits,a)
		}
	}
}

//删除链表的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil{
		return nil
	}
	quick,slow := head,head
	i:=0
	for slow.Next != nil && i<n {
		slow = slow.Next
		i++
	}
	if i<n{
		return head.Next
	}
	for slow.Next!=nil{
		quick = quick.Next
		slow = slow.Next
	}
	quick.Next = quick.Next.Next
	return head
}

//有效的括号
func isValid(s string) bool {
	keyMap := make(map[string]string)
	keyMap["("] = ")"
	keyMap["{"] = "}"
	keyMap["["] = "]"
	arr := make([]string,0)
	for i:=0;i<len(s);i++{
		c := string(s[i])
		if _,ok := keyMap[c];!ok{
			if len(arr)<1{
				return false
			}else if c != keyMap[arr[len(arr)-1]]{
				return false
			}else {
				arr = arr[:len(arr)-1]
			}
		}else {
			arr = append(arr,c)
		}
	}
	if len(arr)!=0{
		return false
	}
	return true
}

//合并两个有序链表
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for list1 != nil && list2 != nil{
		if list1.Val<=list2.Val{
			tmp := &ListNode{list1.Val,nil}
			head.Next = tmp
			head = head.Next
			list1 = list1.Next
		}else{
			tmp := &ListNode{list2.Val,nil}
			head.Next = tmp
			head = head.Next
			list2 = list2.Next
		}
	}
	for list1!=nil{
		tmp := &ListNode{list1.Val,nil}
		head.Next = tmp
		head = head.Next
		list1 = list1.Next
	}
	for list2!=nil{
		tmp := &ListNode{list2.Val,nil}
		head.Next = tmp
		head = head.Next
		list2 = list2.Next
	}
	return p.Next
}

func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	mergeHelp(list1,list2,p)
	return head.Next
}

func mergeHelp(list1 *ListNode,list2 *ListNode,head *ListNode)  {
	if list1!=nil || list2!=nil {
		if list1 != nil {
			if list2==nil || list1.Val<=list2.Val {
				tmp := ListNode{list1.Val,nil}
				list1 = list1.Next
				head.Next = &tmp
				head = head.Next
			}
		}
		if list2 != nil{
			if list1==nil || list2.Val<list1.Val {
				tmp := ListNode{list2.Val,nil}
				list2 = list2.Next
				head.Next = &tmp
				head = head.Next
			}
		}
		mergeHelp(list1,list2,head)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil{
		return list2
	}else if list2 == nil {
		return list1
	}else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next,list2)
		return list1
	}else {
		list2.Next = mergeTwoLists(list1,list2.Next)
		return list2
	}
}

//括号生成
func generateParenthesis(n int) []string {
	res := make([]string,0)
	generateParenthesisHelp(n-1,n,&res,"","(")
	return res
}

func generateParenthesisHelp(ln int,rn int,res *[]string,str string,c string){
	if ln==0 && rn==0{
		str = str + c
		*res = append(*res,str)
		return
	}
	if ln<0 || rn<0{
		return
	}
	if ln > rn{
		return
	}
	str = str + c
	generateParenthesisHelp(ln-1,rn,res,str,"(")
	generateParenthesisHelp(ln,rn-1,res,str,")")
}

//合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeListNode(lists,0,len(lists)-1)
}

func mergeListNode(lists []*ListNode,l,r int) *ListNode {
	if (l==r){
		return lists[l]
	}
	if l>r{
		return nil
	}
	mid := (l+r)>>1
	return mergeTwolist(mergeListNode(lists,l,mid),mergeListNode(lists,mid+1,r))
}

func mergeTwolist(a,b *ListNode) *ListNode{
	if a == nil || b == nil{
		if a != nil{
			return a
		}else {
			return b
		}
	}
	head := &ListNode{}
	tail := head
	aPtr,bPtr := a,b
	for aPtr != nil && bPtr != nil{
		if aPtr.Val<bPtr.Val{
			tail.Next = aPtr
			aPtr = aPtr.Next
		}else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr != nil{
		tail.Next = aPtr
	}else {
		tail.Next = bPtr
	}
	return head.Next
}

//子集
func subsets(nums []int) [][]int {
	res := make([][]int,0)
	arr := make([]int,0)
	dfs(nums,0,&res,arr)
	return res
}

func dfs(nums []int,step int, res *[][]int,arr []int){
	*res = append(*res,arr)
	if step == len(nums){
		return
	}
	for j:=step;j<len(nums);j++{
		arr = append(arr,nums[j])
		tmpArr := make([]int,len(arr))
		copy(tmpArr,arr)
		dfs(nums,j+1,res,tmpArr)
		arr = arr[:len(arr)-1]
	}
}

//下一个排列
func nextPermutation(nums []int)  {
	index := len(nums)-1
	for index>=1 && nums[index-1]>=nums[index]{
		index--
	}
	if index == 0{
		reverse(nums,0,len(nums)-1)
		return
	}
	i:=index
	for i<len(nums) && nums[index-1]<nums[i]{
		i++
	}
	nums[index-1],nums[i-1] = nums[i-1],nums[index-1]
	reverse(nums,index,len(nums)-1)
}

func reverse(nums []int,left,right int) {
	for i:=left;i<=(right+left)/2;i++{
		nums[i],nums[left+right-i] = nums[left+right-i],nums[i]
	}
}

//最长有效括号
func longestValidParentheses1(s string) int {
	stackLeft,stackRight,maxLength := 0,0,0
	for i:=0;i<len(s);i++{
		if s[i] == '('{
			stackLeft++
		}else {
			stackRight++
		}
		if stackLeft == stackRight{
			maxLength = max(maxLength,2*stackRight)
		} else if stackRight>stackLeft{
			stackRight,stackLeft = 0,0
		}
	}
	stackRight,stackLeft = 0,0
	for i:=len(s)-1;i>=0;i--{
		if s[i] == '('{
			stackLeft++
		}else {
			stackRight++
		}
		if stackLeft == stackRight{
			maxLength = max(maxLength,2*stackRight)
		} else if stackLeft>stackRight{
			stackRight,stackLeft = 0,0
		}
	}
	return maxLength
}

func longestValidParentheses(s string) int{
	maxAns := 0
	stack := []int{}
	stack = append(stack,-1)
	for i:=0;i<len(s);i++{
		if s[i] =='('{
			stack = append(stack,i)
		}else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0{
				stack = append(stack,i)
			}else {
				maxAns = max(maxAns,i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}



func search(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for left<=right{
		mid := left + (right-left)>>1
		if nums[mid] == target{
			return mid
		}
		if nums[0]<=nums[mid]{
			if target<=nums[mid] && target>=nums[0]{
				right = mid-1
			}else{
				left = mid+1
			}
		}else{
			if target<=nums[len(nums)-1] && target>=nums[mid]{
				left=mid + 1
			}else {
				right = mid-1
			}
		}
	}
	return -1
}

//在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left,right := 0,len(nums)-1
	var mid int
	for left<=right{
		mid = left + (right-left)>>1
		if nums[mid] == target{
			indexS := mid
			for indexS>=0 && nums[indexS]==target{
				indexS--
			}
			indexE := mid
			for indexE<= len(nums)-1 && nums[indexE]==target {
				indexE++
			}
			return []int{indexS+1,indexE-1}
		}else if nums[mid] < target{
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return []int{-1,-1}
}

//组合总和
func combinationSum(candidates []int, target int) [][]int {
	resComList := make([][]int,0)
	res := make([]int,0)
	findCombHelp(0,candidates,target,&resComList,res)
	return resComList
}

func findCombHelp(index int,nums []int,target int,resComList *[][]int,res []int)  {
	if target == 0{
		list := make([]int,len(res))
		copy(list,res)
		*resComList = append(*resComList,list)
		return
	}
	if target < 0{
		return
	}
	for i:=index;i<len(nums);i++{
		if target - nums[i]<0{
			continue
		}
		res = append(res,nums[i])
		findCombHelp(i,nums,target-nums[i],resComList,res)
		res = res[:len(res)-1]
	}
}

//接雨水
func trap(height []int) int {
	sum := 0
	maxLeft := make([]int,len(height))
	maxRight := make([]int,len(height))
	for i:=1;i<len(height);i++{
		maxLeft[i] = max(maxLeft[i-1],height[i-1])
	}
	for i:=len(height)-2;i>=0;i--{
		maxRight[i] = max(maxRight[i+1],height[i+1])
	}
	for i:=1;i<len(height);i++{
		min := findMin(maxRight[i],maxLeft[i])
		if min>height[i]{
			sum += (min - height[i])
		}
	}
	return sum
}

func findMin(a,b int) int{
	if a<b{
		return a
	}else {
		return b
	}
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//全排列
func permute(nums []int) [][]int {
	combinations := make([][]int,0)
	haveUsed := make([]bool,len(nums))
	res := make([]int,0)
	permuteHelp(0,haveUsed,nums,&combinations,res)
	return combinations
}

func permuteHelp(index int,haveUsed []bool,nums []int,combination *[][]int,res []int)  {
	if index==len(nums){
		list := make([]int,len(res))
		copy(list,res)
		*combination = append(*combination,list)
		return
	}

	for i:=0;i<len(nums);i++{
		if haveUsed[i]{
			continue
		}
		res = append(res,nums[i])
		haveUsed[i] = true
		permuteHelp(index+1,haveUsed,nums,combination,res)
		res = res[:len(res)-1]
		haveUsed[i] = false
	}
	return
}

//旋转图像
/*func rotate(matrix [][]int)  {
	length := len(matrix)
	roundLen := length
	for i:=0;i<length/2;i++{
		round := 0

	}
}*/

func main(){
	num := []int{5,4,6,2}
	fmt.Println(permute(num))
}