package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	var n int
	fmt.Scanln(&n)
	strArr := make([]string,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&strArr[i])
	}
	res := make([]int,len(strArr[0]))
	for i:=0;i<len(strArr[0]);i++{
		var str string
		for j:=0;j<n;j++{
			str = str + string(strArr[j][i])
		}
		num,_ := strconv.Atoi(str)
		res[i] = num
	}
	sort.Ints(res)
	for i:=0;i<len(res)-1;i++{
		fmt.Printf("%d ",res[i])
	}
	fmt.Println(res[len(res)-1])
}

/*func getNumber( a []int ) int {
	// write code here
	len := len(a)
	index := 0
	for i:=len;i>=1;i--{
		if judge(i){
			index = i
			break
		}
	}
	return a[index-1]
}*/

func getNumber( a []int ) int {
	// write code here
	n := len(a)
	primeL := judge(n+1)
	primeL[0],primeL[1] = false,false
	for len(a)>1 {
		tmp := make([]int,0)
		for i:=1;i<=len(a);i++{
			if primeL[i]{
				tmp = append(tmp,a[i-1])
			}
		}
		a = tmp
	}
	return a[0]
}


func judge(n int) []bool {
	// write code here
	isPrime := make([]bool,n)
	for i := range isPrime{
		isPrime[i] = true
	}
	for i:=2;i<n;i++{
		if isPrime[i]{
			for j:=2*i;j<n;j+=i{
				isPrime[j] = false
			}
		}
	}
	return isPrime
}
