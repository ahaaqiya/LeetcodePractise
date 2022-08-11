package main

import "fmt"

func main()  {
	var n,m int
	fmt.Scanln(&n,&m)
	A := make([]int,n)
	B := make([]int,m)
	for i:=0;i<n;i++{
		fmt.Scan(&A[i])
	}
	for i:=0;i<m;i++ {
		fmt.Scan(&B[i])
	}
	count := 0
	contA := 0
	contA2 := 0
	contB2 := 0
	for i:=0;i<n;i++{
		if A[i]%4==0 {
			count += m
		}else{
			contA++
			if A[i]%4==2{
				contA2++
			}
		}
	}
	for i:=0;i<m;i++{
		if B[i]%4==0{
			count += contA
		}else if B[i]%4==2{
			contB2++
		}
	}
	count = count + contB2*contA2
	fmt.Println(count)
}