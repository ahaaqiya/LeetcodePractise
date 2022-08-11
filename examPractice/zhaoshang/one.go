package main

import (
	"fmt"
	"math"
)

func main()  {
	var n int
	fmt.Scan(&n)
	var m int
	fmt.Scan(&m)
	price := make([][]int,m)
	for i:=0;i<m;i++{
		price[i] = make([]int,3)
		for j:=0;j<3;j++{
			fmt.Scan(&price[i][j])
		}
	}
	bug := make([]int,n)
	for i:=0;i<n;i++{
		bug[i] = math.MaxInt
	}
	for i:=0;i<m;i++{
		for j:=price[i][0]-1;j<=price[i][1]-1;j++{
			if bug[j]>price[i][2]{
				bug[j] = price[i][2]
			}
		}
	}
	sum := 0
	for i:=0;i<len(bug);i++{
		sum += bug[i]
	}
	fmt.Println(sum)
}



func findMax(a,b int) int {
	if a>b{
		return a
	}else {
		return b
	}
}