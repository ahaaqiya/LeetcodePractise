package main

import "fmt"

func main()  {
	nums := []int{25,30,125,64}
	fmt.Println(maxZero(nums,4))
}

func findMin(a,b int) int{
	if a<=b{
		return a
	}else{
		return b
	}
}

func abs(a int) int {
	if a<=0{
		return -a
	} else {
		return a
	}
}

func maxZero(nums []int,n int) int{
	dp := make([][]int,n)
	count := make([]int,n)
	for i:=0;i<n;i++{
		dp[i] = make([]int,3)
	}
	dp[0][0],dp[0][1],dp[0][2] = yinshu1025(nums[0])
	count[0] = dp[0][0]
	dp[1][0],dp[1][1],dp[1][2] = yinshu1025(nums[1])
	count[1] = dp[1][0]
	for i:=2;i<n;i++{
		zero,two,five := yinshu1025(nums[i])
		shu := dp[i-2][0]+zero+findMin(dp[i-2][1],five)+findMin(dp[i-2][2],two)
		if shu>count[i-1]{
			count[i] = shu
			dp[i][0] = shu
			dp[i][1] = abs(dp[i-2][1]-five)
			dp[i][2] = abs(dp[i-2][2]-two)
		}else{
			count[i] = count[i-1]
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1]
			dp[i][2] = dp[i-1][2]
		}
	}
	return count[n-1]
}

func yinshu1025(num int)  (int,int,int){
	zero := 0
	two := 0
	five := 0
	for num%10==0{
		zero++
		num = num/10
	}
	for num%2==0{
		two++
		num = num/2
	}
	for num%5==0{
		five++
		num = num/5
	}
	return zero,two,five
}

func findMinThree(a,b,c int) int {
	if a<=b && a<=c {
		return a
	}else if b<=a && b<=c {
		return b
	}else {
		return c
	}
}

func minDistance(word1 string, word2 string) int {
	dp := make([]int,len(word2)+1)
	for i:=1;i<=len(word2);i++{
		dp[i] = i
	}
	tmp := make([]int,len(word2)+1)
	for i:=1;i<=len(word1);i++{
		copy(tmp,dp)
		front := i
		dp[0] = i
		for j:=1;j<=len(word2);j++{
			if word1[i-1]==word2[j-1]{
				dp[j] = tmp[j-1]
			}else{
				dp[j] = findMinThree(tmp[j-1]+1,tmp[j]+1,front+1)
			}
			front = dp[j]
		}
	}
	return dp[len(word2)]
	/*dp := make([][]int,len(word1)+1)
	for i:=0;i<=len(word1);i++{
		dp[i] = make([]int,len(word2)+1)
		if i==0{
			for j:=1;j<=len(word2);j++{
				dp[i][j] = j
			}
		}
		dp[i][0] = i
	}
	for i:=1;i<=len(word1);i++{
		for j:=1;j<=len(word2);j++{
			if word1[i-1]==word2[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = findMinThree(dp[i-1][j-1]+1,dp[i][j-1]+1,dp[i-1][j]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]*/
}

func minDistance2(word1 string, word2 string) int {
	//压缩
	dp := make([]int,len(word2)+1)
	for i:=1;i<=len(word2);i++{
		dp[i] = i
	}
	tmp := make([]int,len(word2)+1)
	for i:=1;i<=len(word1);i++{
		copy(tmp,dp)
		front := i
		dp[0] = i
		for j:=1;j<=len(word2);j++{
			if word1[i-1]==word2[j-1]{
				dp[j] = tmp[j-1]
			}else{
				dp[j] = findMinThree(tmp[j-1]+2,tmp[j]+1,front+1)
			}
			front = dp[j]
		}
	}
	return dp[len(word2)]
	/*dp := make([][]int,len(word1)+1)
	for i:=0;i<=len(word1);i++{
		dp[i] = make([]int,len(word2)+1)
		if i==0{
			for j:=1;j<=len(word2);j++{
				dp[i][j] = j
			}
		}
		dp[i][0] = i
	}
	for i:=1;i<=len(word1);i++{
		for j:=1;j<=len(word2);j++{
			if word1[i-1]==word2[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = findMinThree(dp[i-1][j-1]+2,dp[i][j-1]+1,dp[i-1][j]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]*/
}

func isSubsequence(s string, t string) bool {
	dp := make([][]bool,len(t)+1)
	for i:=0;i<=len(t);i++{
		dp[i] = make([]bool,len(s)+1)
	}
	if len(s)>len(t){
		return false
	}
	for i:=0;i<=len(t);i++{
		dp[i][0] = true
	}
	for i:=1;i<=len(t);i++{
		for j:=1;j<=len(s);j++{
			if t[i-1]==s[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(t)][len(s)]
}

func numDistinct(s string, t string) int {
	if len(t)>len(s){
		return 0
	}
	//压缩
	dp := make([]int,len(t)+1)
	dp[0] = 1
	tmp := make([]int,len(t)+1)
	for i:=1;i<=len(s);i++ {
		tmp = dp
		for j:=1;j<=len(t);j++{
			if s[i-1]==t[j-1]{
				dp[j] = tmp[j-1]+tmp[j]
			}
		}
	}
	return dp[len(t)]
	/*dp := make([][]int,len(t)+1)
	for i:=0;i<=len(t);i++{
		dp[i] = make([]int,len(s)+1)
		if i==0{
			for j:=0;j<=len(s);j++{
				dp[i][j] = 1
			}
		}
	}
	for i:=1;i<=len(t);i++{
		for j:=1;j<=len(s);j++ {
			if t[i-1]==s[j-1]{
				dp[i][j] = dp[i][j-1]+dp[i-1][j-1]
			}else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]*/
}

func maxSubArray(nums []int) int {
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i:=1;i<len(nums);i++{
		if nums[i]>dp[i-1]+nums[i]{
			dp[i] = nums[i]
		}else {
			dp[i] = dp[i-1]+nums[i]
		}
		if res<dp[i]{
			res = dp[i]
		}
	}
	return res
}

func findLengthOfLCIS(nums []int) int {
	dp := make([]int,len(nums))
	dp[0]  = 1
	res := dp[0]
	for i:=1;i<len(nums);i++{
		if nums[i]>nums[i-1]{
			dp[i] = dp[i-1]+1
		}else {
			dp[i] = 1
		}
		if res<dp[i]{
			res = dp[i]
		}
	}
	return res
}

/*func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int,len(nums1)+1)
	res := 0
	for i:=0;i<=len(nums1);i++{
		dp[i] = make([]int,len(nums2)+1)
	}
	for i:=1;i<=len(nums1);i++{
		for j:=1;j<=len(nums2);j++{
			if nums1[i-1]==nums2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}
			if res<dp[i][j]{
				res = dp[i][j]
			}
		}
	}
	return res
}*/

func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	dp[0]  = 1
	res := dp[0]
	for i:=1;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				dp[i] = findMax(dp[i],dp[j]+1)
			}
		}
		if res<dp[i]{
			res = dp[i]
		}
	}
	return res
}
func findMax(x,y int) int {
	if x>y{
		return x
	}else {
		return y
	}
}
