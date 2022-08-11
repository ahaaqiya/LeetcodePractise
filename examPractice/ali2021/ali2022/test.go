package main

import (
	"fmt"
	"math"
)

func step(n,m int) int{
	mod := int(1E9 + 7)
	dp := make([][][]int,n+1)
	for i:=0;i<=n;i++{
		dp[i] = make([][]int,m+1)
		for j:=0;j<=m;j++{
			dp[i][j] = make([]int,m+1)
		}
	}
	dp[0][0][0] = 1
	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			for k:=0;k<=m;k++{
				for l:=0;l<=m;l++{
					if (j>i) || (j==k && k!=0) || (k==l && k!=0) || (j==l && l!=0){
						continue
					}
					dp[i][j][k] += dp[i-j][k][l]
					dp[i][j][k] = dp[i][j][k] % mod
				}
			}
		}
	}
	ans := 0
	for i:=0;i<=m;i++{
		for j:=0;j<=m;j++ {
			ans = ans + dp[n][i][j]
			ans = ans%mod
		}
	}
	return ans
}

func numSubmat(mat [][]int) int {
	m,n := len(mat),len(mat[0])
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	res := 0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if mat[i][j] ==0{
				continue
			}
			if mat[i][j]==1{
				if j==0{
					dp[i][j] = 1
				}else {
					dp[i][j] = dp[i][j-1]+1
				}
			}
			min := math.MaxInt
			for h:=0;h<=i;h++{
				if mat[i-h][j]==0{
					break
				}
				min = findMin(min,dp[i-h][j])
				res = res + min
			}
		}
	}
	return res
}

func findMin(x,y int) int {
	if x>y{
		return y
	}else {
		return x
	}
}
func main() {
	var n,m int
	fmt.Scanln(&n,&m)
	fmt.Println(step(n,m))
}
