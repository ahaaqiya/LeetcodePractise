package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	right := maxDepth(root.Right)+1
	left := maxDepth(root.Left)+1
	return findMax(right,left)
}

func main()  {
	monthArr := make([]string,12)
	costArr := make([]float64,12)
	outArr := make([]float64,12)
	max := float64(math.MinInt)
	min := float64(math.MaxInt)
	index1,index2 := 0,0
	for i:=0;i<12;i++{
		var mn string
		fmt.Scanln(&mn)
		month,num :=findMandN(mn)
		monthArr[i] = month
		sum := 0.00
		outSum := 0.00
		//fmt.Println("输入支出")
		for j:=0;j<int(num);j++{
			var cost string
			fmt.Scan(&cost)
			_,s := findMandN(cost)
			sum += s
			if s<0{
				outSum += s
			}
		}
		if sum>max{
			max = sum
			index1 = i
		}
		if outSum<min{
			min = outSum
			index2 = i
		}
		costArr[i] = sum
		outArr[i] = outSum
	}
	for i:=0;i<12;i++{
		if costArr[i]>0{
			fmt.Printf("%s:%.2f",monthArr[i],costArr[i])
		}else {
			fmt.Printf("%s:%.2f",monthArr[i],costArr[i])
		}
	}
	fmt.Println(monthArr[index2],monthArr[index1])
}

func findMandN(s string) (string,float64) {
	i:=0
	for i=0;i<len(s);i++{
		if s[i]==':'{
			break
		}
	}
	name := s[:i]
	num,_ := strconv.ParseFloat(s[i+1:],64)
	return name,float64(num)
}

