package main

func findMin(x,y int) int {
	if x>y{
		return y
	}else {
		return x
	}
}

func findMax(x,y int) int {
	if x>y{
		return x
	}else {
		return y
	}
}


func maxProfit1(prices []int) int {
	dp := make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i:=1;i<len(prices);i++{
		dp[i][0] = findMax(dp[i-1][0],-prices[i])
		dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit11(prices []int) int {
	keep,sold := -prices[0],0
	for i:=1;i<len(prices);i++{
		keep = findMax(keep,-prices[i])
		sold = findMax(sold,keep+prices[i])
	}
	return sold
}

func maxProfit2(prices []int) int {
	dp := make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i:=1;i<len(prices);i++{
		dp[i][0] = findMax(dp[i-1][0],dp[i-1][1]-prices[i])
		dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	return dp[len(dp)-1][1]
}

func maxProfit222(prices []int) int {
	keep,sold := -prices[0],0
	for i:=1;i<len(prices);i++{
		keep = findMax(keep,sold-prices[i])
		sold = findMax(sold,keep+prices[i])
	}
	return sold
}

func maxProfit444(prices []int) int {
	dp := make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,4)
	}
	dp[0][0],dp[0][2] = -prices[0],-prices[0]
	for i:=1;i<len(prices);i++{
		dp[i][0] = findMax(dp[i-1][0],-prices[i])
		dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]+prices[i])
		dp[i][2] = findMax(dp[i-1][2],dp[i-1][1]-prices[i])
		dp[i][3] = findMax(dp[i-1][3],dp[i-1][2]+prices[i])
	}
	return dp[len(prices) - 1][3]
}

func maxProfit(prices []int, fee int) int {
	dp := make([][]int,len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int,2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i:=1;i<len(prices);i++{
		dp[i][0] = findMax(dp[i-1][0],dp[i-1][1]-prices[i])
		dp[i][1] = findMax(dp[i-1][1],dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}