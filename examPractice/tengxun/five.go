package main

import (
	"fmt"
)

func main()  {
	var n,m int
	fmt.Scanln(&n,&m)
	price := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&price[i])
	}
	dp := make([][]int,n)
	money := make([][]bool,n)
	for i:=0;i<=n;i++{
		dp[i] = make([]int,n)
		money[i] = make([]bool,n)
	}
	dp[0][0] = m
	money[0][0] = true
	if m-price[0] > 0{
		money[0][1] = true
	}
	if money[0][1] == true{
		dp[0][1] = m - price[0]
	}

	for i:=1;i<n;i++{
		for j:=0;j<=i;j++{
			if money[i-1][j+1]==true{
				dp[i][j] = findMaxThree(dp[i-1][j],dp[i-1][j-1]-price[j],dp[i-1][j+1]+price[j])
			}else if dp[i-1][j-1]-price[i]>0 {
				dp[i][j] = findMax(dp[i-1][j],dp[i-1][j-1]-price[i])
			}
			if dp[i][j]>0{
				money[i][j] = true
			}
		}
	}
	fmt.Println()
}

func findMax(a,b int) int{
	if a>b{
		return a
	}else {
		return b
	}
}

func findMaxThree(a,b,c int) int{
	if a>b && a>c{
		return a
	}else if b>a && b>c {
		return b
	}else {
		return c
	}
}