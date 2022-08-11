package main

import (
	"fmt"
)

func main()  {
	weight := []int{2,3,1,1}
	skip := []int{2,1,1,2}
	cutWater(weight,skip,4)
}

func findMax(a,b int) int {
	if a>b{
		return a
	}else {
		return b
	}
}

func comBag(vol,cost []int,N,V int) int {
	dp := make([][]int,N+1)
	for i:=0;i<=N;i++{
		dp[i] = make([]int,V+1)
	}
	dp[0][0] = 0
	for i:=1;i<=N;i++{
		for j:=1;j<=V;j++{
			/*for k:=0;k<=j/vol[i-1];k++{
				dp[i][j] = findMax(dp[i-1][j],dp[i-1][j-vol[i-1]*k]+cost[i-1]*k)
			}*/
			dp[i][j] = dp[i-1][j]
			if j>=vol[i-1]{
				dp[i][j] = findMax(dp[i][j],dp[i][j-vol[i-1]]+cost[i-1])
			}
		}
	}
	return dp[N][V]
}

func bag(vol,cost []int,N,V int) int{
	dp := make([][]int,N+1)
	for i:=0;i<=N;i++{
		dp[i] = make([]int,V+1)
	}
	for i:=1;i<=N;i++{
		for j:=1;j<=V;j++{
			if vol[i-1]>j{
				dp[i][j] = dp[i-1][j]
			}else {
				dp[i][j] = findMax(dp[i-1][j],dp[i-1][j-vol[i-1]]+cost[i-1])
			}
		}
	}
	return dp[N][V]
}

/*func maxScore(SLA,score []int) []int {
	dp := make([][]int,len(SLA))
	haveChoose := make([]bool,len(SLA))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int,len(SLA))
	}
	for i:=0;i<len(SLA);i++{
		for j:=0;j<SLA[i];j++{
			dp[i][j] = score[i]
		}
	}
	//每一列判断
	for j:=0;j<len(dp)-1;j++{
		max := math.MinInt
		index := 0
		flag := false
		for i:=0;i<len(dp);i++{
			if dp[i][j] != -1 && dp[i][j]>=max{
				if dp[i][j]==max{
					if SLA[i]<SLA[index]{
						index = i
					}
				}else{
					max = dp[i][j]
					index = i
				}
			}
			if dp[i][j]>0{
				flag = true
			}
		}
		for i:=0;i<len(dp);i++{
			if dp[i][j+1]>0{
				dp[i][j+1] = dp[i][j+1]+max
			}
		}
		dp[index][j+1] = -1
		sum = sum + max
		if flag==false{
			break
		}
		res = append(res,index)
		haveChoose[index] = true
	}

	return res
}
*/

func maxScore(SLA,score []int) int {
	weight := make([][]int,len(SLA))
	for i:=0;i<len(weight);i++{
		weight[i] = make([]int,len(SLA))
	}
	for i:=0;i<len(SLA);i++{
		for j:=0;j<SLA[i];j++{
			weight[i][j] = score[i]
		}
	}
	return 0
}

func test(weight,skip []int,n int){
	dp := make([][]int,n)
	for i:=0;i<n;i++{
		dp[i] = make([]int,n)
	}
	dp[0][0] = weight[0]
	for i:=1;i<n;i++{
		for j:=0;j<n;j++{
			if i==j {
				if j-skip[i]-1<0{
					dp[i][j] = findMax(dp[i-1][j],weight[i])
				}else{
					dp[i][j] = findMax(dp[i-1][j],dp[i-1][j-skip[i]-1]+weight[i])
				}
			}else{
				if j-skip[i]-1<0{
					dp[i][j] = dp[i-1][j]
				}else{
					dp[i][j] = findMax(dp[i-1][j],dp[i-1][j-skip[i]-1]+weight[i])
				}
			}
		}
	}
	res := 0
	for i:=0;i<len(dp[n-1]);i++{
		if res<dp[n-1][i]{
			res = dp[n-1][i]
		}
	}
	fmt.Println(res)
}

func cutWater(weight,skip []int,n int){
	dp := make([]int,n)
	for i:=0;i<n;i++{
		dp[i] = weight[i]
	}
	for i:=0;i<n;i++{
		step:=i
		for step<n{
			if step+skip[i]+1<n{
				dp[i] = dp[i]+weight[step+skip[i]+1]
			}
			step = step+skip[i]+1
		}
	}
	res := dp[0]
	for i:=0;i<len(dp);i++{
		if res<dp[i]{
			res = dp[i]
		}
	}
	fmt.Println(res)
}