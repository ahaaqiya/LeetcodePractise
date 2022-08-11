package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)
	var str string
	fmt.Scanln(&str)
	latt := 0
	ratt := 0
	rdefend := 0
	for i,v := range str{
		if v =='0'{
			ratt += i+1
		}else {
			rdefend += i+1
		}
	}
	res := rdefend
	for i:=0;i<n;i++ {
		if str[i]=='0'{
			latt += i+1
		}else {
			rdefend -=i+1
		}
		res = findMin(res,abs(latt-rdefend))
	}
	fmt.Println(res)
}

func findMin(a,b int) int {
	if a>b{
		return b
	}else {
		return a
	}
}

func abs(a int) int  {
	if a>=0{
		return a
	}else {
		return -a
	}
}