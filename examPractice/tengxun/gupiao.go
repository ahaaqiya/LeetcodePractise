package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*func dpdp(gupiao *[]int, start, end, budget int) (int, int) {
	dp := make([][]int, end-start+1)
	for i := 0; i < end-start+1; i++ {
		dp[i] = make([]int, end-start+1)
	}
	dp[0][0] = budget
	var res, resM int
	if budget >= (*gupiao)[start] {
		dp[0][1] = budget - (*gupiao)[start]
		res = 1
		resM = dp[0][1]
	} else {
		res = 0
		resM = 0
	}
	// dp[i-start][j]表示买入第i个，当前持有j个时的预算
	for i := start + 1; i < end; i++ {
		for j := 0; j < i-start+1; j++ {
			// 没买
			dp[i-start][j] = dp[i-start-1][j]
			//买了
			if dp[i-start-1][j] >= (*gupiao)[i] {
				dp[i-start][j+1] = dp[i-start-1][j] - (*gupiao)[i]
				if j+1 > res {
					res = j + 1
					resM = dp[i-start][j+1]
				}
			}
		}
	}
	return res, resM
}
*/
func main() {
	// var m, n int
	// fmt.Scan(&n, &m)
	// gupiao := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&gupiao[i])
	// }
	n, m := 6, 2
	gupiao := []int{2, 3, 1, 1, 1, 2}

	// // 定义一个数组，保存某个价格之后股票是否上涨
	// max := gupiao[len(gupiao)-1]
	// huizhang := make([]bool, n)
	// for i := n - 2; i >= 0; i-- {
	// 	if gupiao[i] > max {
	// 		max = gupiao[i]
	// 	} else {
	// 		huizhang[i] = true
	// 	}
	// }

	// // 在会涨的区间尽量入手多个股票
	// // 闭区间
	// start, end := 0, 0
	// curM := m
	// gpNum := 0
	// for i := 0; i < n; i++ {
	// 	if huizhang[i] == false {
	// 		// 左闭右开
	// 		end = i
	// 		// dp(i,j)
	// 		gpNum, curM = dpdp(&gupiao, start, end, curM)
	// 		curM += gpNum * gupiao[i]
	// 		start = end + 1
	// 	}
	// }
	// fmt.Println(curM)

	//传统方法，dp[i][j]表示在i天持有j个股票的本金
	var sell int
	dp := make([][]int, n+1)
	canreach := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		canreach[i] = make([]bool, n+1)
		dp[i][0] = m
		canreach[i][0] = true
		// dp[i][i] = m
	}

	for i := 1; i <= n; i++ {
		// 更新全卖出去的dp
		sell = dp[i-1][0]
		for k := 1; k < i; k++ {
			if canreach[i-1][k] {
				sell = max(dp[i-1][k]+(k)*gupiao[i-1], sell)
			}
		}
		dp[i][0] = sell
		for j := 1; j < i+1; j++ {
			canreach[i][j] = canreach[i-1][j]
			// 三个：不买dp[i][j-1]买了dp[i][j-1]-gupiao[i-1]卖了max(dp[i][j...])
			// 不买
			nobuy := dp[i-1][j]
			// 卖了
			sell = 0
			for k := j + 1; k < i; k++ {
				if canreach[i-1][k] {
					sell = max(dp[i-1][k]+(k-j)*gupiao[i-1], sell)
				}
			}
			// 买了
			buy := -1
			if gupiao[i-1] <= dp[i][j-1] {
				buy = dp[i-1][j-1] - gupiao[i-1]
				canreach[i][j] = true
			}
			dp[i][j] = max(max(nobuy, sell), buy)

		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(canreach[i])
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	fmt.Println(dp[len(dp)-1][0])

}
