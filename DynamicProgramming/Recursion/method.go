package main

import (
	"math"
	"sort"
	"strconv"
)

func generate(numRows int) [][]int  {
	dp := make([][]int,numRows)
	for i:=0;i<numRows;i++{
		dp[i] = make([]int,i+1)
		dp[i][0] = 1
		dp[i][i] = 1
	}
	for i:=0;i<numRows;i++{
		for j:=1;j<len(dp[i])-1;j++{
			dp[i][j] = dp[i-1][j]+dp[i-1][j-1]
		}
	}
	return dp
}

func getRow(rowIndex int) []int {
	ans := make([]int,rowIndex)
	ans[0] = 1
	for i:=1;i<rowIndex;i++{
		for j:=i;j>0;j--{
			ans[j] = ans[j] + ans[j-1]
		}
	}
	return ans
}

func minPathSum(grid [][]int) int {
	m,n := len(grid),len(grid[0])
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	dp[0][0] = grid[0][0]
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0 && j==0{
				continue
			}else if i==0 && j>0{
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}else if j==0 && i>0{
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}else {
				dp[i][j] = findMin(dp[i-1][j]+grid[i][j],dp[i][j-1]+grid[i][j])
			}

		}
	}
	return dp[m-1][n-1]
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int,len(triangle))
	for i:=0;i<m;i++{
		dp[i] = make([]int,len(triangle))
	}
	dp[0][0] = triangle[0][0]
	for i:=0;i<m;i++{
		for j:=0;j<len(triangle[i]);j++{
			if i==0 && j==0{
				continue
			}else if i>0 && j==0{
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			}else if i>0 && j==i{
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			}else {
				dp[i][j] = findMin(triangle[i][j]+ dp[i-1][j],triangle[i][j]+dp[i-1][j-1])

			}
		}
	}
	min := math.MaxInt
	for i:=0;i<len(dp[m-1]);i++{
		if dp[m-1][i]<min{
			min = dp[m-1][i]
		}
	}
	return min
}

func findMinThree(x,y,z int) int {
	if x<=y && x<=z{
		return x
	}else if y<=x && y<=z{
		return y
	}else {
		return z
	}
}

//下降路径最小和
//https://leetcode-cn.com/problems/minimum-falling-path-sum-ii/
func minFallingPathSum2(matrix [][]int) int {
	dp := make([][]int,len(matrix))
	m := len(matrix)
	n := len(matrix[0])
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	min := math.MaxInt
	index1,index2 := -1,-1
	for i:=0;i<n;i++{
		dp[0][i] = matrix[0][i]
		var rule1,rule2 int
		if index1 == -1{
			rule1 = min
		}else {
			rule1 = matrix[0][index1]
		}
		if index2 == -1{
			rule2 = min
		}else {
			rule2 = matrix[0][index2]
		}
		if matrix[0][i]<rule1{
			index2 = index1
			index1 = i
		}else if matrix[0][i]<rule2{
			index2 = i
		}
	}
	for i:=1;i<m;i++{
		tmpIndex1,tmpIndex2 := -1,-1
		for j:=0;j<n;j++{
			if index1 != j{
				dp[i][j] = matrix[i][j] + dp[i-1][index1]
			}else {
				dp[i][j] = matrix[i][j] + dp[i-1][index2]
			}
			var rule1,rule2 int
			if tmpIndex1 == -1{
				rule1 = min
			}else {
				rule1 = dp[i][tmpIndex1]
			}
			if tmpIndex2 == -1{
				rule2 = min
			}else {
				rule2 = dp[i][tmpIndex2]
			}
			if dp[i][j]<rule1{
				tmpIndex2 = tmpIndex1
				tmpIndex1 = j
			}else if dp[i][j]<rule2{
				tmpIndex2 = j
			}
		}
		index1 = tmpIndex1
		index2 = tmpIndex2
	}
	min = math.MaxInt
	for i:=0;i<n;i++{
		if dp[m-1][i]<min{
			min = dp[m-1][i]
		}
	}
	return min
}

// 最大得分的路径数目
//https://leetcode-cn.com/problems/number-of-paths-with-max-score/
func pathsWithMaxScore(board []string) []int {
	n := len(board)
	dp := make([][]int,n)
	count := make([][]int,n)
	for i:=0;i<n;i++{
		dp[i] = make([]int,n)
		count[i] = make([]int,n)
	}
	INF := math.MinInt
	dp[n-1][n-1] = 0
	count[n-1][n-1] = 1
	for i:=n-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if i==n-1 && j == n-1{
				continue
			}
			if string(board[i][j])=="X"{
				dp[i][j] = INF
				count[i][j] = 0
				continue
			}
			val,_ := strconv.Atoi(string(board[i][j]))
			var cur,cnt int
			u,t := INF,0
			//试探向上走
			if i<n-1 {
				cur = dp[i+1][j]+val
				cnt = count[i+1][j]
				res := update(cur,cnt,u,t)
				u,t = res[0],res[1]
			}
			//试探向左走
			if j<n-1{
				cur = dp[i][j+1]+val
				cnt = count[i][j+1]
				res := update(cur,cnt,u,t)
				u,t = res[0],res[1]
			}
			//试探向左上走
			if j<n-1 && i<n-1{
				cur = dp[i+1][j+1] + val
				cnt = count[i+1][j+1]
				res := update(cur,cnt,u,t)
				u,t = res[0],res[1]
			}
			if u<0{
				u = INF
			}
			dp[i][j] = u
			count[i][j] = t
		}
	}
	if dp[0][0]==INF{
		return []int{0,0}
	}else{
		return []int{dp[0][0],count[0][0]}
	}
}

func update(cur,cnt,u,t int) []int {
	mod := int(1e9+7)
	ans := []int{u,t}
	if cur > u{
		ans[0] = cur
		ans[1] = cnt
	}else if cur==u && cur != math.MaxInt {
		ans[1] += cnt
	}
	ans[1] = ans[1]%mod
	return ans
}

func fib(n int) int {
	dp := make([]int,n+1)
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	if n==2{
		return 1
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i:=3;i<=n;i++{
		dp[i] = dp[i-1]+dp[i-2]+dp[i-3]
	}
	return dp[n]
}

func climbStairs(n int) int {
	if n<=2{
		return n
	}
	dp := make([]int,n+1)
	dp[0] = 0
	dp[1] = 1
	for i:=2;i<=n;i++{
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int,length)
	dp[0] = 0
	dp[1] = 0
	for i:=2;i<length;i++{
		dp[i] = findMin(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	}
	return findMin(dp[length-2]+cost[length-2],dp[length-1]+cost[length-1])
}

func findMin(a,b int) int {
	if a<b{
		return a
	}else {
		return b
	}
}

//打家劫舍
//https://leetcode-cn.com/problems/house-robber-ii/
func rob1(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	dp := make([]int,len(nums))
	maxArr := make([]int,len(nums)+2)
	dp[0] = nums[0]
	dp[1] = nums[1]
	max := math.MinInt
	for i:=0;i<len(nums);i++{
		if i>=2{
			dp[i] = dp[maxArr[i]]+nums[i]
		}
		if dp[i]>max{
			max = dp[i]
			maxArr[i+2] = i
		}else {
			maxArr[i+2] = maxArr[i+1]
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

func rob2(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	dp[1] = findMax(nums[0],nums[1])
	for i:=0;i<len(nums);i++{
		if i>=2{
			dp[i] = findMax(dp[i-1],dp[i-2]+nums[i])
		}
	}
	return dp[len(nums)-1]
}

func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	a := robhelp(nums[1:])
	b := robhelp(nums[:len(nums)-1])
	return findMax(a,b)
}

func robhelp(nums []int) int {
	pre,cur := 0,0
	for i:=0;i<len(nums);i++{
		tmp := cur
		if pre+nums[i] > cur{
			cur = pre+nums[i]
		}
		pre = tmp
	}
	return cur
}

//单词拆分
//https://leetcode-cn.com/problems/word-break-ii/
func wordBreak1(s string, wordDict []string) bool{
	dp := make([]bool,len(s)+1)
	dp[0] = true
	for i:=1;i<=len(s);i++{
		for _,word := range wordDict{
			if len(word)>i{
				continue
			}else if !dp[i] {
				dp[i] = dp[i-len(word)] && s[i-len(word):i]==word
			}
		}
	}
	return dp[len(s)]
}

func wordBreak(s string, wordDict []string) []string {
	dp := make([][]string,len(s)+1)
	dp[0] = []string{""}
	for i:=1;i<=len(s);i++{
		for _,word := range wordDict{
			if i<len(word){
				continue
			}
			if word == s[i-len(word):i]{
				for _,k:= range dp[i-len(word)]{
					tmp := word
					if k!=""{
						tmp = k + " " + word
					}
					dp[i] = append(dp[i],tmp)
				}
			}
		}
	}
	return dp[len(dp)-1]
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int,len(matrix))
	for i:=0;i<len(matrix);i++{
		dp[i] = make([]int,len(matrix[i]))
	}
	ans,tmp := 0,0
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			tmp = dfs(i,j,matrix,dp)
			if tmp>ans{
				ans = tmp
			}
		}
	}
	return ans
}

func dfs(i, j int,matrix [][]int,dp [][]int) int {
	var up,down,left,right int
	if dp[i][j] != 0{
		return dp[i][j]
	}
	dp[i][j]++
	if i+1<len(matrix) && matrix[i+1][j]>matrix[i][j]{
		down = dfs(i+1,j,matrix,dp) + 1
		if dp[i][j]<down{
			dp[i][j] = down
		}
	}
	if i-1>=0 && matrix[i-1][j]>matrix[i][j]{
		up = dfs(i-1,j,matrix,dp)+1
		if dp[i][j]<up{
			dp[i][j] = up
		}
	}
	if j+1<len(matrix[i]) && matrix[i][j+1]>matrix[i][j]{
		left = dfs(i,j+1,matrix,dp)+1
		if dp[i][j]<left{
			dp[i][j] = left
		}
	}
	if j-1>=0 && matrix[i][j-1]>matrix[i][j]{
		right = dfs(i,j-1,matrix,dp)+1
		if dp[i][j]<right{
			dp[i][j] = right
		}
	}
	return dp[i][j]
}

func canPartition(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	if sum%2 != 0 || len(nums)<=1{
		return false
	}
	target := sum/2
	//创建二维状态数组，行：物品索引，列：容量
	dp := make([][]bool,len(nums))
	for i:=0;i<len(nums);i++{
		dp[i] = make([]bool,target+1)
	}
	dp[0][0] = false
	if nums[0] < target{
		dp[0][nums[0]] = true
	}
	for i:=1;i<len(nums);i++{
		for j:=0;j<=target;j++{
			dp[i][j] = dp[i-1][j]
			if nums[i] == j{
				dp[i][j] = true
				continue
			}
			if nums[i]<j{
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[j]]
			}
		}
	}
	return dp[len(nums)-1][target]
}

func abs(x int) int {
	if x>=0{
		return x
	}else {
		return -x
	}
}

func findTargetSumWays(nums []int, target int) int {
	dp := make([][]int,len(nums)+1)
	sum := 0
	for i:=0;i<len(nums);i++{
		sum = sum + abs(nums[i])
	}
	if sum<abs(target){
		return 0
	}
	for i:=0;i<=len(nums);i++{
		dp[i] = make([]int,2*sum+1)
	}
	dp[0][sum] = 1
	for i:=1;i<=len(nums);i++{
		for j:=0;j<=2*sum;j++{
			if j-nums[i-1]>=0 && j+nums[i-1]<=2*sum{
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j+nums[i-1]]
			}
			if j-nums[i-1]<0{
				dp[i][j] = dp[i-1][j+nums[i-1]]
			}else if j+nums[i-1]>2*sum{
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum+target]
}

func wiggleSort(nums []int)  {
	index := 1
	for index<len(nums){
		if index%2==0{
			if nums[index]<nums[index-1] {
				tmp := nums[index]
				nums[index] = nums[index-1]
				nums[index-1] = tmp
			}
		}else {
			if nums[index]>nums[index-1]{
				tmp := nums[index]
				nums[index] = nums[index-1]
				nums[index-1] = tmp
			}
		}
		index++
	}
}

func merge(A []int, m int, B []int, n int)  {
	i,j := m-1,n-1
	for i>=0 && j>=0{
		if A[i] > B[j]{
			A[i+j+1] = A[i]
			i--
		}else {
			A[i+j+1] = B[j]
			j--
		}
	}
	for j>=0{
		A[j] = B[j]
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int,len(text2)+1)
	for i:=0;i<=len(text2);i++{
		dp[i] = make([]int,len(text1)+1)
		if i==0{
			for j:=0;j<=len(text1);j++{
				dp[i][j] = 0
			}
		}
		dp[i][0] = 0
	}
	for i:=1;i<=len(text2);i++{
		for j:=1;j<=len(text1);j++{
			if text2[i-1]==text1[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}else {
				dp[i][j] = findMax(dp[i][j-1],dp[i-1][j])
			}
		}
	}
	return dp[len(text2)][len(text1)]
}

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int,len(matrix))
	for i:=0;i<len(matrix);i++{
		dp[i] = make([]int,len(matrix[i]))
	}
	for i:=0;i<len(matrix[0]);i++{
		dp[0][i] = matrix[0][i]
	}
	for i:=1;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if j>=0 && j<len(matrix[0])-1{
				dp[i][j] = matrix[i][j] + findMinThree(dp[i-1][j-1],dp[i-1][j],dp[i-1][j+1])
			}else if j==0{
				dp[i][j] = matrix[i][j] + findMin(dp[i-1][j+1],dp[i-1][j])
			}else if j==len(matrix[i])-1{
				dp[i][j] = matrix[i][j] + findMin(dp[i-1][j-1],dp[i-1][j])
			}
		}
	}
	res := math.MaxInt
	for i:=0;i<len(matrix[0]);i++{
		if dp[len(matrix)-1][i]<res{
			res = dp[len(matrix)-1][i]
		}
	}
	return res
}

//石子游戏
func stoneGame1(piles []int) bool {
	if len(piles)<=1{
		return true
	}
	dp:= make([][]int,len(piles))
	for i:=0;i<len(piles);i++{
		dp[i] = make([]int,len(piles))
		dp[i][i] = piles[i]
	}
	for i:=len(piles)-2;i>=0;i--{
		for j:=i+1;j<len(piles);j++{
			dp[i][j] = findMax(piles[i]-dp[i+1][j],piles[j]-dp[i][j-1])
		}
	}
	if dp[0][len(piles)-1]<0{
		return false
	}else {
		return true
	}
}
//空间优化
func stoneGame(piles []int) bool {
	if len(piles)<=1{
		return true
	}
	dp := make([]int,len(piles))
	for i:=0;i<len(piles);i++{
		dp[i] = piles[i]
	}
	for i:=len(piles)-2;i>=0;i--{
		for j:=i+1;j<len(piles);j++{
			dp[j] = findMax(piles[i]-dp[j],piles[j]-dp[j-1])
		}
	}
	return dp[len(dp)-1]>0
}

func findNthDigit(n int) int {
	round := 1
	for n>10{
		n = n - 9*(round-1)*round
	}
}


func main()  {
	arr := []int{5,3,4,5}
	stoneGame(arr)
}