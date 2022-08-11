package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

func main()  {
	board := [][]int{{0,1},{2,1},{0,0}}
	fmt.Println(isBoomerang(board))

}

func isBoomerang(points [][]int) bool {
	if reflect.DeepEqual(points[0],points[1]) || reflect.DeepEqual(points[1],points[2]) || reflect.DeepEqual(points[0],points[2]) {
		return false
	}
	xie1 := xie(points[0],points[1])
	xie2 := xie(points[1],points[2])
	return xie1!=xie2
}

func xie(a,b []int) float64 {
	if a[0]-b[0]==0{
		return math.MaxInt
	}else if a[1]-b[1]==0{
		return 0
	}else {
		return (float64(a[1]-b[1]))/(float64(a[0]-b[0]))
	}
}


func solveSudoku(board [][]byte)  {
	var backward func([][]byte) bool
	numsstr := "123456789"
	nums := []byte(numsstr)
	n := len(board)
	backward = func(board [][]byte) bool {
		for i:=0;i<n;i++{
			for j:=0;j<n;j++{
				if board[i][j]!='.'{
					continue
				}
				for index:=0;index<len(nums);index++{
					if isVaildSudoku(n,i,j,board,nums[index]){
						board[i][j]=nums[index]
						if backward(board){
							return true
						}
						board[i][j]='.'
					}
				}
				return false
			}
		}
		return true
	}
	backward(board)
}

func isVaildSudoku(n,row,col int,board [][]byte,num byte) bool{
	//查看行
	for i:=0;i<n;i++{
		if board[row][i]==num && i!=col{
			return false
		}
	}
	//查看列
	for i:=0;i<n;i++{
		if board[i][col]==num && i!=row{
			return false
		}
	}
	//查看九宫格
	startCol := col - col%3
	startRow := row - row%3
	for i:=startRow;i<startRow+3;i++{
		for j:=startCol;j<startCol+3;j++{
			if board[i][j]==num{
				return false
			}
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	checkboard := make([][]string,n)
	for i:=0;i<n;i++{
		checkboard[i] = make([]string,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			checkboard[i][j] = "."
		}
	}
	var backward func(int)
	result := make([][]string,0)
	backward = func(row int) {
		if row==n{
			temp := make([]string,n)
			for i,str := range checkboard{
				temp[i] = strings.Join(str,"")
			}
			result = append(result,temp)
			return
		}
		for i:=0;i<n;i++{
			if isValid(n,row,i,checkboard){
				checkboard[row][i]="Q"
				backward(row+1)
				checkboard[row][i]="."
			}
		}
	}
	backward(0)
	return result
}

func isValid(n,row,col int,checkboard [][]string) bool{
	for i:=0;i<row;i++{
		if checkboard[i][col]=="Q"{
			return false
		}
	}
	for i,j:=row-1,col-1;i>=0&&j>=0;i,j = i-1,j-1{
		if checkboard[i][j]=="Q"{
			return false
		}
	}
	for i,j:=row-1,col+1;i>=0&&j<n;i,j = i-1,j+1{
		if checkboard[i][j]=="Q"{
			return false
		}
	}
	return true
}


type pair struct {
	airport string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int{
	return len(p)
}
func (p pairs) Swap(i,j int) {
	p[i],p[j] = p[j],p[i]
}
func (p pairs) Less(i,j int) bool {
	return p[i].airport < p[j].airport
}

func findItinerary(tickets [][]string) []string {
	targets := make(map[string]pairs)
	for _,fly := range tickets{
		start := fly[0]
		if _,have := targets[start];!have{
			targets[fly[0]] = make(pairs,0)
		}
		targets[fly[0]] = append(targets[fly[0]],&pair{airport: fly[1],visited: false})
	}
	for k,_ := range targets{
		sort.Sort(targets[k])
	}
	result := make([]string,0)
	result = append(result,"JFK")
	var backward func() bool
	backward = func() bool {
		if len(result) == len(tickets)+1{
			return true
		}
		for _,pair := range targets[result[len(result)-1]]{
			if !pair.visited{
				result = append(result,pair.airport)
				pair.visited = true
				if backward(){
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}

		}
		return false
	}
	backward()
	return result
}

/*func findItinerary(tickets [][]string) []string {
	mapPlatform := make(map[string][]string)
	for i:=0;i<len(tickets);i++{
		to,have := mapPlatform[tickets[i][0]]
		newTo := tickets[i][1]
		if !have{
			arr := make([]string,0)
			arr = append(arr,newTo)
			mapPlatform[tickets[i][0]] = arr
		}else{
			to = insert(newTo,to)
			mapPlatform[tickets[i][0]] = to
		}
	}

	return nil
}

func backward(mapPlatform map[string][]string,start string,used []string,res *[]string)  {
	if len(used)==0{
		return
	}
	if _,have := mapPlatform[start];!have{
		return
	}

}
*/
func insert(a string,arr []string) []string{
	for i:=0;i<len(arr);i++{
		if compare(a,arr[i]){
			arr = append(arr[:i],append([]string{a},arr[i:]...)...)
			break
		}
	}
	return arr
}

func compare(a,b string) bool{
	for i:=0;i<len(a);i++{
		if a[i]<b[i]{
			return true
		}else if a[i]>b[i]{
			return false
		}
	}
	return true
}

func permuteUnique(nums []int) [][]int {
	combine := make([]int,0)
	res := make([][]int,0)
	used := make([]bool,len(nums))
	sort.Ints(nums)
	permuteUniqueHelp(nums,combine,used,&res)
	return res
}

func permuteUniqueHelp(nums []int,combine []int,used []bool,res *[][]int){
	if len(combine)==len(nums){
		tmp := make([]int,len(nums))
		copy(tmp,combine)
		*res = append(*res,tmp)
		return
	}
	//have := make(map[int]bool)
	for i:=0;i<len(nums);i++{
		if used[i]{
			continue
		}
		if i>0 && nums[i]==nums[i-1] && !used[i-1]{
			continue
		}
		/*if _,ok := have[nums[i]];ok{
			continue
		}*/
		used[i] = true
		//have[nums[i]]=true
		combine = append(combine,nums[i])
		permuteUniqueHelp(nums,combine,used,res)
		combine = combine[:len(combine)-1]
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	combine := make([]int,0)
	res := make([][]int,0)
	used := make([]bool,len(nums))
	permuteHelp(nums,combine,used,&res)
	return res
}

func permuteHelp(nums []int,combine []int,used []bool,res *[][]int){
	if len(combine)==len(nums){
		tmp := make([]int,len(nums))
		copy(tmp,combine)
		*res = append(*res,tmp)
		return
	}
	for i:=0;i<len(nums);i++{
		if used[i]{
			continue
		}
		used[i] = true
		combine = append(combine,nums[i])
		permuteHelp(nums,combine,used,res)
		combine = combine[:len(combine)-1]
		used[i] = false
	}
}



func findSubsequences(nums []int) [][]int {
	path := make([]int,0)
	res := make([][]int,0)
	findSubsequencesHelp(nums,0,path,&res)
	return res
}

func findSubsequencesHelp(nums []int,startIndex int,path []int,res *[][]int){
	if startIndex == len(nums){
		return
	}
	used := make(map[int]bool)
	for i:=startIndex;i<len(nums);i++{
		if len(path)!=0 && path[len(path)-1]>nums[i]{
			continue
		}
		if _,have := used[nums[i]];have{
			continue
		}
		path = append(path,nums[i])
		used[nums[i]] = true
		if len(path)>1{
			tmp := make([]int,len(path))
			copy(tmp,path)
			*res = append(*res,tmp)
		}
		findSubsequencesHelp(nums,i+1,path,res)
		path = path[:len(path)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	count := make([]int,0)
	res := make([][]int,0)
	res = append(res,[]int{})
	subsetsWithDupHelp(nums,0,count,&res)
	return res
}

func subsetsWithDupHelp(nums []int,startIndex int,path []int,res *[][]int) {
	if startIndex==len(nums){
		return
	}
	for i:=startIndex;i<len(nums);i++{
		if i>startIndex && nums[i]==nums[i-1]{
			continue
		}
		path = append(path,nums[i])
		tmp := make([]int,len(path))
		copy(tmp,path)
		*res = append(*res,tmp)
		subsetsWithDupHelp(nums,i+1,path,res)
		path = path[:len(path)-1]
	}
}

func subsets(nums []int) [][]int {
	count := make([]int,0)
	res := make([][]int,0)
	res = append(res,[]int{})
	subsetsHelp(nums,0,count,&res)
	return res
}

func subsetsHelp(nums []int,startIndex int,count []int,res *[][]int) {
	if startIndex==len(nums){
		return
	}
	for i:=startIndex;i<len(nums);i++{
		count = append(count,nums[i])
		tmp := make([]int,len(count))
		copy(tmp,count)
		*res = append(*res,tmp)
		subsetsHelp(nums,i+1,count,res)
		count = count[:len(count)-1]
	}
}

func restoreIpAddresses(s string) []string {
	combine := make([]string,0)
	res := make([][]string,0)
	restoreIpHelp(s,3,1,combine,&res)
	resIp := ipChage(res)
	return resIp
}

func restoreIpHelp(s string,order int,index int,combine []string,res *[][]string)  {
	if len(s)==0{
		if len(combine)==4{
			tmp := make([]string,len(combine))
			copy(tmp,combine)
			*res = append(*res,tmp)
		}
		return
	}
	//if s[index]
	for i:=1;i<=len(s) && i<=3;i++{
		str1,str2 := s[:i],s[i:]
		if len(str2)>3*order{
			continue
		}
		if len(str1)>1 && str1[0]=='0'{
			return
		}
		if len(str1)==3{
			if str1[0]>'2'{
				return
			}else if str1[0]=='2' && str1[1]>'5' {
				return
			}else if str1[0]=='2' && str1[1]=='5' && str1[2]>'5'{
				return
			}
		}
		combine = append(combine,str1)
		restoreIpHelp(s[i:],order-1,i,combine,res)
		combine = combine[:len(combine)-1]

	}
}

func ipChage(res [][]string) []string {
	if len(res)==0{
		return nil
	}
	resIp := make([]string,0)
	for i:=0;i<len(res);i++{
		var str string
		ip := res[i]
		for j:=0;j<len(ip);j++{
			if j==len(ip)-1{
				str += ip[j]
			}else{
				str += ip[j]+"."
			}
		}
		resIp = append(resIp,str)
	}
	return resIp
}

func partition(s string) [][]string {
	if len(s)==0{
		return nil
	}
	combine := make([]string,0)
	res := make([][]string,0)
	partitionHelp(s,1,combine,&res)
	return res
}

func partitionHelp(s string,index int,combine []string,res *[][]string){
	if index>len(s){
		tmp := make([]string,len(combine))
		copy(tmp,combine)
		*res = append(*res,tmp)
		return
	}
	for i:=1;i<=len(s);i++{
		str1,str2 := s[:i],s[i:]
		if !judgeP(str1){
			continue
		}
		combine = append(combine,str1)
		partitionHelp(str2,1,combine,res)
		combine = combine[:len(combine)-1]
	}
}

func judgeP(str string) bool {
	if len(str)==1{
		return true
	}
	start,end := 0,len(str)-1
	for start<=end{
		if str[start]!=str[end]{
			return false
		}
		start++
		end--
	}
	return true
}

func combinationSum2(candidates []int, target int) [][]int {
	count := make([]int,0)
	res := make([][]int,0)
	//used := make([]bool,len(candidates))
	sort.Ints(candidates)
	combinationSum2Help(candidates,target,0,0,count,&res)
	return res
}
func combinationSum2Help(candidates []int, target int,index,sum int,count []int,res *[][]int){
	if sum==target{
		tmp := make([]int,len(count))
		copy(tmp,count)
		*res = append(*res,tmp)
		return
	}
	if sum>target{
		return
	}
	for i:=index;i<len(candidates);i++{
		if i>index && candidates[i]==candidates[i-1]{
			continue
		}
		sum = sum+candidates[i]
		count = append(count,candidates[i])
		combinationSum2Help(candidates,target,i+1,sum,count,res)
		count = count[:len(count)-1]
		sum = sum - candidates[i]
	}
}

/*func combinationSum2Help(candidates []int, target int,index,sum int,count []int,res *[][]int,used []bool){
	if sum==target{
		tmp := make([]int,len(count))
		copy(tmp,count)
		*res = append(*res,tmp)
		return
	}
	if sum>target{
		return
	}
	for i:=index;i<len(candidates);i++{
		if i>0 && candidates[i]==candidates[i-1]{
			if used[i-1] == false{
				continue
			}
		}
		sum = sum+candidates[i]
		count = append(count,candidates[i])
		used[i] = true
		combinationSum2Help(candidates,target,i+1,sum,count,res,used)
		count = count[:len(count)-1]
		sum = sum - candidates[i]
		used[i] = false
	}
}*/

func combinationSum(candidates []int, target int) [][]int {
	count := make([]int,0)
	res := make([][]int,0)
	combinationSumHelp(candidates,target,0,0,count,&res)
	return res
}

func combinationSumHelp(candidates []int, target int,index,sum int,count []int,res *[][]int){
	if sum==target{
		tmp := make([]int,len(count))
		copy(tmp,count)
		*res = append(*res,tmp)
		return
	}
	if sum>target{
		return
	}
	for i:=index;i<len(candidates);i++{
		sum = sum+candidates[i]
		count = append(count,candidates[i])
		combinationSumHelp(candidates,target,i,sum,count,res)
		count = count[:len(count)-1]
		sum = sum - candidates[i]
	}
}

func letterCombinations(digits string) []string {
	vocab := map[string]string{
		"2":"abc",
		"3":"def",
		"4":"ghi",
		"5":"jkl",
		"6":"mno",
		"7":"pqs",
		"8":"tuv",
		"9":"wxyz",
	}
	arrDigit := []byte(digits)
	res := make([]string,0)
	if len(arrDigit)==0{
		return nil
	}
	letterHelp(0,arrDigit,vocab,"",&res)
	return res
}

func letterHelp(index int,arrDigit []byte,vocab map[string]string,count string,res *[]string)  {
	if index==len(arrDigit){
		var tmp string
		tmp = count
		*res = append(*res,tmp)
		return
	}
	str,have := vocab[string(arrDigit[index])]; if !have{
		return
	}
	for i:=0;i<len(str);i++{
		count = count+str[i:i+1]
		letterHelp(index+1,arrDigit,vocab,count,res)
		count = count[:len(count)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	sum := 0
	for i:=1;i<=k;i++{
		sum += i
	}
	if sum>n{
		return nil
	}
	count := make([]int,0)
	res := make([][]int,0)
	combinationSum3Help(k,n,0,1,count,&res)
	return res
}

func combinationSum3Help(k,n,sum int,index int,count []int,res *[][]int)  {
	if sum==n && len(count)==k{
		tmp := make([]int,k)
		copy(tmp,count)
		*res = append(*res,tmp)
		return
	}
	if index>n || sum+index>n || len(count)>k{
		return
	}
	for i:=index;i<=9;i++{
		sum = sum+i
		count = append(count,i)
		combinationSum3Help(k,n,sum,i+1,count,res)
		count = count[:len(count)-1]
		sum = sum-i
	}
}

/*var res [][]int
func combine(n int, k int) [][]int {
	res=[][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}
func backtrack(n,k,start int,track []int){
	if len(track)==k{
		temp:=make([]int,k)
		copy(temp,track)
		res=append(res,temp)
	}
	if len(track)+n-start+1 < k {
		return
	}
	for i:=start;i<=n;i++{
		track=append(track,i)
		backtrack(n,k,i+1,track)
		track=track[:len(track)-1]
	}
}*/

func combine(n int, k int) [][]int {
	res := make([][]int,0)
	count := make([]int,0)
	combineHelp(n,k,1,count,&res)
	return res
}

func combineHelp(n,k,index int,count []int,res *[][]int){
	if len(count)==k{
		tmp := make([]int,k)
		copy(tmp,count)
		*res = append(*res,tmp)
		return
	}
	if len(count)+n-index+1<k{
		return
	}
	for i:=index;i<=n;i++{
		count = append(count,i)
		combineHelp(n,k,i+1,count,res)
		count = count[:len(count)-1]
	}
}
