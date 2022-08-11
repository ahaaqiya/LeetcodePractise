package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	n := 332
	//fmt.Println(n[0])
	fmt.Println(monotoneIncreasingDigits(n))
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

func maxProfit(prices []int, fee int) int {
	res := 0
	minPrice := prices[0]
	for i:=1;i<len(prices);i++{
		if prices[i]<minPrice{
			minPrice = prices[i]
		}
		if prices[i]>=minPrice && prices[i]-minPrice-fee<=0{
			continue
		}
		if prices[i]-minPrice-fee>0{
			res += prices[i]-minPrice-fee
			minPrice = prices[i]-fee
		}
	}
	return res
}

func monotoneIncreasingDigits(n int) int {
	strS := strconv.Itoa(n)
	arr := []byte(strS)
	for i:=len(arr)-1;i>0;i--{
		if arr[i-1]>arr[i]{
			arr[i-1] = arr[i-1]-1
			for j:=i;j<len(arr);j++{
				arr[j] = '9'
			}
		}
	}
	res,_ := strconv.Atoi(string(arr))
	return res
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	res := make([][]int,0)
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]<=intervals[i-1][1]{
			intervals[i][1] = findMax(intervals[i][1],intervals[i-1][1])
			intervals[i][0] = findMin(intervals[i][0],intervals[i-1][0])
		}else {
			res = append(res,[]int{intervals[i-1][0],intervals[i-1][1]})
		}
	}
	res = append(res,[]int{intervals[len(intervals)-1][0],intervals[len(intervals)-1][1]})
	return res
}

func partitionLabels(s string) []int {
	letterMap := make(map[byte]int)
	arrS := []byte(s)
	for i:=0;i<len(arrS);i++{
		letterMap[arrS[i]] = i
	}
	bound := 0
	start := 0
	res := make([]int,0)
	for i:=0;i<len(arrS);i++{
		bound = findMax(letterMap[arrS[i]],bound)
		start++
		if i==bound{
			res = append(res,start)
			start = 0
		}
	}
	return res
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0]==intervals[j][0]{
			return intervals[i][1]<intervals[j][1]
		}
		return intervals[i][0]<intervals[j][0]
	})
	minLeft := intervals[0][1]
	res := 0
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]<minLeft{
			res++
		}else {
			minLeft = intervals[i][1]
		}
	}
	return res
}

func findMinArrowShots(points [][]int) int {
	if len(points)==1{
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]<points[j][0]
	})
	res := 1
	minLeft := points[0][1]
	for i:=1;i<len(points);i++{
		if points[i][0]>minLeft{
			res++
			minLeft = points[i][1]
		}
		if minLeft>points[i][1]{
			minLeft = points[i][1]
		}
	}
	return res
}

func reconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0]==people[j][0]{
			return people[i][1]<people[j][1]
		}
		return people[i][0]>people[j][1]
	})
	result := make([][]int,0)
	for _,person := range people{
		result = append(result,person)
		copy(result[person[1]+1:],result[person[1]:])
		result[person[1]] = person
	}
	return result
}
func reconstructQueue(people [][]int) [][]int {
	//先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people,func(i,j int)bool{
		if people[i][0]==people[j][0]{
			return people[i][1]<people[j][1]//这个才是当身高相同时，将K按照从小到大排序
		}
		return people[i][0]>people[j][0]//这个只是确保身高按照由大到小的顺序来排，并不确定K是按照从小到大排序的
	})
	//再按照K进行插入排序，优先插入K小的
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1] +1:], result[info[1]:])//将插入位置之后的元素后移动一位（意思是腾出空间）
		result[info[1]] = info//将插入元素位置插入元素
	}
	return result
}

func lemonadeChange(bills []int) bool {
	five,ten,twenty := 0,0,0
	for i:=0;i<len(bills);i++{
		if bills[i] == 5{
			five++
		}else if bills[i]==10{
			if five<=0{
				return false
			}
			five--
			ten++
		}
		if bills[i] == 20{
			if ten>0 && five>0{
				ten--
				five--
				twenty++
			}else if five>=3{
				five = five-3
				twenty++
			}else{
				return false
			}
		}
	}
	return true
}

func candy(ratings []int) int {
	candyArr := make([]int,len(ratings))
	candyArr[0] = 1
	for i:=1;i<len(ratings);i++{
		if ratings[i]>ratings[i-1]{
			candyArr[i] = candyArr[i-1]+1
		}else{
			candyArr[i] = 1
		}
	}
	for i:=len(ratings)-2;i>=0;i--{
		if ratings[i]>ratings[i+1]{
			candyArr[i] = findMax(candyArr[i+1]+1,candyArr[i])
		}
	}
	sum := 0
	for i:=0;i<len(candyArr);i++{
		sum += candyArr[i]
	}
	return sum
}

func canCompleteCircuit(gas []int, cost []int) int {
	allSum,curSum := 0,0
	res := 0
	for i:=0;i<len(gas);i++{
		curSum += gas[i] - cost[i]
		allSum += gas[i] - cost[i]
		if curSum<0{
			curSum = 0
			res = i+1
		}
	}
	if allSum<0{
		return -1
	}else {
		return res
	}
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	for i:=0;i<len(nums);i++{
		if nums[i]<0 && k>0{
			k--
			nums[i] = -nums[i]
		}
		if nums[i]==0{
			k=0
		}
	}
	sort.Ints(nums)
	for k>0{
		nums[0] = -nums[0]
		k--
	}
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	return sum
}

func jump(nums []int) int {
	if len(nums)==1{
		return 1
	}
	curDistance := 0
	nextDistance := 0
	ans := 0
	for i:=0;i<len(nums);i++{
		nextDistance = findMax(i+nums[i],nextDistance)
		if i==curDistance{
			if curDistance != len(nums)-1{
				ans++
				curDistance = nextDistance
				if curDistance>len(nums)-1{
					break
				}
			}else {
				break
			}
		}
	}
	return ans
}

func canJump1(nums []int) bool {
	n := len(nums)
	if n==1{
		return true
	}
	cover := 0
	for i:=0;i<cover;i++{
		cover = findMax(cover,i+nums[i])
		if cover>=n-1{
			return true
		}
	}
	return false
}

func jump1(nums []int) int {
	dp := make([]int,len(nums))
	for i:=1;i<len(nums);i++{
		dp[i] = i
		for j:=0;j<i;j++{
			if j+nums[j]>=i{
				dp[i] = findMax(dp[j]+1,dp[i])
			}
		}
	}
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool,n)
	dp[0]=true
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			if dp[j] && j+nums[j]>=i{
				dp[i]=true
				break
			}
		}
	}
	return dp[n-1]
}


func maxProfit1(prices []int) int {
	dp := make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,2)
	}
	//0不持有，1持有一支
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i:=1;i<len(prices);i++{
		dp[i][0] = findMax(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]-prices[i])
	}
	return findMax(dp[len(prices)-1][1],dp[len(prices)-1][0])
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n==0{
		return 0
	}
	res := nums[0]
	count := nums[0]
	for i:=1;i<n;i++{
		if count+nums[i]<nums[i]{
			count = nums[i]
		}else {
			count = count + nums[i]
		}
		if count>res{
			res = count
		}
	}
	return res
}
/*func maxSubArray(nums []int) int {
	n := len(nums)
	if n<1{
		return 0
	}
	dp := make([]int,n)
	dp[0] = nums[0]
	res := dp[0]
	for i:=1;i<n;i++{
		dp[i] = findMax(dp[i-1]+nums[i],nums[i])
		if dp[i]>res{
			res = dp[i]
		}
	}
	return res
}*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := len(s)-1
	result := 0
	for i:=len(g)-1;i>=0;i--{
		if index >= 0 && s[index]>=g[i]{
			index--
			result++
		}
	}
	return result
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	/*up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1*/
	dp := make([][]int,n)
	for i:=0;i<n;i++{
		dp[i] = make([]int,2)
	}
	dp[0][0],dp[0][1] = 1,1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			/*up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]*/
			dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]+1)
			dp[i][0] = dp[i-1][0]
		} else if nums[i] < nums[i-1] {
			/*up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])*/
			dp[i][1] = dp[i-1][1]
			dp[i][0] = findMax(dp[i-1][1]+1,dp[i-1][0])
		} else {
			/*up[i] = up[i-1]
			down[i] = down[i-1]*/
			dp[i][1] = dp[i-1][1]
			dp[i][0] = dp[i-1][0]
		}
	}
	//return max(up[n-1], down[n-1])
	return findMax(dp[n-1][0],dp[n-1][1])
}

func wiggleMaxLength1(nums []int) int {
	if len(nums)<=1{
		return 1
	}
	result := 1
	preOff,curOff := 0,0
	for i:=1;i<len(nums);i++{
		curOff = nums[i]-nums[i-1]
		if (curOff>0 && preOff<=0) || (curOff<0 && preOff>=0){
			result++
			preOff = curOff
		}
	}
	return result
}