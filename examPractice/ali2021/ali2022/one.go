package main

import (
	"fmt"
)

func recurMethod(i int) int{
	if i<=1{
		return 1
	}
	res := recurMethod(i-1) + recurMethod(i-2)
	fmt.Println(res)
	return res
}

func resur(i int) {
	if i>4{
		return
	}
	fmt.Println("start recur call",i)
	resur(i+1)
	fmt.Println("end recur call",i)
}

func main()  {
	var a int
	a =(1 << 63) - 1
	fmt.Printf("%t",a)
}